package main

import (
	"log"
	"os"

	"github.com/koron-go/z80/internal/opcode"
)

func main() {
	err := opcode.DumpDecodeLayer(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
