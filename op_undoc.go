package z80

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
