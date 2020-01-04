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

func tStepNoIO(t *testing.T, states States, memory tMemory, afterStates States, afterMemory tMemory) {
	t.Helper()
	cpu := &CPU{
		States: states,
		Memory: memory,
		IO:     &tForbiddenIO{},
	}
	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}
	if cpu.States != afterStates {
		diff := cmp.Diff(afterStates, cpu.States)
		t.Fatalf("unexpected states: -want +got\n%s", diff)
	}
	if !memory.Equal(afterMemory) {
		diff := cmp.Diff(afterMemory, memory)
		t.Fatalf("memory unmatch: -want +got\n%s", diff)
	}
	// IO won't be called in tStepNoIO()
}
