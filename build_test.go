package z80_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/koron-go/z80"
)

func TestBuild(t *testing.T) {
	got := z80.Build()
	if got == nil {
		t.Fatal("failed with nil")
	}
	if _, ok := got.Memory.(z80.DumbMemory); !ok {
		t.Errorf("field Memory doesn't be initialized: got=%#v", got.Memory)
	}
	d := cmp.Diff(&z80.CPU{}, got, cmpopts.IgnoreFields(z80.CPU{}, "Memory"))
	if d != "" {
		t.Errorf("unexpected fields: -want +got\n%s", d)
	}
}

func TestBuildWithMemory(t *testing.T) {
	mem := z80.DumbMemory{0x12, 0x34, 0x56, 0x78, 0x9a, 0xbc, 0xde, 0xf0}
	got := z80.Build(z80.WithMemory(mem))
	if _, ok := got.Memory.(z80.DumbMemory); !ok {
		t.Errorf("field Memory doesn't be initialized: got=%#v", got.Memory)
	}
	if d := cmp.Diff(&z80.CPU{Memory: mem}, got); d != "" {
		t.Errorf("unexpected fields: -want +got\n%s", d)
	}
}

func TestBuildWithIO(t *testing.T) {
	io := make(z80.DumbIO, 256)
	got := z80.Build(z80.WithIO(io))
	d := cmp.Diff(&z80.CPU{IO: io}, got,
		cmpopts.IgnoreFields(z80.CPU{}, "Memory"),
	)
	if d != "" {
		t.Errorf("unexpected fields: -want +got\n%s", d)
	}
}
