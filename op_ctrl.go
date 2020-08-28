package z80

import "math/bits"

// port from WebMSX.
// See https://github.com/ppeccin/WebMSX/blob/654e3aa303e84404fba4a89d5fa21fae32753cf5/src/main/msx/cpu/CPU.js#L1010-L1030
func oopDAA(cpu *CPU) {
	r := cpu.AF.Hi
	c := cpu.flag(C)
	if cpu.flag(N) {
		if cpu.flag(H) || (cpu.AF.Hi&0x0f) > 9 {
			r -= 0x06
		}
		if c || (cpu.AF.Hi > 0x99) {
			r -= 0x60
		}
	} else {
		if cpu.flag(H) || (cpu.AF.Hi&0x0f) > 9 {
			r += 0x06
		}
		if c || (cpu.AF.Hi > 0x99) {
			r += 0x60
		}
	}
	cpu.flagUpdate(FlagOp{}.
		Put(S, r&0x80 != 0).
		Put(Z, r == 0).
		Put(H, (cpu.AF.Hi^r)&0x10 != 0).
		Put(PV, bits.OnesCount8(r)%2 == 0).
		Keep(N).
		Put(C, c || cpu.AF.Hi > 0x99))
	cpu.AF.Hi = r
}

func oopHALT(cpu *CPU) {
	// nothing todo.
}

func oopEI(cpu *CPU) {
	cpu.IFF1 = true
	cpu.IFF2 = true
}

func oopCPL(cpu *CPU) {
	cpu.AF.Hi = ^cpu.AF.Hi
	cpu.flagUpdate(FlagOp{}.Set(H).Set(N))
}

func opNEG(cpu *CPU, codes []uint8) {
	a := cpu.AF.Hi
	v := ^a + 1
	cpu.AF.Hi = v
	cpu.flagUpdate(FlagOp{}.
		Put(S, v&0x80 != 0).
		Put(Z, v == 0).
		Put(H, a&0x0f != 0).
		Put(PV, a == 0x80).
		Set(N).
		Put(C, a != 0))
}

func oopCCF(cpu *CPU) {
	c := cpu.flag(C)
	cpu.flagUpdate(FlagOp{}.Put(H, c).Reset(N).Put(C, !c))
}

func oopSCF(cpu *CPU) {
	cpu.flagUpdate(FlagOp{}.Reset(H).Reset(N).Set(C))
}

func oopNOP(cpu *CPU) {
	// do nothing.
}

func oopDI(cpu *CPU) {
	cpu.IFF1 = false
	cpu.IFF2 = false
}

func opIM0(cpu *CPU, codes []uint8) {
	cpu.IM = 0
}

func opIM1(cpu *CPU, codes []uint8) {
	cpu.IM = 1
}

func opIM2(cpu *CPU, codes []uint8) {
	cpu.IM = 2
}
