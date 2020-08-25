package opcode

var rotateshift = []*OPCode{

	{
		N: "RLCA",
		C: []Code{
			{0x07, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "RLA",
		C: []Code{
			{0x17, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "RRCA",
		C: []Code{
			{0x0f, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "RRA",
		C: []Code{
			{0x1f, 0x00, nil},
		},
		T: []int{4},
	},

	{
		N: "RLC r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x00, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "RLC (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x06, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "RLC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x06, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RLC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x06, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RL r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x10, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "RL (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x16, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "RL (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x16, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RL (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x16, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RRC r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x08, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "RRC (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x0e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "RRC (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x0e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RRC (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x0e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RR r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x18, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "RR (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x1e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "RR (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x1e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RR (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x1e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "SLA r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x20, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "SLA (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x26, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "SLA (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x26, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "SLA (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x26, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "SRA r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x28, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "SRA (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x2e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "SRA (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x2e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "SRA (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x2e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "SRL r",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x38, 0x07, vReg8},
		},
		T: []int{4, 4},
	},

	{
		N: "SRL (HL)",
		C: []Code{
			{0xcb, 0x00, nil},
			{0x3e, 0x00, nil},
		},
		T: []int{4, 4, 4, 3},
	},

	{
		N: "SRL (IX+d)",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x3e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "SRL (IY+d)",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xcb, 0x00, nil},
			{0x00, 0xff, nil},
			{0x3e, 0x00, nil},
		},
		T: []int{4, 4, 3, 5, 4, 3},
	},

	{
		N: "RLD",
		C: []Code{
			{0xed, 0x00, nil},
			{0x6f, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3},
	},

	{
		N: "RRD",
		C: []Code{
			{0xed, 0x00, nil},
			{0x67, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3},
	},
}
