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

func opADDIXpp(cpu *CPU, codes []uint8) {
	a := cpu.IX
	x := cpu.reg16pp(codes[1] >> 4)
	cpu.IX = cpu.addU16(a, x)
}

func opADDIYrr(cpu *CPU, codes []uint8) {
	a := cpu.IY
	x := cpu.reg16rr(codes[1] >> 4)
	cpu.IY = cpu.addU16(a, x)
}

func opINCIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.incU16(cpu.IX)
}

func opINCIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.incU16(cpu.IY)
}

func opDECIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.decU16(cpu.IX)
}

func opDECIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.decU16(cpu.IY)
}
