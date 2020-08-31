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

func xopBITbHLP(cpu *CPU, b uint8) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.bitres8(b, x)
	cpu.Memory.Set(p, x)
}

func xopSETbHLP(cpu *CPU, b uint8) {
	p := cpu.HL.U16()
	x := cpu.Memory.Get(p)
	x = cpu.bitset8(b, x)
	cpu.Memory.Set(p, x)
}
