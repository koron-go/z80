package opcode

var opcHALT = &OPCode{
	N: "HALT",
	C: []Code{
		{0x76, 0x00, nil},
	},
	T: []int{4},
}

var opcEI = &OPCode{
	N: "EI",
	C: []Code{
		{0xfb, 0x00, nil},
	},
	T: []int{4},
}

var ctrl = []*OPCode{

	{
		N: "DAA",
		C: []Code{
			{0x27, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "CPL",
		C: []Code{
			{0x2F, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "NEG",
		C: []Code{
			{0xed, 0x00, nil},
			{0x44, 0x00, nil},
		},
		T: []int{4, 4},
	},

	{
		N: "CCF",
		C: []Code{
			{0x3f, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "SCF",
		C: []Code{
			{0x37, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "NOP",
		C: []Code{
			{0x00, 0x00, nil},
		},
		T: []int{4},
	},

	opcHALT,

	{
		N: "DI",
		C: []Code{
			{0xf3, 0x00, nil},
		},
		T: []int{4},
	},

	opcEI,

	{
		N: "IM 0",
		C: []Code{
			{0xed, 0x00, nil},
			{0x46, 0x00, nil},
		},
		T: []int{4, 4},
	},

	{
		N: "IM 1",
		C: []Code{
			{0xed, 0x00, nil},
			{0x56, 0x00, nil},
		},
		T: []int{4, 4},
	},

	{
		N: "IM 2",
		C: []Code{
			{0xed, 0x00, nil},
			{0x5e, 0x00, nil},
		},
		T: []int{4, 4},
	},
}
