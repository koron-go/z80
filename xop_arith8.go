package z80

func xopADDAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADCAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopSUBAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSBCAb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopANDAb(cpu *CPU) {
	cpu.AF.Hi &= cpu.BC.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
}

func xopANDAc(cpu *CPU) {
	cpu.AF.Hi &= cpu.BC.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
}

func xopANDAd(cpu *CPU) {
	cpu.AF.Hi &= cpu.DE.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
}

func xopANDAe(cpu *CPU) {
	cpu.AF.Hi &= cpu.DE.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
}

func xopANDAh(cpu *CPU) {
	cpu.AF.Hi &= cpu.HL.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
}

func xopANDAl(cpu *CPU) {
	cpu.AF.Hi &= cpu.HL.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
}

func xopANDAa(cpu *CPU) {
	cpu.AF.Hi &= cpu.AF.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, true)
}

func xopXORb(cpu *CPU) {
	cpu.AF.Hi ^= cpu.BC.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORc(cpu *CPU) {
	cpu.AF.Hi ^= cpu.BC.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORd(cpu *CPU) {
	cpu.AF.Hi ^= cpu.DE.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORe(cpu *CPU) {
	cpu.AF.Hi ^= cpu.DE.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORh(cpu *CPU) {
	cpu.AF.Hi ^= cpu.HL.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORl(cpu *CPU) {
	cpu.AF.Hi ^= cpu.HL.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopXORa(cpu *CPU) {
	cpu.AF.Hi ^= cpu.AF.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopORb(cpu *CPU) {
	cpu.AF.Hi |= cpu.BC.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopORc(cpu *CPU) {
	cpu.AF.Hi |= cpu.BC.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopORd(cpu *CPU) {
	cpu.AF.Hi |= cpu.DE.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopORe(cpu *CPU) {
	cpu.AF.Hi |= cpu.DE.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopORh(cpu *CPU) {
	cpu.AF.Hi |= cpu.HL.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopORl(cpu *CPU) {
	cpu.AF.Hi |= cpu.HL.Lo
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopORa(cpu *CPU) {
	cpu.AF.Hi |= cpu.AF.Hi
	cpu.updateFlagLogic8(cpu.AF.Hi, false)
}

func xopCPb(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Hi
	cpu.subU8(a, x)
}

func xopCPc(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.BC.Lo
	cpu.subU8(a, x)
}

func xopCPd(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Hi
	cpu.subU8(a, x)
}

func xopCPe(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.DE.Lo
	cpu.subU8(a, x)
}

func xopCPh(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Hi
	cpu.subU8(a, x)
}

func xopCPl(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.HL.Lo
	cpu.subU8(a, x)
}

func xopCPa(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.AF.Hi
	cpu.subU8(a, x)
}
