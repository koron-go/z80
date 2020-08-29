package z80

func opADCHLss(cpu *CPU, codes []uint8) {
	a := cpu.HL.U16()
	x := cpu.reg16ss(codes[1] >> 4).U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
}

func opSBCHLss(cpu *CPU, codes []uint8) {
	a := cpu.HL.U16()
	x := cpu.reg16ss(codes[1] >> 4).U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
}

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
