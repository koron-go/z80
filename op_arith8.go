package z80

func oopADDAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.addU8(a, n)
}

func oopADDAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.addU8(a, x)
}

func oopADDAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func oopADDAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func oopADCAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.adcU8(a, n)
}

func oopADCAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func oopADCAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func oopADCAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func oopSUBAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.subU8(a, n)
}

func oopSUBAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.subU8(a, x)
}

func oopSUBAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func oopSUBAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func oopSBCAn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.sbcU8(a, n)
}

func oopSBCAHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func oopSBCAIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func oopSBCAIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func oopANDn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.andU8(a, n)
}

func oopANDHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.andU8(a, x)
}

func oopANDIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func oopANDIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func oopORn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.orU8(a, n)
}

func oopORHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.orU8(a, x)
}

func oopORIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func oopORIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func oopXORn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.AF.Hi = cpu.xorU8(a, n)
}

func oopXORHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func oopXORIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func oopXORIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func oopCPn(cpu *CPU, n uint8) {
	a := cpu.AF.Hi
	cpu.subU8(a, n)
}

func oopCPHLP(cpu *CPU) {
	a := cpu.AF.Hi
	x := cpu.Memory.Get(cpu.HL.U16())
	cpu.subU8(a, x)
}

func oopCPIXdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.subU8(a, x)
}

func oopCPIYdP(cpu *CPU, d uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.subU8(a, x)
}

func oopINCHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func oopINCIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func oopINCIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func oopDECHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}

func oopDECIXdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IX, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}

func oopDECIYdP(cpu *CPU, d uint8) {
	p := addrOff(cpu.IY, d)
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}

func oopINCIXH(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
}

func oopDECIXH(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IX >> 8))
	cpu.IX = uint16(v)<<8 | cpu.IX&0xff
}

func oopINCIXL(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
}

func oopDECIXL(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IX))
	cpu.IX = uint16(v) | cpu.IX&0xff00
}

func oopINCIYH(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
}

func oopDECIYH(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IY >> 8))
	cpu.IY = uint16(v)<<8 | cpu.IY&0xff
}

func oopINCIYL(cpu *CPU) {
	v := cpu.incU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
}

func oopDECIYL(cpu *CPU) {
	v := cpu.decU8(uint8(cpu.IY))
	cpu.IY = uint16(v) | cpu.IY&0xff00
}

//////////////////////////////////////////////////////////////////////////////
// eXpanded OPration codes

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

func xopADDAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func xopADDAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
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

func xopADCAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func xopADCAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
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

func xopSUBAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func xopSUBAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
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

func xopSBCAixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func xopSBCAiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
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

func xopANDixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func xopANDixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func xopANDiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func xopANDiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.andU8(a, x)
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

func xopXORixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func xopXORixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func xopXORiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func xopXORiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.xorU8(a, x)
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

func xopORixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func xopORixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func xopORiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func xopORiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.AF.Hi = cpu.orU8(a, x)
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

func xopCPixh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX >> 8)
	cpu.subU8(a, x)
}

func xopCPixl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IX)
	cpu.subU8(a, x)
}

func xopCPiyh(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY >> 8)
	cpu.subU8(a, x)
}

func xopCPiyl(cpu *CPU) {
	a := cpu.AF.Hi
	x := uint8(cpu.IY)
	cpu.subU8(a, x)
}

func xopINCb(cpu *CPU) {
	cpu.BC.Hi = cpu.incU8(cpu.BC.Hi)
}

func xopINCc(cpu *CPU) {
	cpu.BC.Lo = cpu.incU8(cpu.BC.Lo)
}

func xopINCd(cpu *CPU) {
	cpu.DE.Hi = cpu.incU8(cpu.DE.Hi)
}

func xopINCe(cpu *CPU) {
	cpu.DE.Lo = cpu.incU8(cpu.DE.Lo)
}

func xopINCh(cpu *CPU) {
	cpu.HL.Hi = cpu.incU8(cpu.HL.Hi)
}

func xopINCl(cpu *CPU) {
	cpu.HL.Lo = cpu.incU8(cpu.HL.Lo)
}

func xopINCa(cpu *CPU) {
	cpu.AF.Hi = cpu.incU8(cpu.AF.Hi)
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
