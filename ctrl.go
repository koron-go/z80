package z80

var ctrl = []*OPCode{

	{
		N: "DAA",
		C: []Code{
			{0x27, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			// TODO: implmenet "DAA"
		},
	},

	{
		N: "CPL",
		C: []Code{
			{0x2F, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.AF.Hi = ^cpu.AF.Hi
			cpu.flagUpdate(FlagOp{}.Set(H).Set(N))
		},
	},

	{
		N: "NEG",
		C: []Code{
			{0xed, 0x00, nil},
			{0x44, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			v := ^a + 1
			cpu.AF.Hi = v
			cpu.flagUpdate(FlagOp{}.
				Put(S, v&0x80 != 0).
				Put(Z, v == 0).
				// FIXME: check and fix H flag behavior.
				Put(H, false).
				Put(PV, a == 0x80).
				Set(N).
				Put(C, a != 0))
		},
	},

	{
		N: "CCF",
		C: []Code{
			{0x3f, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			c := cpu.flag(C)
			cpu.flagUpdate(FlagOp{}.Put(H, c).Reset(N).Put(C, !c))
		},
	},

	{
		N: "SCF",
		C: []Code{
			{0x37, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.flagUpdate(FlagOp{}.Reset(H).Reset(N).Set(C))
		},
	},

	{
		N: "NOP",
		C: []Code{
			{0x00, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			// do nothing.
		},
	},

	{
		N: "HALT",
		C: []Code{
			{0x76, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.PC--
		},
	},

	{
		N: "DI",
		C: []Code{
			{0xf3, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IFF1 = false
			cpu.IFF2 = false
		},
	},

	{
		N: "EI",
		C: []Code{
			{0xfb, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IFF1 = true
			cpu.IFF2 = true
		},
	},

	{
		N: "IM 0",
		C: []Code{
			{0xed, 0x00, nil},
			{0x46, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IM = 0
		},
	},

	{
		N: "IM 1",
		C: []Code{
			{0xed, 0x00, nil},
			{0x56, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IM = 1
		},
	},

	{
		N: "IM 2",
		C: []Code{
			{0xed, 0x00, nil},
			{0x5e, 0x00, nil},
		},
		T: []int{4, 4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.IM = 2
		},
	},
}
