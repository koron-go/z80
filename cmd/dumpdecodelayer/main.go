package main

import (
	"log"
	"os"

	"github.com/koron-go/z80"
)

func main() {
	err := z80.DumpDecodeLayer(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
