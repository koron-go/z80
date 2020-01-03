package z80

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
