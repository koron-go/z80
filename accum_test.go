package z80

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	// C is an index for carry flag.
	C = 0

	// N is an index for add/subtract flag.
	N = 1

	// PV is an index for parity/overflow flag.
	PV = 2

	// X3 is reserved index for unused flag.
	//X3 = 3

	// H is an index for half carry flag.
	H = 4

	// X5 is reserved index for unused flag.
	//X5 = 5

	// Z is an index for zero flag.
	Z = 6

	// S is an index for sign flag.
	S = 7
)

// flagOp provides flag operation.  At initial this will keep all bits.
type flagOp struct {
	Nand uint8
	Or   uint8
}

// ApplyOn applies flag operation on uint8 in place.
func (fo flagOp) ApplyOn(v *uint8) {
	*v = *v&^fo.Nand | fo.Or
}

// Keep marks bit-n as keeping.
func (fo flagOp) Keep(n int) flagOp {
	var m uint8 = ^(0x01 << n)
	fo.Nand &= m
	fo.Or &= m
	return fo
}

// Set marks bit-n as being 1.
func (fo flagOp) Set(n int) flagOp {
	var b uint8 = 0x01 << n
	fo.Nand |= b
	fo.Or |= b
	return fo
}

// Reset marks bit-n as being 0.
func (fo flagOp) Reset(n int) flagOp {
	var b uint8 = 0x01 << n
	fo.Nand |= b
	fo.Or &= ^b
	return fo
}

// Put modify bit-n with boolean value.
func (fo flagOp) Put(n int, v bool) flagOp {
	if v {
		return fo.Set(n)
	}
	return fo.Reset(n)
}

func TestFlagOp_ApplyOn(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		op  flagOp
		v   []uint8
		exp []uint8
	}{
		{flagOp{0x00, 0x00}, []uint8{0x01, 0x00}, []uint8{0x01, 0x00}},
		{flagOp{0x01, 0x00}, []uint8{0x01, 0x00}, []uint8{0x00, 0x00}},
		{flagOp{0x01, 0x01}, []uint8{0x01, 0x00}, []uint8{0x01, 0x01}},
		// undefined behavior
		{flagOp{0x00, 0x01}, []uint8{0x01, 0x00}, []uint8{0x01, 0x01}},

		{flagOp{0x00, 0x00}, []uint8{0x10, 0x00}, []uint8{0x10, 0x00}},
		{flagOp{0x10, 0x00}, []uint8{0x10, 0x00}, []uint8{0x00, 0x00}},
		{flagOp{0x10, 0x10}, []uint8{0x10, 0x00}, []uint8{0x10, 0x10}},
		// undefined behavior
		{flagOp{0x00, 0x10}, []uint8{0x10, 0x00}, []uint8{0x10, 0x10}},

		{flagOp{0x00, 0x00}, []uint8{0x11}, []uint8{0x11}},
		{
			flagOp{0x11, 0x00},
			[]uint8{0x11, 0x10, 0x01, 0x00},
			[]uint8{0x00, 0x00, 0x00, 0x00},
		},
		{
			flagOp{0x11, 0x11},
			[]uint8{0x11, 0x10, 0x01, 0x00},
			[]uint8{0x11, 0x11, 0x11, 0x11},
		},
		// undefined behavior
		{
			flagOp{0x00, 0x11},
			[]uint8{0x11, 0x10, 0x01, 0x00},
			[]uint8{0x11, 0x11, 0x11, 0x11},
		},
	} {
		for i, v := range c.v {
			c.op.ApplyOn(&v)
			if v != c.exp[i] {
				t.Fatalf("failed %#v.ApplyOn(0x%02x): expect=0x%02x actual=0x%02x", c.op, c.v[i], c.exp[i], v)
			}
		}
	}
}

func carry2n(c bool) int {
	if c {
		return 1
	}
	return 0
}

func tCheckFlags(t *testing.T, name string, cpu *CPU, fo flagOp) {
	t.Helper()
	exp := fo.Or & ^uint8(maskDefault)
	act := cpu.States.GPR.AF.Lo & ^uint8(maskDefault)
	if act != exp {
		t.Fatalf("%s: unexpected flags: want=%02x got=%02x", name, exp, act)
	}
}

func tCheckFlags2(t *testing.T, name func() string, cpu *CPU, fo flagOp) {
	t.Helper()
	exp := fo.Or & ^uint8(maskDefault)
	act := cpu.States.GPR.AF.Lo & ^uint8(maskDefault)
	if act != exp {
		t.Fatalf("%s: unexpected flags: want=%02x got=%02x", name(), exp, act)
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
			fo := flagOp{}.
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
				fo := flagOp{}.
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
			fo := flagOp{}.
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
				fo := flagOp{}.
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
			fo := flagOp{}.
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
			fo := flagOp{}.
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
			fo := flagOp{}.
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
			if gotZ := cpu.flagZ(); gotZ != wantZ {
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

func TestAccum_addU16(t *testing.T) {
	t.Parallel()
	for _, ua := range u16casesSummary {
		a := int(ua)
		for b := 0; b <= 0xffff; b++ {
			name := func() string {
				return fmt.Sprintf("addU16(%04x, %04x)", a, b)
			}
			r := a + b
			hc := a&0x0fff+b&0x0fff > 0x0fff
			fo := flagOp{}.
				Put(H, hc).        // H is set if carry for bit11
				Reset(N).          // N is reset
				Put(C, r > 0xffff) // C is set if carry from bit 15
			cpu := &CPU{}
			act := cpu.addU16(uint16(a), uint16(b))
			if act != uint16(r) {
				t.Fatalf("%s: failed: want=%04x got=%04x", name(), uint16(r), act)
			}
			tCheckFlags2(t, name, cpu, fo)
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

func Benchmark_addU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.addU8(uint8(i>>8), uint8(i))
	}
}

func Benchmark_adcU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Lo = 0x00
		cpu.adcU8(uint8(i>>8), uint8(i))
	}
}

func Benchmark_subU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.subU8(uint8(i>>8), uint8(i))
	}
}

func Benchmark_sbcU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Lo = 0x01
		cpu.sbcU8(uint8(i>>8), uint8(i))
	}
}

func Benchmark_andU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.andU8(uint8(i>>8), uint8(i))
	}
}

func Benchmark_orU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.orU8(uint8(i>>8), uint8(i))
	}
}

func Benchmark_xorU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.xorU8(uint8(i>>8), uint8(i))
	}
}

func Benchmark_incU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.incU8(uint8(i))
	}
}

func Benchmark_decU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.decU8(uint8(i))
	}
}

func Benchmark_addU16(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.addU16(uint16(i>>16), uint16(i))
	}
}

func Benchmark_adcU16(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Lo = 0x01
		cpu.adcU16(uint16(i>>16), uint16(i))
	}
}

func Benchmark_sbcU16(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Lo = 0x01
		cpu.sbcU16(uint16(i>>16), uint16(i))
	}
}

func Benchmark_rlcU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.rlcU8(uint8(i))
	}
}

func Benchmark_rlU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Lo = uint8(i >> 8 & 0x01)
		cpu.rlU8(uint8(i))
	}
}

func Benchmark_rrcU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.rrcU8(uint8(i))
	}
}

func Benchmark_rrU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Lo = uint8(i >> 8 & 0x01)
		cpu.rrU8(uint8(i))
	}
}

func Benchmark_slaU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.slaU8(uint8(i))
	}
}

func Benchmark_sl1U8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.sl1U8(uint8(i))
	}
}

func Benchmark_sraU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.sraU8(uint8(i))
	}
}

func Benchmark_srlU8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.srlU8(uint8(i))
	}
}

func Benchmark_bitchk8(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.bitchk8(uint8(i>>8&0x7), uint8(i))
	}
}
