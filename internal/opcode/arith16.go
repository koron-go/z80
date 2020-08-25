package opcode

var arith16 = []*OPCode{

	{
		N: "ADD HL, ss",
		C: []Code{
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 3},
	},

	{
		N: "ADC HL, ss",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4a, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "SBC HL, ss",
		C: []Code{
			{0xed, 0x00, nil},
			{0x42, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "ADD IX, pp",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "ADD IY, rr",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x09, 0x30, vReg16_4},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "INC ss",
		C: []Code{
			{0x03, 0x30, vReg16_4},
		},
		T: []int{6},
	},

	{
		N: "INC IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x23, 0x00, nil},
		},
		T: []int{4, 6},
	},

	{
		N: "INC IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x23, 0x00, nil},
		},
		T: []int{4, 6},
	},

	{
		N: "DEC ss",
		C: []Code{
			{0x0b, 0x30, vReg16_4},
		},
		T: []int{6},
	},

	{
		N: "DEC IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2b, 0x00, nil},
		},
		T: []int{4, 6},
	},

	{
		N: "DEC IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2b, 0x00, nil},
		},
		T: []int{4, 6},
	},
}
