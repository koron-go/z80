/*
Package tinycpm provides minimal CP/M compatible BIOS to run Z80 Exerciser
tests.
*/
package tinycpm

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Start is an address where a program starts.
const Start = 0x0100

// Memory provides 64K bytes array memory.
type Memory struct {
	buf [65536]uint8
}

// NewMemory creates a new memory which includes minimal CP/M.
func NewMemory() *Memory {
	m := new(Memory)
	m.put(0x0000, bios0000...)
	m.put(0xfe06, biosFE06...)
	m.put(0xff03, biosFF03...)
	return m
}

// Page 0:
// ref. http://ngs.no.coocan.jp/doc/wiki.cgi/datapack?page=12%BE%CF+%B3%B0%C9%F4%A5%D7%A5%ED%A5%B0%A5%E9%A5%E0%A4%CE%B4%C4%B6%AD#p2
var bios0000 = []byte{
	0xc3, 0x03, 0xff, 0x00, 0x00, 0xc3, 0x06, 0xfe,
}

// source: _z80/minibios.asm
var biosFE06 = []byte{
	0x79, 0xfe, 0x02, 0x28, 0x05, 0xfe, 0x09, 0x28, 0x05, 0x76, 0x7b, 0xd3,
	0x00, 0xc9, 0x1a, 0xfe, 0x24, 0xc8, 0xd3, 0x00, 0x13, 0x18, 0xf7,
}

// page for stop code.
var biosFF03 = []byte{
	0x76,
}

// Get gets a byte at addr of memory.
func (m *Memory) Get(addr uint16) uint8 {
	return m.buf[addr]
}

// Set sets a byte at addr of memory.
func (m *Memory) Set(addr uint16, value uint8) {
	m.buf[addr] = value
}

// put puts "data" block from addr.
func (m *Memory) put(addr uint16, data ...uint8) {
	copy(m.buf[int(addr):int(addr)+len(data)], data)
}

// LoadFile loads a file from "Start" (0x0100) as program.
func (m *Memory) LoadFile(name string) error {
	prog, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	m.put(Start, prog...)
	return nil
}

// IO creates a new I/O, which used with minimal CP/M
type IO struct {
	stdout io.Writer
	warnl  *log.Logger
}

// NewIO creates a new I/O, which used with minimal CP/M
func NewIO() *IO {
	return &IO{
		stdout: os.Stdout,
		warnl:  log.New(os.Stderr, "[WARN][IO]", 0),
	}
}

// In inputs a value from "addr" port.
// However this shows warning message always.
func (io *IO) In(addr uint8) uint8 {
	io.warnl.Printf("not impl. I/O In addr=0x%02x", addr)
	return 0
}

// Out outputs "value" to "addr" port.
// Only supports addr=0x0000, otherwise show warning message.
func (io *IO) Out(addr uint8, value uint8) {
	if addr != 0 {
		io.warnl.Printf("not impl. I/O Out addr=0x%02x value=0x%02x", addr, value)
		return
	}
	b := []byte{value}
	io.stdout.Write(b)
}

// SetStdout overrides stdout for this I/O.
func (io *IO) SetStdout(w io.Writer) {
	io.stdout = w
}

// SetWarnLogger overrides a warning logger.
func (io *IO) SetWarnLogger(l *log.Logger) {
	io.warnl = l
}

// New creates Memory and IO.
func New() (*Memory, *IO) {
	return NewMemory(), NewIO()
}
