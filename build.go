package z80

// Build builds a Z80 CPU emulator with given options.
func Build(options ...Option) *CPU {
	cpu := new(CPU)
	for _, opt := range options {
		if opt != nil {
			opt.apply(cpu)
		}
	}
	// Memory isn't nil-guarded, so we should set a default instance.
	if cpu.Memory == nil {
		cpu.Memory = make(DumbMemory, 65536)
	}
	return cpu
}

// Option is an optional element for Build.
type Option interface {
	apply(*CPU)
}

type optionFunc func(*CPU)

func (fn optionFunc) apply(cpu *CPU) {
	fn(cpu)
}

// WithMemory is an option to setup with Memory.
func WithMemory(v Memory) Option {
	return optionFunc(func(cpu *CPU) {
		cpu.Memory = v
	})
}

// WithIO is an option to setup with IO.
func WithIO(v IO) Option {
	return optionFunc(func(cpu *CPU) {
		cpu.IO = v
	})
}

// WithINT is an option to setup with INT.
func WithINT(v INT) Option {
	return optionFunc(func(cpu *CPU) {
		cpu.INT = v
	})
}

// WithNMI is an option to setup with NMI.
func WithNMI(v NMI) Option {
	return optionFunc(func(cpu *CPU) {
		cpu.NMI = v
	})
}

// WithInterruptMonitor is an option to setup with InterruptMonitor.
func WithInterruptMonitor(v InterruptMonitor) Option {
	return optionFunc(func(cpu *CPU) {
		cpu.IMon = v
	})
}
