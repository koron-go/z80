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

func opLDddnnP(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[1] >> 4)
	nn := toU16(codes[2], codes[3])
	dd.SetU16(cpu.readU16(nn))
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

func opLDnnPdd(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[1] >> 4)
	nn := toU16(codes[2], codes[3])
	cpu.writeU16(nn, dd.U16())
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
