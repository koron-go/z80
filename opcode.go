package z80

// Code is definition of an operation code.  It is consist from V (value) and M
// (mask).  0 bits in M are for constant bits, 1 bits are variable bits for
// operation code.
type Code struct {
	V uint8
	M uint8
}

// OPCode defines opration code and its function.
type OPCode struct {
	// N is string presentation (=label) of opcode.
	N string

	// C is codes definition.
	C []Code

	// T is T cycle and its length is M cycle.
	T []int

	// F is the function of opcode.
	F func(*CPU, []uint8)
}

// addrOff apply offset to address.
func addrOff(addr uint16, off uint8) uint16 {
	return addr + uint16(int16(int8(off)))
}

func toU16(l, u uint8) uint16 {
	return (uint16(u) << 8) | uint16(l)
}
