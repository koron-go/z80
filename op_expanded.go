package z80

func opPOPbc(cpu *CPU, codes []uint8) {
	u16 := cpu.readU16(cpu.SP)
	cpu.BC.SetU16(u16)
	cpu.SP += 2
}

func opPOPde(cpu *CPU, codes []uint8) {
	u16 := cpu.readU16(cpu.SP)
	cpu.DE.SetU16(u16)
	cpu.SP += 2
}

func opPOPhl(cpu *CPU, codes []uint8) {
	u16 := cpu.readU16(cpu.SP)
	cpu.HL.SetU16(u16)
	cpu.SP += 2
}

func opPOPaf(cpu *CPU, codes []uint8) {
	u16 := cpu.readU16(cpu.SP)
	cpu.AF.SetU16(u16)
	cpu.SP += 2
}
