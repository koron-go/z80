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
		F: func(cpu *CPU, codes []uint8) {
			cpu.PC = toU16(codes[1], codes[2])
		},
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
		F: func(cpu *CPU, codes []uint8) {
			if cpu.flagCC(codes[0] >> 3) {
				cpu.PC = toU16(codes[1], codes[2])
			}
		},
	},

	{
		N: "JR e",
		C: []Code{
			{0x18, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			cpu.PC = addrOff(cpu.PC, codes[1])
		},
	},

	{
		N: "JR C, e",
		C: []Code{
			{0x38, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			if cpu.flag(C) {
				cpu.PC = addrOff(cpu.PC, codes[1])
			}
		},
	},

	{
		N: "JR NC, e",
		C: []Code{
			{0x30, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			if !cpu.flag(C) {
				cpu.PC = addrOff(cpu.PC, codes[1])
			}
		},
	},

	{
		N: "JR Z, e",
		C: []Code{
			{0x28, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			if cpu.flag(Z) {
				cpu.PC = addrOff(cpu.PC, codes[1])
			}
		},
	},

	{
		N: "JR NZ, e",
		C: []Code{
			{0x20, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			if !cpu.flag(Z) {
				cpu.PC = addrOff(cpu.PC, codes[1])
			}
		},
	},

	{
		N: "JP (HL)",
		C: []Code{
			{0xe9, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.HL.U16()
			cpu.PC = p
		},
	},

	{
		N: "JP (IX)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe9, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.IX
			cpu.PC = p
		},
	},

	{
		N: "JP (IY)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe9, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.IY
			cpu.PC = p
		},
	},

	{
		N: "DJNZ e",
		C: []Code{
			{0x10, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{5, 3, 5},
		T2: []int{5, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.BC.Hi--
			if cpu.BC.Hi != 0 {
				cpu.PC = addrOff(cpu.PC, codes[1])
			}
		},
	},
}
