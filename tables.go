package z80

import "math/bits"

var sz53p [256]uint8

func init() {
	for i := 0; i < 256; i++ {
		v := uint8(i)
		f := v & maskS53
		if v == 0 {
			f |= maskZ
		}
		if bits.OnesCount8(v)%2 == 0 {
			f |= maskPV
		}
		sz53p[i] = f
	}
}
