package z80

func opLDr1r2(cpu *CPU, codes []uint8) {
	*cpu.regP(codes[0] >> 3) = *cpu.regP(codes[0])
}

func opLDrn(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[0] >> 3)
	*r = codes[1]
}

func opLDrHLP(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[0] >> 3)
	*r = cpu.Memory.Get(cpu.HL.U16())
}

func opLDrIXdP(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1] >> 3)
	p := addrOff(cpu.IX, codes[2])
	*r = cpu.Memory.Get(p)
}

func opLDrIYdP(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1] >> 3)
	p := addrOff(cpu.IY, codes[2])
	*r = cpu.Memory.Get(p)
}

func opLDHLPr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[0])
	cpu.Memory.Set(cpu.HL.U16(), *r)
}

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

func opLDHLPn(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, codes[1])
}

func opLDIXdPn(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, codes[3])
}

func opLDIYdPn(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	cpu.Memory.Set(p, codes[3])
}

func opLDABCP(cpu *CPU, codes []uint8) {
	p := cpu.BC.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func opLDADEP(cpu *CPU, codes []uint8) {
	p := cpu.DE.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func opLDAnnP(cpu *CPU, codes []uint8) {
	p := toU16(codes[1], codes[2])
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func opLDBCPA(cpu *CPU, codes []uint8) {
	p := cpu.BC.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func opLDDEPA(cpu *CPU, codes []uint8) {
	p := cpu.DE.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func opLDnnPA(cpu *CPU, codes []uint8) {
	p := toU16(codes[1], codes[2])
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func opLDAI(cpu *CPU, codes []uint8) {
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

func opLDAR(cpu *CPU, codes []uint8) {
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

func opLDIA(cpu *CPU, codes []uint8) {
	cpu.IR.Hi = cpu.AF.Hi
}

func opLDRA(cpu *CPU, codes []uint8) {
	cpu.IR.Lo = cpu.AF.Hi
}
