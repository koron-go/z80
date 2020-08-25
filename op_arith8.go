package z80

func opADDAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.addU8(a, x)
}

func opADDAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.addU8(a, n)
}

func opADDAHLP(cpu *CPU, codes []uint8) {
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

func opADCAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.adcU8(a, x)
}

func opADCAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.adcU8(a, n)
}

func opADCAHLP(cpu *CPU, codes []uint8) {
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

func opSUBAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.subU8(a, x)
}

func opSUBAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.subU8(a, n)
}

func opSUBAHLP(cpu *CPU, codes []uint8) {
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

func opSBCAr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.sbcU8(a, x)
}

func opSBCAn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.sbcU8(a, n)
}

func opSBCAHLP(cpu *CPU, codes []uint8) {
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

func opANDr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.andU8(a, x)
}

func opANDn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.andU8(a, n)
}

func opANDHLP(cpu *CPU, codes []uint8) {
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

func opORr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.orU8(a, x)
}

func opORn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.orU8(a, n)
}

func opORHLP(cpu *CPU, codes []uint8) {
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

func opXORr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.AF.Hi = cpu.xorU8(a, x)
}

func opXORn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.AF.Hi = cpu.xorU8(a, n)
}

func opXORHLP(cpu *CPU, codes []uint8) {
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

func opCPr(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	x := *cpu.regP(codes[0])
	cpu.subU8(a, x)
}

func opCPn(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	n := codes[1]
	cpu.subU8(a, n)
}

func opCPHLP(cpu *CPU, codes []uint8) {
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

func opINCr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[0] >> 3)
	*r = cpu.incU8(*r)
}

func opINCHLP(cpu *CPU, codes []uint8) {
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

func opDECr(cpu *CPU, codes []uint8) {
	r := cpu.regP(codes[0] >> 3)
	*r = cpu.decU8(*r)
}

func opDECHLP(cpu *CPU, codes []uint8) {
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
