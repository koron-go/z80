package z80

import "reflect"

// MapRAM implements Memory interface with a map.
type MapRAM map[uint16]uint8

// Get gets a byte at addr of memory.
func (mm MapRAM) Get(addr uint16) uint8 {
	v, ok := mm[addr]
	if !ok {
		return 0
	}
	return v
}

// Set sets a byte at addr of memory.
func (mm MapRAM) Set(addr uint16, v uint8) {
	mm[addr] = v
}

// Equal checks two MapRAM same or not.
func (mm MapRAM) Equal(a0 interface{}) bool {
	a, ok := a0.(MapRAM)
	if !ok {
		return false
	}
	return reflect.DeepEqual(mm, a)
}

// Clone creates a clone of this.
func (mm MapRAM) Clone() MapRAM {
	cl := MapRAM{}
	for k, v := range mm {
		cl[k] = v
	}
	return cl
}

// Clear removes all.
func (mm MapRAM) Clear() {
	for k := range mm {
		delete(mm, k)
	}
}

// Put puts "data" block from addr.
func (mm MapRAM) Put(addr uint16, data ...uint8) MapRAM {
	for _, v := range data {
		mm[addr] = v
		addr++
	}
	return mm
}
