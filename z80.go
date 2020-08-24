/*
Package z80 emulates Zilog's Z80 CPU.
*/
package z80

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
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

	onceInit    sync.Once
	decodeLayer *decodeLayer
	decodeBuf   []uint8
}

func (cpu *CPU) failf(msg string, args ...interface{}) {
	log.Printf("Z80 failed: "+msg, args...)
}

func (cpu *CPU) debugf(msg string, args ...interface{}) {
	if !cpu.Debug {
		return
	}
	log.Printf("Z80 debug: "+msg, args...)
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

func (cpu *CPU) reg16qq(n uint8) *Register {
	switch n & 0x03 {
	case 0x00:
		return &cpu.BC
	case 0x01:
		return &cpu.DE
	case 0x02:
		return &cpu.HL
	case 0x03:
		return &cpu.AF
	default:
		cpu.failf("invalid reg16qq: %02x", n)
		return nil
	}
}

// Flag gets a bit of flag register (F)
func (cpu *CPU) Flag(n int) bool {
	return cpu.flag(n)
}

func (cpu *CPU) flag(n int) bool {
	return Flag(cpu.AF.Lo, n)
}

func (cpu *CPU) flagUpdate(fo FlagOp) {
	fo.ApplyOn(&cpu.AF.Lo)
}

func (cpu *CPU) flagCC(n uint8) bool {
	switch n & 0x07 {
	case 0x00:
		return !cpu.flag(Z)
	case 0x01:
		return cpu.flag(Z)
	case 0x02:
		return !cpu.flag(C)
	case 0x03:
		return cpu.flag(C)
	case 0x04:
		return !cpu.flag(PV)
	case 0x05:
		return cpu.flag(PV)
	case 0x06:
		return !cpu.flag(S)
	case 0x07:
		return cpu.flag(S)
	default:
		cpu.failf("invalid flagCC: %02x", n)
		return false
	}
}

func (cpu *CPU) exec(op *OPCode, args []uint8) {
	for i, c := range op.C {
		args[i] &= c.M
	}
	op.F(cpu, args)
}

func (cpu *CPU) init() {
	cpu.onceInit.Do(func() {
		cpu.decodeLayer = defaultDecodeLayer()
		cpu.decodeBuf = make([]uint8, 8)
	})
}

// Run executes instructions till HALT or error.
func (cpu *CPU) Run(ctx context.Context) error {
	cpu.init()

	var ctxErr error
	var canceled int32
	ctx2, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-ctx2.Done()
		ctxErr = ctx.Err()
		atomic.StoreInt32(&canceled, 1)
	}()

	for !cpu.HALT {
		if atomic.LoadInt32(&canceled) != 0 {
			return ctxErr
		}
		err := cpu.step(cpu, true)
		if err != nil {
			return err
		}
		if cpu.BreakPoints != nil {
			if _, ok := cpu.BreakPoints[cpu.PC]; ok {
				return ErrBreakPoint
			}
		}
	}
	return nil
}

// Step executes an instruction.
func (cpu *CPU) Step() error {
	cpu.init()
	return cpu.step(cpu, true)
}

func (cpu *CPU) step(f fetcher, enableInt bool) error {
	return cpu.step2(f, enableInt)
}

func (cpu *CPU) step1(f fetcher, enableInt bool) error {
	afterEI := false
	if !cpu.HALT {
		// fetch an OPCode and increase refresh register.
		op, buf, err := cpu.decode(f)
		if err != nil {
			label := f.fetchLabel()
			return fmt.Errorf("decode failed %X at %s: %w", buf, label, err)
		}
		rr := cpu.IR.Lo
		cpu.IR.Lo = rr&0x80 | (rr+1)&0x7f
		// execute an OPCode.
		if cpu.Debug {
			label := f.fetchLabel()
			cpu.debugf("execute OPCode:%s with %X at %s", op.N, buf, label)
		}
		cpu.exec(op, buf)
		switch op {
		case opcHALT:
			cpu.HALT = true
		case opcEI:
			afterEI = true
		case opcRETI:
			if cpu.INT != nil {
				cpu.INT.ReturnINT()
			}
		case opcRETN:
			cpu.InNMI = false
		}
	}
	// try interruptions.
	if enableInt {
		oldPC := cpu.PC
		ok, err := cpu.tryInterrupt(afterEI)
		if err != nil {
			return err
		}
		if ok && cpu.IMon != nil {
			cpu.IMon.OnInterrupt(cpu.InNMI, oldPC, cpu.PC)
		}
	}
	return nil
}

func (cpu *CPU) tryInterrupt(suppressINT bool) (bool, error) {
	// check non-maskable interrupt.
	if cpu.InNMI {
		return false, nil
	}
	if cpu.NMI != nil && cpu.NMI.CheckNMI() {
		cpu.HALT = false
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = 0x0066
		cpu.IFF2 = cpu.IFF1
		cpu.IFF1 = false
		cpu.InNMI = true
		return true, nil
	}
	// check maskable interrupt.
	if suppressINT || cpu.INT == nil {
		return false, nil
	}
	d := cpu.INT.CheckINT()
	if d == nil {
		return false, nil
	}
	switch cpu.IM {
	case 1:
		cpu.HALT = false
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = 0x0038
		return true, nil
	case 2:
		if n := len(d); n != 1 {
			cpu.failf("interruption data should be 1 byte in IM 2")
			if n == 0 {
				return false, nil
			}
		}
		cpu.HALT = false
		cpu.SP -= 2
		cpu.writeU16(cpu.SP, cpu.PC)
		cpu.PC = toU16(d[0]&0xfe, cpu.IR.Hi)
		return true, nil
	}
	// interrupt with IM 0
	if len(d) == 0 {
		cpu.failf("interruption data should be longer 1 byte in IM 0")
		return false, nil
	}
	cpu.HALT = false
	ms := memSrc(d)
	err := cpu.step(&ms, false)
	if err != nil {
		return true, err
	}
	return true, nil
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
