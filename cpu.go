package z80

import (
	"context"
	"log"
	"sync/atomic"
)

const (
	maskNone = 0x00
	maskC    = 0x01
	maskN    = 0x02
	maskPV   = 0x04
	maskH    = 0x10
	maskZ    = 0x40
	maskS    = 0x80

	mask3 = 0x08
	mask5 = 0x20

	maskS53 = maskS | mask5 | mask3
	mask53  = mask5 | mask3
)

// addrOff apply offset to address.
func addrOff(addr uint16, off uint8) uint16 {
	return addr + uint16(int16(int8(off)))
}

func toU16(l, h uint8) uint16 {
	return (uint16(h) << 8) | uint16(l)
}

func fromU16(v uint16) (l, h uint8) {
	return uint8(v & 0xff), uint8(v >> 8)
}

// im0data is pseudo Memory module be used when IM0 interrupt occurred.
type im0data struct {
	start uint16
	end   uint16
	data  []uint8

	base Memory
}

func newIm0data(pc uint16, d []uint8, base Memory) *im0data {
	return &im0data{
		start: pc,
		end:   pc + uint16(len(d)-1),
		data:  d,
		base:  base,
	}
}

func (im0 *im0data) Get(addr uint16) uint8 {
	if addr < im0.start || addr > im0.end {
		// delegate to base Memory for out of range.
		return im0.base.Get(addr)
	}
	return im0.data[addr-im0.start]
}

func (im0 *im0data) Set(addr uint16, value uint8) {
	if addr >= im0.start && addr <= im0.end {
		// invalid opepration, nothing to do.
		return
	}
	// delegate to base Memory for out of range.
	im0.base.Set(addr, value)
}

//func (cpu *CPU) failf(msg string, args ...interface{}) {
//	log.Printf("Z80 fail: "+msg, args...)
//}

func (cpu *CPU) warnf(msg string, args ...interface{}) {
	log.Printf("Z80 warn: "+msg, args...)
}

// not used for now
//func (cpu *CPU) debugf(msg string, args ...interface{}) {
//	if !cpu.Debug {
//		return
//	}
//	log.Printf("Z80 debug: "+msg, args...)
//}

func (cpu *CPU) flagC() bool {
	return cpu.AF.Lo&maskC != 0
}

func (cpu *CPU) flagN() bool {
	return cpu.AF.Lo&maskN != 0
}

func (cpu *CPU) flagH() bool {
	return cpu.AF.Lo&maskH != 0
}

func (cpu *CPU) flagZ() bool {
	return cpu.AF.Lo&maskZ != 0
}

func (cpu *CPU) flagPV() bool {
	return cpu.AF.Lo&maskPV != 0
}

func (cpu *CPU) flagS() bool {
	return cpu.AF.Lo&maskS != 0
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

func (cpu *CPU) fetch2() (l, h uint8) {
	l = cpu.Memory.Get(cpu.PC)
	cpu.PC++
	h = cpu.Memory.Get(cpu.PC)
	cpu.PC++
	return l, h
}

func (cpu *CPU) fetch16() uint16 {
	l := cpu.Memory.Get(cpu.PC)
	cpu.PC++
	h := cpu.Memory.Get(cpu.PC)
	cpu.PC++
	return (uint16(h) << 8) | uint16(l)
}

// fetchM1 fetches a byte for M1 cycle.
func (cpu *CPU) fetchM1() uint8 {
	c := cpu.Memory.Get(cpu.PC)
	cpu.PC++
	// increment refresh counter
	rc := cpu.IR.Lo
	cpu.IR.Lo = rc&0x80 | (rc+1)&0x7f
	return c
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

// Run executes instructions till HALT or error.
func (cpu *CPU) Run(ctx context.Context) error {
	var ctxErr error
	var canceled int32
	ctx2, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-ctx2.Done()
		ctxErr = ctx.Err()
		atomic.StoreInt32(&canceled, 1)
	}()

	cpu.HALT = false
	for {
		if atomic.LoadInt32(&canceled) != 0 {
			return ctxErr
		}
		cpu.Step()
		if cpu.BreakPoints != nil {
			if _, ok := cpu.BreakPoints[cpu.PC]; ok {
				return ErrBreakPoint
			}
		}
		if cpu.HALT {
			break
		}
	}
	return nil
}

// Step executes an instruction.
func (cpu *CPU) Step() {
	// try interruptions.
	if cpu.Interrupt != nil && cpu.processInterrupt() {
		cpu.Interrupt = nil
		return
	}
	// execute an op-code.
	cpu.executeOne()
}

func (cpu *CPU) processInterrupt() bool {
	if cpu.Interrupt.Type == NMIType {
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = 0x0066
		cpu.IFF2 = cpu.IFF1
		cpu.IFF1 = false
		return true
	}
	// check maskable interrupt.
	if !cpu.IFF1 {
		return false
	}

	switch cpu.IM {
	case 0:
		// Interrupt with IM 0
		if len(cpu.Interrupt.Data) > 0 {
			savedMemory := cpu.Memory
			cpu.Memory = newIm0data(cpu.PC, cpu.Interrupt.Data, savedMemory)
			cpu.executeOne()
			cpu.Memory = savedMemory
			cpu.IFF1 = false
		}
		return true
	case 1:
		// Interrupt with IM 1
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = 0x0038
		cpu.IFF1 = false
		return true
	case 2:
		// Interrupt with IM 2
		if len(cpu.Interrupt.Data) > 0 {
			cpu.SP -= 2
			cpu.writeU16(cpu.SP, cpu.PC)
			// The LSB of interruption vector is ignored in IM 2
			cpu.PC = cpu.readU16(toU16(cpu.Interrupt.Data[0]&0xfe, cpu.IR.Hi))
			cpu.IFF1 = false
		}
		return true
	default:
		return false
	}
}

func (cpu *CPU) invalidCode(code ...uint8) {
	cpu.warnf("detect invalid code, ignored: %X", code)
}
