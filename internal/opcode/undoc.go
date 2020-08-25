package opcode

var undoc = []*OPCode{
	{
		N: "INC IXH",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x24, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "DEC IXH",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x25, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "INC IXL",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2c, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "DEC IXL",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2d, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "INC IYH",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x24, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "DEC IYH",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x25, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "INC IYL",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2c, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "DEC IYL",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2d, 0x00, nil},
		},
		T: []int{4, 6}, // not verified
	},

	{
		N: "LD IXH, n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x26, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
	},

	{
		N: "LD IXL, n",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x2e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
	},

	{
		N: "LD IYH, n",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x26, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
	},

	{
		N: "LD IYL, n",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x2e, 0x00, nil},
			{0x00, 0xff, nil},
		},
		T: []int{4, 3}, // not verified
	},

	{
		N: "SL1 (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x36, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3}, // not verified
	},

	{
		N: "SL1 (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x36, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3}, // not verified
	},

	{
		N: "SL1 r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x30, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "SL1 (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x36, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "LD rx1, rx2",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
	},

	{
		N: "LD ry1, ry2",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x40, 0x3f, vReg88},
		},
		T: []int{4},
	},

	{
		N: "ADD A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x80, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "ADD A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x80, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "ADC A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x88, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "ADC A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x88, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "SUB A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x90, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "SUB A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x90, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "SBC A, rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0x98, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "SBC A, ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0x98, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "AND rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xa0, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "AND ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xa0, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "XOR rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xa8, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "XOR ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xa8, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "OR rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xb0, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "OR ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xb0, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "CP rx",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xb8, 0x07, vReg8},
		},
		T: []int{4},
	},

	{
		N: "CP ry",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xb8, 0x07, vReg8},
		},
		T: []int{4},
	},
}
