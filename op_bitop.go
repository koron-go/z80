package z80

func xopBITbIXdP(cpu *CPU, b, d uint8) {
	p := addrOff(cpu.IX, d)
	v := cpu.Memory.Get(p)
	cpu.bitchk8(b, v)
}

func xopBITbIYdP(cpu *CPU, b, d uint8) {
	p := addrOff(cpu.IY, d)
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

func xopSETbIXdP(cpu *CPU, b, d uint8) {
	p := addrOff(cpu.IX, d)
	v := cpu.Memory.Get(p)
	v = cpu.bitset8(b, v)
	cpu.Memory.Set(p, v)
}

func xopSETbIYdP(cpu *CPU, b, d uint8) {
	p := addrOff(cpu.IY, d)
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

func xopRESbIXdP(cpu *CPU, b, d uint8) {
	p := addrOff(cpu.IX, d)
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}

func xopRESbIYdP(cpu *CPU, b, d uint8) {
	p := addrOff(cpu.IY, d)
	v := cpu.Memory.Get(p)
	v = cpu.bitres8(b, v)
	cpu.Memory.Set(p, v)
}
