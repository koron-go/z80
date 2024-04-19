package z80

func oopJPnn(cpu *CPU) {
	cpu.PC = cpu.fetch16()
}

func oopJRe(cpu *CPU) {
	off := cpu.fetch()
	cpu.PC = addrOff(cpu.PC, off)
}

func oopJRCe(cpu *CPU) {
	off := cpu.fetch()
	if cpu.flagC() {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRNCe(cpu *CPU) {
	off := cpu.fetch()
	if !cpu.flagC() {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRZe(cpu *CPU) {
	off := cpu.fetch()
	if cpu.flagZ() {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

func oopJRNZe(cpu *CPU) {
	off := cpu.fetch()
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

func oopDJNZe(cpu *CPU) {
	off := cpu.fetch()
	cpu.BC.Hi--
	if cpu.BC.Hi != 0 {
		cpu.PC = addrOff(cpu.PC, off)
	}
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func copJPxnn(cpu *CPU, xflag bool) {
	nn := cpu.fetch16()
	if xflag {
		cpu.PC = nn
	}
}

func xopJPnZnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskZ == 0)
}

func xopJPfZnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskZ != 0)
}

func xopJPnCnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskC == 0)
}

func xopJPfCnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskC != 0)
}

func xopJPnPVnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskPV == 0)
}

func xopJPfPVnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskPV != 0)
}

func xopJPnSnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskS == 0)
}

func xopJPfSnn(cpu *CPU) {
	copJPxnn(cpu, cpu.AF.Lo&maskS != 0)
}
