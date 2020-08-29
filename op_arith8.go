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
