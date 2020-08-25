package opcode

var load16 = []*OPCode{

	{
		N: "LD dd, nn",
		C: []Code{
			{0x01, 0x30, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3},
	},

	{
		N: "LD IX, nn",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x21, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3},
	},

	{
		N: "LD IY, nn",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x21, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3},
	},

	{
		N: "LD HL, (nn)",
		C: []Code{
			{0x2a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3, 3},
	},

	{
		N: "LD dd, (nn)",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4b, 0x30, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
	},

	{
		N: "LD IX, (nn)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
	},

	{
		N: "LD IY, (nn)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
	},

	{
		N: "LD (nn), HL",
		C: []Code{
			{0x22, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3, 3},
	},

	{
		N: "LD (nn), dd",
		C: []Code{
			{0xed, 0x00, nil},
			{0x43, 0x30, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
	},

	{
		N: "LD (nn), IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x22, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
	},

	{
		N: "LD (nn), IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x22, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 3, 3, 3},
	},

	{
		N: "LD SP, HL",
		C: []Code{
			{0xf9, 0x00, nil},
		},
		T: []int{6},
	},

	{
		N: "LD SP, IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xf9, 0x00, nil},
		},
		T: []int{4, 6},
	},

	{
		N: "LD SP, IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xf9, 0x00, nil},
		},
		T: []int{4, 6},
	},

	{
		N: "PUSH qq",
		C: []Code{
			{0xc5, 0x30, nil},
		},
		T: []int{5, 3, 3},
	},

	{
		N: "PUSH IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe5, 0x00, nil},
		},
		T: []int{4, 5, 3, 3},
	},

	{
		N: "PUSH IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe5, 0x00, nil},
		},
		T: []int{4, 5, 3, 3},
	},

	{
		N: "POP qq",
		C: []Code{
			{0xc1, 0x30, nil},
		},
		T: []int{4, 3, 3},
	},

	{
		N: "POP IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe1, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
	},

	{
		N: "POP IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe1, 0x00, nil},
		},
		T: []int{4, 4, 3, 3},
	},
}
