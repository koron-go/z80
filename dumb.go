package z80

type DumbMemory []uint8

func (dm DumbMemory) Get(addr uint16) uint8 {
	if int(addr) >= len(dm) {
		return 0
	}
	return dm[addr]
}

func (dm DumbMemory) Set(addr uint16, value uint8) {
	if int(addr) >= len(dm) {
		return
	}
	dm[addr] = value
}

var _ Memory = DumbMemory(nil)

type DumbIO []uint8

func (dio DumbIO) In(addr uint8) uint8 {
	if int(addr) >= len(dio) {
		return 0
	}
	return dio[addr]
}

func (dio DumbIO) Out(addr uint8, value uint8) {
	if int(addr) >= len(dio) {
		return
	}
	dio[addr] = value
}

var _ IO = DumbIO(nil)
