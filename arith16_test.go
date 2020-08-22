package z80

import (
	"fmt"
	"testing"
)

var rArith16_ss = []tReg{
	{"BC", 0}, {"DE", 1}, {"HL", 2}, {"SP", 3},
}

// isOverflowS16 checks overflow as signed 8 bits variable.
func isOverflowS16(v int32) bool {
	return v > 32767 || v < -32768
}

func tADChlss(t *testing.T, r tReg, hl, ss uint16, c bool) {
	hl32, ss32 := uint32(hl), uint32(ss)
	var c32 uint32

	mem := MapMemory{}.Put(0, 0xed, 0x4a|(uint8(r.Code)<<4))

	preGPR := testInitGPR
	preGPR.HL.SetU16(hl)
	preSPR := SPR{}
	switch r.Code {
	case 0:
		preGPR.BC.SetU16(ss)
	case 1:
		preGPR.DE.SetU16(ss)
	case 2:
		if hl != ss {
			t.Logf("warning ss not matches with HL for ADC HL, ss(=HL)")
			ss = hl
		}
	case 3:
		preSPR.SP = ss
	default:
		t.Fatalf("unsupported register: %d", r.Code)
	}
	if c {
		c32 = 1
		// set carry flag
		preGPR.AF.Lo = 0x01
	} else {
		preGPR.AF.Lo = 0x00
	}

	sum := hl32 + ss32 + c32
	var flags uint8
	if sum&0x8000 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xffff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if hl32&0xfff+ss32&0xfff+c32 > 0xfff {
		flags |= 0x10 // H: is set if carry from bit 3
	}
	if isOverflowS16(int32(int16(hl)) + int32(int16(ss)) + int32(c32)) {
		flags |= 0x04 // PV is set if overflow (-32768~+32767)
	}
	// N is reset
	if sum > 0xffff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	postGPR := preGPR
	postGPR.HL.SetU16(uint16(sum))
	postGPR.AF.Lo = flags
	postSPR := preSPR
	postSPR.PC = 0x0002
	postSPR.IR = Register{Lo: 0x01}

	tSteps(t,
		fmt.Sprintf("ADC HL %[1]s (HL=%04[2]x %[1]s=%04[3]x C=%[4]t", r.Label, hl, ss, c),
		States{GPR: preGPR, SPR: preSPR}, mem, 1,
		States{GPR: postGPR, SPR: postSPR}, mem, maskDefault)
}

func TestAirth16_adc16_ADChlss(t *testing.T) {
	t.Parallel()
	for _, r := range rArith16_ss {
		for _, hl := range u16casesSummary {
			if r.Label == "HL" {
				tADChlss(t, r, uint16(hl), uint16(hl), false)
				tADChlss(t, r, uint16(hl), uint16(hl), true)
				continue
			}
			for _, ss := range u16casesSummary {
				tADChlss(t, r, uint16(hl), uint16(ss), false)
				tADChlss(t, r, uint16(hl), uint16(ss), true)
			}
		}
	}
}

func tSBChlss(t *testing.T, r tReg, hl, ss uint16, c bool) {
	hl32, ss32 := uint32(hl), uint32(ss)
	var c32 uint32

	mem := MapMemory{}.Put(0, 0xed, 0x42|(uint8(r.Code)<<4))

	preGPR := testInitGPR
	preGPR.HL.SetU16(hl)
	preSPR := SPR{}
	switch r.Code {
	case 0:
		preGPR.BC.SetU16(ss)
	case 1:
		preGPR.DE.SetU16(ss)
	case 2:
		if hl != ss {
			t.Logf("warning ss not matches with HL for SBC HL, ss(=HL)")
			ss = hl
		}
	case 3:
		preSPR.SP = ss
	default:
		t.Fatalf("unsupported register: %d", r.Code)
	}
	if c {
		c32 = 1
		// set carry flag
		preGPR.AF.Lo = 0x01
	} else {
		preGPR.AF.Lo = 0x00
	}

	sum := hl32 - ss32 - c32
	var flags uint8
	if sum&0x8000 != 0 {
		flags |= 0x80 // S is set if result is negative
	}
	if sum&0xffff == 0 {
		flags |= 0x40 // Z is set if result is 0
	}
	if hl32&0xfff < (ss32+c32)&0xfff {
		flags |= 0x10 // H: is set if carry from bit 3
	}
	if isOverflowS16(int32(int16(hl)) - int32(int16(ss)) - int32(c32)) {
		flags |= 0x04 // PV is set if overflow (-32768~+32767)
	}
	flags |= 0x02 // N is set
	if sum > 0xffff {
		// C is set if carry from bit 7
		flags |= 0x01
	}

	postGPR := preGPR
	postGPR.HL.SetU16(uint16(sum))
	postGPR.AF.Lo = flags
	postSPR := preSPR
	postSPR.PC = 0x0002
	postSPR.IR = Register{Lo: 0x01}

	tSteps(t,
		fmt.Sprintf("SBC HL %[1]s (HL=%04[2]x %[1]s=%04[3]x C=%[4]t", r.Label, hl, ss, c),
		States{GPR: preGPR, SPR: preSPR}, mem, 1,
		States{GPR: postGPR, SPR: postSPR}, mem, maskDefault)
}

func TestArith16_adc16_SBChlss(t *testing.T) {
	t.Parallel()
	for _, r := range rArith16_ss {
		for _, hl := range u16casesSummary {
			if r.Label == "HL" {
				tSBChlss(t, r, uint16(hl), uint16(hl), false)
				tSBChlss(t, r, uint16(hl), uint16(hl), true)
				continue
			}
			for _, ss := range u16casesSummary {
				tSBChlss(t, r, uint16(hl), uint16(ss), false)
				tSBChlss(t, r, uint16(hl), uint16(ss), true)
			}
		}
	}
}
