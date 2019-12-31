package z80

var load16 = []*OPCode{

	{
		N: "",
		C: []Code{
			{0x01, 0x30, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			dd := cpu.reg16dd(codes[0] >> 4)
			nn := toU16(codes[1], codes[2])
			dd.SetU16(nn)
		},
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
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[2], codes[3])
			cpu.IX = nn
		},
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
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[2], codes[3])
			cpu.IY = nn
		},
	},

	{
		N: "LD HL, (nn)",
		C: []Code{
			{0x2a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[1], codes[2])
			cpu.HL.SetU16(cpu.readU16(nn))
		},
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
		F: func(cpu *CPU, codes []uint8) {
			dd := cpu.reg16dd(codes[0] >> 4)
			nn := toU16(codes[2], codes[3])
			dd.SetU16(cpu.readU16(nn))
		},
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
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[2], codes[3])
			cpu.IX = cpu.readU16(nn)
		},
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
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[2], codes[3])
			cpu.IY = cpu.readU16(nn)
		},
	},

	{
		N: "LD (nn), HL",
		C: []Code{
			{0x22, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[1], codes[2])
			cpu.writeU16(nn, cpu.HL.U16())
		},
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
		F: func(cpu *CPU, codes []uint8) {
			dd := cpu.reg16dd(codes[1] >> 4)
			nn := toU16(codes[2], codes[3])
			cpu.writeU16(nn, dd.U16())
		},
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
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[2], codes[3])
			cpu.writeU16(nn, cpu.IX)
		},
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
		F: func(cpu *CPU, codes []uint8) {
			nn := toU16(codes[2], codes[3])
			cpu.writeU16(nn, cpu.IY)
		},
	},

	{
		N: "LD SP, HL",
		C: []Code{
			{0xf9, 0x00, nil},
		},
		T: []int{6},
		F: func(cpu *CPU, codes []uint8) {
			cpu.SP = cpu.HL.U16()
		},
	},

	{
		N: "LD SP, IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xf9, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			cpu.SP = cpu.IX
		},
	},

	{
		N: "LD SP, IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xf9, 0x00, nil},
		},
		T: []int{4, 6},
		F: func(cpu *CPU, codes []uint8) {
			cpu.SP = cpu.IY
		},
	},

	{
		N: "PUSH qq",
		C: []Code{
			{0xc5, 0x30, nil},
		},
		T: []int{5, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			qq := cpu.reg16qq(codes[0] >> 4)
			cpu.SP -= 2
			cpu.writeU16(cpu.SP, qq.U16())
		},
	},

	{
		N: "PUSH IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe5, 0x00, nil},
		},
		T: []int{4, 5, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.SP -= 2
			cpu.writeU16(cpu.SP, cpu.IX)
		},
	},

	{
		N: "PUSH IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe5, 0x00, nil},
		},
		T: []int{4, 5, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.SP -= 2
			cpu.writeU16(cpu.SP, cpu.IY)
		},
	},

	{
		N: "POP qq",
		C: []Code{
			{0xc1, 0x30, nil},
		},
		T: []int{4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			qq := cpu.reg16qq(codes[0] >> 4)
			qq.SetU16(cpu.readU16(cpu.SP))
			cpu.SP += 2
		},
	},

	{
		N: "POP IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe1, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IX = cpu.readU16(cpu.SP)
			cpu.SP += 2
		},
	},

	{
		N: "POP IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe1, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IY = cpu.readU16(cpu.SP)
			cpu.SP += 2
		},
	},
}
