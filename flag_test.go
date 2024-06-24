package z80_test

import (
	"testing"

	"github.com/koron-go/z80"
)

func TestGetFlag(t *testing.T) {
	for i, c := range []struct {
		lo   uint8
		flag z80.Flag
		want bool
	}{
		{0b00000001, z80.FlagC, true},
		{0b00000000, z80.FlagC, false},
		{0b11111110, z80.FlagC, false},

		{0b00000010, z80.FlagN, true},
		{0b00000000, z80.FlagN, false},
		{0b11111101, z80.FlagN, false},

		{0b00000100, z80.FlagPV, true},
		{0b00000000, z80.FlagPV, false},
		{0b11111011, z80.FlagPV, false},

		{0b00010000, z80.FlagH, true},
		{0b00000000, z80.FlagH, false},
		{0b11100111, z80.FlagH, false},

		{0b01000000, z80.FlagZ, true},
		{0b00000000, z80.FlagZ, false},
		{0b10111111, z80.FlagZ, false},

		{0b10000000, z80.FlagS, true},
		{0b00000000, z80.FlagS, false},
		{0b01111111, z80.FlagS, false},
	} {
		var cpu z80.CPU
		cpu.AF.Lo = c.lo
		got := cpu.GetFlag(c.flag)
		if got != c.want {
			t.Errorf("unexpected value: %d %+v: got=%t", i, c, got)
		}
	}
}

func TestSetFlag(t *testing.T) {
	for i, c := range []struct {
		prev uint8
		flag z80.Flag
		want uint8
	}{
		{0b00000000, z80.FlagC, 0b00000001},
		{0b00000001, z80.FlagC, 0b00000001},
		{0b11111110, z80.FlagC, 0b11111111},

		{0b00000000, z80.FlagN, 0b00000010},
		{0b00000010, z80.FlagN, 0b00000010},
		{0b11111101, z80.FlagN, 0b11111111},

		{0b00000000, z80.FlagPV, 0b00000100},
		{0b00000100, z80.FlagPV, 0b00000100},
		{0b11111011, z80.FlagPV, 0b11111111},

		{0b00000000, z80.FlagH, 0b00010000},
		{0b00010000, z80.FlagH, 0b00010000},
		{0b11101111, z80.FlagH, 0b11111111},

		{0b00000000, z80.FlagZ, 0b01000000},
		{0b01000000, z80.FlagZ, 0b01000000},
		{0b10111111, z80.FlagZ, 0b11111111},

		{0b00000000, z80.FlagS, 0b10000000},
		{0b10000000, z80.FlagS, 0b10000000},
		{0b01111111, z80.FlagS, 0b11111111},
	} {
		var cpu z80.CPU
		cpu.AF.Lo = c.prev
		cpu.SetFlag(c.flag)
		got := cpu.AF.Lo
		if got != c.want {
			t.Errorf("unexpected value: %d {prev:%08b flag:0x%02x want:%08b}: got=%08b", i, c.prev, c.flag, c.want, got)
		}
	}
}

func TestResetFlag(t *testing.T) {
	for i, c := range []struct {
		prev uint8
		flag z80.Flag
		want uint8
	}{
		{0b00000001, z80.FlagC, 0b00000000},
		{0b00000000, z80.FlagC, 0b00000000},
		{0b11111111, z80.FlagC, 0b11111110},

		{0b00000010, z80.FlagN, 0b00000000},
		{0b00000000, z80.FlagN, 0b00000000},
		{0b11111111, z80.FlagN, 0b11111101},

		{0b00000100, z80.FlagPV, 0b00000000},
		{0b00000000, z80.FlagPV, 0b00000000},
		{0b11111111, z80.FlagPV, 0b11111011},

		{0b00010000, z80.FlagH, 0b00000000},
		{0b00000000, z80.FlagH, 0b00000000},
		{0b11111111, z80.FlagH, 0b11101111},

		{0b01000000, z80.FlagZ, 0b00000000},
		{0b00000000, z80.FlagZ, 0b00000000},
		{0b10111111, z80.FlagZ, 0b10111111},

		{0b10000000, z80.FlagS, 0b00000000},
		{0b00000000, z80.FlagS, 0b00000000},
		{0b11111111, z80.FlagS, 0b01111111},
	} {
		var cpu z80.CPU
		cpu.AF.Lo = c.prev
		cpu.ResetFlag(c.flag)
		got := cpu.AF.Lo
		if got != c.want {
			t.Errorf("unexpected value: %d {prev:%08b flag:0x%02x want:%08b}: got=%08b", i, c.prev, c.flag, c.want, got)
		}
	}
}
