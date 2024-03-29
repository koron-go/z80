package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
)

var cim string
var bin string
var off0 uint

func writeU16(w io.Writer, u16 uint16) error {
	var buf [2]byte
	buf[0] = uint8(u16)
	buf[1] = uint8(u16 >> 8)
	_, err := w.Write(buf[:])
	return err
}

func run() error {
	flag.StringVar(&cim, "cim", "input.cim", `input CIM file name`)
	flag.StringVar(&bin, "bin", "output.bin", `input BIN file name`)
	flag.UintVar(&off0, "off", 0xa000, `offset to load`)
	flag.Parse()
	var off = uint16(off0)

	// reaad CIM
	b, err := os.ReadFile(cim)
	if err != nil {
		return err
	}

	// open file
	f, err := os.Create(bin)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	// write header
	err = w.WriteByte(0xFE)
	if err != nil {
		return err
	}
	err = writeU16(w, off)
	if err != nil {
		return err
	}
	err = writeU16(w, off+uint16(len(b))-1)
	if err != nil {
		return err
	}
	err = writeU16(w, off)
	if err != nil {
		return err
	}

	// write body
	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return w.Flush()
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
