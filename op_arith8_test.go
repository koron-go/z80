package z80

import (
	"fmt"
	"math/bits"
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

// isOverflowS8 checks overflow as signed 8 bits variable.
func isOverflowS8(v int32) bool {
	return v > 127 || v < -128
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

func TestArith8_ADDar(t *testing.T) {
	t.Parallel()
	for _, r := range rArith8 {
		for av := uint16(0); av <= 0xff; av++ {
			for rv := uint16(0); rv <= 0xff; rv++ {
				tADDar(t, r, av, rv)
			}
		}
	}
}

func tADDan(t *testing.T, av uint8, nv uint8) {
	mem := MapMemory{}.Put(0, 0xc6, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	sum := uint16(av) + uint16(nv)
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if av&0xf+nv&0xf > 0xf {
		flags |= 0x10 // H: is set if carry from bit 3
	}
	if isOverflowS8(int32(int8(av)) + int32(int8(nv))) {
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
		fmt.Sprintf("ADD A, %02[2]x (A=%02[1]x)", av, nv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_ADDan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tADDan(t, uint8(av), uint8(nv))
		}
	}
}

func tADCan(t *testing.T, av uint8, nv uint8, c bool) {
	mem := MapMemory{}.Put(0, 0xce, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	var cv uint8
	if c {
		cv = 1
		// set carry flag
		preGPR.AF.Lo |= 0x01
	} else {
		// reset carry flag
		preGPR.AF.Lo &= 0xfe
	}

	sum := uint16(av) + uint16(nv) + uint16(cv)
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if (av&0xf)+(nv&0xf)+cv > 0xf {
		flags |= 0x10 // H: is set if carry from bit 3
	}
	if isOverflowS8(int32(int8(av)) + int32(int8(nv)) + int32(cv)) {
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
		fmt.Sprintf("ADC A, %02[2]x (A=%02[1]x C=%[3]t)", av, nv, c),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_ADCan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tADCan(t, uint8(av), uint8(nv), false)
			tADCan(t, uint8(av), uint8(nv), true)
		}
	}
}

func tSUBan(t *testing.T, av uint8, nv uint8) {
	mem := MapMemory{}.Put(0, 0xd6, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	sum := uint16(av) - uint16(nv)
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if av&0xf < nv&0xf {
		flags |= 0x10 // H: is set if borrow from bit 4
	}
	if isOverflowS8(int32(int8(av)) - int32(int8(nv))) {
		flags |= 0x04 // PV is set if overflow (-128~+127)
	}
	flags |= 0x02 // N is set
	if sum > 0xff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	postGPR := preGPR
	postGPR.AF.Hi = uint8(sum)
	postGPR.AF.Lo = flags

	tSteps(t,
		fmt.Sprintf("SUB A, %02[2]x (A=%02[1]x)", av, nv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_SUBan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tSUBan(t, uint8(av), uint8(nv))
		}
	}
}

func tSBCan(t *testing.T, av uint8, nv uint8, c bool) {
	mem := MapMemory{}.Put(0, 0xde, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	var cv uint8
	if c {
		cv = 1
		// set carry flag
		preGPR.AF.Lo |= 0x01
	} else {
		// reset carry flag
		preGPR.AF.Lo &= 0xfe
	}

	sum := uint16(av) - uint16(nv) - uint16(cv)
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if av&0xf < nv&0xf+cv {
		flags |= 0x10 // H: is set if borrow from bit 4
	}
	if isOverflowS8(int32(int8(av)) - int32(int8(nv)) - int32(cv)) {
		flags |= 0x04 // PV is set if overflow (-128~+127)
	}
	flags |= 0x02 // N is set
	if sum > 0xff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	postGPR := preGPR
	postGPR.AF.Hi = uint8(sum)
	postGPR.AF.Lo = flags

	tSteps(t,
		fmt.Sprintf("SBC A, %02[2]x (A=%02[1]x)", av, nv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_SBCan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tSBCan(t, uint8(av), uint8(nv), false)
			tSBCan(t, uint8(av), uint8(nv), true)
		}
	}
}

func tANDan(t *testing.T, av uint8, nv uint8) {
	mem := MapMemory{}.Put(0, 0xe6, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	sum := av & nv
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	flags |= 0x10 // H is set
	if bits.OnesCount8(sum)%2 == 0 {
		flags |= 0x04 // PV is set if parity even
	}
	// N is reset
	// C is reset

	postGPR := preGPR
	postGPR.AF.Hi = sum
	postGPR.AF.Lo = flags

	tSteps(t,
		fmt.Sprintf("AND A, %02[2]x (A=%02[1]x)", av, nv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_ANDan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tANDan(t, uint8(av), uint8(nv))
		}
	}
}

func tXORan(t *testing.T, av uint8, nv uint8) {
	mem := MapMemory{}.Put(0, 0xee, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	sum := av ^ nv
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	// H is reset
	if bits.OnesCount8(sum)%2 == 0 {
		flags |= 0x04 // PV is set if parity even
	}
	// N is reset
	// C is reset

	postGPR := preGPR
	postGPR.AF.Hi = sum
	postGPR.AF.Lo = flags

	tSteps(t,
		fmt.Sprintf("XOR A, %02[2]x (A=%02[1]x)", av, nv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_XORan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tXORan(t, uint8(av), uint8(nv))
		}
	}
}

func tORan(t *testing.T, av uint8, nv uint8) {
	mem := MapMemory{}.Put(0, 0xf6, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	sum := av | nv
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	// H is reset
	if bits.OnesCount8(sum)%2 == 0 {
		flags |= 0x04 // PV is set if parity even
	}
	// N is reset
	// C is reset

	postGPR := preGPR
	postGPR.AF.Hi = sum
	postGPR.AF.Lo = flags

	tSteps(t,
		fmt.Sprintf("OR A, %02[2]x (A=%02[1]x)", av, nv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_ORan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tORan(t, uint8(av), uint8(nv))
		}
	}
}

func tCPan(t *testing.T, av uint8, nv uint8) {
	mem := MapMemory{}.Put(0, 0xfe, nv)

	preGPR := testInitGPR
	preGPR.AF.Hi = av

	sum := uint16(av) - uint16(nv)
	var flags uint8
	if sum&0x80 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if av&0xf < nv&0xf {
		flags |= 0x10 // H: is set if borrow from bit 4
	}
	if isOverflowS8(int32(int8(av)) - int32(int8(nv))) {
		flags |= 0x04 // PV is set if overflow (-128~+127)
	}
	flags |= 0x02 // N is set
	if sum > 0xff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	postGPR := preGPR
	postGPR.AF.Lo = flags

	tSteps(t,
		fmt.Sprintf("CP %02[2]x (A=%02[1]x)", av, nv),
		States{GPR: preGPR}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{PC: 0x0002, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_alu8i_CPan(t *testing.T) {
	t.Parallel()
	for av := 0; av <= 0xff; av++ {
		for nv := 0; nv <= 0xff; nv++ {
			tCPan(t, uint8(av), uint8(nv))
		}
	}
}

func tADDAIXd(t *testing.T, av, dv uint8, ix uint16, mem MapMemory) {
	mem.Put(0, 0xdd, 0x86, dv)
	preGPR := testInitGPR
	preGPR.AF.Hi = av
	xv := mem.Get(ix + uint16(int16(int8(dv))))
	r := av + xv
	hc := av&0x0f+xv&0x0f > 0x0f
	pv := isOverflowS8(int32(int8(av)) + int32(int8(xv)))
	fo := flagOp{}.
		Put(S, r&0x80 != 0). // S is set if result is negative
		Put(Z, r&0xff == 0). // Z is set if result is 0
		Put(H, hc).          // H is set if carry for bit3
		Put(PV, pv).         // PV is set if overflow (-128~+127)
		Reset(N).            // N is reset
		Put(C, av > ^xv)     // C is set if carry from bit 7
	postGPR := preGPR
	postGPR.AF = Register{Hi: r, Lo: fo.Or}
	tSteps(t,
		fmt.Sprintf("ADD A, (IX+d) (A=%02[1]x d=%02[2]x IX=%04[3]x (IX)=%02[4]x)", av, dv, ix, xv),
		States{
			GPR: preGPR,
			SPR: SPR{IX: ix},
		}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{IX: ix, PC: 0x0003, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_ADDAIXd(t *testing.T) {
	t.Parallel()
	var base uint16 = 0x4000
	mem := MapMemory{}
	for i := 0; i < 256; i++ {
		mem.Set(base+uint16(i), uint8(0x80+i))
	}
	ix := base + 0x80
	for av := 0; av <= 0xff; av++ {
		for dv := 0; dv <= 0xff; dv++ {
			tADDAIXd(t, uint8(av), uint8(dv), ix, mem)
		}
	}
}

func tADDAIYd(t *testing.T, av, dv uint8, iy uint16, mem MapMemory) {
	mem.Put(0, 0xfd, 0x86, dv)
	preGPR := testInitGPR
	preGPR.AF.Hi = av
	yv := mem.Get(iy + uint16(int16(int8(dv))))
	r := av + yv
	hc := av&0x0f+yv&0x0f > 0x0f
	pv := isOverflowS8(int32(int8(av)) + int32(int8(yv)))
	fo := flagOp{}.
		Put(S, r&0x80 != 0). // S is set if result is negative
		Put(Z, r&0xff == 0). // Z is set if result is 0
		Put(H, hc).          // H is set if carry for bit3
		Put(PV, pv).         // PV is set if overflow (-128~+127)
		Reset(N).            // N is reset
		Put(C, av > ^yv)     // C is set if carry from bit 7
	postGPR := preGPR
	postGPR.AF = Register{Hi: r, Lo: fo.Or}
	tSteps(t,
		fmt.Sprintf("ADD A, (IY+d) (A=%02[1]x d=%02[2]x IY=%04[3]x (IY)=%02[4]x)", av, dv, iy, yv),
		States{
			GPR: preGPR,
			SPR: SPR{IY: iy},
		}, mem, 1,
		States{
			GPR: postGPR,
			SPR: SPR{IY: iy, PC: 0x0003, IR: Register{Lo: 0x01}},
		}, mem, maskDefault)
}

func TestArith8_ADDAIYd(t *testing.T) {
	t.Parallel()
	var base uint16 = 0x4000
	mem := MapMemory{}
	for i := 0; i < 256; i++ {
		mem.Set(base+uint16(i), uint8(0x80+i))
	}
	iy := base + 0x80
	for av := 0; av <= 0xff; av++ {
		for dv := 0; dv <= 0xff; dv++ {
			tADDAIYd(t, uint8(av), uint8(dv), iy, mem)
		}
	}
}
