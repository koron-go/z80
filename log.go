package z80

import "log"

func (cpu *CPU) failf(msg string, args ...interface{}) {
	log.Printf("Z80 failed: "+msg, args...)
}

// not used for now
//func (cpu *CPU) debugf(msg string, args ...interface{}) {
//	if !cpu.Debug {
//		return
//	}
//	log.Printf("Z80 debug: "+msg, args...)
//}
