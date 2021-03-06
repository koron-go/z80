package z80

import (
	"fmt"
	"math/bits"
	"testing"
)

func toBCD(n int) uint8 {
	n %= 100
	return uint8((n/10)<<4) | uint8(n%10)
}

func TestDAA_INC(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		curr := toBCD(i)
		next := toBCD(i + 1)
		// compute expected flag
		var flag uint8
		if i+1 >= 100 {
			flag |= 0x01 // C
		}
		if bits.OnesCount8(next)%2 == 0 {
			flag |= 0x04 // P/V: parity flag
		}
		if (curr^next)&0x10 != 0 {
			flag |= 0x10 // H
		}
		if next == 0 {
			flag |= 0x40 // Z
		}
		if next&0x80 != 0 {
			flag |= 0x80 // S: sign flag
		}
		// LD A, n ; INC A ; DAA
		mem := MapMemory{}.Put(0, 0x3e, curr, 0x3c, 0x27)
		wantMem := mem.Clone()
		n := fmt.Sprintf("LD A, 0x%02x ; INC A ; DAA", curr)
		tSteps(t, n, States{}, mem, 3, States{
			GPR: GPR{AF: Register{
				Hi: next,
				Lo: flag,
			}},
			SPR: SPR{PC: 0x04, IR: Register{Lo: 0x03}},
		}, wantMem, maskDefault)
	}
}

func TestDAA_ADD(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			a, b := toBCD(i), toBCD(j)
			res := toBCD(i + j)
			// compute expected flag
			var flag uint8
			if i+j >= 100 {
				flag |= 0x01 // C: carry flag
			}
			if bits.OnesCount8(res)%2 == 0 {
				flag |= 0x04 // P/V: parity flag
			}
			if (a+b)&0x0f > 9 {
				flag |= 0x10 // H
			}
			if res == 0 {
				flag |= 0x40 // Z: zero flag
			}
			if res&0x80 != 0 {
				flag |= 0x80 // S: sign flag
			}
			// LD A, n ; ADD A, n ; DAA
			mem := MapMemory{}.Put(0, 0x3e, a, 0xc6, b, 0x27)
			wantMem := mem.Clone()
			n := fmt.Sprintf("LD A, 0x%02x ; ADD A, 0x%02x ; DAA", a, b)
			tSteps(t, n, States{}, mem, 3, States{
				GPR: GPR{AF: Register{
					Hi: res,
					Lo: flag,
				}},
				SPR: SPR{PC: 0x05, IR: Register{Lo: 0x03}},
			}, wantMem, maskDefault)
		}
	}
}

func TestDAA_ADC_0(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			a, b := toBCD(i), toBCD(j)
			res := toBCD(i + j)
			// compute expected flag
			var flag uint8
			if i+j >= 100 {
				flag |= 0x01 // C: carry flag
			}
			if bits.OnesCount8(res)%2 == 0 {
				flag |= 0x04 // P/V: parity flag
			}
			if (a+b)&0x0f > 9 {
				flag |= 0x10 // H
			}
			if res == 0 {
				flag |= 0x40 // Z: zero flag
			}
			if res&0x80 != 0 {
				flag |= 0x80 // S: sign flag
			}
			// LD A, n ; ADD A, n ; DAA
			mem := MapMemory{}.Put(0, 0x3e, a, 0xce, b, 0x27)
			wantMem := mem.Clone()
			n := fmt.Sprintf("LD A, 0x%02x ; ADC A, 0x%02x ; DAA", a, b)
			tSteps(t, n, States{}, mem, 3, States{
				GPR: GPR{AF: Register{
					Hi: res,
					Lo: flag,
				}},
				SPR: SPR{PC: 0x05, IR: Register{Lo: 0x03}},
			}, wantMem, maskDefault)
		}
	}
}

func TestDAA_ADC_1(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			a, b := toBCD(i), toBCD(j)
			res := toBCD(i + j + 1)
			// compute expected flag
			var flag uint8
			if i+j+1 >= 100 {
				flag |= 0x01 // C: carry flag
			}
			if bits.OnesCount8(res)%2 == 0 {
				flag |= 0x04 // P/V: parity flag
			}
			if (a+b+1)&0x0f > 9 {
				flag |= 0x10 // H
			}
			if res == 0 {
				flag |= 0x40 // Z: zero flag
			}
			if res&0x80 != 0 {
				flag |= 0x80 // S: sign flag
			}
			// LD A, n ; ADD A, n ; DAA
			mem := MapMemory{}.Put(0, 0x3e, a, 0xce, b, 0x27)
			wantMem := mem.Clone()
			n := fmt.Sprintf("LD A, 0x%02x ; ADC A, 0x%02x ; DAA", a, b)
			tSteps(t, n, States{
				GPR: GPR{AF: Register{Lo: 0x01}},
			}, mem, 3, States{
				GPR: GPR{AF: Register{
					Hi: res,
					Lo: flag,
				}},
				SPR: SPR{PC: 0x05, IR: Register{Lo: 0x03}},
			}, wantMem, maskDefault)
		}
	}
}

func TestDAA_DEC(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		curr := toBCD(i)
		next := toBCD(i + 99)
		// compute expected flag
		var flag uint8
		if next == 0x99 {
			flag |= 0x01 // C
		}
		flag |= 0x02 // N: subtract flag
		if bits.OnesCount8(next)%2 == 0 {
			flag |= 0x04 // P/V: parity flag
		}
		if next == 0 {
			flag |= 0x40 // Z
		}
		if next&0x80 != 0 {
			flag |= 0x80 // S: sign flag
		}
		// LD A, n ; DEC A ; DAA
		mem := MapMemory{}.Put(0, 0x3e, curr, 0x3d, 0x27)
		wantMem := mem.Clone()
		n := fmt.Sprintf("LD A, 0x%02x ; DEC A ; DAA", curr)
		tSteps(t, n, States{}, mem, 3, States{
			GPR: GPR{AF: Register{
				Hi: next,
				Lo: flag,
			}},
			SPR: SPR{PC: 0x04, IR: Register{Lo: 0x03}},
		}, wantMem, maskDefault|maskH)
	}
}

func TestDAA_SUB(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			a, b := toBCD(i), toBCD(j)
			res := toBCD(i - j + 100)
			// compute expected flag
			var flag uint8
			if i-j < 0 {
				flag |= 0x01 // C: carry flag
			}
			flag |= 0x02 // N: subtract flag
			if bits.OnesCount8(res)%2 == 0 {
				flag |= 0x04 // P/V: parity flag
			}
			if res == 0 {
				flag |= 0x40 // Z: zero flag
			}
			if res&0x80 != 0 {
				flag |= 0x80 // S: sign flag
			}
			// LD A, n ; ADD A, n ; DAA
			mem := MapMemory{}.Put(0, 0x3e, a, 0xd6, b, 0x27)
			wantMem := mem.Clone()
			n := fmt.Sprintf("LD A, 0x%02x ; SUB A, 0x%02x ; DAA", a, b)
			tSteps(t, n, States{}, mem, 3, States{
				GPR: GPR{AF: Register{
					Hi: res,
					Lo: flag,
				}},
				SPR: SPR{PC: 0x05, IR: Register{Lo: 0x03}},
			}, wantMem, maskDefault|maskH)
		}
	}
}

func TestDAA_SBC_0(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			a, b := toBCD(i), toBCD(j)
			res := toBCD(i - j + 100)
			// compute expected flag
			var flag uint8
			if i-j < 0 {
				flag |= 0x01 // C: carry flag
			}
			flag |= 0x02 // N: subtract flag
			if bits.OnesCount8(res)%2 == 0 {
				flag |= 0x04 // P/V: parity flag
			}
			if res == 0 {
				flag |= 0x40 // Z: zero flag
			}
			if res&0x80 != 0 {
				flag |= 0x80 // S: sign flag
			}
			// LD A, n ; ADD A, n ; DAA
			mem := MapMemory{}.Put(0, 0x3e, a, 0xde, b, 0x27)
			wantMem := mem.Clone()
			n := fmt.Sprintf("LD A, 0x%02x ; SBC A, 0x%02x ; DAA", a, b)
			tSteps(t, n, States{}, mem, 3, States{
				GPR: GPR{AF: Register{
					Hi: res,
					Lo: flag,
				}},
				SPR: SPR{PC: 0x05, IR: Register{Lo: 0x03}},
			}, wantMem, maskDefault|maskH)
		}
	}
}

func TestDAA_SBC_1(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			a, b := toBCD(i), toBCD(j)
			res := toBCD(i - j - 1 + 100)
			// compute expected flag
			var flag uint8
			if i-j-1 < 0 {
				flag |= 0x01 // C: carry flag
			}
			flag |= 0x02 // N: subtract flag
			if bits.OnesCount8(res)%2 == 0 {
				flag |= 0x04 // P/V: parity flag
			}
			if res == 0 {
				flag |= 0x40 // Z: zero flag
			}
			if res&0x80 != 0 {
				flag |= 0x80 // S: sign flag
			}
			// LD A, n ; ADD A, n ; DAA
			mem := MapMemory{}.Put(0, 0x3e, a, 0xde, b, 0x27)
			wantMem := mem.Clone()
			n := fmt.Sprintf("LD A, 0x%02x ; SBC A, 0x%02x ; DAA", a, b)
			tSteps(t, n, States{
				GPR: GPR{AF: Register{Lo: 0x01}},
			}, mem, 3, States{
				GPR: GPR{AF: Register{
					Hi: res,
					Lo: flag,
				}},
				SPR: SPR{PC: 0x05, IR: Register{Lo: 0x03}},
			}, wantMem, maskDefault|maskH)
		}
	}
}
