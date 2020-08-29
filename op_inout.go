package z80

func oopINAnP(cpu *CPU, n uint8) {
	cpu.AF.Hi = cpu.ioIn(n)
}

func oopINI(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func oopINIR(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func oopIND(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func oopINDR(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func oopOUTnPA(cpu *CPU, n uint8) {
	cpu.ioOut(n, cpu.AF.Hi)
}

func oopOUTI(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func oopOTIR(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func oopOUTD(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func oopOTDR(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}
