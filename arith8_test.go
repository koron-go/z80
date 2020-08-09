package z80

import (
	"fmt"
	"testing"
)

var rArith8 = []tReg{
	{"B", 0}, {"C", 1}, {"D", 2}, {"E", 3}, {"H", 4}, {"L", 5}, {"A", 7},
}

func isOverflow(a, b uint16) bool {
	sa := a&0x80 != 0
	sb := b&0x80 != 0
	sv := (a+b)&0x80 != 0
	return sa == sb && sa != sv
}

func tADDar(t *testing.T, r tReg, av uint16, rv uint16) {
	mem := MapMemory{}.Put(0, 0x80|uint8(r.Code))

	preGPR := testInitGPR
	preGPR.AF.Hi = uint8(av)
	*rGet(t, &preGPR, r.Code) = uint8(rv)
	if r.Code == 7 {
		av = rv
	}

	sum := av + rv
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if av&0xf+rv&0xf > 0xf {
		flags |= 0x10 // H: is set if carry from bit 3
	}
	if isOverflow(av, rv) {
		flags |= 0x04 // PV is set if overflow (-128~+127)
	}
	// N is reset
	if sum > 0xff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	postGPR := preGPR
	postGPR.AF.Hi = uint8(sum)
	postGPR.AF.Lo = flags

	tSteps(t,
		fmt.Sprintf("ADD A, %[1]s (A=%02[2]x %[1]s=%02[3]x)", r.Label, av, rv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0001, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestAdd8_ADDar(t *testing.T) {
	t.Parallel()
	for _, r := range rArith8 {
		for av := uint16(0); av <= 0xff; av++ {
			for rv := uint16(0); rv <= 0xff; rv++ {
				tADDar(t, r, av, rv)
			}
		}
	}
}
