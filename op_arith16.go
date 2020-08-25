package z80

func opADDHLss(cpu *CPU, codes []uint8) {
	a := cpu.HL.U16()
	x := cpu.reg16ss(codes[0] >> 4).U16()
	cpu.HL.SetU16(cpu.addU16(a, x))
}

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

func opINCss(cpu *CPU, codes []uint8) {
	ss := cpu.reg16ss(codes[0] >> 4)
	ss.SetU16(cpu.incU16(ss.U16()))
}

func opINCIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.incU16(cpu.IX)
}

func opINCIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.incU16(cpu.IY)
}

func opDECss(cpu *CPU, codes []uint8) {
	ss := cpu.reg16ss(codes[0] >> 4)
	ss.SetU16(cpu.decU16(ss.U16()))
}

func opDECIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.decU16(cpu.IX)
}

func opDECIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.decU16(cpu.IY)
}
