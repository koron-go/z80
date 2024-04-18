package z80

func oopLDIXnn(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.IX = nn
}

func oopLDIYnn(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.IY = nn
}

func oopLDHLnnP(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.HL.Lo = cpu.Memory.Get(nn)
	cpu.HL.Hi = cpu.Memory.Get(nn + 1)
}

func oopLDIXnnP(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.IX = cpu.readU16(nn)
}

func oopLDIYnnP(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.IY = cpu.readU16(nn)
}

func oopLDnnPHL(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.HL.U16())
}

func oopLDnnPIX(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.IX)
}

func oopLDnnPIY(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.IY)
}

func oopLDSPHL(cpu *CPU) {
	cpu.SP = cpu.HL.U16()
}

func oopLDSPIX(cpu *CPU) {
	cpu.SP = cpu.IX
}

func oopLDSPIY(cpu *CPU) {
	cpu.SP = cpu.IY
}

func oopPUSHIX(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IX)
}

func oopPUSHIY(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IY)
}

func oopPOPIX(cpu *CPU) {
	cpu.IX = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

func oopPOPIY(cpu *CPU) {
	cpu.IY = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopPUSHreg(cpu *CPU, reg Register) {
	// Code before optimization
	//cpu.SP -= 2
	//cpu.Memory.Set(cpu.SP, reg.Lo)
	//cpu.Memory.Set(cpu.SP+1, reg.Hi)
	cpu.SP--
	cpu.Memory.Set(cpu.SP, reg.Hi)
	cpu.SP--
	cpu.Memory.Set(cpu.SP, reg.Lo)
}

func xopPOPreg(cpu *CPU, reg *Register) {
	// Code before optimization
	//reg.Lo = cpu.Memory.Get(cpu.SP)
	//reg.Hi = cpu.Memory.Get(cpu.SP + 1)
	//cpu.SP += 2
	reg.Lo = cpu.Memory.Get(cpu.SP)
	cpu.SP++
	reg.Hi = cpu.Memory.Get(cpu.SP)
	cpu.SP++
}

func xopLDbcnn(cpu *CPU) {
	cpu.BC.Lo = cpu.fetch()
	cpu.BC.Hi = cpu.fetch()
}

func xopLDdenn(cpu *CPU) {
	cpu.DE.Lo = cpu.fetch()
	cpu.DE.Hi = cpu.fetch()
}

func xopLDhlnn(cpu *CPU) {
	cpu.HL.Lo = cpu.fetch()
	cpu.HL.Hi = cpu.fetch()
}

func xopLDspnn(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	cpu.SP = toU16(l, h)
}

func xopLDnnPbc(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.BC.U16())
}

func xopLDnnPde(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.DE.U16())
}

func xopLDnnPhl(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.HL.U16())
}

func xopLDnnPsp(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.SP)
}

func xopLDbcnnP(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.BC.SetU16(cpu.readU16(nn))
}

func xopLDdennP(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.DE.SetU16(cpu.readU16(nn))
}

func xopLDhlnnP(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.HL.SetU16(cpu.readU16(nn))
}

func xopLDspnnP(cpu *CPU) {
	l := cpu.fetch()
	h := cpu.fetch()
	nn := toU16(l, h)
	cpu.SP = cpu.readU16(nn)
}
