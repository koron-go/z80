package opcode

var jump = []*OPCode{

	{
		N: "JP nn",
		C: []Code{
			{0xc3, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3},
	},

	{
		N: "JP cc, nn",
		C: []Code{
			{0xc2, 0x38, vCC3_3},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 4, 3, 3},
		T2: []int{4, 3, 3},
	},

	{
		N: "JR e",
		C: []Code{
			{0x18, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 5},
	},

	{
		N: "JR C, e",
		C: []Code{
			{0x38, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
	},

	{
		N: "JR NC, e",
		C: []Code{
			{0x30, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
	},

	{
		N: "JR Z, e",
		C: []Code{
			{0x28, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
	},

	{
		N: "JR NZ, e",
		C: []Code{
			{0x20, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{4, 3, 5},
		T2: []int{4, 3},
	},

	{
		N: "JP (HL)",
		C: []Code{
			{0xe9, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "JP (IX)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe9, 0x00, nil},
		},
		T: []int{4, 4},
	},

	{
		N: "JP (IY)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe9, 0x00, nil},
		},
		T: []int{4, 4},
	},

	{
		N: "DJNZ e",
		C: []Code{
			{0x10, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T:  []int{5, 3, 5},
		T2: []int{5, 3},
	},
}
