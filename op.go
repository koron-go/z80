package z80

// addrOff apply offset to address.
func addrOff(addr uint16, off uint8) uint16 {
	return addr + uint16(int16(int8(off)))
}

func toU16(l, h uint8) uint16 {
	return (uint16(h) << 8) | uint16(l)
}

func fromU16(v uint16) (l, h uint8) {
	return uint8(v & 0xff), uint8(v >> 8)
}
