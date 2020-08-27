package z80

import "math/bits"

func opRLCA(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	a2 := a<<1 | a>>7
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Reset(N).
		Put(C, a&0x80 != 0))
}

func opRLA(cpu *CPU, codes []uint8) {
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

func opRRCA(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	a2 := a>>1 | a<<7
	cpu.AF.Hi = a2
	cpu.flagUpdate(FlagOp{}.
		Reset(H).
		Reset(N).
		Put(C, a&0x01 != 0))
}

func opRRA(cpu *CPU, codes []uint8) {
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

func opRLCr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.rlcU8(*r)
}

func opRLCHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
}

func opRLCIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, cpu.rlcU8(cpu.Memory.Get(p)))
}

func opRLCIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
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

func opRLIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, cpu.rlU8(cpu.Memory.Get(p)))
}

func opRLIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	cpu.Memory.Set(p, cpu.rlU8(cpu.Memory.Get(p)))
}

func opRRCr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.rrcU8(*r)
}

func opRRCHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.rrcU8(cpu.Memory.Get(p)))
}

func opRRCIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, cpu.rrcU8(cpu.Memory.Get(p)))
}

func opRRCIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	cpu.Memory.Set(p, cpu.rrcU8(cpu.Memory.Get(p)))
}

func opRRr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.rrU8(*r)
}

func opRRHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.rrU8(cpu.Memory.Get(p)))
}

func opRRIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, cpu.rrU8(cpu.Memory.Get(p)))
}

func opRRIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	cpu.Memory.Set(p, cpu.rrU8(cpu.Memory.Get(p)))
}

func opSLAr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.slaU8(*r)
}

func opSLAHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.slaU8(cpu.Memory.Get(p)))
}

func opSLAIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, cpu.slaU8(cpu.Memory.Get(p)))
}

func opSLAIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	cpu.Memory.Set(p, cpu.slaU8(cpu.Memory.Get(p)))
}

func opSRAr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.sraU8(*r)
}

func opSRAHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.sraU8(cpu.Memory.Get(p)))
}

func opSRAIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, cpu.sraU8(cpu.Memory.Get(p)))
}

func opSRAIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	cpu.Memory.Set(p, cpu.sraU8(cpu.Memory.Get(p)))
}

func opSRLr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1])
	*r = cpu.srlU8(*r)
}

func opSRLHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	cpu.Memory.Set(p, cpu.srlU8(cpu.Memory.Get(p)))
}

func opSRLIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	cpu.Memory.Set(p, cpu.srlU8(cpu.Memory.Get(p)))
}

func opSRLIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
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