package main

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/koron-go/z80"
)

//go:generate zmac -o zexdoc.cim -o zexdoc.lst zexdoc.asm

func newMemory() z80.MapMemory {
	m := z80.MapMemory{}
	m.Put(0x0000, bios_0000...)
	m.Put(0xfe06, bios_fe06...)
	m.Put(0xff03, bios_ff03...)
	return m
}

// Page 0:
// ref. http://ngs.no.coocan.jp/doc/wiki.cgi/datapack?page=12%BE%CF+%B3%B0%C9%F4%A5%D7%A5%ED%A5%B0%A5%E9%A5%E0%A4%CE%B4%C4%B6%AD#p2
var bios_0000 = []byte{
	0xc3, 0x03, 0xff, 0x00, 0x00, 0xc3, 0x06, 0xfe,
}

// source: _z80/minibios.asm
var bios_fe06 = []byte{
	0x79, 0xfe, 0x02, 0x28, 0x05, 0xfe, 0x09, 0x28, 0x05, 0x76, 0x7b, 0xd3,
	0x00, 0xc9, 0x1a, 0xfe, 0x24, 0xc8, 0xd3, 0x00, 0x13, 0x18, 0xf7,
}

// page for stop code.
var bios_ff03 = []byte{
	0x76,
}

type IO struct {
	out  io.Writer
	warn *log.Logger
}

func newIO() *IO {
	return &IO{
		out:  os.Stdout,
		warn: log.New(os.Stderr, "[WARN]", 0),
	}
}

func (io *IO) In(addr uint8) uint8 {
	io.warn.Printf("not impl. I/O In addr=0x%02x", addr)
	return 0
}

func (io *IO) Out(addr uint8, value uint8) {
	if addr != 0 {
		io.warn.Printf("not impl. I/O Out addr=0x%02x value=0x%02x", addr, value)
		return
	}
	b := []byte{value}
	io.out.Write(b)
}

func main() {
	err := runZexdoc()
	if err != nil {
		log.Fatal(err)
	}
}

func runZexdoc() error {
	prog, err := ioutil.ReadFile("zexdoc.cim")
	if err != nil {
		return err
	}

	stt := z80.States{SPR: z80.SPR{PC: 0x0100}}
	mem := newMemory()
	mem.Put(0x100, prog...)
	io := newIO()
	cpu := z80.CPU{
		States: stt,
		Memory: mem,
		IO:     io,
	}

	for {
		err := cpu.Run(context.Background())
		if err != nil {
			if err == z80.ErrBreakPoint {
				// TODO:
				continue
			}
			return err
		}
		break
	}
	if cpu.PC == 0xff04 {
		return errors.New("halted on 0xff03+1")
	}
	return nil
}
