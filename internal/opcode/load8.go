package opcode

var load8 = []*OPCode{

	{
		N: "LD r1, r2",
		C: []Code{
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
	},

	{
		N: "LD r, n",
		C: []Code{
			{0x06, 0x38, vReg88},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "LD r, (HL)",
		C: []Code{
			{0x46, 0x38, vReg8_3},
		},
		T: []int{4, 3},
	},

	{
		N: "LD r, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x46, 0x38, vReg8_3},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "LD r, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x46, 0x38, vReg8_3},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "LD (HL), r",
		C: []Code{
			{0x70, 0x07, vReg8},
		},
		T: []int{4, 3},
	},

	{
		N: "LD (IX+d), r",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x70, 0x07, vReg8},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "LD (IY+d), r",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x70, 0x07, vReg8},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "LD (HL), n",
		C: []Code{
			{0x36, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3},
	},

	{
		N: "LD (IX+d), n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x36, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "LD (IY+d), n",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x36, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 4, 3, 5, 3},
	},

	{
		N: "LD A, (BC)",
		C: []Code{
			{0x0a, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "LD A, (DE)",
		C: []Code{
			{0x1a, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "LD A, (nn)",
		C: []Code{
			{0x3a, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3},
	},

	{
		N: "LD (BC), A",
		C: []Code{
			{0x02, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "LD (DE), A",
		C: []Code{
			{0x12, 0x00, nil},
		},
		T: []int{4, 3},
	},

	{
		N: "LD (nn), A",
		C: []Code{
			{0x32, 0x00, nil},
			{0x00, 0xff, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3, 3, 3},
	},

	{
		N: "LD A, I",
		C: []Code{
			{0xed, 0x00, nil},
			{0x57, 0x00, nil},
		},
		T: []int{4, 5},
	},

	{
		N: "LD A, R",
		C: []Code{
			{0xed, 0x00, nil},
			{0x5f, 0x00, nil},
		},
		T: []int{4, 5},
	},

	{
		N: "LD I, A",
		C: []Code{
			{0xed, 0x00, nil},
			{0x47, 0x00, nil},
		},
		T: []int{4, 5},
	},

	{
		N: "LD R, A",
		C: []Code{
			{0xed, 0x00, nil},
			{0x4f, 0x00, nil},
		},
		T: []int{4, 5},
	},
}
