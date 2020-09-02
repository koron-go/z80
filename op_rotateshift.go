package z80

import "math/bits"

func (cpu *CPU) updateFlagRL(r uint8) {
	var nand uint8 = maskH | maskN | maskC
	var or = (r >> 7) & maskC
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopRLCA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a<<1 | a>>7
	cpu.AF.Hi = a2
	cpu.updateFlagRL(a)
}

func oopRLA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a << 1
	if cpu.flagC() {
		a2 |= 0x01
	}
	cpu.AF.Hi = a2
	cpu.updateFlagRL(a)
}

func (cpu *CPU) updateFlagRR(r uint8) {
	var nand uint8 = maskH | maskN | maskC
	var or = r & maskC
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopRRCA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a>>1 | a<<7
	cpu.AF.Hi = a2
	cpu.updateFlagRR(a)
}

func oopRRA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a >> 1
	if cpu.flagC() {
		a2 |= 0x80
	}
	cpu.AF.Hi = a2
	cpu.updateFlagRR(a)
}

func oopRLCIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
}

func oopRLCIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
}

func oopRLIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rlU8(cpu.Memory.Get(p)))
}

func oopRLIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rlU8(cpu.Memory.Get(p)))
}

func oopRRCIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rrcU8(cpu.Memory.Get(p)))
}

func oopRRCIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rrcU8(cpu.Memory.Get(p)))
}

func oopRRIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rrU8(cpu.Memory.Get(p)))
}

func oopRRIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rrU8(cpu.Memory.Get(p)))
}

func oopSLAIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.slaU8(cpu.Memory.Get(p)))
}

func oopSLAIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.slaU8(cpu.Memory.Get(p)))
}

func oopSRAIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.sraU8(cpu.Memory.Get(p)))
}

func oopSRAIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.sraU8(cpu.Memory.Get(p)))
}

func oopSRLIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.srlU8(cpu.Memory.Get(p)))
}

func oopSRLIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.srlU8(cpu.Memory.Get(p)))
}

func (cpu *CPU) updateFlagRxD(r uint8) {
	var nand uint8 = maskStd | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= uint8(r) & maskStd
	if uint8(r) == 0 {
		or |= maskZ
	}
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & maskPV
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopRLD(cpu *CPU) {
	p := cpu.HL.U16()
	a := cpu.AF.Hi
	b := cpu.Memory.Get(p)
	a2 := a&0xf0 | b>>4
	b2 := b<<4 | a&0x0f
	cpu.Memory.Set(p, b2)
	cpu.AF.Hi = a2
	cpu.updateFlagRxD(a2)
}

func oopRRD(cpu *CPU) {
	p := cpu.HL.U16()
	a := cpu.AF.Hi
	b := cpu.Memory.Get(p)
	a2 := a&0xf0 | b&0x0f
	b2 := a<<4 | b>>4
	cpu.Memory.Set(p, b2)
	cpu.AF.Hi = a2
	cpu.updateFlagRxD(a2)
}

func oopSL1IXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
}

func oopSL1IYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	v := cpu.sl1U8(cpu.Memory.Get(p))
	cpu.Memory.Set(p, v)
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

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
