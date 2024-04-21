/*
Package z80 emulates Zilog's Z80 CPU.
*/
package z80

import "errors"

// ErrBreakPoint shows PC is reached to one of break points.
var ErrBreakPoint = errors.New("break point reached")

// CPU is the core of Z80 emulator.
type CPU struct {
	States

	Memory Memory
	IO     IO

	// RETNHandle is called when CPU execute a RETN op.
	RETNHandler RETNHandler
	// RETIHandle is called when CPU execute a RETI op.
	RETIHandler RETIHandler

	// Interrupt is a signal to interrupt.  When you set non-nil value, then
	// CPU.Step and CPU.Run treat it as one of Z80 interruptions.
	Interrupt *Interrupt

	BreakPoints map[uint16]struct{}

	// HALT indicates whether the last Run() is terminated with HALT op.
	HALT bool
}

// States is collection of Z80's internal state.
type States struct {
	GPR
	SPR

	Alternate GPR

	IFF1 bool
	IFF2 bool
	IM   int
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

	// ReturnNMI is called when "RETN" op is executed.
	ReturnNMI()
}

// InterruptType is type of interruption.
type InterruptType int

const (
	// NMIType is a type of NMI interruption.
	NMIType InterruptType = iota
	// IMType is a type of normal interruptions.
	IMType
)

// Interrupt is interruption signal.  Put a point of Interrupt to
// CPU.Interrupt, when you want to make an interrupt.
type Interrupt struct {
	Type InterruptType
	Data []uint8
}

// NMIInterrupt creates a Interrupt objct for NMI.
func NMIInterrupt() *Interrupt {
	return &Interrupt{Type: NMIType}
}

// IM0Interrupt creates an Interrupt object for IM0.
func IM0Interrupt(d uint8, others ...uint8) *Interrupt {
	data := make([]uint8, len(others)+1)
	data[0] = d
	copy(data[1:], others)
	return &Interrupt{
		Type: IMType,
		Data: data,
	}
}

// IM1Interrupt creates an Interrupt object for IM1.
func IM1Interrupt() *Interrupt {
	return &Interrupt{Type: IMType}
}

// IM2Interrupt creates an Interrupt object for IM2.
func IM2Interrupt(n uint8) *Interrupt {
	return &Interrupt{
		Type: IMType,
		Data: []uint8{n},
	}
}

// RETNHandler will be called before execute RETN opcode.
type RETNHandler interface {
	RETNHandle()
}

// RETIHandler will be called before execute RETI opcode.
type RETIHandler interface {
	RETIHandle()
}
