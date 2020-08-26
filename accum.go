package z80

func (cpu *CPU) addU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	r := a16 + b16
	cpu.updateFlagArith8(r, a16, b16, false)
	return uint8(r)
}

func (cpu *CPU) adcU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	r := a16 + b16
	if cpu.flag(C) {
		r++
	}
	cpu.updateFlagArith8(r, a16, b16, false)
	return uint8(r)
}

func (cpu *CPU) subU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	r := a16 - b16
	cpu.updateFlagArith8(r, a16, b16, true)
	return uint8(r)
}

func (cpu *CPU) sbcU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	r := a16 - b16
	if cpu.flag(C) {
		r--
	}
	cpu.updateFlagArith8(r, a16, b16, true)
	return uint8(r)
}

func (cpu *CPU) andU8(a, b uint8) uint8 {
	r := a & b
	cpu.updateFlagLogic8(r, true)
	return uint8(r)
}

func (cpu *CPU) orU8(a, b uint8) uint8 {
	r := a | b
	cpu.updateFlagLogic8(r, false)
	return uint8(r)
}

func (cpu *CPU) xorU8(a, b uint8) uint8 {
	r := a ^ b
	cpu.updateFlagLogic8(r, false)
	return uint8(r)
}

func (cpu *CPU) incU8(a uint8) uint8 {
	r := a + 1
	var nand uint8 = maskStd | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= r & maskStd
	if r == 0 {
		or |= maskZ
	}
	if r&0x0f == 0 {
		or |= maskH
	}
	if a == 0x7f {
		or |= maskPV
	}
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return r
}

func (cpu *CPU) decU8(a uint8) uint8 {
	r := a - 1
	var nand uint8 = maskStd | maskZ | maskH | maskPV | maskN
	var or uint8
	or |= r & maskStd
	if r == 0 {
		or |= maskZ
	}
	if r&0x0f == 0x0f {
		or |= maskH
	}
	if a == 0x80 {
		or |= maskPV
	}
	or |= maskN
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return r
}

func (cpu *CPU) addU16(a, b uint16) uint16 {
	a32, b32 := uint32(a), uint32(b)
	r := a32 + b32
	c := r ^ a32 ^ b32
	var nand uint8 = maskH | maskN | maskC
	var or uint8
	or |= uint8(c>>8) & maskH
	or |= uint8(r>>16) & maskC
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return uint16(r)
}

func (cpu *CPU) adcU16(a, b uint16) uint16 {
	a32, b32 := uint32(a), uint32(b)
	r := a32 + b32
	if cpu.flag(C) {
		r++
	}
	c := r ^ a32 ^ b32
	var nand uint8 = maskStd | maskZ | maskH | maskPV | maskN | maskC
	var or uint8
	or |= uint8(r>>8) & maskStd
	if uint16(r) == 0 {
		or |= maskZ
	}
	or |= uint8(c>>8) & maskH
	or |= uint8((c>>14)^(c>>13)) & maskPV
	or |= uint8(r>>16) & maskC
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return uint16(r)
}

func (cpu *CPU) sbcU16(a, b uint16) uint16 {
	a32, b32 := uint32(a), uint32(b)
	r := a32 - b32
	if cpu.flag(C) {
		r--
	}
	c := r ^ a32 ^ b32
	var nand uint8 = maskStd | maskZ | maskH | maskPV | maskN | maskC
	var or uint8
	or |= uint8(r>>8) & maskStd
	if uint16(r) == 0 {
		or |= maskZ
	}
	or |= uint8(c>>8) & maskH
	or |= uint8((c>>14)^(c>>13)) & maskPV
	or |= maskN
	or |= uint8(r>>16) & maskC
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
	return uint16(r)
}

func (cpu *CPU) incU16(a uint16) uint16 {
	r := a + 1
	return r
}

func (cpu *CPU) decU16(a uint16) uint16 {
	r := a - 1
	return r
}

func (cpu *CPU) rlcU8(a uint8) uint8 {
	r := a<<1 | a>>7
	cpu.updateFlagBitop(r, a>>7)
	return r
}

func (cpu *CPU) rlU8(a uint8) uint8 {
	r := a << 1
	if cpu.flag(C) {
		r |= 0x01
	}
	cpu.updateFlagBitop(r, a>>7)
	return r
}

func (cpu *CPU) rrcU8(a uint8) uint8 {
	r := a>>1 | a<<7
	cpu.updateFlagBitop(r, a)
	return r
}

func (cpu *CPU) rrU8(a uint8) uint8 {
	r := a >> 1
	if cpu.flag(C) {
		r |= 0x80
	}
	cpu.updateFlagBitop(r, a)
	return r
}

func (cpu *CPU) slaU8(a uint8) uint8 {
	r := a << 1
	cpu.updateFlagBitop(r, a>>7)
	return r
}

func (cpu *CPU) sl1U8(a uint8) uint8 {
	r := a<<1 + 1
	cpu.updateFlagBitop(r, a>>7)
	return r
}

func (cpu *CPU) sraU8(a uint8) uint8 {
	r := a&0x80 | a>>1
	cpu.updateFlagBitop(r, a)
	return r
}

func (cpu *CPU) srlU8(a uint8) uint8 {
	r := a >> 1
	cpu.updateFlagBitop(r, a)
	return r
}

func (cpu *CPU) bitchk8(b, v uint8) {
	r := v&(0x01<<b) != 0
	var nand uint8 =  maskZ | maskH | maskN
	var or uint8
	if !r {
		or |= maskZ
	}
	or |= maskH
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func (cpu *CPU) bitset8(b, v uint8) uint8 {
	return v | 0x01<<b
}

func (cpu *CPU) bitres8(b, v uint8) uint8 {
	return v &^ (0x01 << b)
}
