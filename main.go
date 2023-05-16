package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"net/http"
	_ "net/http/pprof"

	"github.com/ChainSafe/gossamer/lib/runtime"
	"github.com/ChainSafe/gossamer/pkg/scale"
	"github.com/klauspost/compress/zstd"
	"github.com/sirupsen/logrus"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Runtime struct {
	wasmerMemory *wasmer.Memory
	mem          *MemoryImpl
	allocator    *FreeingBumpHeapAllocator
	instance     *wasmer.Instance
}

func main() {
	// Server for pprof
	go func() {
		fmt.Println(http.ListenAndServe(":6060", nil))
	}()

	wasmBytes, err := ioutil.ReadFile("./westend_runtime-v9381.compact.compressed.wasm")
	if err != nil {
		logrus.Fatalf("Failed to read wasm file: %s", err)
	}

	wasmBytes, err = decompressWasm(wasmBytes)
	if err != nil {
		panic(err)
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		logrus.Fatalf("failed to compile webassembly binary: %s", err)
	}

	moduleImports := module.Imports()
	var memType *wasmer.MemoryType
	for _, im := range moduleImports {
		fmt.Printf("import -> %s\n", im.Name())
		if im.Name() == "memory" {
			memType = im.Type().IntoMemoryType()
		}
	}

	memory := wasmer.NewMemory(store, memType)
	memoryImpl := NewMemImplementation(memory)

	rt := &Runtime{
		wasmerMemory: memory,
		mem:          memoryImpl,
	}

	imports := importsNodeRuntime(store, rt)
	if err != nil {
		panic(err)
	}

	instance, err := wasmer.NewInstance(module, imports)
	if err != nil {
		logrus.Fatalf("Failed to instantiate module: %s", err)
	}

	//set heap base for allocator, start allocating at heap base
	heapBase, err := instance.Exports.Get("__heap_base")
	if err != nil {
		logrus.Fatalf("failed to get __heap_base from instance: %s", err)
	}

	hb, err := heapBase.IntoGlobal().Get()
	if err != nil {
		logrus.Fatalf("failed to heapBase.IntoGlobal: %s", err)
	}

	rt.allocator = NewAllocator(memoryImpl, uint32(hb.(int32)))
	rt.instance = instance

	input := []byte{}
	for {
		dataLength := uint32(len(input))
		inputPtr, err := rt.allocator.Allocate(dataLength)
		if err != nil {
			logrus.Fatalf("allocating input memory: %w", err)
		}

		// Store the data into memory
		mdata := memory.Data()
		copy(mdata[inputPtr:inputPtr+dataLength], input)

		castedInputPointer, err := safeCastInt32(inputPtr)
		if err != nil {
			panic(err)
		}

		castedDataLength, err := safeCastInt32(dataLength)
		if err != nil {
			panic(err)
		}

		coreVersion, err := rt.instance.Exports.GetFunction("Core_version")
		if err != nil {
			// The wasm is missing `transform` and is therefore invalid
			logrus.Fatalf("transform function not found: %s", err)
		}

		returned, err := coreVersion(castedInputPointer, castedDataLength)
		if err != nil {
			logrus.Fatalf("failed to transform: %s", err)
		}

		wasmValueAsI64 := wasmer.NewI64(returned)
		outputPtr, outputLength := splitPointerSize(wasmValueAsI64.I64())
		mdata = memory.Data() // call Data() again to get larger slice

		allocatedData := make([]byte, outputLength)
		copy(allocatedData[:], mdata[outputPtr:outputPtr+outputLength])

		v := new(runtime.Version)
		err = scale.Unmarshal(allocatedData, v)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%+v", v)
		rt.allocator.Clear()
	}
}

func splitPointerSize(pointerSize int64) (ptr, size uint32) {
	return uint32(pointerSize), uint32(pointerSize >> 32)
}

func decompressWasm(code []byte) ([]byte, error) {
	compressionFlag := []byte{82, 188, 83, 118, 70, 219, 142, 5}
	if !bytes.HasPrefix(code, compressionFlag) {
		return code, nil
	}

	decoder, err := zstd.NewReader(nil)
	if err != nil {
		return nil, fmt.Errorf("creating zstd reader: %s", err)
	}

	return decoder.DecodeAll(code[len(compressionFlag):], nil)
}
