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

func opADDAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	//fmt.Printf("a=%02x p=%04x x=%02x ix=%04x\n", a, p, x, cpu.IX)
	cpu.AF.Hi = cpu.addU8(a, x)
}

func opADDAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
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

func opADCAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func opADCAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
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

func opSUBAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.subU8(a, x)
}

func opSUBAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
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

func opSBCAIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func opSBCAIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
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

func opANDIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.andU8(a, x)
}

func opANDIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
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

func opORIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.orU8(a, x)
}

func opORIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
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

func opXORIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func opXORIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
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

func opCPIXdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.subU8(a, x)
}

func opCPIYdP(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.subU8(a, x)
}

func oopINCHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func opINCIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func opINCIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.incU8(x))
}

func oopDECHLP(cpu *CPU) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}

func opDECIXdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IX, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}

func opDECIYdP(cpu *CPU, codes []uint8) {
	p := addrOff(cpu.IY, codes[2])
	x := cpu.Memory.Get(p)
	cpu.Memory.Set(p, cpu.decU8(x))
}
