package z80

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func carry2n(c bool) int {
	if c {
		return 1
	}
	return 0
}

func tCheckFlags(t *testing.T, name string, cpu *CPU, fo FlagOp) {
	t.Helper()
	exp := fo.Or & ^uint8(maskDefault)
	act := cpu.States.GPR.AF.Lo & ^uint8(maskDefault)
	if act != exp {
		t.Fatalf("%s: unexpected flags: want=%02x got=%02x", name, exp, act)
	}
}

func TestAccum_addU8(t *testing.T) {
	t.Parallel()
	for a := 0; a <= 0xff; a++ {
		for b := 0; b <= 0xff; b++ {
			name := fmt.Sprintf("addU8(%02x, %02x)", a, b)
			r := a + b
			hc := a&0x0f+b&0x0f > 0x0f
			pv := isOverflowS8(int32(int8(a)) + int32(int8(b)))
			fo := FlagOp{}.
				Put(S, r&0x80 != 0). // S is set if result is negative
				Put(Z, r&0xff == 0). // Z is set if result is 0
				Put(H, hc).          // H is set if carry for bit3
				Put(PV, pv).         // PV is set if overflow (-128~+127)
				Reset(N).            // N is reset
				Put(C, r > 0xff)     // C is set if carry from bit 7
			cpu := &CPU{}
			act := cpu.addU8(uint8(a), uint8(b))
			if act != uint8(r) {
				t.Fatalf("%s: failed: want=%02x got=%02x", name, uint8(r), act)
			}
			tCheckFlags(t, name, cpu, fo)
		}
	}
}

func TestAccum_adcU8(t *testing.T) {
	t.Parallel()
	for a := 0; a <= 0xff; a++ {
		for b := 0; b <= 0xff; b++ {
			for _, carry := range []bool{false, true} {
				name := fmt.Sprintf("adcU8(%02x, %02x, C=%t)", a, b, carry)
				c := carry2n(carry)
				r := a + b + c
				hc := a&0x0f+b&0x0f+c > 0x0f
				pv := isOverflowS8(int32(int8(a)) + int32(int8(b)) + int32(c))
				fo := FlagOp{}.
					Put(S, r&0x80 != 0). // S is set if result is negative
					Put(Z, r&0xff == 0). // Z is set if result is 0
					Put(H, hc).          // H is set if carry for bit3
					Put(PV, pv).         // PV is set if overflow (-128~+127)
					Reset(N).            // N is reset
					Put(C, r > 0xff)     // C is set if carry from bit 7
				cpu := &CPU{}
				cpu.AF.Lo = uint8(c)
				act := cpu.adcU8(uint8(a), uint8(b))
				if act != uint8(r) {
					t.Fatalf("%s: failed: want=%02x got=%02x", name, uint8(r), act)
				}
				tCheckFlags(t, name, cpu, fo)
			}
		}
	}
}

func TestAccum_subU8(t *testing.T) {
	t.Parallel()
	for a := 0; a <= 0xff; a++ {
		for b := 0; b <= 0xff; b++ {
			name := fmt.Sprintf("subU8(%02x, %02x)", a, b)
			r := a - b
			hc := a&0x0f < b&0x0f
			pv := isOverflowS8(int32(int8(a)) - int32(int8(b)))
			fo := FlagOp{}.
				Put(S, r&0x80 != 0). // S is set if result is negative
				Put(Z, r&0xff == 0). // Z is set if result is 0
				Put(H, hc).          // H is set if carry for bit3
				Put(PV, pv).         // PV is set if overflow (-128~+127)
				Set(N).              // N is reset
				Put(C, r < 0)        // C is set if carry from bit 7
			cpu := &CPU{}
			act := cpu.subU8(uint8(a), uint8(b))
			if act != uint8(r) {
				t.Fatalf("%s: failed: want=%02x got=%02x", name, uint8(r), act)
			}
			tCheckFlags(t, name, cpu, fo)
		}
	}
}

func TestAccum_sbcU8(t *testing.T) {
	t.Parallel()
	for a := 0; a <= 0xff; a++ {
		for b := 0; b <= 0xff; b++ {
			for _, carry := range []bool{false, true} {
				name := fmt.Sprintf("sbcU8(%02x, %02x, C=%t)", a, b, carry)
				c := carry2n(carry)
				r := a - b - c
				hc := a&0x0f < b&0x0f+c
				pv := isOverflowS8(int32(int8(a)) - int32(int8(b)) - int32(c))
				fo := FlagOp{}.
					Put(S, r&0x80 != 0). // S is set if result is negative
					Put(Z, r&0xff == 0). // Z is set if result is 0
					Put(H, hc).          // H is set if carry for bit3
					Put(PV, pv).         // PV is set if overflow (-128~+127)
					Set(N).              // N is reset
					Put(C, r < 0)        // C is set if carry from bit 7
				cpu := &CPU{}
				cpu.AF.Lo = uint8(c)
				act := cpu.sbcU8(uint8(a), uint8(b))
				if act != uint8(r) {
					t.Fatalf("%s: failed: want=%02x got=%02x", name, uint8(r), act)
				}
				tCheckFlags(t, name, cpu, fo)
			}
		}
	}
}

func TestAccum_andU8(t *testing.T) {
	t.Parallel()
	for a := 0; a <= 0xff; a++ {
		for b := 0; b <= 0xff; b++ {
			name := fmt.Sprintf("andU8(%02x, %02x)", a, b)
			r := a & b
			pv := bits.OnesCount8(uint8(r))%2 == 0
			fo := FlagOp{}.
				Put(S, r&0x80 != 0). // S is set if result is negative
				Put(Z, r&0xff == 0). // Z is set if result is 0
				Set(H).              // H is set
				Put(PV, pv).         // PV is set if parity even
				Reset(N).            // N is reset
				Reset(C)             // C is reset
			cpu := &CPU{}
			act := cpu.andU8(uint8(a), uint8(b))
			if act != uint8(r) {
				t.Fatalf("%s: failed: want=%02x got=%02x", name, uint8(r), act)
			}
			tCheckFlags(t, name, cpu, fo)
		}
	}
}

func TestAccum_orU8(t *testing.T) {
	t.Parallel()
	for a := 0; a <= 0xff; a++ {
		for b := 0; b <= 0xff; b++ {
			name := fmt.Sprintf("orU8(%02x, %02x)", a, b)
			r := a | b
			pv := bits.OnesCount8(uint8(r))%2 == 0
			fo := FlagOp{}.
				Put(S, r&0x80 != 0). // S is set if result is negative
				Put(Z, r&0xff == 0). // Z is set if result is 0
				Reset(H).            // H is reset
				Put(PV, pv).         // PV is set if parity even
				Reset(N).            // N is reset
				Reset(C)             // C is reset
			cpu := &CPU{}
			act := cpu.orU8(uint8(a), uint8(b))
			if act != uint8(r) {
				t.Fatalf("%s: failed: want=%02x got=%02x", name, uint8(r), act)
			}
			tCheckFlags(t, name, cpu, fo)
		}
	}
}

func TestAccum_xorU8(t *testing.T) {
	t.Parallel()
	for a := 0; a <= 0xff; a++ {
		for b := 0; b <= 0xff; b++ {
			name := fmt.Sprintf("xorU8(%02x, %02x)", a, b)
			r := a ^ b
			pv := bits.OnesCount8(uint8(r))%2 == 0
			fo := FlagOp{}.
				Put(S, r&0x80 != 0). // S is set if result is negative
				Put(Z, r&0xff == 0). // Z is set if result is 0
				Reset(H).            // H is reset
				Put(PV, pv).         // PV is set if parity even
				Reset(N).            // N is reset
				Reset(C)             // C is reset
			cpu := &CPU{}
			act := cpu.xorU8(uint8(a), uint8(b))
			if act != uint8(r) {
				t.Fatalf("%s: failed: want=%02x got=%02x", name, uint8(r), act)
			}
			tCheckFlags(t, name, cpu, fo)
		}
	}
}

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
	if a32&0xfff < b32&0xfff+c32 {
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
