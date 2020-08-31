package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/koron-go/z80"
	"github.com/koron-go/z80/internal/tinycpm"
)

var cpuprof string
var memprof string

//go:generate zmac -o zexdoc.cim -o zexdoc.lst ../../_z80/zexdoc.asm

func main() {
	flag.StringVar(&cpuprof, "cpuprof", "", "profile CPU, output filename")
	flag.StringVar(&memprof, "memprof", "", "profile memory, output filename")
	flag.Parse()
	err := runZexdoc()
	if err != nil {
		log.Fatal(err)
	}
}

func runZexdoc() error {
	mem, io := tinycpm.New()
	err := mem.LoadFile("zexdoc.cim")
	if err != nil {
		return err
	}

	stt := z80.States{SPR: z80.SPR{PC: 0x0100}}
	cpu := z80.CPU{
		States: stt,
		Memory: mem,
		IO:     io,
	}

	if cpuprof != "" {
		f, err := os.Create(cpuprof)
		if err != nil {
			return fmt.Errorf("could not create CPU profile: %w", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			return fmt.Errorf("could not start CPU profile: %w", err)
		}
		defer pprof.StopCPUProfile()
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

	if memprof != "" {
		f, err := os.Create(memprof)
		if err != nil {
			return fmt.Errorf("could not create memory profile: %w", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.Lookup("heap").WriteTo(f, 0); err != nil {
			return fmt.Errorf("could not write memory profile: %w", err)
		}
	}

	if cpu.PC != 0xff04 {
		return fmt.Errorf("halted on unexpected PC: %04x", cpu.PC)
	}
	return nil
}
