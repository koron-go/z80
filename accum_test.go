package z80

import "testing"

func TestAccum_bitchk8(t *testing.T) {
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
