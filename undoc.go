package z80

func (cpu *CPU) getRX(n uint8) uint8 {
	switch n & 0x07 {
	case 0x00:
		return cpu.BC.Hi
	case 0x01:
		return cpu.BC.Lo
	case 0x02:
		return cpu.DE.Hi
	case 0x03:
		return cpu.DE.Lo
	case 0x04:
		return uint8(cpu.IX >> 8)
	case 0x05:
		return uint8(cpu.IX & 0xff)
	case 0x07:
		return cpu.AF.Hi
	default:
		cpu.failf("getRX invalid register: %02x", n)
		return 0
	}
}

func (cpu *CPU) setRX(n uint8, v uint8) {
	switch n & 0x07 {
	case 0x00:
		cpu.BC.Hi = v
	case 0x01:
		cpu.BC.Lo = v
	case 0x02:
		cpu.DE.Hi = v
	case 0x03:
		cpu.DE.Lo = v
	case 0x04:
		cpu.IX = uint16(v)<<8 | cpu.IX&0x00ff
	case 0x05:
		cpu.IX = uint16(v) | cpu.IX&0xff00
	case 0x07:
		cpu.AF.Hi = v
	default:
		cpu.failf("setRX invalid register: %02x", n)
	}
}

func (cpu *CPU) getRY(n uint8) uint8 {
	switch n & 0x07 {
	case 0x00:
		return cpu.BC.Hi
	case 0x01:
		return cpu.BC.Lo
	case 0x02:
		return cpu.DE.Hi
	case 0x03:
		return cpu.DE.Lo
	case 0x04:
		return uint8(cpu.IY >> 8)
	case 0x05:
		return uint8(cpu.IY & 0xff)
	case 0x07:
		return cpu.AF.Hi
	default:
		cpu.failf("getRY invalid register: %02x", n)
		return 0
	}
}

func (cpu *CPU) setRY(n uint8, v uint8) {
	switch n & 0x07 {
	case 0x00:
		cpu.BC.Hi = v
	case 0x01:
		cpu.BC.Lo = v
	case 0x02:
		cpu.DE.Hi = v
	case 0x03:
		cpu.DE.Lo = v
	case 0x04:
		cpu.IY = uint16(v)<<8 | cpu.IY&0x00ff
	case 0x05:
		cpu.IY = uint16(v) | cpu.IY&0xff00
	case 0x07:
		cpu.AF.Hi = v
	default:
		cpu.failf("setRY invalid register: %02x", n)
	}
}

var undoc = []*OPCode{
	{
		N: "INC IXH",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x24, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
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
		T: []int{4, 6}, // not verified
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
		T: []int{4, 6}, // not verified
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
		T: []int{4, 6}, // not verified
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
		T: []int{4, 6}, // not verified
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
		T: []int{4, 6}, // not verified
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
		T: []int{4, 6}, // not verified
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
		T: []int{4, 6}, // not verified
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.decU8(uint8(cpu.IY))
			cpu.IY = uint16(v) | cpu.IY&0xff00
		},
	},

	{
		N: "LD IXH, n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x26, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: func(cpu *CPU, codes []uint8) {
			v := codes[2]
			cpu.IX = uint16(v)<<8 | cpu.IX&0xff
		},
	},

	{
		N: "LD IXL, n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: func(cpu *CPU, codes []uint8) {
			v := codes[2]
			cpu.IX = uint16(v) | cpu.IX&0xff00
		},
	},

	{
		N: "LD IYH, n",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x26, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: func(cpu *CPU, codes []uint8) {
			v := codes[2]
			cpu.IY = uint16(v)<<8 | cpu.IY&0xff
		},
	},

	{
		N: "LD IYL, n",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: func(cpu *CPU, codes []uint8) {
			v := codes[2]
			cpu.IY = uint16(v) | cpu.IY&0xff00
		},
	},

	{
		N: "SL1 (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x36, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3}, // not verified
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IX, codes[2])
			v := cpu.sl1U8(cpu.Memory.Get(p))
			cpu.Memory.Set(p, v)
		},
	},

	{
		N: "SL1 (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x36, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3}, // not verified
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IY, codes[2])
			v := cpu.sl1U8(cpu.Memory.Get(p))
			cpu.Memory.Set(p, v)
		},
	},

	{
		N: "SL1 r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x30, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[1])
			*r = cpu.sl1U8(*r)
		},
	},

	{
		N: "SL1 (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x36, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.HL.U16()
			cpu.Memory.Set(p, cpu.sl1U8(cpu.Memory.Get(p)))
		},
	},

	{
		N: "LD rx1, rx2",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.getRX(codes[1])
			cpu.setRX(codes[1]>>3, v)
		},
	},

	{
		N: "LD ry1, ry2",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.getRY(codes[1])
			cpu.setRY(codes[1]>>3, v)
		},
	},
}
