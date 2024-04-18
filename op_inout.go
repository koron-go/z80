package z80

import "math/bits"

func oopINAnP(cpu *CPU) {
	n := cpu.fetch()
	cpu.AF.Hi = cpu.ioIn(n)
}

func (cpu *CPU) updateFlagIObZ() {
	var nand uint8 = maskZ
	var or uint8
	if cpu.BC.Hi == 0 {
		or |= maskZ
	}
	or |= maskN
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopINI(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
}

func oopINIR(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func oopIND(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
}

func oopINDR(cpu *CPU) {
	cpu.Memory.Set(cpu.HL.U16(), cpu.ioIn(cpu.BC.Hi))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func oopOUTnPA(cpu *CPU) {
	n := cpu.fetch()
	cpu.ioOut(n, cpu.AF.Hi)
}

func oopOUTI(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
}

func oopOTIR(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() + 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

func oopOUTD(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
}

func oopOTDR(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.Memory.Get(cpu.HL.U16()))
	cpu.BC.Hi--
	cpu.HL.SetU16(cpu.HL.U16() - 1)
	cpu.updateFlagIObZ()
	if cpu.BC.Hi != 0 {
		cpu.PC -= 2
	}
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func (cpu *CPU) updateIOIn(r uint8) {
	var nand uint8 = maskS53 | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= r & maskS53
	if r == 0 {
		or |= maskZ
	}
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & maskPV
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func xopINbCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.BC.Hi = r
	cpu.updateIOIn(r)
}

func xopINcCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.BC.Lo = r
	cpu.updateIOIn(r)
}

func xopINdCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.DE.Hi = r
	cpu.updateIOIn(r)
}

func xopINeCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.DE.Lo = r
	cpu.updateIOIn(r)
}

func xopINhCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.HL.Hi = r
	cpu.updateIOIn(r)
}

func xopINlCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.HL.Lo = r
	cpu.updateIOIn(r)
}

func xopINaCP(cpu *CPU) {
	r := cpu.ioIn(cpu.BC.Lo)
	cpu.AF.Hi = r
	cpu.updateIOIn(r)
}

func xopOUTCPb(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.BC.Hi)
}

func xopOUTCPc(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.BC.Lo)
}

func xopOUTCPd(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.DE.Hi)
}

func xopOUTCPe(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.DE.Lo)
}

func xopOUTCPh(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.HL.Hi)
}

func xopOUTCPl(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.HL.Lo)
}

func xopOUTCPa(cpu *CPU) {
	cpu.ioOut(cpu.BC.Lo, cpu.AF.Hi)
}
