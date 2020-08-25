package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"

	"github.com/koron-go/z80/internal/opcode"
	"golang.org/x/tools/imports"
)

var name string
var nofmt bool

func main() {
	flag.StringVar(&name, "name", "", "filename to output")
	flag.BoolVar(&nofmt, "nofmt", false, "suppress to format")
	flag.Parse()

	bb := &bytes.Buffer{}
	err := opcode.WriteSwitchDecoder(bb)
	if err != nil {
		log.Fatalf("failed to generate: %s", err)
	}
	b := bb.Bytes()

	if !nofmt {
		b, err = imports.Process(name, b, nil)
		if err != nil {
			log.Fatalf("failed to formatting: %s", err)
		}
	}

	var w io.Writer = os.Stdout
	if name != "" {
		f, err := os.Create(name)
		if err != nil {
			log.Fatalf("failed to create a file: %s", err)
		}
		defer f.Close()
		w = f
	}
	_, err = w.Write(b)
	if err != nil {
		log.Fatalf("failed to write: %s", err)
	}
}
