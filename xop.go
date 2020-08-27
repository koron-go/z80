package z80

func xopPOPbc(cpu *CPU) {
	cpu.BC.Lo = cpu.Memory.Get(cpu.SP)
	cpu.BC.Hi = cpu.Memory.Get(cpu.SP+1)
	cpu.SP += 2
}

func xopPOPde(cpu *CPU) {
	cpu.DE.Lo = cpu.Memory.Get(cpu.SP)
	cpu.DE.Hi = cpu.Memory.Get(cpu.SP+1)
	cpu.SP += 2
}

func xopPOPhl(cpu *CPU) {
	cpu.HL.Lo = cpu.Memory.Get(cpu.SP)
	cpu.HL.Hi = cpu.Memory.Get(cpu.SP+1)
	cpu.SP += 2
}

func xopPOPaf(cpu *CPU) {
	cpu.AF.Lo = cpu.Memory.Get(cpu.SP)
	cpu.AF.Hi = cpu.Memory.Get(cpu.SP+1)
	cpu.SP += 2
}

func xopPUSHbc(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.BC.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.BC.Hi)
}

func xopPUSHde(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.DE.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.DE.Hi)
}

func xopPUSHhl(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.HL.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.HL.Hi)
}

func xopPUSHaf(cpu *CPU) {
	cpu.SP -= 2
	cpu.Memory.Set(cpu.SP, cpu.AF.Lo)
	cpu.Memory.Set(cpu.SP+1, cpu.AF.Hi)
}

func xopINCbc(cpu *CPU) {
	cpu.BC.Lo++
	if cpu.BC.Lo == 0 {
		cpu.BC.Hi++
	}
}

func xopINCde(cpu *CPU) {
	cpu.DE.Lo++
	if cpu.DE.Lo == 0 {
		cpu.DE.Hi++
	}
}

func xopINChl(cpu *CPU) {
	cpu.HL.Lo++
	if cpu.HL.Lo == 0 {
		cpu.HL.Hi++
	}
}

func xopINCsp(cpu *CPU) {
	cpu.SP++
}

func xopLDbchHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Hi = cpu.Memory.Get(p)
}

func xopLDbclHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.BC.Lo = cpu.Memory.Get(p)
}

func xopLDdehHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Hi = cpu.Memory.Get(p)
}

func xopLDdelHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.DE.Lo = cpu.Memory.Get(p)
}

func xopLDhlhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Hi = cpu.Memory.Get(p)
}

func xopLDhllHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.HL.Lo = cpu.Memory.Get(p)
}

func xopLDafhHLP(cpu *CPU) {
	p := cpu.HL.U16()
	cpu.AF.Hi = cpu.Memory.Get(p)
}

func xopDECb(cpu *CPU) {
	cpu.decP8(&cpu.BC.Hi)
}

func xopDECc(cpu *CPU) {
	cpu.decP8(&cpu.BC.Lo)
}

func xopDECd(cpu *CPU) {
	cpu.decP8(&cpu.DE.Hi)
}

func xopDECe(cpu *CPU) {
	cpu.decP8(&cpu.DE.Lo)
}

func xopDECh(cpu *CPU) {
	cpu.decP8(&cpu.HL.Hi)
}

func xopDECl(cpu *CPU) {
	cpu.decP8(&cpu.HL.Lo)
}

func xopDECa(cpu *CPU) {
	cpu.decP8(&cpu.AF.Hi)
}