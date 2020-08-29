package z80

func oopRETI(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2

	if cpu.INT != nil {
		cpu.INT.ReturnINT()
	}
}

func oopRETN(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.IFF1 = cpu.IFF2

	cpu.InNMI = false
}

func oopRET(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
}
