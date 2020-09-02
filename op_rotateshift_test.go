package z80

import "testing"

func Benchmark_oopRLA(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Hi = uint8(i)
		cpu.AF.Lo |= uint8(i>>8) & maskC
		oopRLA(cpu)
	}
}

func Benchmark_oopRRA(b *testing.B) {
	cpu := &CPU{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cpu.AF.Hi = uint8(i)
		cpu.AF.Lo |= uint8(i>>8) & maskC
		oopRRA(cpu)
	}
}
