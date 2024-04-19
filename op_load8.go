package z80

func oopLDHLPn(cpu *CPU) {
	n := cpu.fetch()
	p := cpu.HL.U16()
	cpu.Memory.Set(p, n)
}

func oopLDIXdPn(cpu *CPU) {
	d := cpu.fetch()
	n := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, n)
}

func oopLDIYdPn(cpu *CPU) {
	d := cpu.fetch()
	n := cpu.fetch()
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

func oopLDAnnP(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.AF.Hi = cpu.Memory.Get(nn)
}

func oopLDBCPA(cpu *CPU) {
	p := cpu.BC.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func oopLDDEPA(cpu *CPU) {
	p := cpu.DE.U16()
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func oopLDnnPA(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.Memory.Set(nn, cpu.AF.Hi)
}

func (cpu *CPU) updateFlagIR(d uint8) {
	var nand uint8 = maskS53 | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= d & maskS53
	if d == 0 {
		or |= maskZ
	}
	if cpu.IFF2 {
		or |= maskPV
	}
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
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
	cpu.updateFlagIR(d)
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
	cpu.updateFlagIR(d)
}

func oopLDIA(cpu *CPU) {
	cpu.IR.Hi = cpu.AF.Hi
}

func oopLDRA(cpu *CPU) {
	cpu.IR.Lo = cpu.AF.Hi
}

func oopLDIXHn(cpu *CPU) {
	n := cpu.fetch()
	cpu.IX = uint16(n)<<8 | cpu.IX&0xff
}

func oopLDIXLn(cpu *CPU) {
	n := cpu.fetch()
	cpu.IX = uint16(n) | cpu.IX&0xff00
}

func oopLDIYHn(cpu *CPU) {
	n := cpu.fetch()
	cpu.IY = uint16(n)<<8 | cpu.IY&0xff
}

func oopLDIYLn(cpu *CPU) {
	n := cpu.fetch()
	cpu.IY = uint16(n) | cpu.IY&0xff00
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopLDbHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDcHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDeHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDlHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDaHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopLDHLPb(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.BC.Hi
	cpu.Memory.Set(p, r)
}

func xopLDHLPc(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.BC.Lo
	cpu.Memory.Set(p, r)
}

func xopLDHLPd(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.DE.Hi
	cpu.Memory.Set(p, r)
}

func xopLDHLPe(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.DE.Lo
	cpu.Memory.Set(p, r)
}

func xopLDHLPh(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.HL.Hi
	cpu.Memory.Set(p, r)
}

func xopLDHLPl(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.HL.Lo
	cpu.Memory.Set(p, r)
}

func xopLDHLPa(cpu *CPU) {
	p := cpu.HL.U16()
	r := cpu.AF.Hi
	cpu.Memory.Set(p, r)
}

func xopLDbn(cpu *CPU) {
	cpu.BC.Hi = cpu.fetch()
}

func xopLDcn(cpu *CPU) {
	cpu.BC.Lo = cpu.fetch()
}

func xopLDdn(cpu *CPU) {
	cpu.DE.Hi = cpu.fetch()
}

func xopLDen(cpu *CPU) {
	cpu.DE.Lo = cpu.fetch()
}

func xopLDhn(cpu *CPU) {
	cpu.HL.Hi = cpu.fetch()
}

func xopLDln(cpu *CPU) {
	cpu.HL.Lo = cpu.fetch()
}

func xopLDan(cpu *CPU) {
	cpu.AF.Hi = cpu.fetch()
}

func xopLDbIXdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDcIXdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdIXdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDeIXdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhIXdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDlIXdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDaIXdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopLDbIYdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDcIYdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdIYdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDeIYdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhIYdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDlIYdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDaIYdP(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopLDIXdPb(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.BC.Hi)
}

func xopLDIXdPc(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.BC.Lo)
}

func xopLDIXdPd(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.DE.Hi)
}

func xopLDIXdPe(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.DE.Lo)
}

func xopLDIXdPh(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.HL.Hi)
}

func xopLDIXdPl(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.HL.Lo)
}

func xopLDIXdPa(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func xopLDIYdPb(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.BC.Hi)
}

func xopLDIYdPc(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.BC.Lo)
}

func xopLDIYdPd(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.DE.Hi)
}

func xopLDIYdPe(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.DE.Lo)
}

func xopLDIYdPh(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.HL.Hi)
}

func xopLDIYdPl(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.HL.Lo)
}

func xopLDIYdPa(cpu *CPU) {
	d := cpu.fetch()
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.AF.Hi)
}
