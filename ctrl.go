package z80

import "math/bits"

var opHALT = &OPCode{
	N: "HALT",
	C: []Code{
		{0x76, 0x00, nil},
	},
	T: []int{4},
	F: func(cpu *CPU, codes []uint8) {
		// nothing todo.
	},
}

var opEI = &OPCode{
	N: "EI",
	C: []Code{
		{0xfb, 0x00, nil},
	},
	T: []int{4},
	F: func(cpu *CPU, codes []uint8) {
		cpu.IFF1 = true
		cpu.IFF2 = true
	},
}

var ctrl = []*OPCode{

	{
		N: "DAA",
		C: []Code{
			{0x27, 0x00, nil},
		},
		T: []int{4},
		F: opDAA,
	},

	{
		N: "CPL",
		C: []Code{
			{0x2F, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.AF.Hi = ^cpu.AF.Hi
			cpu.flagUpdate(FlagOp{}.Set(H).Set(N))
		},
	},

	{
		N: "NEG",
		C: []Code{
			{0xed, 0x00, nil},
			{0x44, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			v := ^a + 1
			cpu.AF.Hi = v
			cpu.flagUpdate(FlagOp{}.
				Put(S, v&0x80 != 0).
				Put(Z, v == 0).
				// FIXME: check and fix H flag behavior.
				Put(H, false).
				Put(PV, a == 0x80).
				Set(N).
				Put(C, a != 0))
		},
	},

	{
		N: "CCF",
		C: []Code{
			{0x3f, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			c := cpu.flag(C)
			cpu.flagUpdate(FlagOp{}.Put(H, c).Reset(N).Put(C, !c))
		},
	},

	{
		N: "SCF",
		C: []Code{
			{0x37, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.flagUpdate(FlagOp{}.Reset(H).Reset(N).Set(C))
		},
	},

	{
		N: "NOP",
		C: []Code{
			{0x00, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			// do nothing.
		},
	},

	opHALT,

	{
		N: "DI",
		C: []Code{
			{0xf3, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IFF1 = false
			cpu.IFF2 = false
		},
	},

	opEI,

	{
		N: "IM 0",
		C: []Code{
			{0xed, 0x00, nil},
			{0x46, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IM = 0
		},
	},

	{
		N: "IM 1",
		C: []Code{
			{0xed, 0x00, nil},
			{0x56, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IM = 1
		},
	},

	{
		N: "IM 2",
		C: []Code{
			{0xed, 0x00, nil},
			{0x5e, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IM = 2
		},
	},
}

func in(v, bottom, up uint8) bool {
	return v >= bottom && v <= up
}

func opDAA(cpu *CPU, codes []uint8) {
	var a, b uint8
	var fo FlagOp
	hc := cpu.flag(H)
	a = cpu.AF.Hi
	h4 := (a >> 4) & 0x0f
	l4 := a & 0x0f
	if !cpu.flag(N) {
		// addition adjustment.
		if !cpu.flag(C) {
			if in(h4, 0, 9) && !hc && in(l4, 0, 9) {
				b = 0x00
				fo = fo.Reset(C)
			} else if in(h4, 0, 8) && !hc && in(l4, 10, 15) {
				b = 0x06
				fo = fo.Reset(C)
			} else if in(h4, 0, 9) && hc && in(l4, 0, 3) {
				b = 0x06
				fo = fo.Reset(C)
			} else if in(h4, 10, 15) && !hc && in(l4, 0, 9) {
				b = 0x60
				fo = fo.Set(C)
			} else if in(h4, 9, 15) && !hc && in(l4, 10, 15) {
				b = 0x66
				fo = fo.Set(C)
			} else if in(h4, 10, 15) && hc && in(l4, 0, 3) {
				b = 0x66
				fo = fo.Set(C)
			}
		} else {
			if in(h4, 0, 2) && !hc && in(l4, 0, 9) {
				b = 0x60
				fo = fo.Set(C)
			} else if in(h4, 0, 2) && !hc && in(l4, 10, 15) {
				b = 0x66
				fo = fo.Set(C)
			} else if in(h4, 0, 3) && hc && in(l4, 0, 3) {
				b = 0x66
				fo = fo.Set(C)
			}
		}
	} else {
		// subtraction adjustment.
		if !cpu.flag(C) {
			if in(h4, 0, 9) && !hc && in(l4, 0, 9) {
				b = 0x00
				fo = fo.Reset(C)
			} else if in(h4, 0, 8) && hc && in(l4, 6, 15) {
				b = 0xfa
				fo = fo.Reset(C)
			}
		} else {
			if in(h4, 7, 15) && !hc && in(l4, 0, 9) {
				b = 0xa0
				fo = fo.Set(C)
			} else if in(h4, 6, 15) && hc && in(l4, 6, 15) {
				// different from the manual, it is `in(h4, 6, 7)`
				b = 0x9a
				fo = fo.Set(C)
			}
		}
	}
	v := uint16(a) + uint16(b)
	// update regsiter && flags.
	cpu.AF.Hi = uint8(v)
	cpu.flagUpdate(fo.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(cpu.AF.Hi)%2 == 0))
}
