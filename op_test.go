package z80

import "testing"

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