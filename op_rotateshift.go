package z80

import "math/bits"

var rotateshift = []*OPCode{

	{
		N: "RLCA",
		C: []Code{
			{0x07, 0x00, nil},
		},
		T: []int{4},
		F: opRLCA,
	},

	{
		N: "RLA",
		C: []Code{
			{0x17, 0x00, nil},
		},
		T: []int{4},
		F: opRLA,
	},

	{
		N: "RRCA",
		C: []Code{
			{0x0f, 0x00, nil},
		},
		T: []int{4},
		F: opRRCA,
	},

	{
		N: "RRA",
		C: []Code{
			{0x1f, 0x00, nil},
		},
		T: []int{4},
		F: opRRA,
	},

	{
		N: "RLC r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x00, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opRLCr,
	},

	{
		N: "RLC (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x06, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opRLCHLP,
	},

	{
		N: "RLC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x06, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRLCIXdP,
	},

	{
		N: "RLC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x06, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRLCIYdP,
	},

	{
		N: "RL r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x10, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opRLr,
	},

	{
		N: "RL (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x16, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opRLHLP,
	},

	{
		N: "RL (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x16, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRLIXdP,
	},

	{
		N: "RL (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x16, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRLIYdP,
	},

	{
		N: "RRC r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x08, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opRRCr,
	},

	{
		N: "RRC (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x0e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opRRCHLP,
	},

	{
		N: "RRC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x0e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRRCIXdP,
	},

	{
		N: "RRC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x0e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRRCIYdP,
	},

	{
		N: "RR r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x18, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opRRr,
	},

	{
		N: "RR (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x1e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opRRHLP,
	},

	{
		N: "RR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x1e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRRIXdP,
	},

	{
		N: "RR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x1e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opRRIYdP,
	},

	{
		N: "SLA r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x20, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opSLAr,
	},

	{
		N: "SLA (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x26, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opSLAHLP,
	},

	{
		N: "SLA (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x26, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opSLAIXdP,
	},

	{
		N: "SLA (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x26, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opSLAIYdP,
	},

	{
		N: "SRA r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x28, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opSRAr,
	},

	{
		N: "SRA (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x2e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opSRAHLP,
	},

	{
		N: "SRA (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x2e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opSRAIXdP,
	},

	{
		N: "SRA (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x2e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opSRAIYdP,
	},

	{
		N: "SRL r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x38, 0x07, vReg8},
		},
		T: []int{4, 4},
		F: opSRLr,
	},

	{
		N: "SRL (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x3e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
		F: opSRLHLP,
	},

	{
		N: "SRL (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x3e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opSRLIXdP,
	},

	{
		N: "SRL (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x3e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opSRLIYdP,
	},

	{
		N: "RLD",
		C: []Code{
			{0xed, 0x00, nil},
			{0x6f, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3},
		F: opRLD,
	},

	{
		N: "RRD",
		C: []Code{
			{0xed, 0x00, nil},
			{0x67, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3},
		F: opRRD,
	},
}

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
