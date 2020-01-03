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

func equalBytes(a, b []uint8) bool {
	n := len(a)
	if n != len(b) {
		return false
	}
	for i := 0; i < n;i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
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
	if after.states != cpu.States {
		diff := cmp.Diff(after.states, cpu.States)
		t.Fatalf("unexpected states: -want +got\n%s", diff)
	}
	if !equalBytes(after.memory, mem) {
		diff := cmp.Diff(after.memory, mem)
		t.Fatalf("memory unmatch: -want +got\n%s", diff)
	}
	if !equalBytes(after.io, io) {
		diff := cmp.Diff(after.io, io)
		t.Fatalf("io unmatch: -want +got\n%s", diff)
	}
}
