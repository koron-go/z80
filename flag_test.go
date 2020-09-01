package z80

import "testing"

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

// flagOp provides flag operation.  At initial this will keep all bits.
type flagOp struct {
	Nand uint8
	Or   uint8
}

// ApplyOn applies flag operation on uint8 in place.
func (fo flagOp) ApplyOn(v *uint8) {
	*v = *v&^fo.Nand | fo.Or
}

// Keep marks bit-n as keeping.
func (fo flagOp) Keep(n int) flagOp {
	var m uint8 = ^(0x01 << n)
	fo.Nand &= m
	fo.Or &= m
	return fo
}

// Set marks bit-n as being 1.
func (fo flagOp) Set(n int) flagOp {
	var b uint8 = 0x01 << n
	fo.Nand |= b
	fo.Or |= b
	return fo
}

// Reset marks bit-n as being 0.
func (fo flagOp) Reset(n int) flagOp {
	var b uint8 = 0x01 << n
	fo.Nand |= b
	fo.Or &= ^b
	return fo
}

// Put modify bit-n with boolean value.
func (fo flagOp) Put(n int, v bool) flagOp {
	if v {
		return fo.Set(n)
	}
	return fo.Reset(n)
}

func TestFlagOp_ApplyOn(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		op  flagOp
		v   []uint8
		exp []uint8
	}{
		{flagOp{0x00, 0x00}, []uint8{0x01, 0x00}, []uint8{0x01, 0x00}},
		{flagOp{0x01, 0x00}, []uint8{0x01, 0x00}, []uint8{0x00, 0x00}},
		{flagOp{0x01, 0x01}, []uint8{0x01, 0x00}, []uint8{0x01, 0x01}},
		// undefined behavior
		{flagOp{0x00, 0x01}, []uint8{0x01, 0x00}, []uint8{0x01, 0x01}},

		{flagOp{0x00, 0x00}, []uint8{0x10, 0x00}, []uint8{0x10, 0x00}},
		{flagOp{0x10, 0x00}, []uint8{0x10, 0x00}, []uint8{0x00, 0x00}},
		{flagOp{0x10, 0x10}, []uint8{0x10, 0x00}, []uint8{0x10, 0x10}},
		// undefined behavior
		{flagOp{0x00, 0x10}, []uint8{0x10, 0x00}, []uint8{0x10, 0x10}},

		{flagOp{0x00, 0x00}, []uint8{0x11}, []uint8{0x11}},
		{
			flagOp{0x11, 0x00},
			[]uint8{0x11, 0x10, 0x01, 0x00},
			[]uint8{0x00, 0x00, 0x00, 0x00},
		},
		{
			flagOp{0x11, 0x11},
			[]uint8{0x11, 0x10, 0x01, 0x00},
			[]uint8{0x11, 0x11, 0x11, 0x11},
		},
		// undefined behavior
		{
			flagOp{0x00, 0x11},
			[]uint8{0x11, 0x10, 0x01, 0x00},
			[]uint8{0x11, 0x11, 0x11, 0x11},
		},
	} {
		for i, v := range c.v {
			c.op.ApplyOn(&v)
			if v != c.exp[i] {
				t.Fatalf("failed %#v.ApplyOn(0x%02x): expect=0x%02x actual=0x%02x", c.op, c.v[i], c.exp[i], v)
			}
		}
	}
}
