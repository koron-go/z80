package z80

import (
	"math/bits"
)

func (cpu *CPU) addU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	v := a16 + b16
	cpu.updateFlagArith8(v, a16, b16, false)
	return uint8(v)
}

func (cpu *CPU) adcU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	v := a16 + b16
	if cpu.flag(C) {
		v++
	}
	cpu.updateFlagArith8(v, a16, b16, false)
	return uint8(v)
}

func (cpu *CPU) subU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	v := a16 - b16
	cpu.updateFlagArith8(v, a16, b16, true)
	return uint8(v)
}

func (cpu *CPU) sbcU8(a, b uint8) uint8 {
	a16, b16 := uint16(a), uint16(b)
	v := a16 - b16
	if cpu.flag(C) {
		v--
	}
	cpu.updateFlagArith8(v, a16, b16, true)
	return uint8(v)
}

func (cpu *CPU) andU8(a, b uint8) uint8 {
	v := a & b
	cpu.updateFlagLogic8(v, true)
	return uint8(v)
}

func (cpu *CPU) orU8(a, b uint8) uint8 {
	v := a | b
	cpu.updateFlagLogic8(v, false)
	return uint8(v)
}

func (cpu *CPU) xorU8(a, b uint8) uint8 {
	v := a ^ b
	cpu.updateFlagLogic8(v, false)
	return uint8(v)
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
	v := uint32(a) + uint32(b)
	cpu.flagUpdate(FlagOp{}.
		Put(H, a&0x0fff+b&0x0fff > 0x0fff).
		Reset(N).
		Put(C, v > 0xffff))
	return uint16(v)
}

func (cpu *CPU) adcU16(a, b uint16) uint16 {
	var c uint16
	if cpu.flag(C) {
		c = 1
	}
	v := uint32(a) + uint32(b) + uint32(c)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x8000 != 0).
		Put(Z, v&0xffff == 0).
		Put(H, a&0x0fff+b&0x0fff+c > 0x0fff).
		Put(PV, a&0x8000 == b&0x8000 && a&0x8000 != uint16(v&0x8000)).
		Reset(N).
		Put(C, v > 0xffff))
	return uint16(v)
}

func (cpu *CPU) sbcU16(a, b uint16) uint16 {
	a32 := uint32(a)
	b32 := uint32(b)
	var c32 uint32
	if cpu.flag(C) {
		c32 = 1
	}
	v := a32 - b32 - c32
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x8000 != 0).
		Put(Z, v&0xffff == 0).
		Put(H, a32&0x0fff < b32&0x0fff+c32).
		Put(PV, a&0x8000 != b&0x8000 && a&0x8000 != uint16(v&0x8000)).
		Set(N).
		Put(C, v > 0xffff))
	return uint16(v)
}

func (cpu *CPU) incU16(a uint16) uint16 {
	v := a + 1
	return v
}

func (cpu *CPU) decU16(a uint16) uint16 {
	v := a - 1
	return v
}

func (cpu *CPU) rlcU8(a uint8) uint8 {
	a2 := a<<1 | a>>7
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x80 != 0))
	return a2
}

func (cpu *CPU) rlU8(a uint8) uint8 {
	a2 := a << 1
	if cpu.flag(C) {
		a2 |= 0x01
	}
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x80 != 0))
	return a2
}

func (cpu *CPU) rrcU8(a uint8) uint8 {
	a2 := a>>1 | a<<7
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}

func (cpu *CPU) rrU8(a uint8) uint8 {
	a2 := a >> 1
	if cpu.flag(C) {
		a2 |= 0x80
	}
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}

func (cpu *CPU) slaU8(a uint8) uint8 {
	a2 := a << 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x80 != 0))
	return a2
}

func (cpu *CPU) sl1U8(a uint8) uint8 {
	a2 := a<<1 + 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x80 != 0))
	return a2
}

func (cpu *CPU) sraU8(a uint8) uint8 {
	a2 := a&0x80 | a>>1
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}

func (cpu *CPU) srlU8(a uint8) uint8 {
	a2 := a >> 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, a2&0x80 != 0).
		Put(Z, a2 == 0).
		Reset(H).
		Put(PV, bits.OnesCount8(a2)%2 == 0).
		Reset(N).
		Put(C, a&0x01 != 0))
	return a2
}

func (cpu *CPU) bitchk8(b, v uint8) {
	f := v&(0x01<<b) != 0
	cpu.flagUpdate(FlagOp{}.
		Put(Z, !f).
		Set(H).
		Reset(N))
}

func (cpu *CPU) bitset8(b, v uint8) uint8 {
	return v | 0x01<<b
}

func (cpu *CPU) bitres8(b, v uint8) uint8 {
	return v &^ (0x01 << b)
}
