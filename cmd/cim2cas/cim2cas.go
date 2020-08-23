package main

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var cim string
var cas string
var nam string
var off0 uint

var header = []byte{0x1f, 0xa6, 0xde, 0xba, 0xcc, 0x13, 0x7d, 0x74}
var typeBin = []byte{0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0, 0xd0}

func writeU16(w io.Writer, u16 uint16) error {
	var buf [2]byte
	buf[0] = uint8(u16)
	buf[1] = uint8(u16 >> 8)
	_, err := w.Write(buf[:])
	return err
}

func writeName(w io.Writer, name []byte) error {
	buf := []byte{0x20, 0x20, 0x20, 0x020, 0x20, 0x20}
	if len(name) > 6 {
		name = name[:6]
	}
	copy(buf, name)
	_, err := w.Write(buf)
	return err
}

func run() error {
	flag.StringVar(&cim, "cim", "input.cim", `input CIM file name`)
	flag.StringVar(&cas, "cas", "output.cas", `input BIN file name`)
	flag.StringVar(&nam, "nam", "", `name in tape`)
	flag.UintVar(&off0, "off", 0xa000, `offset to load`)
	flag.Parse()
	var off = uint16(off0)
	if nam == "" {
		nam = cim
	}

	// reaad CIM
	b, err := ioutil.ReadFile(cim)
	if err != nil {
		return err
	}

	// open file
	f, err := os.Create(cas)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	// write header
	_, err = w.Write(header)
	if err != nil {
		return err
	}
	_, err = w.Write(typeBin)
	if err != nil {
		return err
	}
	err = writeName(w, []byte(nam))
	if err != nil {
		return err
	}
	_, err = w.Write(header)
	if err != nil {
		return err
	}

	// begin, end and start
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
