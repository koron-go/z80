package z80

import (
	"testing"

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
