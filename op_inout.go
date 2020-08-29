package z80

import "math/bits"

func oopINAnP(cpu *CPU, n uint8) {
	cpu.AF.Hi = cpu.ioIn(n)
}

func opINrCP(cpu *CPU, codes []uint8) {
	v := cpu.ioIn(cpu.BC.Lo)
	// FIXME: support 0x06 to apply flags only.
	r := cpu.regP(codes[1] >> 3)
	*r = v
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N))
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

func opOUTCPr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1] >> 3)
	cpu.ioOut(cpu.BC.Lo, *r)
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
