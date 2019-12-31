package z80

var arith8 = []*OPCode{

	{
		N: "ADD A, r",
		C: []Code{
			{0x80, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.AF.Hi = cpu.addU8(a, x)
		},
	},

	{
		N: "ADD A, n",
		C: []Code{
			{0xc6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.AF.Hi = cpu.addU8(a, n)
		},
	},

	{
		N: "ADD A, (HL)",
		C: []Code{
			{0x86, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.AF.Hi = cpu.addU8(a, x)
		},
	},

	{
		N: "ADD A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x86, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.addU8(a, x)
		},
	},

	{
		N: "ADD A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x86, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.addU8(a, x)
		},
	},

	{
		N: "ADC A, r",
		C: []Code{
			{0x88, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.AF.Hi = cpu.adcU8(a, x)
		},
	},

	{
		N: "ADC A, n",
		C: []Code{
			{0xce, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.AF.Hi = cpu.adcU8(a, n)
		},
	},

	{
		N: "ADC A, (HL)",
		C: []Code{
			{0x8e, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.AF.Hi = cpu.adcU8(a, x)
		},
	},

	{
		N: "ADC A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x8e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.adcU8(a, x)
		},
	},

	{
		N: "ADC A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x8e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.adcU8(a, x)
		},
	},

	{
		N: "SUB A, r",
		C: []Code{
			{0x90, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.AF.Hi = cpu.subU8(a, x)
		},
	},

	{
		N: "SUB A, n",
		C: []Code{
			{0xd6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.AF.Hi = cpu.subU8(a, n)
		},
	},

	{
		N: "SUB A, (HL)",
		C: []Code{
			{0x96, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.AF.Hi = cpu.subU8(a, x)
		},
	},

	{
		N: "SUB A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x96, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.subU8(a, x)
		},
	},

	{
		N: "SUB A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x96, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.subU8(a, x)
		},
	},

	{
		N: "SBC A, r",
		C: []Code{
			{0x98, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.AF.Hi = cpu.sbcU8(a, x)
		},
	},

	{
		N: "SBC A, n",
		C: []Code{
			{0xde, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.AF.Hi = cpu.sbcU8(a, n)
		},
	},

	{
		N: "SBC A, (HL)",
		C: []Code{
			{0x9e, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.AF.Hi = cpu.sbcU8(a, x)
		},
	},

	{
		N: "SBC A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x9e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.sbcU8(a, x)
		},
	},

	{
		N: "SBC A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x9e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.sbcU8(a, x)
		},
	},

	{
		N: "AND r",
		C: []Code{
			{0xa0, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.AF.Hi = cpu.andU8(a, x)
		},
	},

	{
		N: "AND n",
		C: []Code{
			{0xe6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.AF.Hi = cpu.andU8(a, n)
		},
	},

	{
		N: "AND (HL)",
		C: []Code{
			{0xa6, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.AF.Hi = cpu.andU8(a, x)
		},
	},

	{
		N: "AND (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xa6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.andU8(a, x)
		},
	},

	{
		N: "AND (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xa6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.andU8(a, x)
		},
	},

	{
		N: "OR r",
		C: []Code{
			{0xb0, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.AF.Hi = cpu.orU8(a, x)
		},
	},

	{
		N: "OR n",
		C: []Code{
			{0xf6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.AF.Hi = cpu.orU8(a, n)
		},
	},

	{
		N: "OR (HL)",
		C: []Code{
			{0xb6, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.AF.Hi = cpu.orU8(a, x)
		},
	},

	{
		N: "OR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xb6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.orU8(a, x)
		},
	},

	{
		N: "OR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xb6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.orU8(a, x)
		},
	},

	{
		N: "XOR r",
		C: []Code{
			{0xa8, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.AF.Hi = cpu.xorU8(a, x)
		},
	},

	{
		N: "XOR n",
		C: []Code{
			{0xee, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.AF.Hi = cpu.xorU8(a, n)
		},
	},

	{
		N: "XOR (HL)",
		C: []Code{
			{0xae, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.AF.Hi = cpu.xorU8(a, x)
		},
	},

	{
		N: "XOR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xae, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.xorU8(a, x)
		},
	},

	{
		N: "XOR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xae, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.AF.Hi = cpu.xorU8(a, x)
		},
	},

	{
		N: "CP r",
		C: []Code{
			{0xb8, 0x07, vReg8},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := *cpu.regP(codes[0])
			cpu.subU8(a, x)
		},
	},

	{
		N: "CP n",
		C: []Code{
			{0xfe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			n := codes[1]
			cpu.subU8(a, n)
		},
	},

	{
		N: "CP (HL)",
		C: []Code{
			{0xbe, 0x00, nil},
		},
		T: []int{4, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			x := cpu.Memory.Get(cpu.HL.U16())
			cpu.subU8(a, x)
		},
	},

	{
		N: "CP (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xbe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.subU8(a, x)
		},
	},

	{
		N: "CP (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xbe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.subU8(a, x)
		},
	},

	{
		N: "INC r",
		C: []Code{
			{0x04, 0x38, vReg8_3},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0])
			*r = cpu.incU8(*r)
		},
	},

	{
		N: "INC (HL)",
		C: []Code{
			{0x34, 0x00, nil},
		},
		T: []int{4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.HL.U16()
			x := cpu.Memory.Get(p)
			cpu.Memory.Set(p, cpu.incU8(x))
		},
	},

	{
		N: "INC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x34, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.Memory.Set(p, cpu.incU8(x))
		},
	},

	{
		N: "INC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x34, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.Memory.Set(p, cpu.incU8(x))
		},
	},

	{
		N: "DEC r",
		C: []Code{
			{0x05, 0x38, vReg8_3},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			r := cpu.regP(codes[0])
			*r = cpu.decU8(*r)
		},
	},

	{
		N: "DEC (HL)",
		C: []Code{
			{0x35, 0x00, nil},
		},
		T: []int{4, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := cpu.HL.U16()
			x := cpu.Memory.Get(p)
			cpu.Memory.Set(p, cpu.decU8(x))
		},
	},

	{
		N: "DEC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x35, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IX, codes[2])
			x := cpu.Memory.Get(p)
			cpu.Memory.Set(p, cpu.decU8(x))
		},
	},

	{
		N: "DEC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x35, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
		F: func(cpu *CPU, codes []uint8) {
			p := addrOff(cpu.IY, codes[2])
			x := cpu.Memory.Get(p)
			cpu.Memory.Set(p, cpu.decU8(x))
		},
	},
}
