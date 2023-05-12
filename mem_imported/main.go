package main

import (
	"fmt"

	"net/http"
	_ "net/http/pprof"

	"github.com/sirupsen/logrus"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	// Server for pprof
	go func() {
		fmt.Println(http.ListenAndServe(":6060", nil))
	}()

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	wat := []byte(`
	(module
		;; Import memory
		(import "env" "memory" (memory 1))
	  
		;; Function to add 2 to a number and store it back in memory
		(func $plusTwo (export "plusTwo") (param $ptr i32) (param $len i32)
		  (local $i i32)
		  (local $value i32)
		  (set_local $i (i32.const 0))
		  (block $loop
			(loop $inner
			  (br_if $loop
				(i32.ge_u (get_local $i) (get_local $len)))
			  (set_local $value
				(i32.add
				  (i32.load8_u
					(i32.add
					  (get_local $ptr)
					  (get_local $i)))
				  (i32.const 2)))
			  (i32.store8
				(i32.add
				  (get_local $ptr)
				  (get_local $i))
				(get_local $value))
			  (set_local $i
				(i32.add
				  (get_local $i)
				  (i32.const 1)))
			  (br $inner)))))
	`)
	module, err := wasmer.NewModule(store, wat)
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

	importsMap := make(map[string]wasmer.IntoExtern)
	importsMap["memory"] = memory

	imports := wasmer.NewImportObject()
	imports.Register("env", importsMap)

	instance, err := wasmer.NewInstance(module, imports)
	if err != nil {
		logrus.Fatalf("Failed to instantiate module: %s", err)
	}

	plustTwoFunc, err := instance.Exports.GetFunction("plusTwo")
	if err != nil {
		// The wasm is missing `transform` and is therefore invalid
		logrus.Fatalf("accumulate function not found: %s", err)
	}

	for {
		fmt.Printf("mem size at begining: %d\n", len(memory.Data()))

		//Store the data into memory
		mem := memory.Data()
		mem[0] = 10

		if plustTwoFunc == nil {
			fmt.Printf("function is nil for some reason...")
			plustTwoFunc, err = instance.Exports.GetFunction("plusTwo")
			if err != nil {
				// The wasm is missing `transform` and is therefore invalid
				logrus.Fatalf("accumulate function not found: %s", err)
			}
		}

		// Calls the exported function with the start and length of the inputJson in memory
		// The return value is an integer containing the start and length of the transformed json in memory
		_, err = plustTwoFunc(int32(0), int32(1))
		if err != nil {
			logrus.Fatalf("failed to transform: %s", err)
		}

		mem = memory.Data()
		fmt.Printf("%d\n", mem[0])

		//time.Sleep(3 * time.Second)
	}
}
