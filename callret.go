package z80

var callret = []*OPCode{

	{
		N: "CALL nn",
		C: []Code{
			{0xcd, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.SP -= 2
			cpu.writeU16(cpu.SP, cpu.PC)
			cpu.PC = toU16(codes[1], codes[2])
		},
	},

	{
		N: "CALL cc, nn",
		C: []Code{
			{0xc4, 0x38, vCC3_3},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 4, 3, 3},
		T2: []int{4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			if cpu.flagCC(codes[0] >> 3) {
				cpu.SP -= 2
				cpu.writeU16(cpu.SP, cpu.PC)
				cpu.PC = toU16(codes[1], codes[2])
			}
		},
	},

	{
		N: "RET",
		C: []Code{
			{0xc9, 0x00, nil},
		},
		T: []int{4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.PC = cpu.readU16(cpu.SP)
			cpu.SP += 2
		},
	},

	{
		N: "RET cc",
		C: []Code{
			{0xc0, 0x38, vCC3_3},
		},
		T:  []int{5, 3, 3},
		T2: []int{5},
		F: func(cpu *CPU, codes []uint8) {
			if cpu.flagCC(codes[0] >> 3) {
				cpu.PC = cpu.readU16(cpu.SP)
				cpu.SP += 2
			}
		},
	},

	{
		N: "RETI",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4d, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			// TODO: implement "RETI"
		},
	},

	{
		N: "RETN",
		C: []Code{
			{0xed, 0x00, nil},
			{0x45, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			// TODO: implement "RETN"
		},
	},

	{
		N: "RST p",
		C: []Code{
			{0xc7, 0x38, vRSTp3_3},
		},
		T: []int{5, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.SP -= 2
			cpu.writeU16(cpu.SP, cpu.PC)
			cpu.PC = uint16(codes[0] & 0x38)
		},
	},
}
