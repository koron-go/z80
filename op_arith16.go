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

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopINCbc(cpu *CPU) {
	cpu.BC.Lo++
	if cpu.BC.Lo == 0 {
		cpu.BC.Hi++
	}
}

func xopINCde(cpu *CPU) {
	cpu.DE.Lo++
	if cpu.DE.Lo == 0 {
		cpu.DE.Hi++
	}
}

func xopINChl(cpu *CPU) {
	cpu.HL.Lo++
	if cpu.HL.Lo == 0 {
		cpu.HL.Hi++
	}
}

func xopINCsp(cpu *CPU) {
	cpu.SP++
}

func xopDECbc(cpu *CPU) {
	cpu.BC.Lo--
	if cpu.BC.Lo == 0xff {
		cpu.BC.Hi--
	}
}

func xopDECde(cpu *CPU) {
	cpu.DE.Lo--
	if cpu.DE.Lo == 0xff {
		cpu.DE.Hi--
	}
}

func xopDEChl(cpu *CPU) {
	cpu.HL.Lo--
	if cpu.HL.Lo == 0xff {
		cpu.HL.Hi--
	}
}

func xopDECsp(cpu *CPU) {
	cpu.SP--
}

func xopADDHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDIXbc(cpu *CPU) {
	a := cpu.IX
	x := cpu.BC.U16()
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIXde(cpu *CPU) {
	a := cpu.IX
	x := cpu.DE.U16()
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIXix(cpu *CPU) {
	a := cpu.IX
	x := cpu.IX
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIXsp(cpu *CPU) {
	a := cpu.IX
	x := cpu.SP
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIYbc(cpu *CPU) {
	a := cpu.IY
	x := cpu.BC.U16()
	cpu.IY = cpu.addU16(a, x)
}

func xopADDIYde(cpu *CPU) {
	a := cpu.IY
	x := cpu.DE.U16()
	cpu.IY = cpu.addU16(a, x)
}

func xopADDIYiy(cpu *CPU) {
	a := cpu.IY
	x := cpu.IY
	cpu.IY = cpu.addU16(a, x)
}

func xopADDIYsp(cpu *CPU) {
	a := cpu.IY
	x := cpu.SP
	cpu.IY = cpu.addU16(a, x)
}

func xopSBCHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
}

func xopSBCHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
}

func xopSBCHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	cpu.HL.SetU16(cpu.sbcU16(a, x))
}

func xopSBCHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	cpu.HL.SetU16(cpu.sbcU16(a, x))
}

func xopADCHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
}

func xopADCHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
}

func xopADCHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	cpu.HL.SetU16(cpu.adcU16(a, x))
}

func xopADCHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	cpu.HL.SetU16(cpu.adcU16(a, x))
}
