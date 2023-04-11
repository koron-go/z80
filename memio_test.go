package z80_test

import (
	"testing"

	"github.com/koron-go/z80"
)

func TestDummyMemoryGet(t *testing.T) {
	dm := z80.DumbMemory{0x01, 0x02, 0x03, 0x04}

	for i, want := range dm {
		if got := dm.Get(uint16(i)); got != want {
			t.Errorf("get failed at %d: want=%02x got=%02x", i, want, got)
		}
	}
	for i := uint16(4); i != 0; i++ {
		if got := dm.Get(i); got != 0 {
			t.Errorf("get out of range (at %d) should returns zero: got=%02x", i, got)
		}
	}
}
