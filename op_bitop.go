package z80

var bitop = []*OPCode{

	{
		N: "BIT b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x40, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[1] >> 3) & 0x07
			r := cpu.regP(codes[1])
			cpu.bitchk8(b, *r)
		},
	},

	{
		N: "BIT b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x46, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[1] >> 3) & 0x07
			p := cpu.HL.U16()
			v := cpu.Memory.Get(p)
			cpu.bitchk8(b, v)
		},
	},

	{
		N: "BIT b, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x46, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[3] >> 3) & 0x07
			p := addrOff(cpu.IX, codes[2])
			v := cpu.Memory.Get(p)
			cpu.bitchk8(b, v)
		},
	},

	{
		N: "BIT b, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x46, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[3] >> 3) & 0x07
			p := addrOff(cpu.IY, codes[2])
			v := cpu.Memory.Get(p)
			cpu.bitchk8(b, v)
		},
	},

	{
		N: "SET b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0xc0, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[1] >> 3) & 0x07
			r := cpu.regP(codes[1])
			*r = cpu.bitset8(b, *r)
		},
	},

	{
		N: "SET b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0xc6, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[1] >> 3) & 0x07
			p := cpu.HL.U16()
			v := cpu.Memory.Get(p)
			v = cpu.bitset8(b, v)
			cpu.Memory.Set(p, v)
		},
	},

	{
		N: "SET b, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0xc6, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[3] >> 3) & 0x07
			p := addrOff(cpu.IX, codes[2])
			v := cpu.Memory.Get(p)
			v = cpu.bitset8(b, v)
			cpu.Memory.Set(p, v)
		},
	},

	{
		N: "SET b, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0xc6, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[3] >> 3) & 0x07
			p := addrOff(cpu.IY, codes[2])
			v := cpu.Memory.Get(p)
			v = cpu.bitset8(b, v)
			cpu.Memory.Set(p, v)
		},
	},

	{
		N: "RES b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x80, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[1] >> 3) & 0x07
			r := cpu.regP(codes[1])
			*r = cpu.bitres8(b, *r)
		},
	},

	{
		N: "RES b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x86, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[1] >> 3) & 0x07
			p := cpu.HL.U16()
			v := cpu.Memory.Get(p)
			v = cpu.bitres8(b, v)
			cpu.Memory.Set(p, v)
		},
	},

	{
		N: "RES b, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x86, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[3] >> 3) & 0x07
			p := addrOff(cpu.IX, codes[2])
			v := cpu.Memory.Get(p)
			v = cpu.bitres8(b, v)
			cpu.Memory.Set(p, v)
		},
	},

	{
		N: "RES b, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x86, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			b := (codes[3] >> 3) & 0x07
			p := addrOff(cpu.IY, codes[2])
			v := cpu.Memory.Get(p)
			v = cpu.bitres8(b, v)
			cpu.Memory.Set(p, v)
		},
	},
}
