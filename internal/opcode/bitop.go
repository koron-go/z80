package opcode

var bitop = []*OPCode{

	{
		N: "BIT b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x40, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
	},

	{
		N: "BIT b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x46, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4},
	},

	{
		N: "BIT b, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x46, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4},
	},

	{
		N: "BIT b, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x46, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4},
	},

	{
		N: "SET b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0xc0, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
	},

	{
		N: "SET b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0xc6, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "SET b, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0xc6, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "SET b, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0xc6, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RES b, r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x80, 0x3f, vBit3Reg8},
		},
		T: []int{4, 4},
	},

	{
		N: "RES b, (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x86, 0x38, vBit3_3},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "RES b, (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x86, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RES b, (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x86, 0x38, vBit3_3},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},
}
