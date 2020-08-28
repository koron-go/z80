package z80

func oopJPnn(cpu *CPU, l, h uint8) {
	cpu.PC = toU16(l, h)
}

func opJPccnn(cpu *CPU, codes []uint8) {
	if cpu.flagCC(codes[0] >> 3) {
		cpu.PC = toU16(codes[1], codes[2])
	}
}

func opJRe(cpu *CPU, codes []uint8) {
	cpu.PC = addrOff(cpu.PC, codes[1])
}

func opJRCe(cpu *CPU, codes []uint8) {
	if cpu.flag(C) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJRNCe(cpu *CPU, codes []uint8) {
	if !cpu.flag(C) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJRZe(cpu *CPU, codes []uint8) {
	if cpu.flag(Z) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJRNZe(cpu *CPU, codes []uint8) {
	if !cpu.flag(Z) {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}

func opJPHLP(cpu *CPU, codes []uint8) {
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

func opDJNZe(cpu *CPU, codes []uint8) {
	cpu.BC.Hi--
	if cpu.BC.Hi != 0 {
		cpu.PC = addrOff(cpu.PC, codes[1])
	}
}
