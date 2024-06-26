package z80

func oopCALLnn(cpu *CPU) {
	nn := cpu.fetch16()
	copCALLnn(cpu, nn)
}

func oopRETI(cpu *CPU) {
	if cpu.RETIHandler != nil {
		cpu.RETIHandler.RETIHandle()
	}

	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

func oopRETN(cpu *CPU) {
	if cpu.RETNHandler != nil {
		cpu.RETNHandler.RETNHandle()
	}

	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.IFF1 = cpu.IFF2
}

func oopRET(cpu *CPU) {
	//cpu.PC = cpu.readU16(cpu.SP)
	//cpu.SP += 2
	l := cpu.Memory.Get(cpu.SP)
	cpu.SP++
	h := cpu.Memory.Get(cpu.SP)
	cpu.SP++
	cpu.PC = (uint16(h) << 8) | uint16(l)
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func copCALLnn(cpu *CPU, nn uint16) {
	cpu.SP--
	cpu.Memory.Set(cpu.SP, uint8(cpu.PC>>8))
	cpu.SP--
	cpu.Memory.Set(cpu.SP, uint8(cpu.PC))
	cpu.PC = nn
}

func xopCALLnZnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagZ() {
		copCALLnn(cpu, nn)
	}
}

func xopCALLfZnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagZ() {
		copCALLnn(cpu, nn)
	}
}

func xopCALLnCnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagC() {
		copCALLnn(cpu, nn)
	}
}

func xopCALLfCnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagC() {
		copCALLnn(cpu, nn)
	}
}

func xopCALLnPVnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagPV() {
		copCALLnn(cpu, nn)
	}
}

func xopCALLfPVnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagPV() {
		copCALLnn(cpu, nn)
	}
}

func xopCALLnSnn(cpu *CPU) {
	nn := cpu.fetch16()
	if !cpu.flagS() {
		copCALLnn(cpu, nn)
	}
}

func xopCALLfSnn(cpu *CPU) {
	nn := cpu.fetch16()
	if cpu.flagS() {
		copCALLnn(cpu, nn)
	}
}

func xopRETnZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ == 0 {
		oopRET(cpu)
	}
}

func xopRETfZ(cpu *CPU) {
	if cpu.AF.Lo&maskZ != 0 {
		oopRET(cpu)
	}
}

func xopRETnC(cpu *CPU) {
	if cpu.AF.Lo&maskC == 0 {
		oopRET(cpu)
	}
}

func xopRETfC(cpu *CPU) {
	if cpu.AF.Lo&maskC != 0 {
		oopRET(cpu)
	}
}

func xopRETnPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV == 0 {
		oopRET(cpu)
	}
}

func xopRETfPV(cpu *CPU) {
	if cpu.AF.Lo&maskPV != 0 {
		oopRET(cpu)
	}
}

func xopRETnS(cpu *CPU) {
	if cpu.AF.Lo&maskS == 0 {
		oopRET(cpu)
	}
}

func xopRETfS(cpu *CPU) {
	if cpu.AF.Lo&maskS != 0 {
		oopRET(cpu)
	}
}

func xopRST00(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0000
}

func xopRST08(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0008
}

func xopRST10(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0010
}

func xopRST18(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0018
}

func xopRST20(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0020
}

func xopRST28(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0028
}

func xopRST30(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0030
}

func xopRST38(cpu *CPU) {
	cpu.SP -= 2
	cpu.writeU16(cpu.SP, cpu.PC)
	cpu.PC = 0x0038
}
