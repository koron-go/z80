package z80

import "fmt"

//go:generate go run ./cmd/gen_switch/ -name switch.go

func (cpu *CPU) step(f fetcher, enableInt bool) error {
	cpu.afterEI = false
	// increment refresh counter
	rc := cpu.IR.Lo
	cpu.IR.Lo = rc&0x80 | (rc+1)&0x7f
	if !cpu.HALT {
		// decode and execute with big switch
		err := decodeExec(cpu, f)
		if err != nil {
			return fmt.Errorf("decode failed at %s: %w", f.fetchLabel(), err)
		}
	}
	// try interruptions.
	if enableInt {
		oldPC := cpu.PC
		ok, err := cpu.tryInterrupt(cpu.afterEI)
		if err != nil {
			return err
		}
		if ok && cpu.IMon != nil {
			cpu.IMon.OnInterrupt(cpu.InNMI, oldPC, cpu.PC)
		}
	}
	return nil
}
