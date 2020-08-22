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
		N: "DEC IXH",
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

	{
		N: "INC IXL",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2c, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.incU8(uint8(cpu.IX))
			cpu.IX = uint16(v) | cpu.IX&0xff00
		},
	},

	{
		N: "DEC IXL",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2d, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.decU8(uint8(cpu.IX))
			cpu.IX = uint16(v) | cpu.IX&0xff00
		},
	},

	{
		N: "INC IYH",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x24, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.incU8(uint8(cpu.IY >> 8))
			cpu.IY = uint16(v)<<8 | cpu.IY&0xff
		},
	},

	{
		N: "DEC IYH",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x25, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.decU8(uint8(cpu.IY >> 8))
			cpu.IY = uint16(v)<<8 | cpu.IY&0xff
		},
	},

	{
		N: "INC IYL",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2c, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.incU8(uint8(cpu.IY))
			cpu.IY = uint16(v) | cpu.IY&0xff00
		},
	},

	{
		N: "DEC IYL",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2d, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.decU8(uint8(cpu.IY))
			cpu.IY = uint16(v) | cpu.IY&0xff00
		},
	},
}
