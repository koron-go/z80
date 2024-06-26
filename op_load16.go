package z80

func oopLDIXnn(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.IX = nn
}

func oopLDIYnn(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.IY = nn
}

func oopLDHLnnP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.HL.Lo = cpu.Memory.Get(nn)
	cpu.HL.Hi = cpu.Memory.Get(nn + 1)
}

func oopLDIXnnP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.IX = cpu.readU16(nn)
}

func oopLDIYnnP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.IY = cpu.readU16(nn)
}

func oopLDnnPHL(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.writeU16(nn, cpu.HL.U16())
}

func oopLDnnPIX(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.writeU16(nn, cpu.IX)
}

func oopLDnnPIY(cpu *CPU) {
	nn := cpu.fetch16()
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

func copPUSHreg(cpu *CPU, reg Register) {
	// Code before optimization
	//cpu.SP -= 2
	//cpu.Memory.Set(cpu.SP, reg.Lo)
	//cpu.Memory.Set(cpu.SP+1, reg.Hi)
	cpu.SP--
	cpu.Memory.Set(cpu.SP, reg.Hi)
	cpu.SP--
	cpu.Memory.Set(cpu.SP, reg.Lo)
}

func xopPUSHbc(cpu *CPU) {
	copPUSHreg(cpu, cpu.BC)
}

func xopPUSHde(cpu *CPU) {
	copPUSHreg(cpu, cpu.DE)
}

func xopPUSHhl(cpu *CPU) {
	copPUSHreg(cpu, cpu.HL)
}

func xopPUSHaf(cpu *CPU) {
	copPUSHreg(cpu, cpu.AF)
}

func copPOPreg(cpu *CPU, reg *Register) {
	// Code before optimization
	//reg.Lo = cpu.Memory.Get(cpu.SP)
	//reg.Hi = cpu.Memory.Get(cpu.SP + 1)
	//cpu.SP += 2
	reg.Lo = cpu.Memory.Get(cpu.SP)
	cpu.SP++
	reg.Hi = cpu.Memory.Get(cpu.SP)
	cpu.SP++
}

func xopPOPbc(cpu *CPU) {
	copPOPreg(cpu, &cpu.BC)
}

func xopPOPde(cpu *CPU) {
	copPOPreg(cpu, &cpu.DE)
}

func xopPOPhl(cpu *CPU) {
	copPOPreg(cpu, &cpu.HL)
}

func xopPOPaf(cpu *CPU) {
	copPOPreg(cpu, &cpu.AF)
}

func xopLDbcnn(cpu *CPU) {
	cpu.BC.Lo, cpu.BC.Hi = cpu.fetch2()
}

func xopLDdenn(cpu *CPU) {
	cpu.DE.Lo, cpu.DE.Hi = cpu.fetch2()
}

func xopLDhlnn(cpu *CPU) {
	cpu.HL.Lo, cpu.HL.Hi = cpu.fetch2()
}

func xopLDspnn(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.SP = nn
}

func xopLDnnPbc(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.writeU16(nn, cpu.BC.U16())
}

func xopLDnnPde(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.writeU16(nn, cpu.DE.U16())
}

func xopLDnnPhl(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.writeU16(nn, cpu.HL.U16())
}

func xopLDnnPsp(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.writeU16(nn, cpu.SP)
}

func xopLDbcnnP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.BC.SetU16(cpu.readU16(nn))
}

func xopLDdennP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.DE.SetU16(cpu.readU16(nn))
}

func xopLDhlnnP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.HL.SetU16(cpu.readU16(nn))
}

func xopLDspnnP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.SP = cpu.readU16(nn)
}
