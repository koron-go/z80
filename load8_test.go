package z80

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func rGet(t *testing.T, gpr *GPR, n int) *uint8 {
	t.Helper()
	switch n {
	case 0:
		return &gpr.BC.Hi
	case 1:
		return &gpr.BC.Lo
	case 2:
		return &gpr.DE.Hi
	case 3:
		return &gpr.DE.Lo
	case 4:
		return &gpr.HL.Hi
	case 5:
		return &gpr.HL.Lo
	case 7:
		return &gpr.AF.Hi
	default:
		t.Fatalf("invalid r index: %d", n)
		return nil
	}
}

var rLabel = []string{"B", "C", "D", "E", "H", "L", "(N/A)", "A"}

var testInitGPR = GPR{
	AF: Register{Hi: 0x12},
	BC: Register{Hi: 0x34, Lo: 0x56},
	DE: Register{Hi: 0x78, Lo: 0x9A},
	HL: Register{Hi: 0xBC, Lo: 0xDE},
}

func TestLoad8_LDr1r2(t *testing.T) {
	t.Parallel()
	for r1 := 0; r1 <= 7; r1++ {
		if r1 == 6 {
			continue
		}
		for r2 := 0; r2 <= 7; r2++ {
			if r2 == 6 {
				continue
			}
			n := fmt.Sprintf("LD %s, %s", rLabel[r1], rLabel[r2])
			c := uint8(0x40 | r1<<3 | r2)
			beforeGPR := testInitGPR
			afterGPR := testInitGPR
			*rGet(t, &afterGPR, r1) = *rGet(t, &beforeGPR, r2)
			t.Run(n, func(t *testing.T) {
				testStep(t,
					&testStates{
						States{GPR: beforeGPR},
						DumbMemory{c},
						DumbIO{}},
					&testStates{
						States{GPR: afterGPR,
							SPR: SPR{PC: 0x0001, IR: Register{Lo: 0x01}}},
						DumbMemory{c},
						DumbIO{}})
			})
		}
	}
}

func TestLoad8_LDrn(t *testing.T) {
	t.Parallel()
	for r := 0; r <= 7; r++ {
		if r == 6 {
			continue
		}
		n := fmt.Sprintf("LD %s, n", rLabel[r])
		c := uint8(0x06 | r<<3)
		beforeGPR := testInitGPR
		var afterGPR GPR
		p := rGet(t, &afterGPR, r)
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			for n := 0; n <= 0xff; n++ {
				afterGPR = testInitGPR
				*p = uint8(n)
				testStep(t,
					&testStates{
						States{GPR: beforeGPR},
						DumbMemory{c, uint8(n)},
						DumbIO{}},
					&testStates{
						States{GPR: afterGPR,
							SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}}},
						DumbMemory{c, uint8(n)},
						DumbIO{}})
			}
		})
	}
}

func TestLoad8_LDrHL(t *testing.T) {
	t.Parallel()
	for r := 0; r <= 7; r++ {
		if r == 6 {
			continue
		}
		n := fmt.Sprintf("LD %s, (HL)", rLabel[r])
		c := uint8(0x46 | r<<3)
		var beforeGPR, afterGPR GPR
		p := rGet(t, &afterGPR, r)
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			rnd := rand.New(rand.NewSource(time.Now().UnixNano() * int64(r)))
			for hl := 0; hl <= 0xffff; hl++ {
				memory := MapMemory{}.Put(0, c)
				if hl != 0 {
					d := uint8(rnd.Intn(255)+1)
					memory.Put(uint16(hl), d)
				}
				wantMem := memory.Clone()
				//wantMem := memory.Clone()
				beforeGPR = testInitGPR
				beforeGPR.HL.SetU16(uint16(hl))
				afterGPR = beforeGPR
				*p = memory.Get(uint16(hl))
				testStepNoIO(t,
					States{GPR: beforeGPR}, memory,
					States{GPR: afterGPR,
						SPR: SPR{PC: 0x0001, IR: Register{Lo: 0x01}}},
					wantMem)
			}
		})
	}
}
