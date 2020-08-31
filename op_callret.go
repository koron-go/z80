package z80

func oopRETI(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2

	if cpu.INT != nil {
		cpu.INT.ReturnINT()
	}
}

func oopRETN(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
	cpu.IFF1 = cpu.IFF2

	cpu.InNMI = false
}

func oopRET(cpu *CPU) {
	cpu.PC = cpu.readU16(cpu.SP)
	cpu.SP += 2
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

func xopCALLnn(cpu *CPU, l, h uint8) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, uint8(cpu.PC))
	cpu.Memory.Set(cpu.SP+1, uint8(cpu.PC>>8))
	cpu.PC = toU16(l, h)
}

func xopCALLnZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfZnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskZ != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfCnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskC != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfPVnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskPV != 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLnSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS == 0 {
		xopCALLnn(cpu, l, h)
	}
}

func xopCALLfSnn(cpu *CPU, l, h uint8) {
	if cpu.AF.Lo&maskS != 0 {
		xopCALLnn(cpu, l, h)
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
