package z80

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
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

// Page 0:
// ref. http://ngs.no.coocan.jp/doc/wiki.cgi/datapack?page=12%BE%CF+%B3%B0%C9%F4%A5%D7%A5%ED%A5%B0%A5%E9%A5%E0%A4%CE%B4%C4%B6%AD#p2
var minibios0000 = []byte{
	0xc3, 0x03, 0xff, 0x00, 0x00, 0xc3, 0x06, 0xfe,
}

// source: _z80/minibios.asm
var minibiosFE06 = []byte{
	0x79, 0xfe, 0x02, 0x28, 0x05, 0xfe, 0x09, 0x28, 0x05, 0x76, 0x7b, 0xd3,
	0x00, 0xc9, 0x1a, 0xfe, 0x24, 0xc8, 0xd3, 0x00, 0x13, 0x18, 0xf7,
}

// page for stop code.
var minibiosFF03 = []byte{
	0x76,
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

func dumpBytes(mem Memory, p uint16, n int) string {
	bb := &bytes.Buffer{}
	for i := n - 1; i >= 0; i-- {
		fmt.Fprintf(bb, "%02X", mem.Get(p+uint16(i)))
	}
	return bb.String()
}

func dumpCounterShifter(mem Memory) {
	fmt.Printf("  counter=%s/%s\n", dumpBytes(mem, 0x15c7, 20), dumpBytes(mem, 0x15db, 20))
	fmt.Printf("  shifter=%s/%s\n", dumpBytes(mem, 0x15ef, 20), dumpBytes(mem, 0x1603, 20))
}

var lastTime = time.Now()

func tRunMinibios(t *testing.T, prog []byte, expOut string, debug bool, breakpoints ...uint16) {
	t.Helper()

	buf := &bytes.Buffer{}
	stt := States{SPR: SPR{PC: 0x100}}
	io := &prerimIO{t: t, out: buf}

	mem := MapMemory{}
	mem.Put(0x0000, minibios0000...)
	mem.Put(0xfe06, minibiosFE06...)
	mem.Put(0xff03, minibiosFF03...)
	mem.Put(0x100, prog...)

	cpu := CPU{
		States: stt,
		Memory: mem,
		IO:     io,
		Debug:  debug,
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
				switch cpu.PC {
				case 0x122:
					p := cpu.HL.U16()
					v := uint16(mem.Get(p+1))<<8 | uint16(mem.Get(p+0))
					fmt.Printf("loop: HL=%04X (HL)=%04X\n", p, v)
				case 0x1431:
					if time.Since(lastTime) < 5*time.Second {
						break
					}
					fmt.Printf("1431:\n")
					dumpCounterShifter(cpu.Memory)
					lastTime = time.Now()
				}
				//t.Logf("break: %#v", cpu.States)
				continue
			}
			t.Fatal(err)
		}
		break
	}
	if cpu.PC == 0xff04 {
		t.Fatalf("halted on 0xff03+1: buf=%q", buf.String())
	}
	actOut := buf.String()
	if actOut != expOut {
		t.Errorf("output mismatch (PC=%04x):\nwant=%q\n got=%q\n", cpu.PC, expOut, actOut)
	}
}

func TestExerciser(t *testing.T) {
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
