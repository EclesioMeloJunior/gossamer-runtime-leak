package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"math"

	"net/http"
	_ "net/http/pprof"

	"github.com/klauspost/compress/zstd"
	"github.com/sirupsen/logrus"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Runtime struct {
	mem       *MemoryImpl
	allocator *FreeingBumpHeapAllocator
	instance  *wasmer.Instance
}

func main() {
	// Server for pprof
	go func() {
		fmt.Println(http.ListenAndServe(":6060", nil))
	}()

	// Retrieve WASM binary
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
	var memImport *wasmer.ImportType
	for _, im := range moduleImports {
		if im.Name() == "memory" {
			memImport = im
			break
		}
	}

	memoryType := memImport.Type().IntoMemoryType()
	memory := wasmer.NewMemory(store, memoryType)
	memoryImpl := &MemoryImpl{memory}

	rt := &Runtime{
		mem: memoryImpl,
	}

	imports := importsNodeRuntime(store, memory, rt)
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
		inputPtr, err := rt.allocator.Allocate(memoryImpl, dataLength)
		if err != nil {
			logrus.Fatalf("allocating input memory: %w", err)
		}

		//Store the data into memory
		memory := rt.mem.Data()
		copy(memory[inputPtr:inputPtr+dataLength], input)

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

		// Calls the exported function with the start and length of the inputJson in memory
		// The return value is an integer containing the start and length of the transformed json in memory
		wasmValue, err := coreVersion(castedInputPointer, castedDataLength)
		if err != nil {
			logrus.Fatalf("failed to transform: %s", err)
		}

		wasmValueAsI64 := wasmer.NewI64(wasmValue)
		outputPtr, outputLength := splitPointerSize(wasmValueAsI64.I64())
		memory = rt.mem.Data() // call Data() again to get larger slice

		allocatedData := make([]byte, outputLength)
		copy(allocatedData[:], memory[outputPtr:outputPtr+outputLength])

		fmt.Printf("%b\n", allocatedData)
		rt.allocator.Clear()
	}
}

func safeCastInt32(value uint32) (int32, error) {
	if value > math.MaxInt32 {
		return 0, errors.New("out of bounds")
	}
	return int32(value), nil
}

// splitPointerSize converts an int64 pointer size to an
// uint32 pointer and an uint32 size.
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
