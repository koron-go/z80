package z80

func oopEXDEHL(cpu *CPU) {
	cpu.HL, cpu.DE = cpu.DE, cpu.HL
}

func oopEXAFAF(cpu *CPU) {
	cpu.AF, cpu.Alternate.AF = cpu.Alternate.AF, cpu.AF
}

func oopEXX(cpu *CPU) {
	cpu.BC, cpu.Alternate.BC = cpu.Alternate.BC, cpu.BC
	cpu.DE, cpu.Alternate.DE = cpu.Alternate.DE, cpu.DE
	cpu.HL, cpu.Alternate.HL = cpu.Alternate.HL, cpu.HL
}

func oopEXSPPHL(cpu *CPU) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.HL.U16())
	cpu.HL.SetU16(v)
}

func opEXSPPIX(cpu *CPU, codes []uint8) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.IX)
	cpu.IX = v
}

func opEXSPPIY(cpu *CPU, codes []uint8) {
	v := cpu.readU16(cpu.SP)
	cpu.writeU16(cpu.SP, cpu.IY)
	cpu.IY = v
}

func opLDI(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de + 1)
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
}

func opLDIR(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de + 1)
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
	if bc != 0 {
		cpu.PC -= 2
	}
}

func opLDD(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de - 1)
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
}

func opLDDR(cpu *CPU, codes []uint8) {
	de := cpu.DE.U16()
	hl := cpu.HL.U16()
	cpu.Memory.Set(de, cpu.Memory.Get(hl))
	cpu.DE.SetU16(de - 1)
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Put(PV, bc != 0).
		Reset(N))
	if bc != 0 {
		cpu.PC -= 2
	}
}

func opCPI(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
}

func opCPIR(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl + 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
	if bc != 0 && v != 0 {
		cpu.PC -= 2
	}
}

func opCPD(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
}

func opCPDR(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	hl := cpu.HL.U16()
	x := cpu.Memory.Get(hl)
	v := a - x
	cpu.HL.SetU16(hl - 1)
	bc := cpu.BC.U16() - 1
	cpu.BC.SetU16(bc)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < x&0x0f).
		Put(PV, bc != 0).
		Set(N))
	if bc != 0 && v != 0 {
		cpu.PC -= 2
	}
}
