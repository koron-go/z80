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

func xopJPnZnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagZ() {
		cpu.PC = nn
	}
}

func xopJPfZnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagZ() {
		cpu.PC = nn
	}
}

func xopJPnCnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagC() {
		cpu.PC = nn
	}
}

func xopJPfCnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagC() {
		cpu.PC = nn
	}
}

func xopJPnPVnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagPV() {
		cpu.PC = nn
	}
}

func xopJPfPVnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagPV() {
		cpu.PC = nn
	}
}

func xopJPnSnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagS() {
		cpu.PC = nn
	}
}

func xopJPfSnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagS() {
		cpu.PC = nn
	}
}
