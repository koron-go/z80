package z80

func opRETI(cpu *CPU, codes []uint8) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

func opRETN(cpu *CPU, codes []uint8) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.IFF1 = cpu.IFF2
}

func opCALLnn(cpu *CPU, codes []uint8) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = toU16(codes[1], codes[2])
}

func opCALLccnn(cpu *CPU, codes []uint8) {
	if cpu.flagCC(codes[0] >> 3) {
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = toU16(codes[1], codes[2])
	}
}

func oopRET(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

func opRETcc(cpu *CPU, codes []uint8) {
	if cpu.flagCC(codes[0] >> 3) {
		cpu.PC = cpu.readU16(cpu.SP)
		cpu.SP += 2
	}
}

func opRSTp(cpu *CPU, codes []uint8) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = uint16(codes[0] & 0x38)
}
