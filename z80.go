package z80

import "log"

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

// RegisterSet is pair of four registers AF, BC, DE and HL.
type RegisterSet struct {
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

// CPU is Z80 emulator.
type CPU struct {
	RegisterSet
	SPR

	Alternate RegisterSet

	IFF1 bool
	IFF2 bool

	Memory Memory
	IO     IO
}

func (cpu *CPU) failf(msg string, args ...interface{}) {
	log.Printf("Z80 failed: "+msg, args...)
}

func (cpu *CPU) readAtPC() uint8 {
	v := cpu.Memory.Get(cpu.PC)
	cpu.PC++
	return v
}

func (cpu *CPU) regP(n uint8) *uint8 {
	switch (n & 0x07) {
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

// Next executes an instruction.
func (cpu *CPU) Next() *CPU {
	// TODO:
	op0 := cpu.readAtPC()
	_ = op0
	return cpu
}
