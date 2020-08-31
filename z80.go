/*
Package z80 emulates Zilog's Z80 CPU.
*/
package z80

import "errors"

// ErrBreakPoint shows PC is reached to one of break points.
var ErrBreakPoint = errors.New("break point reached")

// CPU is Z80 emulator.
type CPU struct {
	States

	Memory Memory
	IO     IO
	INT    INT
	NMI    NMI
	IMon   InterruptMonitor

	Debug       bool
	BreakPoints map[uint16]struct{}
}

// States is collection of Z80's internal state.
type States struct {
	GPR
	SPR

	Alternate GPR

	IFF1 bool
	IFF2 bool
	IM   int

	HALT  bool
	InNMI bool
}

// GPR is general purpose reigsters, pair of four registers AF, BC, DE and HL.
type GPR struct {
	AF Register
	BC Register
	DE Register
	HL Register
}

// SPR is special purpose registers.
type SPR struct {
	IR Register
	IX uint16
	IY uint16
	SP uint16
	PC uint16
}

// Register is 16 bits register.
type Register struct {
	Hi uint8
	Lo uint8
}

// U16 gets 16 bits value of this register.
func (r Register) U16() uint16 {
	return (uint16(r.Hi) << 8) | uint16(r.Lo)
}

// SetU16 updates 16 bits value of this register.
func (r *Register) SetU16(v uint16) {
	r.Hi = uint8(v >> 8)
	r.Lo = uint8(v & 0x00ff)
}

// Memory is requirements interface for memory.
type Memory interface {
	Get(addr uint16) uint8
	Set(addr uint16, value uint8)
}

// IO is requirements interface for I/O.
type IO interface {
	In(addr uint8) uint8
	Out(addr uint8, value uint8)
}

// INT is interface for maskable interrupt.
type INT interface {
	// CheckINT should return valid interruption data if maskable interruption
	// made. The data is used for interruption codes or as a vector depending
	// on interruption mode.
	CheckINT() []uint8

	// ReturnINT is called when "RETI" op is executed.
	ReturnINT()
}

// NMI is interruption for non-maskable interrupt.
type NMI interface {
	// CheckNMI should return true if non-maskable interruption made.
	CheckNMI() bool
}

// InterruptMonitor monitors interruptions.
type InterruptMonitor interface {
	OnInterrupt(maskable bool, oldPC, newPC uint16)
}
