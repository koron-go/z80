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

// port from WebMSX.
// See https://github.com/ppeccin/WebMSX/blob/654e3aa303e84404fba4a89d5fa21fae32753cf5/src/main/msx/cpu/CPU.js#L1010-L1030
func opDAA(cpu *CPU, codes []uint8) {
	r := cpu.AF.Hi
	c := cpu.flag(C)
	if cpu.flag(N) {
		if cpu.flag(H) || (cpu.AF.Hi&0x0f) > 9 {
			r -= 0x06
		}
		if c || (cpu.AF.Hi > 0x99) {
			r -= 0x60
		}
	} else {
		if cpu.flag(H) || (cpu.AF.Hi&0x0f) > 9 {
			r += 0x06
		}
		if c || (cpu.AF.Hi > 0x99) {
			r += 0x60
		}
	}
	cpu.flagUpdate(FlagOp{}.
		Put(S, r&0x80 != 0).
		Put(Z, r == 0).
		Put(H, (cpu.AF.Hi^r)&0x10 != 0).
		Put(PV, bits.OnesCount8(r)%2 == 0).
		Put(C, c || cpu.AF.Hi > 0x99))
	cpu.AF.Hi = r
}
