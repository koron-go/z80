package opcode

var inout = []*OPCode{

	{
		N: "IN A, (n)",
		C: []Code{
			{0xdb, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 4},
	},

	{
		N: "IN r (C)",
		C: []Code{
			{0xed, 0x00, nil},
			{0x40, 0x38, vReg8_3},
		},
		T: []int{4, 4, 4},
	},

	{
		N: "INI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa2, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
	},

	{
		N: "INIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb2, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
	},

	{
		N: "IND",
		C: []Code{
			{0xed, 0x00, nil},
			{0xaa, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
	},

	{
		N: "INDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xba, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
	},

	{
		N: "OUT (n), A",
		C: []Code{
			{0xd3, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 4},
	},

	{
		N: "OUT (C), r",
		C: []Code{
			{0xed, 0x00, nil},
			{0x41, 0x38, vReg8_3},
		},
		T: []int{4, 4, 4},
	},

	{
		N: "OUTI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa3, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
	},

	{
		N: "OTIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb3, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
	},

	{
		N: "OUTD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xab, 0x00, nil},
		},
		T: []int{4, 5, 3, 4},
	},

	{
		N: "OTDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xbb, 0x00, nil},
		},
		T:  []int{4, 5, 3, 4, 5},
		T2: []int{4, 5, 3, 4},
	},
}
