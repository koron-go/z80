package z80

var arith8 = []*OPCode{

	{
		N: "ADD A, r",
		C: []Code{
			{0x80, 0x07, vReg8},
		},
		T: []int{4},
		F: opADDAr,
	},

	{
		N: "ADD A, n",
		C: []Code{
			{0xc6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opADDAn,
	},

	{
		N: "ADD A, (HL)",
		C: []Code{
			{0x86, 0x00, nil},
		},
		T: []int{4, 3},
		F: opADDAHLP,
	},

	{
		N: "ADD A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x86, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opADDAIXdP,
	},

	{
		N: "ADD A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x86, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opADDAIYdP,
	},

	{
		N: "ADC A, r",
		C: []Code{
			{0x88, 0x07, vReg8},
		},
		T: []int{4},
		F: opADCAr,
	},

	{
		N: "ADC A, n",
		C: []Code{
			{0xce, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opADCAn,
	},

	{
		N: "ADC A, (HL)",
		C: []Code{
			{0x8e, 0x00, nil},
		},
		T: []int{4, 3},
		F: opADCAHLP,
	},

	{
		N: "ADC A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x8e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opADCAIXdP,
	},

	{
		N: "ADC A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x8e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opADCAIYdP,
	},

	{
		N: "SUB A, r",
		C: []Code{
			{0x90, 0x07, vReg8},
		},
		T: []int{4},
		F: opSUBAr,
	},

	{
		N: "SUB A, n",
		C: []Code{
			{0xd6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opSUBAn,
	},

	{
		N: "SUB A, (HL)",
		C: []Code{
			{0x96, 0x00, nil},
		},
		T: []int{4, 3},
		F: opSUBAHLP,
	},

	{
		N: "SUB A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x96, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opSUBAIXdP,
	},

	{
		N: "SUB A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x96, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opSUBAIYdP,
	},

	{
		N: "SBC A, r",
		C: []Code{
			{0x98, 0x07, vReg8},
		},
		T: []int{4},
		F: opSBCAr,
	},

	{
		N: "SBC A, n",
		C: []Code{
			{0xde, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opSBCAn,
	},

	{
		N: "SBC A, (HL)",
		C: []Code{
			{0x9e, 0x00, nil},
		},
		T: []int{4, 3},
		F: opSBCAHLP,
	},

	{
		N: "SBC A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x9e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opSBCAIXdP,
	},

	{
		N: "SBC A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x9e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opSBCAIYdP,
	},

	{
		N: "AND r",
		C: []Code{
			{0xa0, 0x07, vReg8},
		},
		T: []int{4},
		F: opANDr,
	},

	{
		N: "AND n",
		C: []Code{
			{0xe6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opANDn,
	},

	{
		N: "AND (HL)",
		C: []Code{
			{0xa6, 0x00, nil},
		},
		T: []int{4, 3},
		F: opANDHLP,
	},

	{
		N: "AND (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xa6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opANDIXdP,
	},

	{
		N: "AND (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xa6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opANDIYdP,
	},

	{
		N: "OR r",
		C: []Code{
			{0xb0, 0x07, vReg8},
		},
		T: []int{4},
		F: opORr,
	},

	{
		N: "OR n",
		C: []Code{
			{0xf6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opORn,
	},

	{
		N: "OR (HL)",
		C: []Code{
			{0xb6, 0x00, nil},
		},
		T: []int{4, 3},
		F: opORHLP,
	},

	{
		N: "OR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xb6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opORIXdP,
	},

	{
		N: "OR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xb6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opORIYdP,
	},

	{
		N: "XOR r",
		C: []Code{
			{0xa8, 0x07, vReg8},
		},
		T: []int{4},
		F: opXORr,
	},

	{
		N: "XOR n",
		C: []Code{
			{0xee, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opXORn,
	},

	{
		N: "XOR (HL)",
		C: []Code{
			{0xae, 0x00, nil},
		},
		T: []int{4, 3},
		F: opXORHLP,
	},

	{
		N: "XOR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xae, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opXORIXdP,
	},

	{
		N: "XOR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xae, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opXORIYdP,
	},

	{
		N: "CP r",
		C: []Code{
			{0xb8, 0x07, vReg8},
		},
		T: []int{4},
		F: opCPr,
	},

	{
		N: "CP n",
		C: []Code{
			{0xfe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: opCPn,
	},

	{
		N: "CP (HL)",
		C: []Code{
			{0xbe, 0x00, nil},
		},
		T: []int{4, 3},
		F: opCPHLP,
	},

	{
		N: "CP (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xbe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opCPIXdP,
	},

	{
		N: "CP (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xbe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: opCPIYdP,
	},

	{
		N: "INC r",
		C: []Code{
			{0x04, 0x38, vReg8_3},
		},
		T: []int{4},
		F: opINCr,
	},

	{
		N: "INC (HL)",
		C: []Code{
			{0x34, 0x00, nil},
		},
		T: []int{4, 4, 3},
		F: opINCHLP,
	},

	{
		N: "INC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x34, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opINCIXdP,
	},

	{
		N: "INC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x34, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opINCIYdP,
	},

	{
		N: "DEC r",
		C: []Code{
			{0x05, 0x38, vReg8_3},
		},
		T: []int{4},
		F: opDECr,
	},

	{
		N: "DEC (HL)",
		C: []Code{
			{0x35, 0x00, nil},
		},
		T: []int{4, 4, 3},
		F: opDECHLP,
	},

	{
		N: "DEC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x35, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opDECIXdP,
	},

	{
		N: "DEC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x35, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: opDECIYdP,
	},
}

func opADDAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.addU8(a, x)
}

func opADDAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.addU8(a, n)
}

func opADDAHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.addU8(a, x)
}

func opADDAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	//fmt.Printf("a=%02x p=%04x x=%02x ix=%04x\n", a, p, x, cpu.IX)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func opADDAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func opADCAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func opADCAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.adcU8(a, n)
}

func opADCAHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func opADCAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func opADCAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func opSUBAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.subU8(a, x)
}

func opSUBAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.subU8(a, n)
}

func opSUBAHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.subU8(a, x)
}

func opSUBAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func opSUBAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func opSBCAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func opSBCAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.sbcU8(a, n)
}

func opSBCAHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func opSBCAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func opSBCAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func opANDr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.andU8(a, x)
}

func opANDn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.andU8(a, n)
}

func opANDHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.andU8(a, x)
}

func opANDIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func opANDIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func opORr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.orU8(a, x)
}

func opORn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.orU8(a, n)
}

func opORHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.orU8(a, x)
}

func opORIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func opORIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func opXORr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func opXORn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.xorU8(a, n)
}

func opXORHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func opXORIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func opXORIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func opCPr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.subU8(a, x)
}

func opCPn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.subU8(a, n)
}

func opCPHLP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.subU8(a, x)
}

func opCPIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.subU8(a, x)
}

func opCPIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.subU8(a, x)
}

func opINCr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[0] >> 3)
	*r = cpu.incU8(*r)
}

func opINCHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func opINCIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func opINCIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func opDECr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[0] >> 3)
	*r = cpu.decU8(*r)
}

func opDECHLP(cpu *CPU, codes []uint8) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}

func opDECIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}

func opDECIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}
