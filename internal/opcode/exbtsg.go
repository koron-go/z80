package opcode

var exbtsg = []*OPCode{

	{
		N: "EX DE, HL",
		C: []Code{
			{0xeb, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "EX AF, AF'",
		C: []Code{
			{0x08, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "EXX",
		C: []Code{
			{0xd9, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "EX (SP), HL",
		C: []Code{
			{0xe3, 0x00, nil},
		},
		T: []int{4, 3, 4, 3, 5},
	},

	{
		N: "EX (SP), IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe3, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3, 5},
	},

	{
		N: "EX (SP), IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe3, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3, 5},
	},

	{
		N: "LDI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa0, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
	},

	{
		N: "LDIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb0, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
	},

	{
		N: "LDD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa8, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
	},

	{
		N: "LDDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb8, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
	},

	{
		N: "CPI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa1, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
	},

	{
		N: "CPIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb1, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
	},

	{
		N: "CPD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa9, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
	},

	{
		N: "CPDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb9, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
	},
}
