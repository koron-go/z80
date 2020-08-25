package opcode

var opcRETI = &OPCode{
	N: "RETI",
	C: []Code{
		{0xed, 0x00, nil},
		{0x4d, 0x00, nil},
	},
	T: []int{4, 4, 3, 3},
}

var opcRETN = &OPCode{
	N: "RETN",
	C: []Code{
		{0xed, 0x00, nil},
		{0x45, 0x00, nil},
	},
	T: []int{4, 4, 3, 3},
}

var callret = []*OPCode{

	{
		N: "CALL nn",
		C: []Code{
			{0xcd, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 4, 3, 3},
	},

	{
		N: "CALL cc, nn",
		C: []Code{
			{0xc4, 0x38, vCC3_3},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 4, 3, 3},
		T2: []int{4, 3, 3},
	},

	{
		N: "RET",
		C: []Code{
			{0xc9, 0x00, nil},
		},
		T: []int{4, 3, 3},
	},

	{
		N: "RET cc",
		C: []Code{
			{0xc0, 0x38, vCC3_3},
		},
		T:  []int{5, 3, 3},
		T2: []int{5},
	},

	opcRETI,

	opcRETN,

	{
		N: "RST p",
		C: []Code{
			{0xc7, 0x38, vRSTp3_3},
		},
		T: []int{5, 3, 3},
	},
}
