package z80

func opLDddnn(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[0] >> 4)
	nn := toU16(codes[1], codes[2])
	dd.SetU16(nn)
}

func opLDIXnn(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IX = nn
}

func opLDIYnn(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IY = nn
}

func opLDHLnnP(cpu *CPU, codes []uint8) {
	nn := toU16(codes[1], codes[2])
	cpu.HL.SetU16(cpu.readU16(nn))
}

func opLDddnnP(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[1] >> 4)
	nn := toU16(codes[2], codes[3])
	dd.SetU16(cpu.readU16(nn))
}

func opLDIXnnP(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IX = cpu.readU16(nn)
}

func opLDIYnnP(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.IY = cpu.readU16(nn)
}

func opLDnnPHL(cpu *CPU, codes []uint8) {
	nn := toU16(codes[1], codes[2])
	cpu.writeU16(nn, cpu.HL.U16())
}

func opLDnnPdd(cpu *CPU, codes []uint8) {
	dd := cpu.reg16dd(codes[1] >> 4)
	nn := toU16(codes[2], codes[3])
	cpu.writeU16(nn, dd.U16())
}

func opLDnnPIX(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.writeU16(nn, cpu.IX)
}

func opLDnnPIY(cpu *CPU, codes []uint8) {
	nn := toU16(codes[2], codes[3])
	cpu.writeU16(nn, cpu.IY)
}

func opLDSPHL(cpu *CPU, codes []uint8) {
	cpu.SP = cpu.HL.U16()
}

func opLDSPIX(cpu *CPU, codes []uint8) {
	cpu.SP = cpu.IX
}

func opLDSPIY(cpu *CPU, codes []uint8) {
	cpu.SP = cpu.IY
}

func opPUSHqq(cpu *CPU, codes []uint8) {
	qq := cpu.reg16qq(codes[0] >> 4)
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, qq.U16())
}

func opPUSHIX(cpu *CPU, codes []uint8) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IX)
}

func opPUSHIY(cpu *CPU, codes []uint8) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.IY)
}

func opPOPIX(cpu *CPU, codes []uint8) {
	cpu.IX = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

func opPOPIY(cpu *CPU, codes []uint8) {
	cpu.IY = cpu.readU16(cpu.SP)
	cpu.SP += 2
}
