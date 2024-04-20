package z80

import "math/bits"

// port from WebMSX.
// See https://github.com/ppeccin/WebMSX/blob/654e3aa303e84404fba4a89d5fa21fae32753cf5/src/main/msx/cpu/CPU.js#L1010-L1030
func oopDAA(cpu *CPU) {
	r := cpu.AF.Hi
	c := cpu.flagC()
	if cpu.flagN() {
		if cpu.flagH() || (cpu.AF.Hi&0x0f) > 9 {
			r -= 0x06
		}
		if c || (cpu.AF.Hi > 0x99) {
			r -= 0x60
		}
	} else {
		if cpu.flagH() || (cpu.AF.Hi&0x0f) > 9 {
			r += 0x06
		}
		if c || (cpu.AF.Hi > 0x99) {
			r += 0x60
		}
	}

	var nand uint8 = maskS53 | maskZ | maskH | maskPV
	var or uint8
	or |= uint8(r) & maskS53
	if uint8(r) == 0 {
		or |= maskZ
	}
	or |= (cpu.AF.Hi ^ r) & maskH
	or |= (uint8(bits.OnesCount8(r)%2) - 1) & maskPV
	if cpu.AF.Hi > 0x99 {
		or |= maskC
	}
	cpu.AF.Lo = cpu.AF.Lo&^nand | or

	cpu.AF.Hi = r
}

func oopHALT(cpu *CPU) {
	// HALT does nothing. Since the program counter (PC) also does not advance, rewind it that was advanced by M1 fetch.
	cpu.PC--
	cpu.HALT = true
}

func oopEI(cpu *CPU) {
	cpu.IFF1 = true
	cpu.IFF2 = true
}

func oopCPL(cpu *CPU) {
	cpu.AF.Hi = ^cpu.AF.Hi
	var nand uint8 = mask53
	var or uint8
	or |= cpu.AF.Hi & mask53
	or |= maskH | maskN
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopNEG(cpu *CPU) {
	a := cpu.AF.Hi
	r := ^a + 1
	cpu.AF.Hi = r

	var nand uint8 = maskS53 | maskZ | maskH | maskPV | maskN | maskC
	var or uint8
	or |= r & maskS53
	if r == 0 {
		or |= maskZ
	}
	if r&0x0f != 0 {
		or |= maskH
	}
	if a == 0x80 {
		or |= maskPV
	}
	or |= maskN
	if a != 0x00 {
		or |= maskC
	}
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopCCF(cpu *CPU) {
	var nand uint8 = mask53 | maskH | maskN | maskC
	var or uint8
	or |= cpu.AF.Hi & mask53
	if cpu.flagC() {
		or |= maskH
	} else {
		or |= maskC
	}
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopSCF(cpu *CPU) {
	var nand uint8 = mask53 | maskH | maskN | maskC
	var or uint8
	or |= cpu.AF.Hi & mask53
	or |= maskC
	cpu.AF.Lo = cpu.AF.Lo&^nand | or
}

func oopNOP(cpu *CPU) {
	// do nothing.
}

func oopDI(cpu *CPU) {
	cpu.IFF1 = false
	cpu.IFF2 = false
}

func oopIM0(cpu *CPU) {
	cpu.IM = 0
}

func oopIM1(cpu *CPU) {
	cpu.IM = 1
}

func oopIM2(cpu *CPU) {
	cpu.IM = 2
}
