package z80

import "fmt"

func (cpu *CPU) step(f fetcher, enableInt bool) error {
	afterEI := false
	if !cpu.HALT {
		// increment refresh counter
		rr := cpu.IR.Lo
		cpu.IR.Lo = rr&0x80 | (rr+1)&0x7f
		// decode and execute with big switch
		cpu.decodeBuf[0] = 0
		cpu.decodeBuf[1] = 0
		err := decodeExec(cpu, f)
		if err != nil {
			return fmt.Errorf("decode failed at %s: %w", f.fetchLabel(), err)
		}
		switch cpu.decodeBuf[0] {
		case 0x76: // HALT
			cpu.HALT = true
		case 0xfb: // EI
			afterEI = true
		case 0xed: // RETI or RETN
			switch cpu.decodeBuf[1] {
			case 0x4d: // RETI
				if cpu.INT != nil {
					cpu.INT.ReturnINT()
				}
			case 0x45: // RETN
				cpu.InNMI = false
			}
		}
	}
	// try interruptions.
	if enableInt {
		oldPC := cpu.PC
		ok, err := cpu.tryInterrupt(afterEI)
		if err != nil {
			return err
		}
		if ok && cpu.IMon != nil {
			cpu.IMon.OnInterrupt(cpu.InNMI, oldPC, cpu.PC)
		}
	}
	return nil
}
