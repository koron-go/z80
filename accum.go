package z80

import "math/bits"

func (cpu *CPU) addU8(a, b uint8) uint8 {
	v := uint16(a) + uint16(b)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		Put(H, a&0x0f+b&0x0f > 0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Reset(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) adcU8(a, b uint8) uint8 {
	a16 := uint16(a)
	x16 := uint16(b)
	if cpu.flag(C) {
		x16++
	}
	v := a16 + x16
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		Put(H, a&0x0f+b&0x0f > 0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Reset(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) subU8(a, b uint8) uint8 {
	v := uint16(a) - uint16(b)
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		Put(H, a&0x0f < b&0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Set(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) sbcU8(a, b uint8) uint8 {
	a16 := uint16(a)
	x16 := uint16(b)
	if cpu.flag(C) {
		x16++
	}
	v := a16 - x16
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v&0xff == 0).
		Put(H, a&0x0f < b&0x0f).
		// TODO: verify PV behavior.
		Put(PV, a&0x80 == b&0x80 && a&0x80 != uint8(v&0x80)).
		Set(N).
		Put(C, v > 0xff))
	return uint8(v)
}

func (cpu *CPU) andU8(a, b uint8) uint8 {
	v := a & b
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Set(H).
		// TODO: verify PV behavior.
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N).
		Reset(C))
	return uint8(v)
}

func (cpu *CPU) orU8(a, b uint8) uint8 {
	v := a | b
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Reset(H).
		// TODO: verify PV behavior.
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N).
		Reset(C))
	return uint8(v)
}

func (cpu *CPU) xorU8(a, b uint8) uint8 {
	v := a ^ b
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Reset(H).
		// TODO: verify PV behavior.
		Put(PV, bits.OnesCount8(v)%2 == 0).
		Reset(N).
		Reset(C))
	return uint8(v)
}

func (cpu *CPU) incU8(a uint8) uint8 {
	v := a + 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f+1 > 0x0f).
		Put(PV, a == 0x7f).
		Reset(N))
	return v
}

func (cpu *CPU) decU8(a uint8) uint8 {
	v := a - 1
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f < 1).
		Put(PV, a == 0x80).
		Reset(N))
	return v
}
