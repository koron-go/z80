package z80

var exbtsg = []*OPCode{

	{
		N: "EX DE, HL",
		C: []Code{
			{0xeb, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.HL, cpu.DE = cpu.DE, cpu.HL
		},
	},

	{
		N: "EX AF, AF'",
		C: []Code{
			{0x08, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.AF, cpu.Alternate.AF = cpu.Alternate.AF, cpu.AF
		},
	},

	{
		N: "EXX",
		C: []Code{
			{0xd9, 0x00, nil},
		},
		T: []int{4},
		F: func(cpu *CPU, codes []uint8) {
			cpu.BC, cpu.Alternate.BC = cpu.Alternate.BC, cpu.BC
			cpu.DE, cpu.Alternate.DE = cpu.Alternate.DE, cpu.DE
			cpu.HL, cpu.Alternate.HL = cpu.Alternate.HL, cpu.HL
		},
	},

	{
		N: "EX (SP), HL",
		C: []Code{
			{0xe3, 0x00, nil},
		},
		T: []int{4, 3, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.readU16(cpu.SP)
			cpu.writeU16(cpu.SP, cpu.HL.U16())
			cpu.HL.SetU16(v)
		},
	},

	{
		N: "EX (SP), IX",
		C: []Code{
			{0xdd, 0x00, nil},
			{0xe3, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.readU16(cpu.SP)
			cpu.writeU16(cpu.SP, cpu.IX)
			cpu.IX = v
		},
	},

	{
		N: "EX (SP), IY",
		C: []Code{
			{0xfd, 0x00, nil},
			{0xe3, 0x00, nil},
		},
		T: []int{4, 4, 3, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			v := cpu.readU16(cpu.SP)
			cpu.writeU16(cpu.SP, cpu.IY)
			cpu.IY = v
		},
	},

	{
		N: "LDI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa0, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			de := cpu.DE.U16()
			hl := cpu.HL.U16()
			cpu.writeU16(de, cpu.readU16(hl))
			cpu.DE.SetU16(de + 1)
			cpu.HL.SetU16(hl + 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Reset(H).
				Put(PV, bc != 0).
				Reset(N))
		},
	},

	{
		N: "LDIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb0, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			de := cpu.DE.U16()
			hl := cpu.HL.U16()
			cpu.writeU16(de, cpu.readU16(hl))
			cpu.DE.SetU16(de + 1)
			cpu.HL.SetU16(hl + 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Reset(H).
				Put(PV, bc != 0).
				Reset(N))
			if bc != 0 {
				cpu.PC -= 2
			}
		},
	},

	{
		N: "LDD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa8, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			de := cpu.DE.U16()
			hl := cpu.HL.U16()
			cpu.writeU16(de, cpu.readU16(hl))
			cpu.DE.SetU16(de - 1)
			cpu.HL.SetU16(hl - 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Reset(H).
				Put(PV, bc != 0).
				Reset(N))
		},
	},

	{
		N: "LDDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb8, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			de := cpu.DE.U16()
			hl := cpu.HL.U16()
			cpu.writeU16(de, cpu.readU16(hl))
			cpu.DE.SetU16(de - 1)
			cpu.HL.SetU16(hl - 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Reset(H).
				Put(PV, bc != 0).
				Reset(N))
			if bc != 0 {
				cpu.PC -= 2
			}
		},
	},

	{
		N: "CPI",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa1, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			hl := cpu.HL.U16()
			x := cpu.Memory.Get(hl)
			v := a - x
			cpu.HL.SetU16(hl + 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Put(S, v&0x80 != 0).
				Put(Z, v == 0).
				Put(H, a&0x0f < x&0x0f).
				Put(PV, bc != 0).
				Set(N))
		},
	},

	{
		N: "CPIR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb1, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			hl := cpu.HL.U16()
			x := cpu.Memory.Get(hl)
			v := a - x
			cpu.HL.SetU16(hl + 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Put(S, v&0x80 != 0).
				Put(Z, v == 0).
				Put(H, a&0x0f < x&0x0f).
				Put(PV, bc != 0).
				Set(N))
			if bc != 0 {
				cpu.PC -= 2
			}
		},
	},

	{
		N: "CPD",
		C: []Code{
			{0xed, 0x00, nil},
			{0xa9, 0x00, nil},
		},
		T: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			hl := cpu.HL.U16()
			x := cpu.Memory.Get(hl)
			v := a - x
			cpu.HL.SetU16(hl - 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Put(S, v&0x80 != 0).
				Put(Z, v == 0).
				Put(H, a&0x0f < x&0x0f).
				Put(PV, bc != 0).
				Set(N))
		},
	},

	{
		N: "CPDR",
		C: []Code{
			{0xed, 0x00, nil},
			{0xb9, 0x00, nil},
		},
		T:  []int{4, 4, 3, 5, 5},
		T2: []int{4, 4, 3, 5},
		F: func(cpu *CPU, codes []uint8) {
			a := cpu.AF.Hi
			hl := cpu.HL.U16()
			x := cpu.Memory.Get(hl)
			v := a - x
			cpu.HL.SetU16(hl - 1)
			bc := cpu.BC.U16() - 1
			cpu.BC.SetU16(bc)
			cpu.flagUpdate(FlagOp{}.
				Put(S, v&0x80 != 0).
				Put(Z, v == 0).
				Put(H, a&0x0f < x&0x0f).
				Put(PV, bc != 0).
				Set(N))
			if bc != 0 {
				cpu.PC -= 2
			}
		},
	},
}
