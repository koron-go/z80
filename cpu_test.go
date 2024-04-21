package z80

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const maskDefault = 0x28

func TestAddrOff(t *testing.T) {
	t.Parallel()
	for _, c := range []struct {
		addr uint16
		off  uint8
		exp  uint16
	}{
		{0x0000, 0x00, 0x0000},
		{0x0000, 0x01, 0x0001},
		{0x0000, 0x7f, 0x007f},
		{0x0000, 0xff, 0xffff},
		{0x0000, 0xfe, 0xfffe},
		{0x0000, 0x80, 0xff80},
		{0x1000, 0x00, 0x1000},
		{0x1000, 0x01, 0x1001},
		{0x1000, 0x7f, 0x107f},
		{0x1000, 0xff, 0x0fff},
		{0x1000, 0xfe, 0x0ffe},
		{0x1000, 0x80, 0x0f80},
	} {
		act := addrOff(c.addr, c.off)
		if act != c.exp {
			t.Fatalf("failed: %+v actual=%04x", c, act)
		}
	}
}

type tMemory interface {
	Memory
	Equal(interface{}) bool
}

type tForbiddenIO struct {
	t *testing.T
}

func (fio *tForbiddenIO) In(addr uint8) uint8 {
	fio.t.Helper()
	fio.t.Fatalf("unexpected call IO.In: addr=%04x", addr)
	return 0
}

func (fio *tForbiddenIO) Out(addr uint8, v uint8) {
	fio.t.Helper()
	fio.t.Fatalf("unexpected call IO.Out: addr=%04x", addr)
}

var _ IO = (*tForbiddenIO)(nil)

// tOneStep executes one step without I/O.
func tOneStep(t *testing.T, states States, memory tMemory, afterStates States, afterMemory tMemory) {
	t.Helper()
	tSteps(t, "", states, memory, 1, afterStates, afterMemory, maskDefault)
}

// tSteps executes N steps without I/O.
func tSteps(t *testing.T, label string, states States, memory tMemory, steps int, afterStates States, afterMemory tMemory, ignoreFlags uint8) {
	t.Helper()
	cpu := &CPU{
		States: states,
		Memory: memory,
		IO:     &tForbiddenIO{},
	}
	for i := 0; i < steps; i++ {
		cpu.Step()
	}

	mask := ^ignoreFlags
	act := maskFlags(cpu.States, mask)
	exp := maskFlags(afterStates, mask)
	if act != exp {
		diff := cmp.Diff(exp, act)
		if label != "" {
			t.Errorf("failed label: %s", label)
		}
		t.Fatalf("unexpected states: -want +got\n%s", diff)
	}

	if !memory.Equal(afterMemory) {
		diff := cmp.Diff(afterMemory, memory)
		if label != "" {
			t.Errorf("failed label: %s", label)
		}
		t.Fatalf("memory unmatch: -want +got\n%s", diff)
	}
	// IO won't be called in tStepNoIO()
}

func maskFlags(s States, mask uint8) States {
	s.GPR.AF.Lo &= mask
	s.Alternate.AF.Lo &= mask
	return s
}

type tReg struct {
	Label string
	Code  int
}

type tRETHandle bool

func (h *tRETHandle) RETNHandle() {
	*h = true
}

func (h *tRETHandle) RETIHandle() {
	*h = true
}

func testIM0(t *testing.T, n uint8) {
	var (
		addr = uint16(n) * 8
		code = 0xC7 + n*0x08
	)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var calledRETI tRETHandle
	cpu := &CPU{
		States: States{SPR: SPR{PC: 0x0100}, IM: 1},
		Memory: MapMemory{}.
			// HALT
			Put(0x0000, 0x76).
			// EI ; RETI
			Put(addr,
				0xfb,
				0xed, 0x4d).
			// IM 0 ; EI ; HALT ; HALT (for return)
			Put(0x0100,
				0xed, 0x46,
				0xfb,
				0x76,
				0x76,
			),
		IO:          &tForbiddenIO{},
		RETIHandler: &calledRETI,
	}

	// Start the program and HALT at 0x0102
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0103 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x103, cpu.PC)
	}
	if cpu.IM != 0 {
		t.Fatalf("unexpected interrupt mode: want=0 got=%d", cpu.IM)
	}

	// Interrupt with IM 0
	cpu.Interrupt = IM0Interrupt(code)
	cpu.Step()
	if cpu.PC != addr {
		t.Fatalf("RST 38H not work: want=%04X got=%04X", addr, cpu.PC)
	}
	if cpu.IFF1 {
		t.Fatal("IFF1 is true, unexpectedly")
	}
	if cpu.Interrupt != nil {
		t.Fatalf("interrupt data not cleaned: %+v", cpu.Interrupt)
	}

	// Return from the interruption.
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0104 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x104, cpu.PC)
	}
	if !cpu.IFF1 {
		t.Fatal("IFF1 is false, unexpectedly")
	}
	if !calledRETI {
		t.Fatalf("RETI is not processed, unexpectedly")
	}
}

func TestInterruptIM0(t *testing.T) {
	for i := 0; i < 8; i++ {
		t.Run(fmt.Sprintf("RST %02XH", i*8), func(t *testing.T) {
			testIM0(t, uint8(i))
		})
	}
}

func TestInterruptIM1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var calledRETI tRETHandle
	cpu := &CPU{
		States: States{SPR: SPR{PC: 0x0100}, IM: 0},
		Memory: MapMemory{}.
			// HALT
			Put(0x0000, 0x76).
			// EI ; RETI
			Put(0x0038,
				0xfb,
				0xed, 0x4d).
			// IM 1 ; EI ; HALT
			Put(0x0100,
				0xed, 0x56,
				0xfb,
				0x76,
			),
		IO:          &tForbiddenIO{},
		RETIHandler: &calledRETI,
	}

	// Start the program and HALT at 0x0102
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0103 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x103, cpu.PC)
	}
	if cpu.IM != 1 {
		t.Fatalf("unexpected interrupt mode: want=1 got=%d", cpu.IM)
	}

	// Interrupt with IM 1: with dummy data
	cpu.Interrupt = IM1Interrupt()
	cpu.Step()
	if cpu.PC != 0x0038 {
		t.Fatalf("IM 1 interruption not work: want=%04X got=%04X", 0x0038, cpu.PC)
	}
	if cpu.IFF1 {
		t.Fatal("IFF1 is true, unexpectedly")
	}
	if cpu.Interrupt != nil {
		t.Fatalf("interrupt data not cleaned: %+v", cpu.Interrupt)
	}

	// Return from the interruption.
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0103 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x103, cpu.PC)
	}
	if !cpu.IFF1 {
		t.Fatal("IFF1 is false, unexpectedly")
	}
	if !calledRETI {
		t.Fatalf("RETI is not processed, unexpectedly")
	}
}

func testIM2(t *testing.T, addr uint16) {
	hi := uint8(addr >> 8)
	lo := uint8(addr)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var calledRETI tRETHandle
	cpu := &CPU{
		States: States{SPR: SPR{PC: 0x0100}, IM: 0},
		Memory: MapMemory{}.
			// HALT
			Put(0x0000, 0x76).
			// EI ; RETI
			Put(0x00C0,
				0xfb, 0xed, 0x4d).
			// IM 2 ; LD A, 20H ; LD I, A ; IM 1 ; EI ; HALT
			Put(0x0100,
				0xed, 0x5e,
				0x3e, hi,
				0xed, 0x47,
				0xfb,
				0x76,
			).
			// vector
			Put(addr-2, 0, 0, 0xc0, 0x00, 0, 0),
		IO:          &tForbiddenIO{},
		RETIHandler: &calledRETI,
	}

	// Start the program and HALT at 0x0102
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0107 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x107, cpu.PC)
	}
	if cpu.IM != 2 {
		t.Fatalf("unexpected interrupt mode: want=2 got=%d", cpu.IM)
	}

	// Interrupt with IM 2: with dummy empty byte array.
	cpu.Interrupt = IM2Interrupt(lo)
	cpu.Step()
	if cpu.PC != 0x00C0 {
		t.Fatalf("IM 2 interruption not work: want=%04X got=%04X", 0x00C0, cpu.PC)
	}
	if cpu.IFF1 {
		t.Fatal("IFF1 is true, unexpectedly")
	}
	if cpu.Interrupt != nil {
		t.Fatalf("interrupt data not cleaned: %+v", cpu.Interrupt)
	}

	// Return from the interruption.
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0107 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x107, cpu.PC)
	}
	if !cpu.IFF1 {
		t.Fatal("IFF1 is false, unexpectedly")
	}
	if !calledRETI {
		t.Fatalf("RETI is not processed, unexpectedly")
	}
}

func TestInterruptIM2(t *testing.T) {
	for _, addr := range []uint16{
		0x3000,
		0x3004,
		0x3080,
		0x4080,
		0x40c0,
		// odd values are invalid because of restriction of IM 2
	} {
		t.Run(fmt.Sprintf("IM 2 with %04X", addr), func(t *testing.T) {
			testIM2(t, addr)
		})
	}
}

func testNMI(t *testing.T, iff1 bool) *CPU {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var calledRETN tRETHandle
	var diOrEi uint8 = 0xf3 // DI
	if iff1 {
		diOrEi = 0xfb
	}
	cpu := &CPU{
		States: States{SPR: SPR{PC: 0x0100}, IM: 0},
		Memory: MapMemory{}.
			// HALT
			Put(0x0000, 0x76).
			Put(0x0038, 0x76).
			// RETN
			Put(0x0066,
				0xed, 0x45).
			// IM 1 ; DI ; HALT
			Put(0x0100,
				0xed, 0x56,
				diOrEi,
				0x76,
			),
		IO: &tForbiddenIO{},

		RETNHandler: &calledRETN,
	}

	// Start the program and HALT at 0x0102
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0103 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x103, cpu.PC)
	}
	if cpu.IM != 1 {
		t.Fatalf("unexpected interrupt mode: want=%d got=%d", 1, cpu.IM)
	}

	// Interrupt with NMI
	cpu.Interrupt = NMIInterrupt()
	cpu.Step()
	if cpu.PC != 0x0066 {
		t.Fatalf("NMI interruption not work: want=%04X got=%04X", 0x0066, cpu.PC)
	}
	if cpu.IFF1 {
		t.Fatal("IFF1 should be false in NMI")
	}
	if cpu.Interrupt != nil {
		t.Fatalf("interrupt data not cleaned: %+v", cpu.Interrupt)
	}

	// Return from the interruption.
	if err := cpu.Run(ctx); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cpu.PC != 0x0103 {
		t.Fatalf("unexpected PC: want=%04X got=%04X", 0x103, cpu.PC)
	}
	if cpu.IFF1 != iff1 {
		t.Fatalf("unexpected IFF1 after NMI: want=%t got=%t", iff1, cpu.IFF1)
	}
	if !calledRETN {
		t.Fatalf("RETN is not processed, unexpectedly")
	}

	return cpu
}

func TestInterruptNMI(t *testing.T) {
	t.Run("DI", func(t *testing.T) {
		cpu := testNMI(t, false)
		// Try to interrupt with IM 1: should be failed.
		cpu.Interrupt = &Interrupt{Type: IMType}
		cpu.Step()
		if cpu.PC != 0x0103 {
			t.Fatalf("IM 1 interruption should be failed: want=%04X got=%04X", 0x0103, cpu.PC)
		}
		if cpu.IFF1 {
			t.Fatal("IFF1 is true, unexpectedly")
		}
	})

	t.Run("EI", func(t *testing.T) {
		cpu := testNMI(t, true)
		// Try to interrupt with IM 1: should be succeeded
		cpu.Interrupt = &Interrupt{Type: IMType}
		cpu.Step()
		if cpu.PC != 0x0038 {
			t.Fatalf("IM 1 interruption should succeeded: want=%04X got=%04X", 0x0038, cpu.PC)
		}
		if cpu.IFF1 {
			t.Fatal("IFF1 is true, unexpectedly")
		}
	})
}

func TestRunCancel(t *testing.T) {
	cpu := &CPU{
		States: States{SPR: SPR{PC: 0x0100}},
		Memory: MapMemory{}.
			// JR -2
			Put(0x0100, 0x18, 0xfe),
		IO: &tForbiddenIO{},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	err := cpu.Run(ctx)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("unexpected error: %+v", err)
	}
}
