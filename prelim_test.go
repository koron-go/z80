package z80

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"
)

type prerimIO struct {
	t   *testing.T
	out *bytes.Buffer
}

func (io *prerimIO) In(addr uint8) uint8 {
	io.t.Helper()
	io.t.Logf("I/O In addr=0x%02x", addr)
	return 0
}

func (io *prerimIO) Out(addr uint8, value uint8) {
	if addr != 0 {
		io.t.Helper()
		io.t.Logf("I/O Out addr=0x%02x value=0x%02x", addr, value)
		return
	}
	io.out.WriteByte(value)
}

// source: _z80/minibios.z80
var minibios = []byte{
	0x79, 0xfe, 0x02, 0x28, 0x05, 0xfe, 0x09, 0x28, 0x05, 0x76, 0x7b, 0xd3,
	0x00, 0xc9, 0x1a, 0xfe, 0x24, 0xc8, 0xd3, 0x00, 0x13, 0x18, 0xf7,
}

// source: _z80/testfn02.z80
var testFn02 = []byte{
	0x0e, 0x02, 0x1e, 0x4f, 0xcd, 0x05, 0x00, 0x1e, 0x4b, 0xcd, 0x05, 0x00,
	0x76,
}

// source: _z80/testfn09.z80
var testFn09 = []byte{
	0x0e, 0x09, 0x11, 0x09, 0x01, 0xcd, 0x05, 0x00, 0x76, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x20, 0x30, 0x39, 0x68, 0x24,
}

func tRunMinibios(t *testing.T, prog []byte, expOut string, debug bool) {
	t.Helper()

	buf := &bytes.Buffer{}
	stt := States{SPR: SPR{PC: 0x100}}
	io := &prerimIO{t: t, out: buf}

	mem := MapMemory{}
	mem.Put(0, 0x76, 0x00, 076, 0x00, 0x76)
	mem.Put(5, minibios...)
	mem.Put(0x100, prog...)

	cpu := CPU{
		States: stt,
		Memory: mem,
		IO:     io,
		Debug: debug,
	}
	err := cpu.Run(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if cpu.PC == 0x0001 {
		t.Fatalf("halted on 0x0000: buf=%q", buf.String())
	}
	actOut := buf.String()
	if actOut != expOut {
		t.Errorf("output mismatch (PC=%04x):\nwant=%s\n got=%s\n", cpu.PC, expOut, actOut)
	}
}

func TestPrelim(t *testing.T) {
	t.Run("test function call 02h", func(t *testing.T) {
		tRunMinibios(t, testFn02, "OK", false)
	})
	t.Run("test function call 09h", func(t *testing.T) {
		tRunMinibios(t, testFn09, "Hello 09h", false)
	})
	t.Run("run testdata/prelim.cim", func(t *testing.T) {
		b, err := ioutil.ReadFile("testdata/prelim.cim")
		if err != nil {
			t.Fatal(err)
		}
		tRunMinibios(t, b, "Preliminary tests complete", false)
	})
}
