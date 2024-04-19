package z80

func oopCALLnn(cpu *CPU) {
	nn := cpu.fetch16()
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, uint8(cpu.PC))
	cpu.Memory.Set(cpu.SP+1, uint8(cpu.PC>>8))
	cpu.PC = nn
}

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

func copCALLxnn(cpu *CPU, xflag bool) {
	nn := cpu.fetch16()
	if xflag {
		cpu.SP -= 2
		cpu.Memory.Set(cpu.SP, uint8(cpu.PC))
		cpu.Memory.Set(cpu.SP+1, uint8(cpu.PC>>8))
		cpu.PC = nn
	}
}

func xopCALLnZnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskZ == 0)
}

func xopCALLfZnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskZ != 0)
}

func xopCALLnCnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskC == 0)
}

func xopCALLfCnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskC != 0)
}

func xopCALLnPVnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskPV == 0)
}

func xopCALLfPVnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskPV != 0)
}

func xopCALLnSnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskS == 0)
}

func xopCALLfSnn(cpu *CPU) {
	copCALLxnn(cpu, cpu.AF.Lo&maskS != 0)
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
