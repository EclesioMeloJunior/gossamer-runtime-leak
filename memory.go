package main

import (
	"errors"
	"math"

	"github.com/wasmerio/wasmer-go/wasmer"
)

// Memory is a thin wrapper around Wasmer memory to support
// Gossamer runtime.Memory interface
type MemoryImpl struct {
	memory *wasmer.Memory
}

func checkBounds(value uint64) (uint32, error) {
	if value > math.MaxUint32 {
		return 0, errors.New("out of bounds")
	}
	return uint32(value), nil
}

// Data returns the memory's data
func (m *MemoryImpl) Data() []byte {
	return m.memory.Data()
}

// Length returns the memory's length
func (m *MemoryImpl) Length() uint32 {
	value, err := checkBounds(uint64(m.memory.DataSize()))
	if err != nil {
		panic(err)
	}
	return value
}

// Grow grows the memory by the given number of pages
func (m *MemoryImpl) Grow(numPages uint32) error {
	ok := m.memory.Grow(wasmer.Pages(numPages))
	if !ok {
		return errors.New("cannot grow mem")
	}
	return nil
}
