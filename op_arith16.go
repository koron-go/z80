package z80

var arith16 = []*OPCode{

	{
		N: "ADD HL, ss",
		C: []Code{
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 3},
		F: opADDHLss,
	},

	{
		N: "ADC HL, ss",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4a, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: opADCHLss,
	},

	{
		N: "SBC HL, ss",
		C: []Code{
			{0xed, 0x00, nil},
			{0x42, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: opSBCHLss,
	},

	{
		N: "ADD IX, pp",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: opADDIXpp,
	},

	{
		N: "ADD IY, rr",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: opADDIYrr,
	},

	{
		N: "INC ss",
		C: []Code{
			{0x03, 0x30, vReg16_4},
		},
		T: []int{6},
		F: opINCss,
	},

	{
		N: "INC IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x23, 0x00, nil},
		},
		T: []int{4, 6},
		F: opINCIX,
	},

	{
		N: "INC IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x23, 0x00, nil},
		},
		T: []int{4, 6},
		F: opINCIY,
	},

	{
		N: "DEC ss",
		C: []Code{
			{0x0b, 0x30, vReg16_4},
		},
		T: []int{6},
		F: opDECss,
	},

	{
		N: "DEC IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2b, 0x00, nil},
		},
		T: []int{4, 6},
		F: opDECIX,
	},

	{
		N: "DEC IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2b, 0x00, nil},
		},
		T: []int{4, 6},
		F: opDECIY,
	},
}

func opADDHLss(cpu *CPU, codes []uint8) {
	a := cpu.HL.U16()
	x := cpu.reg16ss(codes[0] >> 4).U16()
	cpu.HL.SetU16(cpu.addU16(a, x))
}

func opADCHLss(cpu *CPU, codes []uint8) {
	a := cpu.HL.U16()
	x := cpu.reg16ss(codes[1] >> 4).U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
}

func opSBCHLss(cpu *CPU, codes []uint8) {
	a := cpu.HL.U16()
	x := cpu.reg16ss(codes[1] >> 4).U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
}

func opADDIXpp(cpu *CPU, codes []uint8) {
	a := cpu.IX
	x := cpu.reg16pp(codes[1] >> 4)
	cpu.IX = cpu.addU16(a, x)
}

func opADDIYrr(cpu *CPU, codes []uint8) {
	a := cpu.IY
	x := cpu.reg16rr(codes[1] >> 4)
	cpu.IY = cpu.addU16(a, x)
}

func opINCss(cpu *CPU, codes []uint8) {
	ss := cpu.reg16ss(codes[0] >> 4)
	ss.SetU16(cpu.incU16(ss.U16()))
}

func opINCIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.incU16(cpu.IX)
}

func opINCIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.incU16(cpu.IY)
}

func opDECss(cpu *CPU, codes []uint8) {
	ss := cpu.reg16ss(codes[0] >> 4)
	ss.SetU16(cpu.decU16(ss.U16()))
}

func opDECIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.decU16(cpu.IX)
}

func opDECIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.decU16(cpu.IY)
}
