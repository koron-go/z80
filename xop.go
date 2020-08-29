package z80

func xopPOPbc(cpu *CPU) {
	cpu.BC.Lo = cpu.Memory.Get(cpu.SP)
	cpu.BC.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPOPde(cpu *CPU) {
	cpu.DE.Lo = cpu.Memory.Get(cpu.SP)
	cpu.DE.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPOPhl(cpu *CPU) {
	cpu.HL.Lo = cpu.Memory.Get(cpu.SP)
	cpu.HL.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPOPaf(cpu *CPU) {
	cpu.AF.Lo = cpu.Memory.Get(cpu.SP)
	cpu.AF.Hi = cpu.Memory.Get(cpu.SP + 1)
	cpu.SP += 2
}

func xopPUSHbc(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.BC.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.BC.Hi)
}

func xopPUSHde(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.DE.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.DE.Hi)
}

func xopPUSHhl(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.HL.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.HL.Hi)
}

func xopPUSHaf(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.AF.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.AF.Hi)
}

func xopINCbc(cpu *CPU) {
	cpu.BC.Lo++
	if cpu.BC.Lo == 0 {
		cpu.BC.Hi++
	}
}

func xopINCde(cpu *CPU) {
	cpu.DE.Lo++
	if cpu.DE.Lo == 0 {
		cpu.DE.Hi++
	}
}

func xopINChl(cpu *CPU) {
	cpu.HL.Lo++
	if cpu.HL.Lo == 0 {
		cpu.HL.Hi++
	}
}

func xopINCsp(cpu *CPU) {
	cpu.SP++
}

func xopDECbc(cpu *CPU) {
	cpu.BC.Lo--
	if cpu.BC.Lo == 0xff {
		cpu.BC.Hi--
	}
}

func xopDECde(cpu *CPU) {
	cpu.DE.Lo--
	if cpu.DE.Lo == 0xff {
		cpu.DE.Hi--
	}
}

func xopDEChl(cpu *CPU) {
	cpu.HL.Lo--
	if cpu.HL.Lo == 0xff {
		cpu.HL.Hi--
	}
}

func xopDECsp(cpu *CPU) {
	cpu.SP--
}

func xopLDbchHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDbclHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdehHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDdelHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhlhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDhllHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDafhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopINCb(cpu *CPU) {
	cpu.BC.Hi = cpu.incU8(cpu.BC.Hi)
}

func xopINCc(cpu *CPU) {
	cpu.BC.Lo = cpu.incU8(cpu.BC.Lo)
}

func xopINCd(cpu *CPU) {
	cpu.DE.Hi = cpu.incU8(cpu.DE.Hi)
}

func xopINCe(cpu *CPU) {
	cpu.DE.Lo = cpu.incU8(cpu.DE.Lo)
}

func xopINCh(cpu *CPU) {
	cpu.HL.Hi = cpu.incU8(cpu.HL.Hi)
}

func xopINCl(cpu *CPU) {
	cpu.HL.Lo = cpu.incU8(cpu.HL.Lo)
}

func xopINCa(cpu *CPU) {
	cpu.AF.Hi = cpu.incU8(cpu.AF.Hi)
}

func xopDECb(cpu *CPU) {
	cpu.decP8(&cpu.BC.Hi)
}

func xopDECc(cpu *CPU) {
	cpu.decP8(&cpu.BC.Lo)
}

func xopDECd(cpu *CPU) {
	cpu.decP8(&cpu.DE.Hi)
}

func xopDECe(cpu *CPU) {
	cpu.decP8(&cpu.DE.Lo)
}

func xopDECh(cpu *CPU) {
	cpu.decP8(&cpu.HL.Hi)
}

func xopDECl(cpu *CPU) {
	cpu.decP8(&cpu.HL.Lo)
}

func xopDECa(cpu *CPU) {
	cpu.decP8(&cpu.AF.Hi)
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

func xopJPnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		cpu.PC = toU16(l, h)
	}
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

func xopCALLnn(cpu *CPU, l, h uint8) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, uint8(cpu.PC))
	cpu.Memory.Set(cpu.SP+1, uint8(cpu.PC>>8))
	cpu.PC = toU16(l, h)
}

func xopCALLnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopADDHLbc(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.BC.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLde(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.DE.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLhl(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.HL.U16()
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopADDHLsp(cpu *CPU) {
	a := cpu.HL.U16()
	x := cpu.SP
	r := cpu.addU16(a, x)
	cpu.HL.SetU16(r)
}

func xopRETnZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ == 0 {
		oopRET(cpu)
	}
}

func xopRETfZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ != 0 {
		oopRET(cpu)
	}
}

func xopRETnC(cpu *CPU) {
	if cpu.AF.Lo&maskC == 0 {
		oopRET(cpu)
	}
}

func xopRETfC(cpu *CPU) {
	if cpu.AF.Lo&maskC != 0 {
		oopRET(cpu)
	}
}

func xopRETnPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV == 0 {
		oopRET(cpu)
	}
}

func xopRETfPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV != 0 {
		oopRET(cpu)
	}
}

func xopRETnS(cpu *CPU) {
	if cpu.AF.Lo&maskS == 0 {
		oopRET(cpu)
	}
}

func xopRETfS(cpu *CPU) {
	if cpu.AF.Lo&maskS != 0 {
		oopRET(cpu)
	}
}

func xopLDbn(cpu *CPU, n uint8) {
	cpu.BC.Hi = n
}

func xopLDcn(cpu *CPU, n uint8) {
	cpu.BC.Lo = n
}

func xopLDdn(cpu *CPU, n uint8) {
	cpu.DE.Hi = n
}

func xopLDen(cpu *CPU, n uint8) {
	cpu.DE.Lo = n
}

func xopLDhn(cpu *CPU, n uint8) {
	cpu.HL.Hi = n
}

func xopLDln(cpu *CPU, n uint8) {
	cpu.HL.Lo = n
}

func xopLDan(cpu *CPU, n uint8) {
	cpu.AF.Hi = n
}

func xopRST00(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0000
}

func xopRST08(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0008
}

func xopRST10(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0010
}

func xopRST18(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0018
}

func xopRST20(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0020
}

func xopRST28(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0028
}

func xopRST30(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0030
}

func xopRST38(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0038
}

func xopRLCb(cpu *CPU) {
	cpu.BC.Hi = cpu.rlcU8(cpu.BC.Hi)
}

func xopRLCc(cpu *CPU) {
	cpu.BC.Lo = cpu.rlcU8(cpu.BC.Lo)
}

func xopRLCd(cpu *CPU) {
	cpu.DE.Hi = cpu.rlcU8(cpu.DE.Hi)
}

func xopRLCe(cpu *CPU) {
	cpu.DE.Lo = cpu.rlcU8(cpu.DE.Lo)
}

func xopRLCh(cpu *CPU) {
	cpu.HL.Hi = cpu.rlcU8(cpu.HL.Hi)
}

func xopRLCl(cpu *CPU) {
	cpu.HL.Lo = cpu.rlcU8(cpu.HL.Lo)
}

func xopRLCa(cpu *CPU) {
	cpu.AF.Hi = cpu.rlcU8(cpu.AF.Hi)
}

func xopRLCHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rlcU8(x)
	cpu.Memory.Set(p, x)
}

func xopRRCb(cpu *CPU) {
	cpu.BC.Hi = cpu.rrcU8(cpu.BC.Hi)
}

func xopRRCc(cpu *CPU) {
	cpu.BC.Lo = cpu.rrcU8(cpu.BC.Lo)
}

func xopRRCd(cpu *CPU) {
	cpu.DE.Hi = cpu.rrcU8(cpu.DE.Hi)
}

func xopRRCe(cpu *CPU) {
	cpu.DE.Lo = cpu.rrcU8(cpu.DE.Lo)
}

func xopRRCh(cpu *CPU) {
	cpu.HL.Hi = cpu.rrcU8(cpu.HL.Hi)
}

func xopRRCl(cpu *CPU) {
	cpu.HL.Lo = cpu.rrcU8(cpu.HL.Lo)
}

func xopRRCa(cpu *CPU) {
	cpu.AF.Hi = cpu.rrcU8(cpu.AF.Hi)
}

func xopRRCHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rrcU8(x)
	cpu.Memory.Set(p, x)
}

func xopRLb(cpu *CPU) {
	cpu.BC.Hi = cpu.rlU8(cpu.BC.Hi)
}

func xopRLc(cpu *CPU) {
	cpu.BC.Lo = cpu.rlU8(cpu.BC.Lo)
}

func xopRLd(cpu *CPU) {
	cpu.DE.Hi = cpu.rlU8(cpu.DE.Hi)
}

func xopRLe(cpu *CPU) {
	cpu.DE.Lo = cpu.rlU8(cpu.DE.Lo)
}

func xopRLh(cpu *CPU) {
	cpu.HL.Hi = cpu.rlU8(cpu.HL.Hi)
}

func xopRLl(cpu *CPU) {
	cpu.HL.Lo = cpu.rlU8(cpu.HL.Lo)
}

func xopRLa(cpu *CPU) {
	cpu.AF.Hi = cpu.rlU8(cpu.AF.Hi)
}

func xopRLHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rlU8(x)
	cpu.Memory.Set(p, x)
}

func xopRRb(cpu *CPU) {
	cpu.BC.Hi = cpu.rrU8(cpu.BC.Hi)
}

func xopRRc(cpu *CPU) {
	cpu.BC.Lo = cpu.rrU8(cpu.BC.Lo)
}

func xopRRd(cpu *CPU) {
	cpu.DE.Hi = cpu.rrU8(cpu.DE.Hi)
}

func xopRRe(cpu *CPU) {
	cpu.DE.Lo = cpu.rrU8(cpu.DE.Lo)
}

func xopRRh(cpu *CPU) {
	cpu.HL.Hi = cpu.rrU8(cpu.HL.Hi)
}

func xopRRl(cpu *CPU) {
	cpu.HL.Lo = cpu.rrU8(cpu.HL.Lo)
}

func xopRRa(cpu *CPU) {
	cpu.AF.Hi = cpu.rrU8(cpu.AF.Hi)
}

func xopRRHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.rrU8(x)
	cpu.Memory.Set(p, x)
}

func xopSLAb(cpu *CPU) {
	cpu.BC.Hi = cpu.slaU8(cpu.BC.Hi)
}

func xopSLAc(cpu *CPU) {
	cpu.BC.Lo = cpu.slaU8(cpu.BC.Lo)
}

func xopSLAd(cpu *CPU) {
	cpu.DE.Hi = cpu.slaU8(cpu.DE.Hi)
}

func xopSLAe(cpu *CPU) {
	cpu.DE.Lo = cpu.slaU8(cpu.DE.Lo)
}

func xopSLAh(cpu *CPU) {
	cpu.HL.Hi = cpu.slaU8(cpu.HL.Hi)
}

func xopSLAl(cpu *CPU) {
	cpu.HL.Lo = cpu.slaU8(cpu.HL.Lo)
}

func xopSLAa(cpu *CPU) {
	cpu.AF.Hi = cpu.slaU8(cpu.AF.Hi)
}

func xopSLAHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.slaU8(x)
	cpu.Memory.Set(p, x)
}

func xopSRAb(cpu *CPU) {
	cpu.BC.Hi = cpu.sraU8(cpu.BC.Hi)
}

func xopSRAc(cpu *CPU) {
	cpu.BC.Lo = cpu.sraU8(cpu.BC.Lo)
}

func xopSRAd(cpu *CPU) {
	cpu.DE.Hi = cpu.sraU8(cpu.DE.Hi)
}

func xopSRAe(cpu *CPU) {
	cpu.DE.Lo = cpu.sraU8(cpu.DE.Lo)
}

func xopSRAh(cpu *CPU) {
	cpu.HL.Hi = cpu.sraU8(cpu.HL.Hi)
}

func xopSRAl(cpu *CPU) {
	cpu.HL.Lo = cpu.sraU8(cpu.HL.Lo)
}

func xopSRAa(cpu *CPU) {
	cpu.AF.Hi = cpu.sraU8(cpu.AF.Hi)
}

func xopSRAHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.sraU8(x)
	cpu.Memory.Set(p, x)
}

func xopSL1b(cpu *CPU) {
	cpu.BC.Hi = cpu.sl1U8(cpu.BC.Hi)
}

func xopSL1c(cpu *CPU) {
	cpu.BC.Lo = cpu.sl1U8(cpu.BC.Lo)
}

func xopSL1d(cpu *CPU) {
	cpu.DE.Hi = cpu.sl1U8(cpu.DE.Hi)
}

func xopSL1e(cpu *CPU) {
	cpu.DE.Lo = cpu.sl1U8(cpu.DE.Lo)
}

func xopSL1h(cpu *CPU) {
	cpu.HL.Hi = cpu.sl1U8(cpu.HL.Hi)
}

func xopSL1l(cpu *CPU) {
	cpu.HL.Lo = cpu.sl1U8(cpu.HL.Lo)
}

func xopSL1a(cpu *CPU) {
	cpu.AF.Hi = cpu.sl1U8(cpu.AF.Hi)
}

func xopSL1HLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.sl1U8(x)
	cpu.Memory.Set(p, x)
}

func xopSRLb(cpu *CPU) {
	cpu.BC.Hi = cpu.srlU8(cpu.BC.Hi)
}

func xopSRLc(cpu *CPU) {
	cpu.BC.Lo = cpu.srlU8(cpu.BC.Lo)
}

func xopSRLd(cpu *CPU) {
	cpu.DE.Hi = cpu.srlU8(cpu.DE.Hi)
}

func xopSRLe(cpu *CPU) {
	cpu.DE.Lo = cpu.srlU8(cpu.DE.Lo)
}

func xopSRLh(cpu *CPU) {
	cpu.HL.Hi = cpu.srlU8(cpu.HL.Hi)
}

func xopSRLl(cpu *CPU) {
	cpu.HL.Lo = cpu.srlU8(cpu.HL.Lo)
}

func xopSRLa(cpu *CPU) {
	cpu.AF.Hi = cpu.srlU8(cpu.AF.Hi)
}

func xopSRLHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.srlU8(x)
	cpu.Memory.Set(p, x)
}

func xopBITbHLP(cpu *CPU, b uint8) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.bitres8(b, x)
	cpu.Memory.Set(p, x)
}

func xopSETbHLP(cpu *CPU, b uint8) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.bitset8(b, x)
	cpu.Memory.Set(p, x)
}

func xopADDIXbc(cpu *CPU) {
	a := cpu.IX
	x := cpu.BC.U16()
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIXde(cpu *CPU) {
	a := cpu.IX
	x := cpu.DE.U16()
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIXix(cpu *CPU) {
	a := cpu.IX
	x := cpu.IX
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIXsp(cpu *CPU) {
	a := cpu.IX
	x := cpu.SP
	cpu.IX = cpu.addU16(a, x)
}

func xopADDIYbc(cpu *CPU) {
	a := cpu.IY
	x := cpu.BC.U16()
	cpu.IY = cpu.addU16(a, x)
}

func xopADDIYde(cpu *CPU) {
	a := cpu.IY
	x := cpu.DE.U16()
	cpu.IY = cpu.addU16(a, x)
}

func xopADDIYiy(cpu *CPU) {
	a := cpu.IY
	x := cpu.IY
	cpu.IY = cpu.addU16(a, x)
}

func xopADDIYsp(cpu *CPU) {
	a := cpu.IY
	x := cpu.SP
	cpu.IY = cpu.addU16(a, x)
}

func xopLDbIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDcIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDeIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDlIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDaIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopLDbIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDcIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDeIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDlIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDaIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopLDIXdPb(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.BC.Hi)
}

func xopLDIXdPc(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.BC.Lo)
}

func xopLDIXdPd(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.DE.Hi)
}

func xopLDIXdPe(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.DE.Lo)
}

func xopLDIXdPh(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.HL.Hi)
}

func xopLDIXdPl(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.HL.Lo)
}

func xopLDIXdPa(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.AF.Hi)
}

func xopLDIYdPb(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.BC.Hi)
}

func xopLDIYdPc(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.BC.Lo)
}

func xopLDIYdPd(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.DE.Hi)
}

func xopLDIYdPe(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.DE.Lo)
}

func xopLDIYdPh(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.HL.Hi)
}

func xopLDIYdPl(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.HL.Lo)
}

func xopLDIYdPa(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.AF.Hi)
}
