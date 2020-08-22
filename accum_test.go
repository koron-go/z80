package z80

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAccum_bitchk8(t *testing.T) {
	t.Parallel()
	for b := 0; b <= 7; b++ {
		for v := 0; v <= 255; v++ {
			cpu := &CPU{}
			cpu.bitchk8(uint8(b), uint8(v))
			wantZ := v&(1<<b) == 0
			if gotZ := cpu.flag(Z); gotZ != wantZ {
				t.Fatalf("flag Z mismatch: want=%t got=%t b=%d v=%02x", wantZ, gotZ, b, v)
			}
		}
	}
}

var u16casesSummary = []uint16{
	0x0000, 0x0001, 0x0002, 0x0ffe, 0x0fff, 0x1000, 0x1001, 0x1002, 0x7ffe,
	0x7fff, 0x8000, 0x8001, 0x8002, 0xfffe, 0xffff,
}

func tADCu16(t *testing.T, a, b uint16, c bool) {
	var c32 uint32
	var pre States

	if c {
		c32 = 1
		pre.GPR.AF.Lo |= 0x01
	}

	a32, b32 := uint32(a), uint32(b)
	sum := a32 + b32 + c32
	var flags uint8
	if sum&0x8000 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xffff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if a32&0xfff+b32&0xfff+c32 > 0xfff {
		flags |= 0x10 // H: is set if carry from bit 3
	}
	if v := int32(int16(a32)) + int32(int16(b32)) + int32(int16(c32)); v > 32767 || v < -32768 {
		flags |= 0x04 // PV is set if overflow (-32768~+32767)
	}
	// N is reset
	if sum > 0xffff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	mem := MapMemory{}
	cpu := CPU{States: pre, Memory: mem, IO: &tForbiddenIO{}}
	act := cpu.adcU16(a, b)
	if act != uint16(sum) {
		t.Fatalf("CPU.adcU16(%04x, %04x) failed: want=%04x got=%04x", a, b, uint16(sum), act)
	}

	post := States{GPR: GPR{AF: Register{Lo: flags}}}
	actS := maskFlags(cpu.States, ^uint8(maskDefault))
	expS := maskFlags(post, ^uint8(maskDefault))
	if actS != expS {
		diff := cmp.Diff(expS, actS)
		t.Fatalf("CPU.adcU16(%04x, %04x): unexpected states: -want +got\n%s", a, b, diff)
	}
}

func TestAccum_adcU16(t *testing.T) {
	t.Parallel()
	for _, a := range u16casesSummary {
		for _, b := range u16casesSummary {
			tADCu16(t, a, b, false)
			tADCu16(t, a, b, true)
		}
	}
	for _, a := range u16casesSummary {
		for b := 0; b <= 0xffff; b++ {
			tADCu16(t, a, uint16(b), false)
			tADCu16(t, a, uint16(b), true)
			tADCu16(t, uint16(b), a, false)
			tADCu16(t, uint16(b), a, true)
		}
	}
}

func tSBCu16(t *testing.T, a, b uint16, c bool) {
	var c32 uint32
	var pre States

	if c {
		c32 = 1
		pre.GPR.AF.Lo |= 0x01
	}

	a32, b32 := uint32(a), uint32(b)
	sum := a32 - b32 - c32
	var flags uint8
	if sum&0x8000 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xffff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if a32&0xfff < (b32+c32)&0xfff {
		flags |= 0x10 // H: is set if carry from bit 3
	}
	if v := int32(int16(a32)) - int32(int16(b32)) - int32(int16(c32)); v > 32767 || v < -32768 {
		flags |= 0x04 // PV is set if overflow (-32768~+32767)
	}
	flags |= 0x02 // N is set
	if sum > 0xffff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	mem := MapMemory{}
	cpu := CPU{States: pre, Memory: mem, IO: &tForbiddenIO{}}
	act := cpu.sbcU16(a, b)
	if act != uint16(sum) {
		t.Fatalf("CPU.sbcU16(%04x, %04x) failed: want=%04x got=%04x", a, b, uint16(sum), act)
	}

	post := States{GPR: GPR{AF: Register{Lo: flags}}}
	actS := maskFlags(cpu.States, ^uint8(maskDefault))
	expS := maskFlags(post, ^uint8(maskDefault))
	if actS != expS {
		diff := cmp.Diff(expS, actS)
		t.Fatalf("CPU.sbcU16(%04x, %04x): unexpected states: -want +got\n%s", a, b, diff)
	}
}

func TestAccum_sbcU16(t *testing.T) {
	t.Parallel()
	for _, a := range u16casesSummary {
		for _, b := range u16casesSummary {
			tSBCu16(t, a, b, false)
			tSBCu16(t, a, b, true)
		}
	}
	for _, a := range u16casesSummary {
		for b := 0; b <= 0xffff; b++ {
			tSBCu16(t, a, uint16(b), false)
			tSBCu16(t, a, uint16(b), true)
			tSBCu16(t, uint16(b), a, false)
			tSBCu16(t, uint16(b), a, true)
		}
	}
}
