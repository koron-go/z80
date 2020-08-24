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
		F: opINCIXH,
	},

	{
		N: "DEC IXH",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x25, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
		F: opDECIXH,
	},

	{
		N: "INC IXL",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2c, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
		F: opINCIXL,
	},

	{
		N: "DEC IXL",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2d, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
		F: opDECIXL,
	},

	{
		N: "INC IYH",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x24, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
		F: opINCIYH,
	},

	{
		N: "DEC IYH",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x25, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
		F: opDECIYH,
	},

	{
		N: "INC IYL",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2c, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
		F: opINCIYL,
	},

	{
		N: "DEC IYL",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2d, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
		F: opDECIYL,
	},

	{
		N: "LD IXH, n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x26, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: opLDIXHn,
	},

	{
		N: "LD IXL, n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: opLDIXLn,
	},

	{
		N: "LD IYH, n",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x26, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: opLDIYHn,
	},

	{
		N: "LD IYL, n",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
		F: opLDIYLn,
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
		F: opSL1IXdP,
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
		F: opSL1IYdP,
	},

	{
		N: "SL1 r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x30, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opSL1r,
	},

	{
		N: "SL1 (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x36, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opSL1HLP,
	},

	{
		N: "LD rx1, rx2",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
		F: opLDrx1rx2,
	},

	{
		N: "LD ry1, ry2",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
		F: opLDry1ry2,
	},

	{
		N: "ADD A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x80, 0x07, vReg8},
		},
		T: []int{4},
		F: opADDArx,
	},

	{
		N: "ADD A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x80, 0x07, vReg8},
		},
		T: []int{4},
		F: opADDAry,
	},

	{
		N: "ADC A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x88, 0x07, vReg8},
		},
		T: []int{4},
		F: opADCArx,
	},

	{
		N: "ADC A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x88, 0x07, vReg8},
		},
		T: []int{4},
		F: opADCAry,
	},

	{
		N: "SUB A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x90, 0x07, vReg8},
		},
		T: []int{4},
		F: opSUBArx,
	},

	{
		N: "SUB A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x90, 0x07, vReg8},
		},
		T: []int{4},
		F: opSUBAry,
	},

	{
		N: "SBC A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x98, 0x07, vReg8},
		},
		T: []int{4},
		F: opSBCArx,
	},

	{
		N: "SBC A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x98, 0x07, vReg8},
		},
		T: []int{4},
		F: opSBCAry,
	},

	{
		N: "AND rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xa0, 0x07, vReg8},
		},
		T: []int{4},
		F: opANDrx,
	},

	{
		N: "AND ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xa0, 0x07, vReg8},
		},
		T: []int{4},
		F: opANDry,
	},

	{
		N: "XOR rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xa8, 0x07, vReg8},
		},
		T: []int{4},
		F: opXORrx,
	},

	{
		N: "XOR ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xa8, 0x07, vReg8},
		},
		T: []int{4},
		F: opXORry,
	},

	{
		N: "OR rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xb0, 0x07, vReg8},
		},
		T: []int{4},
		F: opORrx,
	},

	{
		N: "OR ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xb0, 0x07, vReg8},
		},
		T: []int{4},
		F: opORry,
	},

	{
		N: "CP rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xb8, 0x07, vReg8},
		},
		T: []int{4},
		F: opCPrx,
	},

	{
		N: "CP ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xb8, 0x07, vReg8},
		},
		T: []int{4},
		F: opCPry,
	},
}

func opINCIXH(cpu *CPU, codes []uint8) {
	v := cpu.incU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
}

func opDECIXH(cpu *CPU, codes []uint8) {
	v := cpu.decU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
}

func opINCIXL(cpu *CPU, codes []uint8) {
	v := cpu.incU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
}

func opDECIXL(cpu *CPU, codes []uint8) {
	v := cpu.decU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
}

func opINCIYH(cpu *CPU, codes []uint8) {
	v := cpu.incU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
}

func opDECIYH(cpu *CPU, codes []uint8) {
	v := cpu.decU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
}

func opINCIYL(cpu *CPU, codes []uint8) {
	v := cpu.incU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
}

func opDECIYL(cpu *CPU, codes []uint8) {
	v := cpu.decU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
}

func opLDIXHn(cpu *CPU, codes []uint8) {
	v := codes[2]
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
}

func opLDIXLn(cpu *CPU, codes []uint8) {
	v := codes[2]
	cpu.IX = uint16(v) | cpu.IX&0xff00
}

func opLDIYHn(cpu *CPU, codes []uint8) {
	v := codes[2]
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
}

func opLDIYLn(cpu *CPU, codes []uint8) {
	v := codes[2]
	cpu.IY = uint16(v) | cpu.IY&0xff00
}

func opSL1IXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
}

func opSL1IYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
}

func opSL1r(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.sl1U8(*r)
}

func opSL1HLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.sl1U8(cpu.Memory.Get(p)))
}

func opLDrx1rx2(cpu *CPU, codes []uint8) {
	v := cpu.getRX(codes[1])
	cpu.setRX(codes[1]>>3, v)
}

func opLDry1ry2(cpu *CPU, codes []uint8) {
	v := cpu.getRY(codes[1])
	cpu.setRY(codes[1]>>3, v)
}

func opADDArx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.AF.Hi = cpu.addU8(a, x)
}

func opADDAry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.AF.Hi = cpu.addU8(a, y)
}

func opADCArx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func opADCAry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.AF.Hi = cpu.adcU8(a, y)
}

func opSUBArx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.AF.Hi = cpu.subU8(a, x)
}

func opSUBAry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.AF.Hi = cpu.subU8(a, y)
}

func opSBCArx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func opSBCAry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.AF.Hi = cpu.sbcU8(a, y)
}

func opANDrx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.AF.Hi = cpu.andU8(a, x)
}

func opANDry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.AF.Hi = cpu.andU8(a, y)
}

func opXORrx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func opXORry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.AF.Hi = cpu.xorU8(a, y)
}

func opORrx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.AF.Hi = cpu.orU8(a, x)
}

func opORry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.AF.Hi = cpu.orU8(a, y)
}

func opCPrx(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.getRX(codes[1])
	cpu.subU8(a, x)
}

func opCPry(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	y := cpu.getRY(codes[1])
	cpu.subU8(a, y)
}
