package z80

func oopINCIX(cpu *CPU) {
	cpu.IX = cpu.incU16(cpu.IX)
}

func oopINCIY(cpu *CPU) {
	cpu.IY = cpu.incU16(cpu.IY)
}

func oopDECIX(cpu *CPU) {
	cpu.IX = cpu.decU16(cpu.IX)
}

func oopDECIY(cpu *CPU) {
	cpu.IY = cpu.decU16(cpu.IY)
}
