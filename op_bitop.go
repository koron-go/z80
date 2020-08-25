package z80

func opBITbr(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	r := cpu.regP(codes[1])
	cpu.bitchk8(b, *r)
}

func opBITbHLP(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	p := cpu.HL.U16()
	v := cpu.Memory.Get(p)
	cpu.bitchk8(b, v)
}

func opBITbIXdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IX, codes[2])
	v := cpu.Memory.Get(p)
	cpu.bitchk8(b, v)
}

func opBITbIYdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IY, codes[2])
	v := cpu.Memory.Get(p)
	cpu.bitchk8(b, v)
}

func opSETbr(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	r := cpu.regP(codes[1])
	*r = cpu.bitset8(b, *r)
}

func opSETbHLP(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	p := cpu.HL.U16()
	v := cpu.Memory.Get(p)
	v = cpu.bitset8(b, v)
	cpu.Memory.Set(p, v)
}

func opSETbIXdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IX, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitset8(b, v)
	cpu.Memory.Set(p, v)
}

func opSETbIYdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IY, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitset8(b, v)
	cpu.Memory.Set(p, v)
}

func opRESbr(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	r := cpu.regP(codes[1])
	*r = cpu.bitres8(b, *r)
}

func opRESbHLP(cpu *CPU, codes []uint8) {
	b := (codes[1] >> 3) & 0x07
	p := cpu.HL.U16()
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}

func opRESbIXdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IX, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}

func opRESbIYdP(cpu *CPU, codes []uint8) {
	b := (codes[3] >> 3) & 0x07
	p := addrOff(cpu.IY, codes[2])
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}
