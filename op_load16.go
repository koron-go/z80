package z80

func oopLDIXnn(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IX = nn
}

func oopLDIYnn(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IY = nn
}

func oopLDHLnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.HL.Lo = cpu.Memory.Get(nn)
	cpu.HL.Hi = cpu.Memory.Get(nn + 1)
}

func oopLDIXnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IX = cpu.readU16(nn)
}

func oopLDIYnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.IY = cpu.readU16(nn)
}

func oopLDnnPHL(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.HL.U16())
}

func oopLDnnPIX(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.IX)
}

func oopLDnnPIY(cpu *CPU, l, h uint8) {
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

func xopLDbcnn(cpu *CPU, l, h uint8) {
	cpu.BC.Lo = l
	cpu.BC.Hi = h
}

func xopLDdenn(cpu *CPU, l, h uint8) {
	cpu.DE.Lo = l
	cpu.DE.Hi = h
}

func xopLDhlnn(cpu *CPU, l, h uint8) {
	cpu.HL.Lo = l
	cpu.HL.Hi = h
}

func xopLDspnn(cpu *CPU, l, h uint8) {
	cpu.SP = toU16(l, h)
}

func xopLDnnPbc(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.BC.U16())
}

func xopLDnnPde(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.DE.U16())
}

func xopLDnnPhl(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.HL.U16())
}

func xopLDnnPsp(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.writeU16(nn, cpu.SP)
}

func xopLDbcnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.BC.SetU16(cpu.readU16(nn))
}

func xopLDdennP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.DE.SetU16(cpu.readU16(nn))
}

func xopLDhlnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.HL.SetU16(cpu.readU16(nn))
}

func xopLDspnnP(cpu *CPU, l, h uint8) {
	nn := toU16(l, h)
	cpu.SP = cpu.readU16(nn)
}
