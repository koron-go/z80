package z80

func oopJPnn(cpu *CPU, l, h uint8) {
	cpu.PC = toU16(l, h)
}

func opJPccnn(cpu *CPU, codes []uint8) {
	if cpu.flagCC(codes[0] >> 3) {
		cpu.PC = toU16(codes[1], codes[2])
	}
}

func oopJRe(cpu *CPU, off uint8) {
	cpu.PC = addrOff(cpu.PC, off)
}

func oopJRCe(cpu *CPU, off uint8) {
	if cpu.flag(C) {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRNCe(cpu *CPU, off uint8) {
	if !cpu.flag(C) {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRZe(cpu *CPU, off uint8) {
	if cpu.flag(Z) {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRNZe(cpu *CPU, off uint8) {
	if !cpu.flag(Z) {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJPHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.PC = p
}

func opJPIXP(cpu *CPU, codes []uint8) {
	p := cpu.IX
	cpu.PC = p
}

func opJPIYP(cpu *CPU, codes []uint8) {
	p := cpu.IY
	cpu.PC = p
}

func oopDJNZe(cpu *CPU, off uint8) {
	cpu.BC.Hi--
	if cpu.BC.Hi != 0 {
		cpu.PC = addrOff(cpu.PC, off)
	}
}
