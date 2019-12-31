package z80

var arith16 = []*OPCode{

	{
		N: "ADD HL, ss",
		C: []Code{
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.HL.U16()
			x := cpu.reg16ss(codes[0] >> 4).U16()
			cpu.HL.SetU16(cpu.addU16(a, x))
		},
	},

	{
		N: "ADC HL, ss",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4a, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.HL.U16()
			x := cpu.reg16ss(codes[0] >> 4).U16()
			cpu.HL.SetU16(cpu.adcU16(a, x))
		},
	},

	{
		N: "SBC HL, ss",
		C: []Code{
			{0xed, 0x00, nil},
			{0x42, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.HL.U16()
			x := cpu.reg16ss(codes[0] >> 4).U16()
			cpu.HL.SetU16(cpu.sbcU16(a, x))
		},
	},

	{
		N: "ADD IX, pp",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.IX
			x := cpu.reg16pp(codes[0] >> 4)
			cpu.IX = cpu.addU16(a, x)
		},
	},

	{
		N: "ADD IY, pp",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.IY
			x := cpu.reg16pp(codes[0] >> 4)
			cpu.IY = cpu.addU16(a, x)
		},
	},

	{
		N: "INC ss",
		C: []Code{
			{0x03, 0x30, vReg16_4},
		},
		T: []int{6},
		F: func(cpu *CPU, codes []uint8) {
			ss := cpu.reg16ss(codes[0] >> 4)
			ss.SetU16(cpu.incU16(ss.U16()))
		},
	},

	{
		N: "INC IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x23, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IX = cpu.incU16(cpu.IX)
		},
	},

	{
		N: "INC IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x23, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IY = cpu.incU16(cpu.IY)
		},
	},

	{
		N: "DEC ss",
		C: []Code{
			{0x0b, 0x30, vReg16_4},
		},
		T: []int{6},
		F: func(cpu *CPU, codes []uint8) {
			ss := cpu.reg16ss(codes[0] >> 4)
			ss.SetU16(cpu.decU16(ss.U16()))
		},
	},

	{
		N: "DEC IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2b, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IX = cpu.decU16(cpu.IX)
		},
	},

	{
		N: "DEC IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2b, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IY = cpu.decU16(cpu.IY)
		},
	},
}
