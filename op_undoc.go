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

func oopINCIXH(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
}

func oopDECIXH(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
}

func oopINCIXL(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
}

func oopDECIXL(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
}

func oopINCIYH(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
}

func oopDECIYH(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
}

func oopINCIYL(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
}

func oopDECIYL(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
}

func oopLDIXHn(cpu *CPU, n uint8) {
	cpu.IX = uint16(n)<<8 | cpu.IX&0xff
}

func oopLDIXLn(cpu *CPU, n uint8) {
	cpu.IX = uint16(n) | cpu.IX&0xff00
}

func oopLDIYHn(cpu *CPU, n uint8) {
	cpu.IY = uint16(n)<<8 | cpu.IY&0xff
}

func oopLDIYLn(cpu *CPU, n uint8) {
	cpu.IY = uint16(n) | cpu.IY&0xff00
}

func oopSL1IXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
}

func oopSL1IYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
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
