package z80

// Flag is the flag value used for GetFlag etc.
// Each value can be combined with the OR (|) operator.
type Flag uint8

const (
	FlagC  Flag = 0x01 // FlagC is carry flag.
	FlagN  Flag = 0x02 // FlagN is subtract flag.
	FlagPV Flag = 0x04 // FlagPV is parity/overflow flag.
	FlagH  Flag = 0x10 // FlagH is half carry flag.
	FlagZ  Flag = 0x40 // FlagZ is zero flag.
	FlagS  Flag = 0x80 // FlagS is sign flag.

	Flag3 Flag = 0x08 // Flag3 is undefined flag, at 3rd from LSB.
	Flag5 Flag = 0x20 // Flag5 is undefined flag, at 5th from LSB.
)

// GetFlag gets a flag status.  For available flags, see Flag type.  When
// multiple flags are combined, if any one of them is true, this returns true.
func (gpr GPR) GetFlag(f Flag) bool {
	return gpr.AF.Lo&uint8(f) != 0
}

// SetFlag sets the specified flag to true.  For available flags, see Flag
// type.  When multiple flags are combined, all flags are set to true.
func (gpr *GPR) SetFlag(f Flag) {
	gpr.AF.Lo |= uint8(f)
}

// ResetFlag resets the specified flag to false. Fo available flags, see Flag
// type. When multiple flags are combined, all flags are set to false.
func (gpr *GPR) ResetFlag(f Flag) {
	gpr.AF.Lo &= ^uint8(f)
}
