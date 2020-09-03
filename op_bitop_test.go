package z80

import (
	"fmt"
	"testing"
)

func TestBitop_BITbr(t *testing.T) {
	t.Parallel()
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 7; j++ {
			if j == 6 {
				continue
			}
			b, r := i, j
			n := fmt.Sprintf("BIT %d, %s", b, rLabel[r])
			t.Run(n, func(t *testing.T) {
				t.Parallel()
				var gpr, wantGPR GPR
				p := rGet(t, &gpr, r)
				for v := 0; v <= 255; v++ {
					gpr = testInitGPR
					*p = uint8(v)
					mem := MapMemory{}.Put(0, 0xcb, uint8(0x40|b<<3|r))
					wantGPR = gpr
					flagOp{}.Put(Z, v&(1<<b) == 0).Set(H).Reset(N).
						ApplyOn(&wantGPR.AF.Lo)
					tSteps(t, "",
						States{GPR: gpr, SPR: SPR{IX: 0x1000}},
						mem,
						1,
						States{
							GPR: wantGPR,
							SPR: SPR{PC: 2, IR: Register{Lo: 0x01},
								IX: 0x1000,
							}},
						mem.Clone(),
						maskS|maskPV|mask5|mask3|maskC,
					)
				}
			})
		}
	}
}

func TestBitop_BITbIXd(t *testing.T) {
	t.Parallel()
	for i := 0; i <= 7; i++ {
		b := i
		n := fmt.Sprintf("BIT %d, (IX+d)", b)
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			for d := -128; d <= 127; d++ {
				for v := 0; v <= 255; v++ {
					mem := MapMemory{}.
						Put(0, 0xdd, 0xcb, uint8(d), uint8(0x46|b<<3)).
						Put(0x1000+uint16(d), uint8(v))
					wantZ := v&(1<<b) == 0
					var flag uint8
					flagOp{}.Put(Z, wantZ).Set(H).Reset(N).ApplyOn(&flag)
					tSteps(t, "",
						States{GPR: GPR{}, SPR: SPR{IX: 0x1000}},
						mem,
						1,
						States{
							GPR: GPR{AF: Register{Lo: flag}},
							SPR: SPR{PC: 4, IR: Register{Lo: 0x01},
								IX: 0x1000,
							}},
						mem.Clone(),
						maskS|maskPV|mask5|mask3|maskC,
					)
				}
			}
		})
	}
}

func TestBitop_BITbIYd(t *testing.T) {
	t.Parallel()
	for i := 0; i <= 7; i++ {
		b := i
		n := fmt.Sprintf("BIT %d, (IY+d)", b)
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			for d := -128; d <= 127; d++ {
				for v := 0; v <= 255; v++ {
					mem := MapMemory{}.
						Put(0, 0xfd, 0xcb, uint8(d), uint8(0x46|b<<3)).
						Put(0x4180+uint16(d), uint8(v))
					wantZ := v&(1<<b) == 0
					var flag uint8
					flagOp{}.Put(Z, wantZ).Set(H).Reset(N).ApplyOn(&flag)
					tSteps(t, "",
						States{GPR: GPR{}, SPR: SPR{IY: 0x4180}},
						mem,
						1,
						States{
							GPR: GPR{AF: Register{Lo: flag}},
							SPR: SPR{PC: 4, IR: Register{Lo: 0x01},
								IY: 0x4180,
							}},
						mem.Clone(),
						maskS|maskPV|mask5|mask3|maskC,
					)
				}
			}
		})
	}
}

func TestBitop_SETbIXd(t *testing.T) {
	t.Parallel()
	const base uint16 = 0x5678
	for i := 0; i <= 7; i++ {
		b := i
		n := fmt.Sprintf("SET %d, (IX+d)", b)
		t.Run(n, func(t *testing.T) {
			for d := -128; d <= 127; d++ {
				mem := MapMemory{}.
					Put(0, 0xdd, 0xcb, uint8(d), uint8(0xc6|b<<3)).
					Put(base+uint16(d), 0)
				wantMem := mem.Clone().Put(base+uint16(d), 1<<b)
				tOneStep(t,
					States{GPR: GPR{}, SPR: SPR{IX: base}},
					mem,
					States{GPR: GPR{},
						SPR: SPR{PC: 4, IR: Register{Lo: 0x01}, IX: base}},
					wantMem)
			}
		})
	}
}

func TestBitop_SETbIYd(t *testing.T) {
	t.Parallel()
	const base uint16 = 0x9abc
	for i := 0; i <= 7; i++ {
		b := i
		n := fmt.Sprintf("SET %d, (IY+d)", b)
		t.Run(n, func(t *testing.T) {
			for d := -128; d <= 127; d++ {
				mem := MapMemory{}.
					Put(0, 0xfd, 0xcb, uint8(d), uint8(0xc6|b<<3)).
					Put(base+uint16(d), 0)
				wantMem := mem.Clone().Put(base+uint16(d), 1<<b)
				tOneStep(t,
					States{GPR: GPR{}, SPR: SPR{IY: base}},
					mem,
					States{GPR: GPR{},
						SPR: SPR{PC: 4, IR: Register{Lo: 0x01}, IY: base}},
					wantMem)
			}
		})
	}
}

func TestBitop_RESbIXd(t *testing.T) {
	t.Parallel()
	const base uint16 = 0x5678
	for i := 0; i <= 7; i++ {
		b := i
		n := fmt.Sprintf("RES %d, (IX+d)", b)
		t.Run(n, func(t *testing.T) {
			for d := -128; d <= 127; d++ {
				mem := MapMemory{}.
					Put(0, 0xdd, 0xcb, uint8(d), uint8(0x86|b<<3)).
					Put(base+uint16(d), 0xff)
				wantMem := mem.Clone().Put(base+uint16(d), 0xff&^(1<<b))
				tOneStep(t,
					States{GPR: GPR{}, SPR: SPR{IX: base}},
					mem,
					States{GPR: GPR{},
						SPR: SPR{PC: 4, IR: Register{Lo: 0x01}, IX: base}},
					wantMem)
			}
		})
	}
}

func TestBitop_RESbIYd(t *testing.T) {
	t.Parallel()
	const base uint16 = 0x9abc
	for i := 0; i <= 7; i++ {
		b := i
		n := fmt.Sprintf("RES %d, (IY+d)", b)
		t.Run(n, func(t *testing.T) {
			for d := -128; d <= 127; d++ {
				mem := MapMemory{}.
					Put(0, 0xfd, 0xcb, uint8(d), uint8(0x86|b<<3)).
					Put(base+uint16(d), 0xff)
				wantMem := mem.Clone().Put(base+uint16(d), 0xff&^(1<<b))
				tOneStep(t,
					States{GPR: GPR{}, SPR: SPR{IY: base}},
					mem,
					States{GPR: GPR{},
						SPR: SPR{PC: 4, IR: Register{Lo: 0x01}, IY: base}},
					wantMem)
			}
		})
	}
}
