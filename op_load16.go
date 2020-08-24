package z80

var load16 = []*OPCode{

	{
		N: "LD dd, nn",
		C: []Code{
			{0x01, 0x30, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3},
		F: opLDddnn,
	},

	{
		N: "LD IX, nn",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x21, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3},
		F: opLDIXnn,
	},

	{
		N: "LD IY, nn",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x21, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3},
		F: opLDIYnn,
	},

	{
		N: "LD HL, (nn)",
		C: []Code{
			{0x2a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3, 3},
		F: opLDHLnnP,
	},

	{
		N: "LD dd, (nn)",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4b, 0x30, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
		F: opLDddnnP,
	},

	{
		N: "LD IX, (nn)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
		F: opLDIXnnP,
	},

	{
		N: "LD IY, (nn)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
		F: opLDIYnnP,
	},

	{
		N: "LD (nn), HL",
		C: []Code{
			{0x22, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3, 3},
		F: opLDnnPHL,
	},

	{
		N: "LD (nn), dd",
		C: []Code{
			{0xed, 0x00, nil},
			{0x43, 0x30, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
		F: opLDnnPdd,
	},

	{
		N: "LD (nn), IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x22, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
		F: opLDnnPIX,
	},

	{
		N: "LD (nn), IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x22, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
		F: opLDnnPIY,
	},

	{
		N: "LD SP, HL",
		C: []Code{
			{0xf9, 0x00, nil},
		},
		T: []int{6},
		F: opLDSPHL,
	},

	{
		N: "LD SP, IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xf9, 0x00, nil},
		},
		T: []int{4, 6},
		F: opLDSPIX,
	},

	{
		N: "LD SP, IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xf9, 0x00, nil},
		},
		T: []int{4, 6},
		F: opLDSPIY,
	},

	{
		N: "PUSH qq",
		C: []Code{
			{0xc5, 0x30, nil},
		},
		T: []int{5, 3, 3},
		F: opPUSHqq,
	},

	{
		N: "PUSH IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe5, 0x00, nil},
		},
		T: []int{4, 5, 3, 3},
		F: opPUSHIX,
	},

	{
		N: "PUSH IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe5, 0x00, nil},
		},
		T: []int{4, 5, 3, 3},
		F: opPUSHIY,
	},

	{
		N: "POP qq",
		C: []Code{
			{0xc1, 0x30, nil},
		},
		T: []int{4, 3, 3},
		F: opPOPqq,
	},

	{
		N: "POP IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe1, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
		F: opPOPIX,
	},

	{
		N: "POP IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe1, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
		F: opPOPIY,
	},
}

func opLDddnn(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[0] >> 4)
	nn := toU16(codes[1], codes[2])
	dd.SetU16(nn)
}

func opLDIXnn(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IX = nn
}

func opLDIYnn(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IY = nn
}

func opLDHLnnP(cpu *CPU, codes []uint8) {
	nn := toU16(codes[1], codes[2])
	cpu.HL.SetU16(cpu.readU16(nn))
}

func opLDddnnP(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[1] >> 4)
	nn := toU16(codes[2], codes[3])
	dd.SetU16(cpu.readU16(nn))
}

func opLDIXnnP(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IX = cpu.readU16(nn)
}

func opLDIYnnP(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IY = cpu.readU16(nn)
}

func opLDnnPHL(cpu *CPU, codes []uint8) {
	nn := toU16(codes[1], codes[2])
	cpu.writeU16(nn, cpu.HL.U16())
}

func opLDnnPdd(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[1] >> 4)
	nn := toU16(codes[2], codes[3])
	cpu.writeU16(nn, dd.U16())
}

func opLDnnPIX(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.writeU16(nn, cpu.IX)
}

func opLDnnPIY(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.writeU16(nn, cpu.IY)
}

func opLDSPHL(cpu *CPU, codes []uint8) {
	cpu.SP = cpu.HL.U16()
}

func opLDSPIX(cpu *CPU, codes []uint8) {
	cpu.SP = cpu.IX
}

func opLDSPIY(cpu *CPU, codes []uint8) {
	cpu.SP = cpu.IY
}

func opPUSHqq(cpu *CPU, codes []uint8) {
	qq := cpu.reg16qq(codes[0] >> 4)
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, qq.U16())
}

func opPUSHIX(cpu *CPU, codes []uint8) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IX)
}

func opPUSHIY(cpu *CPU, codes []uint8) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IY)
}

func opPOPqq(cpu *CPU, codes []uint8) {
	qq := cpu.reg16qq(codes[0] >> 4)
	qq.SetU16(cpu.readU16(cpu.SP))
	cpu.SP += 2
}

func opPOPIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

func opPOPIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.readU16(cpu.SP)
	cpu.SP += 2
}
