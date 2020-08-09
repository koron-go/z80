package z80

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
		err := cpu.Step()
		if err != nil {
			t.Fatal(err)
		}
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

const (
	maskNone    = 0x00
	maskDefault = 0x28
	maskC       = 0x01
	maskN       = 0x02
	maskPV      = 0x04
	maskH       = 0x10
	maskZ       = 0x40
	maskS       = 0x80
)

type tReg struct {
	Label string
	Code  int
}
