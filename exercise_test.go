package z80

import (
	"bytes"
	"context"
	"log"
	"testing"

	"github.com/koron-go/z80/internal/tinycpm"
)

//go:generate zmac -o testdata/testfn02.cim -o testdata/testfn02.lst _z80/testfn02.asm
//go:generate zmac -o testdata/testfn09.cim -o testdata/testfn09.lst _z80/testfn09.asm
//go:generate zmac -o testdata/prelim.cim -o testdata/prelim.lst _z80/prelim.asm

func tRunMinibios(t *testing.T, name, expOut string, breakpoints ...uint16) {
	t.Helper()

	mem, io := tinycpm.New()
	err := mem.LoadFile(name)
	if err != nil {
		t.Fatalf("faield to load program: %s", err)
	}

	// setup I/O
	outbuf := &bytes.Buffer{}
	io.SetStdout(outbuf)
	warnbuf := &bytes.Buffer{}
	io.SetWarnLogger(log.New(warnbuf, "[WARN]", 0))

	stt := States{SPR: SPR{PC: 0x100}}

	cpu := CPU{
		States: stt,
		Memory: mem,
		IO:     io,
	}
	if len(breakpoints) > 0 {
		cpu.BreakPoints = map[uint16]struct{}{}
		for _, v := range breakpoints {
			cpu.BreakPoints[v] = struct{}{}
		}
	}

	for {
		err := cpu.Run(context.Background())
		if err != nil {
			if err == ErrBreakPoint {
				// do something?
				continue
			}
			t.Fatalf("stop with error: %s", err)
		}
		break
	}
	if cpu.PC != 0xff04 {
		t.Errorf("halted on unexpected PC: %04x", cpu.PC)
	}
	if s := warnbuf.String(); s != "" {
		t.Errorf("detect warning:\n%s", s)
	}
	actOut := outbuf.String()
	if actOut != expOut {
		t.Errorf("output mismatch:\nwant=%q\n got=%q\n", expOut, actOut)
	}
}

func TestExerciser(t *testing.T) {
	t.Run("test function call 02h", func(t *testing.T) {
		tRunMinibios(t, "testdata/testfn02.cim", "OK")
	})
	t.Run("test function call 09h", func(t *testing.T) {
		tRunMinibios(t, "testdata/testfn09.cim", "Hello 09h")
	})
	t.Run("run testdata/prelim.cim", func(t *testing.T) {
		tRunMinibios(t, "testdata/prelim.cim", "Preliminary tests complete")
	})
}
