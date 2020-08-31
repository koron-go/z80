/*
Package zex provides test cases of Z80 Exerciser.

http://mdfs.net/Software/Z80/Exerciser/
*/
package zex

import (
	"fmt"
	"math/bits"
)

// CRC32 has inverted CRC32 (IEEE) value.
type CRC32 uint32

// Status is Z80 machine status for test.
type Status struct {
	Inst0 uint8
	Inst1 uint8
	Inst2 uint8
	Inst3 uint8

	MemOP uint16
	IY    uint16
	IX    uint16
	HL    uint16
	DE    uint16
	BC    uint16

	Flags uint8
	Accum uint8

	SP uint16
}

func toU16(l, h uint8) uint16 {
	return uint16(l) | uint16(h)<<8
}

func fromU16(v uint16) (l, h uint8) {
	return uint8(v), uint8(v >> 8)
}

func newStatus(b []uint8) Status {
	b = b[:20]
	return Status{
		Inst0: b[0],
		Inst1: b[1],
		Inst2: b[2],
		Inst3: b[3],
		MemOP: toU16(b[4], b[5]),
		IY:    toU16(b[6], b[7]),
		IX:    toU16(b[8], b[9]),
		HL:    toU16(b[10], b[11]),
		DE:    toU16(b[12], b[13]),
		BC:    toU16(b[14], b[15]),
		Flags: b[16],
		Accum: b[17],
		SP:    toU16(b[18], b[19]),
	}
}

// Bytes returns []byte representaion of status.
func (s Status) Bytes() []byte {
	buf := make([]byte, 20)
	buf[0] = s.Inst0
	buf[1] = s.Inst1
	buf[2] = s.Inst2
	buf[3] = s.Inst3
	buf[4], buf[5] = fromU16(s.MemOP)
	buf[6], buf[7] = fromU16(s.IY)
	buf[8], buf[9] = fromU16(s.IX)
	buf[10], buf[11] = fromU16(s.HL)
	buf[12], buf[13] = fromU16(s.DE)
	buf[14], buf[15] = fromU16(s.BC)
	buf[16] = s.Flags
	buf[17] = s.Accum
	buf[18], buf[19] = fromU16(s.SP)
	return buf
}

func (s Status) String() string {
	return fmt.Sprintf("Inst:%02x%02x%02x%02x MemOP:%04x IY:%04x IX:%04x "+
		"HL:%04x DE:%04x BC:%04x F:%02x A:%02x SP:%04x",
		s.Inst0, s.Inst1, s.Inst2, s.Inst3, s.MemOP, s.IY, s.IX, s.HL, s.DE,
		s.BC, s.Flags, s.Accum, s.SP)
}

// OnesCount returns count of 1 bits in Status.
func (s Status) OnesCount() int {
	return bits.OnesCount8(s.Inst0) +
		bits.OnesCount8(s.Inst1) +
		bits.OnesCount8(s.Inst2) +
		bits.OnesCount8(s.Inst3) +
		bits.OnesCount16(s.MemOP) +
		bits.OnesCount16(s.IY) +
		bits.OnesCount16(s.IX) +
		bits.OnesCount16(s.HL) +
		bits.OnesCount16(s.DE) +
		bits.OnesCount16(s.BC) +
		bits.OnesCount8(s.Flags) +
		bits.OnesCount8(s.Accum) +
		bits.OnesCount16(s.SP)
}

// Case defines a test case.
type Case struct {
	FlagMask uint8
	BaseCase Status
	IncVec   Status
	ShiftVec Status
	Expect   CRC32
	Desc     string
}

// Maxes returns max valuies of shifter and counter.
func (c Case) Maxes() (shiftMax, countMax uint64) {
	shiftMax = uint64(c.ShiftVec.OnesCount())
	countMax = uint64(1) << c.IncVec.OnesCount()
	return
}

// Iter creates new states iterator.
func (c Case) Iter() Iter {
	return Iter{
		base:  c.BaseCase.Bytes(),
		inc:   c.IncVec.Bytes(),
		shift: c.ShiftVec.Bytes(),
	}
}

// Iter is states iterator/generator.
type Iter struct {
	base  []uint8
	inc   []uint8
	shift []uint8
}

// Status generate before test stauts for a case.
func (it Iter) Status(shift, count uint64) Status {
	code := make([]uint8, len(it.base))
	copy(code, it.base)
	mi := it.inc[:len(code)]
	ms := it.shift[:len(code)]
	for i := range code {
		if m := mi[i]; m != 0 && count != 0 {
			for j := uint8(0x01); j != 0; j <<= 1 {
				if m&j == 0 {
					continue
				}
				if count&1 != 0 {
					code[i] ^= j
				}
				count >>= 1
			}
		}
		if m := ms[i]; m != 0 && shift != 0 {
			for j := uint8(0x01); j != 0; j <<= 1 {
				if m&j == 0 {
					continue
				}
				if shift == 1 {
					code[i] ^= j
				}
				shift--
			}
		}
	}
	return newStatus(code)
}

// Msbt means machine state before test.
const Msbt = 0x0103

// MsbtLo has low 8 bits of Msbt.
const MsbtLo = Msbt & 0xff

// MsbtHi has high 8 bits of Msbt.
const MsbtHi = (Msbt >> 8) & 0xff
