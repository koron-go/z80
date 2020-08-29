/*
Package z80 emulates Zilog's Z80 CPU.
*/
package z80

import (
	"fmt"
)

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

func (cpu *CPU) readU16(addr uint16) uint16 {
	l := cpu.Memory.Get(addr)
	h := cpu.Memory.Get(addr + 1)
	return toU16(l, h)
}

func (cpu *CPU) writeU16(addr uint16, v uint16) {
	l, h := fromU16(v)
	cpu.Memory.Set(addr, l)
	cpu.Memory.Set(addr+1, h)
}

func (cpu *CPU) fetch() uint8 {
	v := cpu.Memory.Get(cpu.PC)
	cpu.PC++
	return v
}

func (cpu *CPU) fetchLabel() string {
	return fmt.Sprintf("0x%04X", cpu.PC)
}

func (cpu *CPU) regP(n uint8) *uint8 {
	switch n & 0x07 {
	case 0x00:
		return &cpu.BC.Hi
	case 0x01:
		return &cpu.BC.Lo
	case 0x02:
		return &cpu.DE.Hi
	case 0x03:
		return &cpu.DE.Lo
	case 0x04:
		return &cpu.HL.Hi
	case 0x05:
		return &cpu.HL.Lo
	case 0x07:
		return &cpu.AF.Hi
	default:
		cpu.failf("invalid register: %02x", n)
		return nil
	}
}

type reg16 interface {
	U16() uint16
	SetU16(uint16)
}

type reg16uint uint16

func (r *reg16uint) U16() uint16 {
	return *((*uint16)(r))
}

func (r *reg16uint) SetU16(v uint16) {
	*((*uint16)(r)) = v
}

func (cpu *CPU) reg16dd(n uint8) reg16 {
	switch n & 0x03 {
	case 0x00:
		return &cpu.BC
	case 0x01:
		return &cpu.DE
	case 0x02:
		return &cpu.HL
	case 0x03:
		return (*reg16uint)(&cpu.SP)
	default:
		cpu.failf("invalid reg16dd: %02x", n)
		return nil
	}
}

func (cpu *CPU) reg16ss(n uint8) reg16 {
	switch n & 0x03 {
	case 0x00:
		return &cpu.BC
	case 0x01:
		return &cpu.DE
	case 0x02:
		return &cpu.HL
	case 0x03:
		return (*reg16uint)(&cpu.SP)
	default:
		cpu.failf("invalid reg16ss: %02x", n)
		return nil
	}
}

func (cpu *CPU) reg16pp(n uint8) uint16 {
	switch n & 0x03 {
	case 0x00:
		return cpu.BC.U16()
	case 0x01:
		return cpu.DE.U16()
	case 0x02:
		return cpu.IX
	case 0x03:
		return cpu.SP
	default:
		cpu.failf("invalid reg16pp: %02x", n)
		return 0
	}
}

func (cpu *CPU) reg16rr(n uint8) uint16 {
	switch n & 0x03 {
	case 0x00:
		return cpu.BC.U16()
	case 0x01:
		return cpu.DE.U16()
	case 0x02:
		return cpu.IY
	case 0x03:
		return cpu.SP
	default:
		cpu.failf("invalid reg16rr: %02x", n)
		return 0
	}
}

// Flag gets a bit of flag register (F)
func (cpu *CPU) Flag(n int) bool {
	return cpu.flag(n)
}

func (cpu *CPU) flag(n int) bool {
	return Flag(cpu.AF.Lo, n)
}

func (cpu *CPU) ioIn(addr uint8) uint8 {
	if cpu.IO == nil {
		return 0
	}
	return cpu.IO.In(addr)
}

func (cpu *CPU) ioOut(addr uint8, value uint8) {
	if cpu.IO == nil {
		return
	}
	cpu.IO.Out(addr, value)
}
