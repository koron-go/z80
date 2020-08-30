/*
Package zex provides test cases of Z80 Exerciser.

http://mdfs.net/Software/Z80/Exerciser/
*/
package zex

type CRC32 [4]uint8

type Status struct {
	Inst0 uint8
	InsN1 uint8
	InsN2 uint8
	InsN3 uint8

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

type Case struct {
	Mask     uint8
	Template Status
	CountS   Status
	ShiftS   Status
	Expect   CRC32
	Message  string
}

// MSBT means machine state before test.
const MSBT = 0xe000
