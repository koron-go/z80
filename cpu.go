package z80

import (
	"context"
	"log"
	"sync/atomic"
)

const (
	maskStd = 0xa8

	maskNone = 0x00
	maskC    = 0x01
	maskN    = 0x02
	maskPV   = 0x04
	maskH    = 0x10
	maskZ    = 0x40
	maskS    = 0x80

	maskDefault = 0x28
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

// memSrc implements fetcher interface.
type memSrc []uint8

func (m *memSrc) fetch() uint8 {
	if len(*m) == 0 {
		return 0
	}
	var b uint8
	b, *m = (*m)[0], (*m)[1:]
	return b
}

func (cpu *CPU) failf(msg string, args ...interface{}) {
	log.Printf("Z80 fail: "+msg, args...)
}

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
	// increment refresh counter
	rc := cpu.IR.Lo
	cpu.IR.Lo = rc&0x80 | (rc+1)&0x7f
	// try interruptions.
	oldPC := cpu.PC
	if cpu.tryInterrupt() {
		if cpu.IMon != nil {
			cpu.IMon.OnInterrupt(cpu.InNMI, oldPC, cpu.PC)
		}
		return
	}
	// execute an op-code.
	cpu.executeOne(cpu)
}

func (cpu *CPU) tryInterrupt() bool {
	// check non-maskable interrupt.
	if cpu.InNMI {
		return false
	}
	if cpu.NMI != nil && cpu.NMI.CheckNMI() {
		cpu.HALT = false
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = 0x0066
		cpu.IFF2 = cpu.IFF1
		cpu.IFF1 = false
		cpu.InNMI = true
		return true
	}
	// check maskable interrupt.
	if cpu.INT == nil {
		return false
	}
	d := cpu.INT.CheckINT()
	if d == nil {
		return false
	}
	switch cpu.IM {
	case 1:
		cpu.HALT = false
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = 0x0038
		return true
	case 2:
		if n := len(d); n != 1 {
			cpu.failf("interruption data should be 1 byte in IM 2")
			if n == 0 {
				return false
			}
		}
		cpu.HALT = false
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = toU16(d[0]&0xfe, cpu.IR.Hi)
		return true
	}
	// interrupt with IM 0
	if len(d) == 0 {
		cpu.failf("interruption data should be longer 1 byte in IM 0")
		return false
	}
	cpu.HALT = false
	ms := memSrc(d)
	cpu.executeOne(&ms)
	return true
}

func (cpu *CPU) invalidCode(code ...uint8) {
	cpu.warnf("detect invalid code, ignored: %X", code)
}
