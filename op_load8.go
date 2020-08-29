package z80

func opLDIXdPr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, *r)
}

func opLDIYdPr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	p := addrOff(cpu.IY, codes[2])
	cpu.Memory.Set(p, *r)
}

func oopLDHLPn(cpu *CPU, n uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, n)
}

func oopLDIXdPn(cpu *CPU, d, n uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, n)
}

func oopLDIYdPn(cpu *CPU, d, n uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, n)
}

func oopLDABCP(cpu *CPU) {
	p := cpu.BC.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func oopLDADEP(cpu *CPU) {
	p := cpu.DE.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func oopLDAnnP(cpu *CPU, l, h uint8) {
	p := toU16(l, h)
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func oopLDBCPA(cpu *CPU) {
	p := cpu.BC.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func oopLDDEPA(cpu *CPU) {
	p := cpu.DE.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func oopLDnnPA(cpu *CPU, l, h uint8) {
	p := toU16(l, h)
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func oopLDAI(cpu *CPU) {
	d := cpu.IR.Hi
	cpu.AF.Hi = d
	// update F by d
	// - S is set if the I Register is negative; otherwise, it is
	//   reset.
	// - Z is set if the I Register is 0; otherwise, it is reset.
	// - H is reset.
	// - P/V contains contents of IFF2.
	// - N is reset.
	// - C is not affected.
	// - If an interrupt occurs during execution of this instruction,
	//   the Parity flag contains a 0.
	cpu.flagUpdate(FlagOp{}.
		Put(S, d&0x80 != 0).
		Put(Z, d == 0).
		Reset(H).
		Put(PV, cpu.IFF2).
		Reset(N).
		Keep(C))
}

func oopLDAR(cpu *CPU) {
	d := cpu.IR.Lo
	cpu.AF.Hi = d
	// update F by d
	// - S is set if, R-Register is negative; otherwise, it is reset.
	// - Z is set if the R Register is 0; otherwise, it is reset.
	// - H is reset.
	// - P/V contains contents of IFF2.
	// - N is reset.
	// - C is not affected.
	// - If an interrupt occurs during execution of this instruction,
	//	 the parity flag contains a 0.
	cpu.flagUpdate(FlagOp{}.
		Put(S, d&0x80 != 0).
		Put(Z, d == 0).
		Reset(H).
		Put(PV, cpu.IFF2).
		Reset(N).
		Keep(C))
}

func oopLDIA(cpu *CPU) {
	cpu.IR.Hi = cpu.AF.Hi
}

func oopLDRA(cpu *CPU) {
	cpu.IR.Lo = cpu.AF.Hi
}
