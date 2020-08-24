package z80

var bitop = []*OPCode{

	{
		N: "BIT b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x40, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
		F: opBITbr,
	},

	{
		N: "BIT b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x46, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4},
		F: opBITbHLP,
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
		F: opBITbIXdP,
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
		F: opBITbIYdP,
	},

	{
		N: "SET b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0xc0, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
		F: opSETbr,
	},

	{
		N: "SET b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0xc6, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4, 3},
		F: opSETbHLP,
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
		F: opSETbIXdP,
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
		F: opSETbIYdP,
	},

	{
		N: "RES b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x80, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
		F: opRESbr,
	},

	{
		N: "RES b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x86, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4, 3},
		F: opRESbHLP,
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
		F: opRESbIXdP,
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
		F: opRESbIYdP,
	},
}

func opBITbr(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	r := cpu.regP(codes[1])
	cpu.bitchk8(b, *r)
}

func opBITbHLP(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	p := cpu.HL.U16()
	v := cpu.Memory.Get(p)
	cpu.bitchk8(b, v)
}

func opBITbIXdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IX, codes[2])
	v := cpu.Memory.Get(p)
	cpu.bitchk8(b, v)
}

func opBITbIYdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IY, codes[2])
	v := cpu.Memory.Get(p)
	cpu.bitchk8(b, v)
}

func opSETbr(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	r := cpu.regP(codes[1])
	*r = cpu.bitset8(b, *r)
}

func opSETbHLP(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	p := cpu.HL.U16()
	v := cpu.Memory.Get(p)
	v = cpu.bitset8(b, v)
	cpu.Memory.Set(p, v)
}

func opSETbIXdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IX, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitset8(b, v)
	cpu.Memory.Set(p, v)
}

func opSETbIYdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IY, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitset8(b, v)
	cpu.Memory.Set(p, v)
}

func opRESbr(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	r := cpu.regP(codes[1])
	*r = cpu.bitres8(b, *r)
}

func opRESbHLP(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	p := cpu.HL.U16()
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}

func opRESbIXdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IX, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}

func opRESbIYdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IY, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}
