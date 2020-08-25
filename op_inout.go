package z80

import "math/bits"

var inout = []*OPCode{

	{
		N: "IN A, (n)",
		C: []Code{
			{0xdb, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 4},
		F: opINAnP,
	},

	{
		N: "IN r (C)",
		C: []Code{
			{0xed, 0x00, nil},
			{0x40, 0x38, vReg8_3},
		},
		T: []int{4, 4, 4},
		F: opINrCP,
	},

	{
		N: "INI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa2, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
		F: opINI,
	},

	{
		N: "INIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb2, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
		F: opINIR,
	},

	{
		N: "IND",
		C: []Code{
			{0xed, 0x00, nil},
			{0xaa, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
		F: opIND,
	},

	{
		N: "INDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xba, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
		F: opINDR,
	},

	{
		N: "OUT (n), A",
		C: []Code{
			{0xd3, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 4},
		F: opOUTnPA,
	},

	{
		N: "OUT (C), r",
		C: []Code{
			{0xed, 0x00, nil},
			{0x41, 0x38, vReg8_3},
		},
		T: []int{4, 4, 4},
		F: opOUTCPr,
	},

	{
		N: "OUTI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa3, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
		F: opOUTI,
	},

	{
		N: "OTIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb3, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
		F: opOTIR,
	},

	{
		N: "OUTD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xab, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
		F: opOUTD,
	},

	{
		N: "OTDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xbb, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
		F: opOTDR,
	},
}

func opINAnP(cpu *CPU, codes []uint8) {
	cpu.AF.Hi = cpu.ioIn(codes[1])
}

func opINrCP(cpu *CPU, codes []uint8) {
	v := cpu.ioIn(cpu.BC.Lo)
	// FIXME: support 0x06 to apply flags only.
	r := cpu.regP(codes[1] >> 3)
	*r = v
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N))
}

func opINI(cpu *CPU, codes []uint8) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opINIR(cpu *CPU, codes []uint8) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func opIND(cpu *CPU, codes []uint8) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opINDR(cpu *CPU, codes []uint8) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func opOUTnPA(cpu *CPU, codes []uint8) {
	cpu.ioOut(codes[1], cpu.AF.Hi)
}

func opOUTCPr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1] >> 3)
	cpu.ioOut(cpu.BC.Lo, *r)
}

func opOUTI(cpu *CPU, codes []uint8) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opOTIR(cpu *CPU, codes []uint8) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func opOUTD(cpu *CPU, codes []uint8) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opOTDR(cpu *CPU, codes []uint8) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}
