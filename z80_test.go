package z80

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testStates struct {
	states States
	memory DumbMemory
	io     DumbIO
}

func testStep(t *testing.T, before *testStates, after *testStates) {
	t.Helper()
	mem := before.memory
	io := before.io
	cpu := &CPU{
		States: before.states,
		Memory: mem,
		IO:     io,
	}
	err := cpu.Step()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(after.states, cpu.States); diff != "" {
		t.Fatalf("unexpected states: -want +got\n%s", diff)
	}
	if diff := cmp.Diff(after.memory, mem); diff != "" {
		t.Fatalf("memory unmatch: -want +got\n%s", diff)
	}
	if diff := cmp.Diff(after.io, io); diff != "" {
		t.Fatalf("io unmatch: -want +got\n%s", diff)
	}
}
