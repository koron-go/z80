package z80

import "reflect"

// Memory is requirements interface for memory.
type Memory interface {
	Get(addr uint16) uint8
	Set(addr uint16, value uint8)
}

// IO is requirements interface for I/O.
type IO interface {
	In(addr uint8) uint8
	Out(addr uint8, value uint8)
}

// DumbMemory provides Memory interface as wrapper of []uint8
type DumbMemory []uint8

// Get gets a byte at addr of memory.
func (dm DumbMemory) Get(addr uint16) uint8 {
	if int(addr) >= len(dm) {
		return 0
	}
	return dm[addr]
}

// Set sets a byte at addr of memory.
func (dm DumbMemory) Set(addr uint16, value uint8) {
	if int(addr) >= len(dm) {
		return
	}
	dm[addr] = value
}

var _ Memory = DumbMemory(nil)

// DumbIO provides IO interface as wrapper of []uint8
type DumbIO []uint8

// In gets a byte at addr of IO
func (dio DumbIO) In(addr uint8) uint8 {
	if int(addr) >= len(dio) {
		return 0
	}
	return dio[addr]
}

// Out sets a byte at addr of IO
func (dio DumbIO) Out(addr uint8, value uint8) {
	if int(addr) >= len(dio) {
		return
	}
	dio[addr] = value
}

var _ IO = DumbIO(nil)

// MapMemory implements Memory interface with a map.
type MapMemory map[uint16]uint8

// Get gets a byte at addr of memory.
func (mm MapMemory) Get(addr uint16) uint8 {
	v, ok := mm[addr]
	if !ok {
		return 0
	}
	return v
}

// Set sets a byte at addr of memory.
func (mm MapMemory) Set(addr uint16, v uint8) {
	mm[addr] = v
}

// Equal checks two MapMemory same or not.
func (mm MapMemory) Equal(a0 interface{}) bool {
	a, ok := a0.(MapMemory)
	if !ok {
		return false
	}
	return reflect.DeepEqual(mm, a)
}

// Clone creates a clone of this.
func (mm MapMemory) Clone() MapMemory {
	cl := MapMemory{}
	for k, v := range mm {
		cl[k] = v
	}
	return cl
}

// Clear removes all.
func (mm MapMemory) Clear() {
	for k := range mm {
		delete(mm, k)
	}
}

// Put puts "data" block from addr.
func (mm MapMemory) Put(addr uint16, data ...uint8) MapMemory {
	for _, v := range data {
		mm[addr] = v
		addr++
	}
	return mm
}
