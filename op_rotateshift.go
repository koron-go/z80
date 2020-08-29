package z80

import "math/bits"

func oopRLCA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a<<1 | a>>7
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Reset(N).
		Put(C, a&0x80 != 0))
}

func oopRLA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a << 1
	if cpu.flag(C) {
		a2 |= 0x01
	}
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Reset(N).
		Put(C, a&0x80 != 0))
}

func oopRRCA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a>>1 | a<<7
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Reset(N).
		Put(C, a&0x01 != 0))
}

func oopRRA(cpu *CPU) {
	a := cpu.AF.Hi
	a2 := a >> 1
	if cpu.flag(C) {
		a2 |= 0x80
	}
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Reset(N).
		Put(C, a&0x01 != 0))
}

func oopRLCIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
}

func oopRLCIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
}

func opRLr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.rlU8(*r)
}

func opRLHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.rlU8(cpu.Memory.Get(p)))
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

func opRLD(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	a := cpu.AF.Hi
	b := cpu.Memory.Get(p)
	a2 := a&0xf0 | b>>4
	b2 := b<<4 | a&0x0f
	cpu.Memory.Set(p, b2)
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N))
}

func opRRD(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	a := cpu.AF.Hi
	b := cpu.Memory.Get(p)
	a2 := a&0xf0 | b&0x0f
	b2 := a<<4 | b>>4
	cpu.Memory.Set(p, b2)
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N))
}
