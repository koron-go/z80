package z80

func oopJPnn(cpu *CPU, l, h uint8) {
	cpu.PC = toU16(l, h)
}

func oopJRe(cpu *CPU, off uint8) {
	cpu.PC = addrOff(cpu.PC, off)
}

func oopJRCe(cpu *CPU, off uint8) {
	if cpu.flagC() {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRNCe(cpu *CPU, off uint8) {
	if !cpu.flagC() {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRZe(cpu *CPU, off uint8) {
	if cpu.flagZ() {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRNZe(cpu *CPU, off uint8) {
	if !cpu.flagZ() {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJPHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.PC = p
}

func oopJPIXP(cpu *CPU) {
	p := cpu.IX
	cpu.PC = p
}

func oopJPIYP(cpu *CPU) {
	p := cpu.IY
	cpu.PC = p
}

func oopDJNZe(cpu *CPU, off uint8) {
	cpu.BC.Hi--
	if cpu.BC.Hi != 0 {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopJPnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		cpu.PC = toU16(l, h)
	}
}

func xopJPfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		cpu.PC = toU16(l, h)
	}
}
