package z80

import (
	"fmt"
	"testing"
)

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
					FlagOp{}.Put(Z, wantZ).Set(H).Reset(N).ApplyOn(&flag)
					tStepNoIO(t,
						States{GPR: GPR{}, SPR: SPR{IX: 0x1000}},
						mem,
						States{
							GPR: GPR{AF: Register{Lo: flag}},
							SPR: SPR{PC: 4, IR: Register{Lo: 0x01},
								IX: 0x1000,
							}},
						mem.Clone())
				}
			}
		})
	}
}
