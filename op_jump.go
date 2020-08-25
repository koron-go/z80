package z80

var jump = []*OPCode{

	{
		N: "JP nn",
		C: []Code{
			{0xc3, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3},
		F: opJPnn,
	},

	{
		N: "JP cc, nn",
		C: []Code{
			{0xc2, 0x38, vCC3_3},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 4, 3, 3},
		T2: []int{4, 3, 3},
		F: opJPccnn,
	},

	{
		N: "JR e",
		C: []Code{
			{0x18, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 5},
		F: opJRe,
	},

	{
		N: "JR C, e",
		C: []Code{
			{0x38, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: opJRCe,
	},

	{
		N: "JR NC, e",
		C: []Code{
			{0x30, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: opJRNCe,
	},

	{
		N: "JR Z, e",
		C: []Code{
			{0x28, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: opJRZe,
	},

	{
		N: "JR NZ, e",
		C: []Code{
			{0x20, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: opJRNZe,
	},

	{
		N: "JP (HL)",
		C: []Code{
			{0xe9, 0x00, nil},
		},
		T: []int{4},
		F: opJPHLP,
	},

	{
		N: "JP (IX)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe9, 0x00, nil},
		},
		T: []int{4, 4},
		F: opJPIXP,
	},

	{
		N: "JP (IY)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe9, 0x00, nil},
		},
		T: []int{4, 4},
		F: opJPIYP,
	},

	{
		N: "DJNZ e",
		C: []Code{
			{0x10, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{5, 3, 5},
		T2: []int{5, 3},
		F: opDJNZe,
	},
}

func opJPnn(cpu *CPU, codes []uint8) {
	cpu.PC = toU16(codes[1], codes[2])
}

func opJPccnn(cpu *CPU, codes []uint8) {
	if cpu.flagCC(codes[0] >> 3) {
		cpu.PC = toU16(codes[1], codes[2])
	}
}

func opJRe(cpu *CPU, codes []uint8) {
	cpu.PC = addrOff(cpu.PC, codes[1])
}

func opJRCe(cpu *CPU, codes []uint8) {
	if cpu.flag(C) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJRNCe(cpu *CPU, codes []uint8) {
	if !cpu.flag(C) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJRZe(cpu *CPU, codes []uint8) {
	if cpu.flag(Z) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJRNZe(cpu *CPU, codes []uint8) {
	if !cpu.flag(Z) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJPHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.PC = p
}

func opJPIXP(cpu *CPU, codes []uint8) {
	p := cpu.IX
	cpu.PC = p
}

func opJPIYP(cpu *CPU, codes []uint8) {
	p := cpu.IY
	cpu.PC = p
}

func opDJNZe(cpu *CPU, codes []uint8) {
	cpu.BC.Hi--
	if cpu.BC.Hi != 0 {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}
