package z80

var undoc = []*OPCode{
	{
		N: "INC IXH",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x24, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.incU8(uint8(cpu.IX >> 8))
			cpu.IX = uint16(v)<<8 | cpu.IX&0xff
		},
	},

	{
		N: "INC IXH",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x25, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.decU8(uint8(cpu.IX >> 8))
			cpu.IX = uint16(v)<<8 | cpu.IX&0xff
		},
	},
}
