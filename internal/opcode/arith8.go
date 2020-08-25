package opcode

var arith8 = []*OPCode{

	{
		N: "ADD A, r",
		C: []Code{
			{0x80, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "ADD A, n",
		C: []Code{
			{0xc6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "ADD A, (HL)",
		C: []Code{
			{0x86, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "ADD A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x86, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "ADD A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x86, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "ADC A, r",
		C: []Code{
			{0x88, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "ADC A, n",
		C: []Code{
			{0xce, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "ADC A, (HL)",
		C: []Code{
			{0x8e, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "ADC A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x8e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "ADC A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x8e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "SUB A, r",
		C: []Code{
			{0x90, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "SUB A, n",
		C: []Code{
			{0xd6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "SUB A, (HL)",
		C: []Code{
			{0x96, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "SUB A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x96, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "SUB A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x96, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "SBC A, r",
		C: []Code{
			{0x98, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "SBC A, n",
		C: []Code{
			{0xde, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "SBC A, (HL)",
		C: []Code{
			{0x9e, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "SBC A, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x9e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "SBC A, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x9e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "AND r",
		C: []Code{
			{0xa0, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "AND n",
		C: []Code{
			{0xe6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "AND (HL)",
		C: []Code{
			{0xa6, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "AND (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xa6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "AND (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xa6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "OR r",
		C: []Code{
			{0xb0, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "OR n",
		C: []Code{
			{0xf6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "OR (HL)",
		C: []Code{
			{0xb6, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "OR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xb6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "OR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xb6, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "XOR r",
		C: []Code{
			{0xa8, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "XOR n",
		C: []Code{
			{0xee, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "XOR (HL)",
		C: []Code{
			{0xae, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "XOR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xae, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "XOR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xae, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "CP r",
		C: []Code{
			{0xb8, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "CP n",
		C: []Code{
			{0xfe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "CP (HL)",
		C: []Code{
			{0xbe, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "CP (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xbe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "CP (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xbe, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "INC r",
		C: []Code{
			{0x04, 0x38, vReg8_3},
		},
		T: []int{4},
	},

	{
		N: "INC (HL)",
		C: []Code{
			{0x34, 0x00, nil},
		},
		T: []int{4, 4, 3},
	},

	{
		N: "INC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x34, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "INC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x34, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "DEC r",
		C: []Code{
			{0x05, 0x38, vReg8_3},
		},
		T: []int{4},
	},

	{
		N: "DEC (HL)",
		C: []Code{
			{0x35, 0x00, nil},
		},
		T: []int{4, 4, 3},
	},

	{
		N: "DEC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x35, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "DEC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x35, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},
}
