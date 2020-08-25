package z80

var exbtsg = []*OPCode{

	{
		N: "EX DE, HL",
		C: []Code{
			{0xeb, 0x00, nil},
		},
		T: []int{4},
		F: opEXDEHL,
	},

	{
		N: "EX AF, AF'",
		C: []Code{
			{0x08, 0x00, nil},
		},
		T: []int{4},
		F: opEXAFAF,
	},

	{
		N: "EXX",
		C: []Code{
			{0xd9, 0x00, nil},
		},
		T: []int{4},
		F: opEXX,
	},

	{
		N: "EX (SP), HL",
		C: []Code{
			{0xe3, 0x00, nil},
		},
		T: []int{4, 3, 4, 3, 5},
		F: opEXSPPHL,
	},

	{
		N: "EX (SP), IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe3, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3, 5},
		F: opEXSPPIX,
	},

	{
		N: "EX (SP), IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe3, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3, 5},
		F: opEXSPPIY,
	},

	{
		N: "LDI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa0, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: opLDI,
	},

	{
		N: "LDIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb0, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: opLDIR,
	},

	{
		N: "LDD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa8, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: opLDD,
	},

	{
		N: "LDDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb8, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: opLDDR,
	},

	{
		N: "CPI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa1, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: opCPI,
	},

	{
		N: "CPIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb1, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: opCPIR,
	},

	{
		N: "CPD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa9, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: opCPD,
	},

	{
		N: "CPDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb9, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: opCPDR,
	},
}

func opEXDEHL(cpu *CPU, codes []uint8) {
	cpu.HL, cpu.DE = cpu.DE, cpu.HL
}

func opEXAFAF(cpu *CPU, codes []uint8) {
	cpu.AF, cpu.Alternate.AF = cpu.Alternate.AF, cpu.AF
}

func opEXX(cpu *CPU, codes []uint8) {
	cpu.BC, cpu.Alternate.BC = cpu.Alternate.BC, cpu.BC
	cpu.DE, cpu.Alternate.DE = cpu.Alternate.DE, cpu.DE
	cpu.HL, cpu.Alternate.HL = cpu.Alternate.HL, cpu.HL
}

func opEXSPPHL(cpu *CPU, codes []uint8) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.HL.U16())
	cpu.HL.SetU16(v)
}

func opEXSPPIX(cpu *CPU, codes []uint8) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.IX)
	cpu.IX = v
}

func opEXSPPIY(cpu *CPU, codes []uint8) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.IY)
	cpu.IY = v
}

func opLDI(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de + 1)
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
}

func opLDIR(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de + 1)
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
	if bc != 0 {
		cpu.PC -= 2
	}
}

func opLDD(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de - 1)
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
}

func opLDDR(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de - 1)
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
	if bc != 0 {
		cpu.PC -= 2
	}
}

func opCPI(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
}

func opCPIR(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
	if bc != 0 && v != 0 {
		cpu.PC -= 2
	}
}

func opCPD(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
}

func opCPDR(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
	if bc != 0 && v != 0 {
		cpu.PC -= 2
	}
}
