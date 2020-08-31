package z80

func (cpu *CPU) flagUpdate(fo FlagOp) {
	fo.ApplyOn(&cpu.AF.Lo)
}

const (
	// C is an index for carry flag.
	C = 0

	// N is an index for add/subtract flag.
	N = 1

	// PV is an index for parity/overflow flag.
	PV = 2

	// X3 is reserved index for unused flag.
	X3 = 3

	// H is an index for half carry flag.
	H = 4

	// X5 is reserved index for unused flag.
	X5 = 5

	// Z is an index for zero flag.
	Z = 6

	// S is an index for sign flag.
	S = 7
)

// Flag gets a bit of flag register value.
func Flag(b uint8, n int) bool {
	return b&(0x01<<n) != 0
}

// FlagOp provides flag operation.  At initial this will keep all bits.
type FlagOp struct {
	Nand uint8
	Or   uint8
}

// ApplyOn applies flag operation on uint8 in place.
func (fo FlagOp) ApplyOn(v *uint8) {
	*v = *v&^fo.Nand | fo.Or
}

// Keep marks bit-n as keeping.
func (fo FlagOp) Keep(n int) FlagOp {
	var m uint8 = ^(0x01 << n)
	fo.Nand &= m
	fo.Or &= m
	return fo
}

// Set marks bit-n as being 1.
func (fo FlagOp) Set(n int) FlagOp {
	var b uint8 = 0x01 << n
	fo.Nand |= b
	fo.Or |= b
	return fo
}

// Reset marks bit-n as being 0.
func (fo FlagOp) Reset(n int) FlagOp {
	var b uint8 = 0x01 << n
	fo.Nand |= b
	fo.Or &= ^b
	return fo
}

// Put modify bit-n with boolean value.
func (fo FlagOp) Put(n int, v bool) FlagOp {
	if v {
		return fo.Set(n)
	}
	return fo.Reset(n)
}
