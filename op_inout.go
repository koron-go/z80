package z80

import "math/bits"

func opINAnP(cpu *CPU, codes []uint8) {
	cpu.AF.Hi = cpu.ioIn(codes[1])
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

func opINI(cpu *CPU, codes []uint8) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opINIR(cpu *CPU, codes []uint8) {
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

func opIND(cpu *CPU, codes []uint8) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opINDR(cpu *CPU, codes []uint8) {
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

func opOUTnPA(cpu *CPU, codes []uint8) {
	cpu.ioOut(codes[1], cpu.AF.Hi)
}

func opOUTCPr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[1] >> 3)
	cpu.ioOut(cpu.BC.Lo, *r)
}

func opOUTI(cpu *CPU, codes []uint8) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opOTIR(cpu *CPU, codes []uint8) {
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

func opOUTD(cpu *CPU, codes []uint8) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.flagUpdate(FlagOp{}.
		Put(Z, cpu.BC.Hi == 0).
		Set(N))
}

func opOTDR(cpu *CPU, codes []uint8) {
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
