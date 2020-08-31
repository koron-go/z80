package z80

import (
	"bytes"
	"context"
	"fmt"
	"hash/crc32"
	"log"
	"testing"

	"github.com/koron-go/z80/internal/tinycpm"
	"github.com/koron-go/z80/internal/zex"
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
	t.Run("run zexdoc", testRunZexdoc)
}

func zexSetStatus(cpu *CPU, s zex.Status) {
	mem := cpu.Memory
	// setup IUT (instruction under test)
	mem.Set(0x1000, s.Inst0)
	mem.Set(0x1001, s.Inst1)
	mem.Set(0x1002, s.Inst2)
	mem.Set(0x1003, s.Inst3)
	mem.Set(0x1004, 0x00)
	// setup machine state
	cpu.IY = s.IY
	cpu.IX = s.IX
	cpu.HL.SetU16(s.HL)
	cpu.DE.SetU16(s.DE)
	cpu.BC.SetU16(s.BC)
	cpu.AF.Lo = s.Flags
	cpu.AF.Hi = s.Accum
	cpu.SP = s.SP
	cpu.PC = 0x1000
	// setup Msbt
	for i, u8 := range s.Bytes()[4:] {
		mem.Set(zex.Msbt+uint16(i), u8)
	}
	mem.Set(zex.Msbt+16, 0x2a)
	mem.Set(zex.Msbt+17, 0x06)
}

func zexGetStatus(cpu *CPU) zex.Status {
	var s zex.Status
	s.MemOP = uint16(cpu.Memory.Get(zex.Msbt)) |
		uint16(cpu.Memory.Get(zex.Msbt+1))<<8
	s.IY = cpu.IY
	s.IX = cpu.IX
	s.HL = cpu.HL.U16()
	s.DE = cpu.DE.U16()
	s.BC = cpu.BC.U16()
	s.Flags = cpu.AF.Lo
	s.Accum = cpu.AF.Hi
	s.SP = cpu.SP
	return s
}

func zexIsHalt(cpu *CPU) bool {
	switch cpu.Memory.Get(0x1000) {
	case 0x76:
		return true
	case 0xdd, 0xfd:
		switch cpu.Memory.Get(0x1001) {
		case 0x76:
			return true
		}
	}
	return false
}

var crcTable = crc32.MakeTable(crc32.IEEE)

func zexUpdateCRC(sum uint32, v uint8) uint32 {
	return crcTable[uint8(sum)^v] ^ (sum >> 8)
}

func zexRunIter(cpu *CPU, iter zex.Iter, shift, count uint64, flagMask uint8, crc uint32) uint32 {
	before := iter.Status(shift, count)
	zexSetStatus(cpu, before)
	// nothing for HALT
	if zexIsHalt(cpu) {
		return crc
	}
	//beforeBytes := before.Bytes()
	//fmt.Printf("\n# %d, %d", shift, count)
	//fmt.Printf("\n%08x %032x\n", beforeBytes[:4], beforeBytes[4:])
	err := cpu.Run(context.Background())
	if err != ErrBreakPoint {
		panic(fmt.Sprintf("unexpected termination: %v", err))
	}
	// capture after status
	after := zexGetStatus(cpu)
	after.Flags &= flagMask
	afterBytes := after.Bytes()[4:] // skip instruction 4 bytes.
	for _, b := range afterBytes {
		crc = zexUpdateCRC(crc, b)
	}
	//fmt.Printf("%08x %032x\n", crc, afterBytes)
	return crc
}

func testRunZexdoc(t *testing.T) {
	for _, c := range zex.DocCases {
		t.Run(c.Desc, func(t *testing.T) {
			t.Parallel()
			testRunZexCase(t, c)
		})
	}
}

func testRunZexCase(t *testing.T, c zex.Case) {
	t.Helper()
	mem, io := tinycpm.New()
	cpu := &CPU{
		Memory: mem,
		IO:     io,
		BreakPoints: map[uint16]struct{}{
			0x1004: {},
		},
	}

	var crc uint32 = 0xffffffff
	iter := c.Iter()

	crc = zexRunIter(cpu, iter, 0, 0, c.FlagMask, crc)

	shiftMax, countMax := c.Maxes()
	for j := uint64(1); j < countMax; j++ {
		crc = zexRunIter(cpu, iter, 1, j, c.FlagMask, crc)
	}
	for i := uint64(2); i < shiftMax+2; i++ {
		for j := uint64(0); j < countMax; j++ {
			crc = zexRunIter(cpu, iter, i, j, c.FlagMask, crc)
		}
	}

	act := crc

	exp := uint32(c.Expect)
	if act != exp {
		t.Errorf("failed %s: want=%08x got=%08x", c.Desc, exp, act)
	}
}
