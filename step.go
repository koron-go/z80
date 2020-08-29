package z80

import (
	"context"
	"sync/atomic"
)

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
