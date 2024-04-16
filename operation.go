package z80

type fetcher interface {
	fetch() uint8
}

// executeOne executes only an op-code.
func (cpu *CPU) executeOne(f fetcher) {
	if cpu.HALT {
		return
	}
	c0 := f.fetch()
	switch c0 {
	case 0x00:
		oopNOP(cpu)

	case 0x01:
		l := f.fetch()
		h := f.fetch()
		xopLDbcnn(cpu, l, h)
	case 0x11:
		l := f.fetch()
		h := f.fetch()
		xopLDdenn(cpu, l, h)
	case 0x21:
		l := f.fetch()
		h := f.fetch()
		xopLDhlnn(cpu, l, h)
	case 0x31:
		l := f.fetch()
		h := f.fetch()
		xopLDspnn(cpu, l, h)

	case 0x02:
		oopLDBCPA(cpu)

	case 0x03:
		xopINCbc(cpu)
	case 0x13:
		xopINCde(cpu)
	case 0x23:
		xopINChl(cpu)
	case 0x33:
		xopINCsp(cpu)

	case 0x04:
		xopINCb(cpu)
	case 0x0c:
		xopINCc(cpu)
	case 0x14:
		xopINCd(cpu)
	case 0x1c:
		xopINCe(cpu)
	case 0x24:
		xopINCh(cpu)
	case 0x2c:
		xopINCl(cpu)
	case 0x3c:
		xopINCa(cpu)

	case 0x05:
		xopDECb(cpu)
	case 0x0d:
		xopDECc(cpu)
	case 0x15:
		xopDECd(cpu)
	case 0x1d:
		xopDECe(cpu)
	case 0x25:
		xopDECh(cpu)
	case 0x2d:
		xopDECl(cpu)
	case 0x3d:
		xopDECa(cpu)

	// LD r, n
	case 0x06:
		n := f.fetch()
		xopLDbn(cpu, n)
	case 0x0e:
		n := f.fetch()
		xopLDcn(cpu, n)
	case 0x16:
		n := f.fetch()
		xopLDdn(cpu, n)
	case 0x1e:
		n := f.fetch()
		xopLDen(cpu, n)
	case 0x26:
		n := f.fetch()
		xopLDhn(cpu, n)
	case 0x2e:
		n := f.fetch()
		xopLDln(cpu, n)
	case 0x3e:
		n := f.fetch()
		xopLDan(cpu, n)

	case 0x07:
		oopRLCA(cpu)

	case 0x08:
		oopEXAFAF(cpu)

	case 0x09:
		xopADDHLbc(cpu)
	case 0x19:
		xopADDHLde(cpu)
	case 0x29:
		xopADDHLhl(cpu)
	case 0x39:
		xopADDHLsp(cpu)

	case 0x0a:
		oopLDABCP(cpu)

	case 0x0b:
		xopDECbc(cpu)
	case 0x1b:
		xopDECde(cpu)
	case 0x2b:
		xopDEChl(cpu)
	case 0x3b:
		xopDECsp(cpu)

	case 0x0f:
		oopRRCA(cpu)

	case 0x10:
		off := f.fetch()
		oopDJNZe(cpu, off)

	case 0x12:
		oopLDDEPA(cpu)

	case 0x17:
		oopRLA(cpu)

	case 0x18:
		off := f.fetch()
		oopJRe(cpu, off)

	case 0x1a:
		oopLDADEP(cpu)

	case 0x1f:
		oopRRA(cpu)

	case 0x20:
		off := f.fetch()
		oopJRNZe(cpu, off)

	case 0x22:
		l := f.fetch()
		h := f.fetch()
		oopLDnnPHL(cpu, l, h)

	case 0x27:
		oopDAA(cpu)

	case 0x28:
		off := f.fetch()
		oopJRZe(cpu, off)

	case 0x2a:
		l := f.fetch()
		h := f.fetch()
		oopLDHLnnP(cpu, l, h)

	case 0x2f:
		oopCPL(cpu)

	case 0x30:
		off := f.fetch()
		oopJRNCe(cpu, off)

	case 0x32:
		l := f.fetch()
		h := f.fetch()
		oopLDnnPA(cpu, l, h)

	case 0x34:
		oopINCHLP(cpu)

	case 0x35:
		oopDECHLP(cpu)

	case 0x36:
		n := f.fetch()
		oopLDHLPn(cpu, n)

	case 0x37:
		oopSCF(cpu)

	case 0x38:
		off := f.fetch()
		oopJRCe(cpu, off)

	case 0x3a:
		l := f.fetch()
		h := f.fetch()
		oopLDAnnP(cpu, l, h)

	case 0x3f:
		oopCCF(cpu)

	// LD r1, r2
	case 0x40:
		//cpu.BC.Hi = cpu.BC.Hi
	case 0x41:
		cpu.BC.Hi = cpu.BC.Lo
	case 0x42:
		cpu.BC.Hi = cpu.DE.Hi
	case 0x43:
		cpu.BC.Hi = cpu.DE.Lo
	case 0x44:
		cpu.BC.Hi = cpu.HL.Hi
	case 0x45:
		cpu.BC.Hi = cpu.HL.Lo
	case 0x47:
		cpu.BC.Hi = cpu.AF.Hi
	case 0x48:
		cpu.BC.Lo = cpu.BC.Hi
	case 0x49:
		//cpu.BC.Lo = cpu.BC.Lo
	case 0x4a:
		cpu.BC.Lo = cpu.DE.Hi
	case 0x4b:
		cpu.BC.Lo = cpu.DE.Lo
	case 0x4c:
		cpu.BC.Lo = cpu.HL.Hi
	case 0x4d:
		cpu.BC.Lo = cpu.HL.Lo
	case 0x4f:
		cpu.BC.Lo = cpu.AF.Hi
	case 0x50:
		cpu.DE.Hi = cpu.BC.Hi
	case 0x51:
		cpu.DE.Hi = cpu.BC.Lo
	case 0x52:
		//cpu.DE.Hi = cpu.DE.Hi
	case 0x53:
		cpu.DE.Hi = cpu.DE.Lo
	case 0x54:
		cpu.DE.Hi = cpu.HL.Hi
	case 0x55:
		cpu.DE.Hi = cpu.HL.Lo
	case 0x57:
		cpu.DE.Hi = cpu.AF.Hi
	case 0x58:
		cpu.DE.Lo = cpu.BC.Hi
	case 0x59:
		cpu.DE.Lo = cpu.BC.Lo
	case 0x5a:
		cpu.DE.Lo = cpu.DE.Hi
	case 0x5b:
		//cpu.DE.Lo = cpu.DE.Lo
	case 0x5c:
		cpu.DE.Lo = cpu.HL.Hi
	case 0x5d:
		cpu.DE.Lo = cpu.HL.Lo
	case 0x5f:
		cpu.DE.Lo = cpu.AF.Hi
	case 0x60:
		cpu.HL.Hi = cpu.BC.Hi
	case 0x61:
		cpu.HL.Hi = cpu.BC.Lo
	case 0x62:
		cpu.HL.Hi = cpu.DE.Hi
	case 0x63:
		cpu.HL.Hi = cpu.DE.Lo
	case 0x64:
		//cpu.HL.Hi = cpu.HL.Hi
	case 0x65:
		cpu.HL.Hi = cpu.HL.Lo
	case 0x67:
		cpu.HL.Hi = cpu.AF.Hi
	case 0x68:
		cpu.HL.Lo = cpu.BC.Hi
	case 0x69:
		cpu.HL.Lo = cpu.BC.Lo
	case 0x6a:
		cpu.HL.Lo = cpu.DE.Hi
	case 0x6b:
		cpu.HL.Lo = cpu.DE.Lo
	case 0x6c:
		cpu.HL.Lo = cpu.HL.Hi
	case 0x6d:
		//cpu.HL.Lo = cpu.HL.Lo
	case 0x6f:
		cpu.HL.Lo = cpu.AF.Hi
	case 0x78:
		cpu.AF.Hi = cpu.BC.Hi
	case 0x79:
		cpu.AF.Hi = cpu.BC.Lo
	case 0x7a:
		cpu.AF.Hi = cpu.DE.Hi
	case 0x7b:
		cpu.AF.Hi = cpu.DE.Lo
	case 0x7c:
		cpu.AF.Hi = cpu.HL.Hi
	case 0x7d:
		cpu.AF.Hi = cpu.HL.Lo
	case 0x7f:
		//cpu.AF.Hi = cpu.AF.Hi

	case 0x46:
		xopLDbHLP(cpu)
	case 0x4e:
		xopLDcHLP(cpu)
	case 0x56:
		xopLDdHLP(cpu)
	case 0x5e:
		xopLDeHLP(cpu)
	case 0x66:
		xopLDhHLP(cpu)
	case 0x6e:
		xopLDlHLP(cpu)
	case 0x7e:
		xopLDaHLP(cpu)

	case 0x70:
		xopLDHLPb(cpu)
	case 0x71:
		xopLDHLPc(cpu)
	case 0x72:
		xopLDHLPd(cpu)
	case 0x73:
		xopLDHLPe(cpu)
	case 0x74:
		xopLDHLPh(cpu)
	case 0x75:
		xopLDHLPl(cpu)
	case 0x77:
		xopLDHLPa(cpu)

	case 0x76:
		oopHALT(cpu)

	// ADD A, r
	case 0x80:
		xopADDAb(cpu)
	case 0x81:
		xopADDAc(cpu)
	case 0x82:
		xopADDAd(cpu)
	case 0x83:
		xopADDAe(cpu)
	case 0x84:
		xopADDAh(cpu)
	case 0x85:
		xopADDAl(cpu)
	case 0x87:
		xopADDAa(cpu)

	// ADD A, (HL)
	case 0x86:
		oopADDAHLP(cpu)

	// ADC A, r
	case 0x88:
		xopADCAb(cpu)
	case 0x89:
		xopADCAc(cpu)
	case 0x8a:
		xopADCAd(cpu)
	case 0x8b:
		xopADCAe(cpu)
	case 0x8c:
		xopADCAh(cpu)
	case 0x8d:
		xopADCAl(cpu)
	case 0x8f:
		xopADCAa(cpu)

	case 0x8e:
		oopADCAHLP(cpu)

	// SUB A, r
	case 0x90:
		xopSUBAb(cpu)
	case 0x91:
		xopSUBAc(cpu)
	case 0x92:
		xopSUBAd(cpu)
	case 0x93:
		xopSUBAe(cpu)
	case 0x94:
		xopSUBAh(cpu)
	case 0x95:
		xopSUBAl(cpu)
	case 0x97:
		xopSUBAa(cpu)

	case 0x96:
		oopSUBAHLP(cpu)

	// SBC A, r
	case 0x98:
		xopSBCAb(cpu)
	case 0x99:
		xopSBCAc(cpu)
	case 0x9a:
		xopSBCAd(cpu)
	case 0x9b:
		xopSBCAe(cpu)
	case 0x9c:
		xopSBCAh(cpu)
	case 0x9d:
		xopSBCAl(cpu)
	case 0x9f:
		xopSBCAa(cpu)

	case 0x9e:
		oopSBCAHLP(cpu)

	// ADD r
	case 0xa0:
		xopANDAb(cpu)
	case 0xa1:
		xopANDAc(cpu)
	case 0xa2:
		xopANDAd(cpu)
	case 0xa3:
		xopANDAe(cpu)
	case 0xa4:
		xopANDAh(cpu)
	case 0xa5:
		xopANDAl(cpu)
	case 0xa7:
		xopANDAa(cpu)

	case 0xa6:
		oopANDHLP(cpu)

	// XOR r
	case 0xa8:
		xopXORb(cpu)
	case 0xa9:
		xopXORc(cpu)
	case 0xaa:
		xopXORd(cpu)
	case 0xab:
		xopXORe(cpu)
	case 0xac:
		xopXORh(cpu)
	case 0xad:
		xopXORl(cpu)
	case 0xaf:
		xopXORa(cpu)

	case 0xae:
		oopXORHLP(cpu)

	// OR r
	case 0xb0:
		xopORb(cpu)
	case 0xb1:
		xopORc(cpu)
	case 0xb2:
		xopORd(cpu)
	case 0xb3:
		xopORe(cpu)
	case 0xb4:
		xopORh(cpu)
	case 0xb5:
		xopORl(cpu)
	case 0xb7:
		xopORa(cpu)

	case 0xb6:
		oopORHLP(cpu)

	// CP r
	case 0xb8:
		xopCPb(cpu)
	case 0xb9:
		xopCPc(cpu)
	case 0xba:
		xopCPd(cpu)
	case 0xbb:
		xopCPe(cpu)
	case 0xbc:
		xopCPh(cpu)
	case 0xbd:
		xopCPl(cpu)
	case 0xbf:
		xopCPa(cpu)

	case 0xbe:
		oopCPHLP(cpu)

	case 0xc0:
		xopRETnZ(cpu)
	case 0xc8:
		xopRETfZ(cpu)
	case 0xd0:
		xopRETnC(cpu)
	case 0xd8:
		xopRETfC(cpu)
	case 0xe0:
		xopRETnPV(cpu)
	case 0xe8:
		xopRETfPV(cpu)
	case 0xf0:
		xopRETnS(cpu)
	case 0xf8:
		xopRETfS(cpu)

	case 0xc1:
		xopPOPreg(cpu, &cpu.BC)
	case 0xd1:
		xopPOPreg(cpu, &cpu.DE)
	case 0xe1:
		xopPOPreg(cpu, &cpu.HL)
	case 0xf1:
		xopPOPreg(cpu, &cpu.AF)

	case 0xc2:
		l := f.fetch()
		h := f.fetch()
		xopJPnZnn(cpu, l, h)
	case 0xca:
		l := f.fetch()
		h := f.fetch()
		xopJPfZnn(cpu, l, h)
	case 0xd2:
		l := f.fetch()
		h := f.fetch()
		xopJPnCnn(cpu, l, h)
	case 0xda:
		l := f.fetch()
		h := f.fetch()
		xopJPfCnn(cpu, l, h)
	case 0xe2:
		l := f.fetch()
		h := f.fetch()
		xopJPnPVnn(cpu, l, h)
	case 0xea:
		l := f.fetch()
		h := f.fetch()
		xopJPfPVnn(cpu, l, h)
	case 0xf2:
		l := f.fetch()
		h := f.fetch()
		xopJPnSnn(cpu, l, h)
	case 0xfa:
		l := f.fetch()
		h := f.fetch()
		xopJPfSnn(cpu, l, h)

	case 0xc3:
		l := f.fetch()
		h := f.fetch()
		oopJPnn(cpu, l, h)

	case 0xc4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnZnn(cpu, l, h)
	case 0xcc:
		l := f.fetch()
		h := f.fetch()
		xopCALLfZnn(cpu, l, h)
	case 0xd4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnCnn(cpu, l, h)
	case 0xdc:
		l := f.fetch()
		h := f.fetch()
		xopCALLfCnn(cpu, l, h)
	case 0xe4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnPVnn(cpu, l, h)
	case 0xec:
		l := f.fetch()
		h := f.fetch()
		xopCALLfPVnn(cpu, l, h)
	case 0xf4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnSnn(cpu, l, h)
	case 0xfc:
		l := f.fetch()
		h := f.fetch()
		xopCALLfSnn(cpu, l, h)

	case 0xc5:
		xopPUSHreg(cpu, cpu.BC)
	case 0xd5:
		xopPUSHreg(cpu, cpu.DE)
	case 0xe5:
		xopPUSHreg(cpu, cpu.HL)
	case 0xf5:
		xopPUSHreg(cpu, cpu.AF)

	case 0xc6:
		n := f.fetch()
		oopADDAn(cpu, n)

	case 0xc7:
		xopRST00(cpu)
	case 0xcf:
		xopRST08(cpu)
	case 0xd7:
		xopRST10(cpu)
	case 0xdf:
		xopRST18(cpu)
	case 0xe7:
		xopRST20(cpu)
	case 0xef:
		xopRST28(cpu)
	case 0xf7:
		xopRST30(cpu)
	case 0xff:
		xopRST38(cpu)

	case 0xc9:
		oopRET(cpu)

	case 0xcd:
		l := f.fetch()
		h := f.fetch()
		xopCALLnn(cpu, l, h)

	case 0xce:
		n := f.fetch()
		oopADCAn(cpu, n)

	case 0xd3:
		n := f.fetch()
		oopOUTnPA(cpu, n)

	case 0xd6:
		n := f.fetch()
		oopSUBAn(cpu, n)

	case 0xd9:
		oopEXX(cpu)

	case 0xdb:
		n := f.fetch()
		oopINAnP(cpu, n)

	case 0xde:
		n := f.fetch()
		oopSBCAn(cpu, n)

	case 0xe3:
		oopEXSPPHL(cpu)

	case 0xe6:
		n := f.fetch()
		oopANDn(cpu, n)

	case 0xe9:
		oopJPHLP(cpu)

	case 0xeb:
		oopEXDEHL(cpu)

	case 0xee:
		n := f.fetch()
		oopXORn(cpu, n)

	case 0xf3:
		oopDI(cpu)

	case 0xf6:
		n := f.fetch()
		oopORn(cpu, n)

	case 0xf9:
		oopLDSPHL(cpu)

	case 0xfb:
		oopEI(cpu)

	case 0xfe:
		n := f.fetch()
		oopCPn(cpu, n)

	case 0xcb:
		c1 := f.fetch()
		switch c1 {

		// RLC r / RLC (HL)
		case 0x00:
			xopRLCb(cpu)
		case 0x01:
			xopRLCc(cpu)
		case 0x02:
			xopRLCd(cpu)
		case 0x03:
			xopRLCe(cpu)
		case 0x04:
			xopRLCh(cpu)
		case 0x05:
			xopRLCl(cpu)
		case 0x06:
			xopRLCHLP(cpu)
		case 0x07:
			xopRLCa(cpu)

		// RRC r / RRC (HL)
		case 0x08:
			xopRRCb(cpu)
		case 0x09:
			xopRRCc(cpu)
		case 0x0a:
			xopRRCd(cpu)
		case 0x0b:
			xopRRCe(cpu)
		case 0x0c:
			xopRRCh(cpu)
		case 0x0d:
			xopRRCl(cpu)
		case 0x0e:
			xopRRCHLP(cpu)
		case 0x0f:
			xopRRCa(cpu)

		// RL r / RL (HL)
		case 0x10:
			xopRLb(cpu)
		case 0x11:
			xopRLc(cpu)
		case 0x12:
			xopRLd(cpu)
		case 0x13:
			xopRLe(cpu)
		case 0x14:
			xopRLh(cpu)
		case 0x15:
			xopRLl(cpu)
		case 0x16:
			xopRLHLP(cpu)
		case 0x17:
			xopRLa(cpu)

		// RR r / RR (HL)
		case 0x18:
			xopRRb(cpu)
		case 0x19:
			xopRRc(cpu)
		case 0x1a:
			xopRRd(cpu)
		case 0x1b:
			xopRRe(cpu)
		case 0x1c:
			xopRRh(cpu)
		case 0x1d:
			xopRRl(cpu)
		case 0x1e:
			xopRRHLP(cpu)
		case 0x1f:
			xopRRa(cpu)

		// SLA r / SLA (HL)
		case 0x20:
			xopSLAb(cpu)
		case 0x21:
			xopSLAc(cpu)
		case 0x22:
			xopSLAd(cpu)
		case 0x23:
			xopSLAe(cpu)
		case 0x24:
			xopSLAh(cpu)
		case 0x25:
			xopSLAl(cpu)
		case 0x26:
			xopSLAHLP(cpu)
		case 0x27:
			xopSLAa(cpu)

		// SRA r / SRA (HL)
		case 0x28:
			xopSRAb(cpu)
		case 0x29:
			xopSRAc(cpu)
		case 0x2a:
			xopSRAd(cpu)
		case 0x2b:
			xopSRAe(cpu)
		case 0x2c:
			xopSRAh(cpu)
		case 0x2d:
			xopSRAl(cpu)
		case 0x2e:
			xopSRAHLP(cpu)
		case 0x2f:
			xopSRAa(cpu)

		// SL1 r / SL1 (HL) (undocumented)
		case 0x30:
			xopSL1b(cpu)
		case 0x31:
			xopSL1c(cpu)
		case 0x32:
			xopSL1d(cpu)
		case 0x33:
			xopSL1e(cpu)
		case 0x34:
			xopSL1h(cpu)
		case 0x35:
			xopSL1l(cpu)
		case 0x36:
			xopSL1HLP(cpu)
		case 0x37:
			xopSL1a(cpu)

		// SRL r / SRL (HL)
		case 0x38:
			xopSRLb(cpu)
		case 0x39:
			xopSRLc(cpu)
		case 0x3a:
			xopSRLd(cpu)
		case 0x3b:
			xopSRLe(cpu)
		case 0x3c:
			xopSRLh(cpu)
		case 0x3d:
			xopSRLl(cpu)
		case 0x3e:
			xopSRLHLP(cpu)
		case 0x3f:
			xopSRLa(cpu)

		// BIT 0, r|(HL)
		case 0x40:
			cpu.bitchk8(0, cpu.BC.Hi)
		case 0x41:
			cpu.bitchk8(0, cpu.BC.Lo)
		case 0x42:
			cpu.bitchk8(0, cpu.DE.Hi)
		case 0x43:
			cpu.bitchk8(0, cpu.DE.Lo)
		case 0x44:
			cpu.bitchk8(0, cpu.HL.Hi)
		case 0x45:
			cpu.bitchk8(0, cpu.HL.Lo)
		case 0x46:
			cpu.bitchk8b(0, cpu.Memory.Get(cpu.HL.U16()))
		case 0x47:
			cpu.bitchk8(0, cpu.AF.Hi)

		// BIT 1, r|(HL)
		case 0x48:
			cpu.bitchk8(1, cpu.BC.Hi)
		case 0x49:
			cpu.bitchk8(1, cpu.BC.Lo)
		case 0x4a:
			cpu.bitchk8(1, cpu.DE.Hi)
		case 0x4b:
			cpu.bitchk8(1, cpu.DE.Lo)
		case 0x4c:
			cpu.bitchk8(1, cpu.HL.Hi)
		case 0x4d:
			cpu.bitchk8(1, cpu.HL.Lo)
		case 0x4e:
			cpu.bitchk8b(1, cpu.Memory.Get(cpu.HL.U16()))
		case 0x4f:
			cpu.bitchk8(1, cpu.AF.Hi)

		// BIT 2, r|(HL)
		case 0x50:
			cpu.bitchk8(2, cpu.BC.Hi)
		case 0x51:
			cpu.bitchk8(2, cpu.BC.Lo)
		case 0x52:
			cpu.bitchk8(2, cpu.DE.Hi)
		case 0x53:
			cpu.bitchk8(2, cpu.DE.Lo)
		case 0x54:
			cpu.bitchk8(2, cpu.HL.Hi)
		case 0x55:
			cpu.bitchk8(2, cpu.HL.Lo)
		case 0x56:
			cpu.bitchk8b(2, cpu.Memory.Get(cpu.HL.U16()))
		case 0x57:
			cpu.bitchk8(2, cpu.AF.Hi)

		// BIT 3, r|(HL)
		case 0x58:
			cpu.bitchk8(3, cpu.BC.Hi)
		case 0x59:
			cpu.bitchk8(3, cpu.BC.Lo)
		case 0x5a:
			cpu.bitchk8(3, cpu.DE.Hi)
		case 0x5b:
			cpu.bitchk8(3, cpu.DE.Lo)
		case 0x5c:
			cpu.bitchk8(3, cpu.HL.Hi)
		case 0x5d:
			cpu.bitchk8(3, cpu.HL.Lo)
		case 0x5e:
			cpu.bitchk8b(3, cpu.Memory.Get(cpu.HL.U16()))
		case 0x5f:
			cpu.bitchk8(3, cpu.AF.Hi)

		// BIT 4, r|(HL)
		case 0x60:
			cpu.bitchk8(4, cpu.BC.Hi)
		case 0x61:
			cpu.bitchk8(4, cpu.BC.Lo)
		case 0x62:
			cpu.bitchk8(4, cpu.DE.Hi)
		case 0x63:
			cpu.bitchk8(4, cpu.DE.Lo)
		case 0x64:
			cpu.bitchk8(4, cpu.HL.Hi)
		case 0x65:
			cpu.bitchk8(4, cpu.HL.Lo)
		case 0x66:
			cpu.bitchk8b(4, cpu.Memory.Get(cpu.HL.U16()))
		case 0x67:
			cpu.bitchk8(4, cpu.AF.Hi)

		// BIT 5, r|(HL)
		case 0x68:
			cpu.bitchk8(5, cpu.BC.Hi)
		case 0x69:
			cpu.bitchk8(5, cpu.BC.Lo)
		case 0x6a:
			cpu.bitchk8(5, cpu.DE.Hi)
		case 0x6b:
			cpu.bitchk8(5, cpu.DE.Lo)
		case 0x6c:
			cpu.bitchk8(5, cpu.HL.Hi)
		case 0x6d:
			cpu.bitchk8(5, cpu.HL.Lo)
		case 0x6e:
			cpu.bitchk8b(5, cpu.Memory.Get(cpu.HL.U16()))
		case 0x6f:
			cpu.bitchk8(5, cpu.AF.Hi)

		// BIT 6, r|(HL)
		case 0x70:
			cpu.bitchk8(6, cpu.BC.Hi)
		case 0x71:
			cpu.bitchk8(6, cpu.BC.Lo)
		case 0x72:
			cpu.bitchk8(6, cpu.DE.Hi)
		case 0x73:
			cpu.bitchk8(6, cpu.DE.Lo)
		case 0x74:
			cpu.bitchk8(6, cpu.HL.Hi)
		case 0x75:
			cpu.bitchk8(6, cpu.HL.Lo)
		case 0x76:
			cpu.bitchk8b(6, cpu.Memory.Get(cpu.HL.U16()))
		case 0x77:
			cpu.bitchk8(6, cpu.AF.Hi)

		// BIT 7, r|(HL)
		case 0x78:
			cpu.bitchk8(7, cpu.BC.Hi)
		case 0x79:
			cpu.bitchk8(7, cpu.BC.Lo)
		case 0x7a:
			cpu.bitchk8(7, cpu.DE.Hi)
		case 0x7b:
			cpu.bitchk8(7, cpu.DE.Lo)
		case 0x7c:
			cpu.bitchk8(7, cpu.HL.Hi)
		case 0x7d:
			cpu.bitchk8(7, cpu.HL.Lo)
		case 0x7e:
			cpu.bitchk8b(7, cpu.Memory.Get(cpu.HL.U16()))
		case 0x7f:
			cpu.bitchk8(7, cpu.AF.Hi)

		// RES 0, r|(HL)
		case 0x80:
			cpu.BC.Hi = cpu.bitres8(0, cpu.BC.Hi)
		case 0x81:
			cpu.BC.Lo = cpu.bitres8(0, cpu.BC.Lo)
		case 0x82:
			cpu.DE.Hi = cpu.bitres8(0, cpu.DE.Hi)
		case 0x83:
			cpu.DE.Lo = cpu.bitres8(0, cpu.DE.Lo)
		case 0x84:
			cpu.HL.Hi = cpu.bitres8(0, cpu.HL.Hi)
		case 0x85:
			cpu.HL.Lo = cpu.bitres8(0, cpu.HL.Lo)
		case 0x86:
			xopBITbHLP(cpu, 0)
		case 0x87:
			cpu.AF.Hi = cpu.bitres8(0, cpu.AF.Hi)

		// RES 1, r|(HL)
		case 0x88:
			cpu.BC.Hi = cpu.bitres8(1, cpu.BC.Hi)
		case 0x89:
			cpu.BC.Lo = cpu.bitres8(1, cpu.BC.Lo)
		case 0x8a:
			cpu.DE.Hi = cpu.bitres8(1, cpu.DE.Hi)
		case 0x8b:
			cpu.DE.Lo = cpu.bitres8(1, cpu.DE.Lo)
		case 0x8c:
			cpu.HL.Hi = cpu.bitres8(1, cpu.HL.Hi)
		case 0x8d:
			cpu.HL.Lo = cpu.bitres8(1, cpu.HL.Lo)
		case 0x8e:
			xopBITbHLP(cpu, 1)
		case 0x8f:
			cpu.AF.Hi = cpu.bitres8(1, cpu.AF.Hi)

		// RES 2, r|(HL)
		case 0x90:
			cpu.BC.Hi = cpu.bitres8(2, cpu.BC.Hi)
		case 0x91:
			cpu.BC.Lo = cpu.bitres8(2, cpu.BC.Lo)
		case 0x92:
			cpu.DE.Hi = cpu.bitres8(2, cpu.DE.Hi)
		case 0x93:
			cpu.DE.Lo = cpu.bitres8(2, cpu.DE.Lo)
		case 0x94:
			cpu.HL.Hi = cpu.bitres8(2, cpu.HL.Hi)
		case 0x95:
			cpu.HL.Lo = cpu.bitres8(2, cpu.HL.Lo)
		case 0x96:
			xopBITbHLP(cpu, 2)
		case 0x97:
			cpu.AF.Hi = cpu.bitres8(2, cpu.AF.Hi)

		// RES 3, r|(HL)
		case 0x98:
			cpu.BC.Hi = cpu.bitres8(3, cpu.BC.Hi)
		case 0x99:
			cpu.BC.Lo = cpu.bitres8(3, cpu.BC.Lo)
		case 0x9a:
			cpu.DE.Hi = cpu.bitres8(3, cpu.DE.Hi)
		case 0x9b:
			cpu.DE.Lo = cpu.bitres8(3, cpu.DE.Lo)
		case 0x9c:
			cpu.HL.Hi = cpu.bitres8(3, cpu.HL.Hi)
		case 0x9d:
			cpu.HL.Lo = cpu.bitres8(3, cpu.HL.Lo)
		case 0x9e:
			xopBITbHLP(cpu, 3)
		case 0x9f:
			cpu.AF.Hi = cpu.bitres8(3, cpu.AF.Hi)

		// RES 4, r|(HL)
		case 0xa0:
			cpu.BC.Hi = cpu.bitres8(4, cpu.BC.Hi)
		case 0xa1:
			cpu.BC.Lo = cpu.bitres8(4, cpu.BC.Lo)
		case 0xa2:
			cpu.DE.Hi = cpu.bitres8(4, cpu.DE.Hi)
		case 0xa3:
			cpu.DE.Lo = cpu.bitres8(4, cpu.DE.Lo)
		case 0xa4:
			cpu.HL.Hi = cpu.bitres8(4, cpu.HL.Hi)
		case 0xa5:
			cpu.HL.Lo = cpu.bitres8(4, cpu.HL.Lo)
		case 0xa6:
			xopBITbHLP(cpu, 4)
		case 0xa7:
			cpu.AF.Hi = cpu.bitres8(4, cpu.AF.Hi)

		// RES 5, r|(HL)
		case 0xa8:
			cpu.BC.Hi = cpu.bitres8(5, cpu.BC.Hi)
		case 0xa9:
			cpu.BC.Lo = cpu.bitres8(5, cpu.BC.Lo)
		case 0xaa:
			cpu.DE.Hi = cpu.bitres8(5, cpu.DE.Hi)
		case 0xab:
			cpu.DE.Lo = cpu.bitres8(5, cpu.DE.Lo)
		case 0xac:
			cpu.HL.Hi = cpu.bitres8(5, cpu.HL.Hi)
		case 0xad:
			cpu.HL.Lo = cpu.bitres8(5, cpu.HL.Lo)
		case 0xae:
			xopBITbHLP(cpu, 5)
		case 0xaf:
			cpu.AF.Hi = cpu.bitres8(5, cpu.AF.Hi)

		// RES 6, r|(HL)
		case 0xb0:
			cpu.BC.Hi = cpu.bitres8(6, cpu.BC.Hi)
		case 0xb1:
			cpu.BC.Lo = cpu.bitres8(6, cpu.BC.Lo)
		case 0xb2:
			cpu.DE.Hi = cpu.bitres8(6, cpu.DE.Hi)
		case 0xb3:
			cpu.DE.Lo = cpu.bitres8(6, cpu.DE.Lo)
		case 0xb4:
			cpu.HL.Hi = cpu.bitres8(6, cpu.HL.Hi)
		case 0xb5:
			cpu.HL.Lo = cpu.bitres8(6, cpu.HL.Lo)
		case 0xb6:
			xopBITbHLP(cpu, 6)
		case 0xb7:
			cpu.AF.Hi = cpu.bitres8(6, cpu.AF.Hi)

		// RES 7, r|(HL)
		case 0xb8:
			cpu.BC.Hi = cpu.bitres8(7, cpu.BC.Hi)
		case 0xb9:
			cpu.BC.Lo = cpu.bitres8(7, cpu.BC.Lo)
		case 0xba:
			cpu.DE.Hi = cpu.bitres8(7, cpu.DE.Hi)
		case 0xbb:
			cpu.DE.Lo = cpu.bitres8(7, cpu.DE.Lo)
		case 0xbc:
			cpu.HL.Hi = cpu.bitres8(7, cpu.HL.Hi)
		case 0xbd:
			cpu.HL.Lo = cpu.bitres8(7, cpu.HL.Lo)
		case 0xbe:
			xopBITbHLP(cpu, 7)
		case 0xbf:
			cpu.AF.Hi = cpu.bitres8(7, cpu.AF.Hi)

		// SET 0, r|(HL)
		case 0xc0:
			cpu.BC.Hi = cpu.bitset8(0, cpu.BC.Hi)
		case 0xc1:
			cpu.BC.Lo = cpu.bitset8(0, cpu.BC.Lo)
		case 0xc2:
			cpu.DE.Hi = cpu.bitset8(0, cpu.DE.Hi)
		case 0xc3:
			cpu.DE.Lo = cpu.bitset8(0, cpu.DE.Lo)
		case 0xc4:
			cpu.HL.Hi = cpu.bitset8(0, cpu.HL.Hi)
		case 0xc5:
			cpu.HL.Lo = cpu.bitset8(0, cpu.HL.Lo)
		case 0xc6:
			xopSETbHLP(cpu, 0)
		case 0xc7:
			cpu.AF.Hi = cpu.bitset8(0, cpu.AF.Hi)

		// SET 1, r|(HL)
		case 0xc8:
			cpu.BC.Hi = cpu.bitset8(1, cpu.BC.Hi)
		case 0xc9:
			cpu.BC.Lo = cpu.bitset8(1, cpu.BC.Lo)
		case 0xca:
			cpu.DE.Hi = cpu.bitset8(1, cpu.DE.Hi)
		case 0xcb:
			cpu.DE.Lo = cpu.bitset8(1, cpu.DE.Lo)
		case 0xcc:
			cpu.HL.Hi = cpu.bitset8(1, cpu.HL.Hi)
		case 0xcd:
			cpu.HL.Lo = cpu.bitset8(1, cpu.HL.Lo)
		case 0xce:
			xopSETbHLP(cpu, 1)
		case 0xcf:
			cpu.AF.Hi = cpu.bitset8(1, cpu.AF.Hi)

		// SET 2, r|(HL)
		case 0xd0:
			cpu.BC.Hi = cpu.bitset8(2, cpu.BC.Hi)
		case 0xd1:
			cpu.BC.Lo = cpu.bitset8(2, cpu.BC.Lo)
		case 0xd2:
			cpu.DE.Hi = cpu.bitset8(2, cpu.DE.Hi)
		case 0xd3:
			cpu.DE.Lo = cpu.bitset8(2, cpu.DE.Lo)
		case 0xd4:
			cpu.HL.Hi = cpu.bitset8(2, cpu.HL.Hi)
		case 0xd5:
			cpu.HL.Lo = cpu.bitset8(2, cpu.HL.Lo)
		case 0xd6:
			xopSETbHLP(cpu, 2)
		case 0xd7:
			cpu.AF.Hi = cpu.bitset8(2, cpu.AF.Hi)

		// SET 3, r|(HL)
		case 0xd8:
			cpu.BC.Hi = cpu.bitset8(3, cpu.BC.Hi)
		case 0xd9:
			cpu.BC.Lo = cpu.bitset8(3, cpu.BC.Lo)
		case 0xda:
			cpu.DE.Hi = cpu.bitset8(3, cpu.DE.Hi)
		case 0xdb:
			cpu.DE.Lo = cpu.bitset8(3, cpu.DE.Lo)
		case 0xdc:
			cpu.HL.Hi = cpu.bitset8(3, cpu.HL.Hi)
		case 0xdd:
			cpu.HL.Lo = cpu.bitset8(3, cpu.HL.Lo)
		case 0xde:
			xopSETbHLP(cpu, 3)
		case 0xdf:
			cpu.AF.Hi = cpu.bitset8(3, cpu.AF.Hi)

		// SET 4, r|(HL)
		case 0xe0:
			cpu.BC.Hi = cpu.bitset8(4, cpu.BC.Hi)
		case 0xe1:
			cpu.BC.Lo = cpu.bitset8(4, cpu.BC.Lo)
		case 0xe2:
			cpu.DE.Hi = cpu.bitset8(4, cpu.DE.Hi)
		case 0xe3:
			cpu.DE.Lo = cpu.bitset8(4, cpu.DE.Lo)
		case 0xe4:
			cpu.HL.Hi = cpu.bitset8(4, cpu.HL.Hi)
		case 0xe5:
			cpu.HL.Lo = cpu.bitset8(4, cpu.HL.Lo)
		case 0xe6:
			xopSETbHLP(cpu, 4)
		case 0xe7:
			cpu.AF.Hi = cpu.bitset8(4, cpu.AF.Hi)

		// SET 5, r|(HL)
		case 0xe8:
			cpu.BC.Hi = cpu.bitset8(5, cpu.BC.Hi)
		case 0xe9:
			cpu.BC.Lo = cpu.bitset8(5, cpu.BC.Lo)
		case 0xea:
			cpu.DE.Hi = cpu.bitset8(5, cpu.DE.Hi)
		case 0xeb:
			cpu.DE.Lo = cpu.bitset8(5, cpu.DE.Lo)
		case 0xec:
			cpu.HL.Hi = cpu.bitset8(5, cpu.HL.Hi)
		case 0xed:
			cpu.HL.Lo = cpu.bitset8(5, cpu.HL.Lo)
		case 0xee:
			xopSETbHLP(cpu, 5)
		case 0xef:
			cpu.AF.Hi = cpu.bitset8(5, cpu.AF.Hi)

		// SET 6, r|(HL)
		case 0xf0:
			cpu.BC.Hi = cpu.bitset8(6, cpu.BC.Hi)
		case 0xf1:
			cpu.BC.Lo = cpu.bitset8(6, cpu.BC.Lo)
		case 0xf2:
			cpu.DE.Hi = cpu.bitset8(6, cpu.DE.Hi)
		case 0xf3:
			cpu.DE.Lo = cpu.bitset8(6, cpu.DE.Lo)
		case 0xf4:
			cpu.HL.Hi = cpu.bitset8(6, cpu.HL.Hi)
		case 0xf5:
			cpu.HL.Lo = cpu.bitset8(6, cpu.HL.Lo)
		case 0xf6:
			xopSETbHLP(cpu, 6)
		case 0xf7:
			cpu.AF.Hi = cpu.bitset8(6, cpu.AF.Hi)

		// SET 7, r|(HL)
		case 0xf8:
			cpu.BC.Hi = cpu.bitset8(7, cpu.BC.Hi)
		case 0xf9:
			cpu.BC.Lo = cpu.bitset8(7, cpu.BC.Lo)
		case 0xfa:
			cpu.DE.Hi = cpu.bitset8(7, cpu.DE.Hi)
		case 0xfb:
			cpu.DE.Lo = cpu.bitset8(7, cpu.DE.Lo)
		case 0xfc:
			cpu.HL.Hi = cpu.bitset8(7, cpu.HL.Hi)
		case 0xfd:
			cpu.HL.Lo = cpu.bitset8(7, cpu.HL.Lo)
		case 0xfe:
			xopSETbHLP(cpu, 7)
		case 0xff:
			cpu.AF.Hi = cpu.bitset8(7, cpu.AF.Hi)

		default:
			cpu.invalidCode(c0, c1)
		}

	case 0xdd:
		c1 := f.fetch()
		switch c1 {

		// ADD IX, pp
		case 0x09:
			xopADDIXbc(cpu)
		case 0x19:
			xopADDIXde(cpu)
		case 0x29:
			xopADDIXix(cpu)
		case 0x39:
			xopADDIXsp(cpu)

		case 0x21:
			l := f.fetch()
			h := f.fetch()
			oopLDIXnn(cpu, l, h)

		case 0x22:
			l := f.fetch()
			h := f.fetch()
			oopLDnnPIX(cpu, l, h)

		case 0x23:
			oopINCIX(cpu)

		case 0x24:
			oopINCIXH(cpu)

		case 0x25:
			oopDECIXH(cpu)

		case 0x26:
			n := f.fetch()
			oopLDIXHn(cpu, n)

		case 0x2a:
			l := f.fetch()
			h := f.fetch()
			oopLDIXnnP(cpu, l, h)

		case 0x2b:
			oopDECIX(cpu)

		case 0x2c:
			oopINCIXL(cpu)

		case 0x2d:
			oopDECIXL(cpu)

		case 0x2e:
			n := f.fetch()
			oopLDIXLn(cpu, n)

		case 0x34:
			d := f.fetch()
			oopINCIXdP(cpu, d)

		case 0x35:
			d := f.fetch()
			oopDECIXdP(cpu, d)

		case 0x36:
			d := f.fetch()
			n := f.fetch()
			oopLDIXdPn(cpu, d, n)

		// LD rx1, rx2
		case 0x40:
			//cpu.BC.Hi = cpu.BC.Hi
		case 0x41:
			cpu.BC.Hi = cpu.BC.Lo
		case 0x42:
			cpu.BC.Hi = cpu.DE.Hi
		case 0x43:
			cpu.BC.Hi = cpu.DE.Lo
		case 0x44:
			cpu.BC.Hi = uint8(cpu.IX >> 8)
		case 0x45:
			cpu.BC.Hi = uint8(cpu.IX)
		case 0x47:
			cpu.BC.Hi = cpu.AF.Hi
		case 0x48:
			cpu.BC.Lo = cpu.BC.Hi
		case 0x49:
			//cpu.BC.Lo = cpu.BC.Lo
		case 0x4a:
			cpu.BC.Lo = cpu.DE.Hi
		case 0x4b:
			cpu.BC.Lo = cpu.DE.Lo
		case 0x4c:
			cpu.BC.Lo = uint8(cpu.IX >> 8)
		case 0x4d:
			cpu.BC.Lo = uint8(cpu.IX)
		case 0x4f:
			cpu.BC.Lo = cpu.AF.Hi
		case 0x50:
			cpu.DE.Hi = cpu.BC.Hi
		case 0x51:
			cpu.DE.Hi = cpu.BC.Lo
		case 0x52:
			//cpu.DE.Hi = cpu.DE.Hi
		case 0x53:
			cpu.DE.Hi = cpu.DE.Lo
		case 0x54:
			cpu.DE.Hi = uint8(cpu.IX >> 8)
		case 0x55:
			cpu.DE.Hi = uint8(cpu.IX)
		case 0x57:
			cpu.DE.Hi = cpu.AF.Hi
		case 0x58:
			cpu.DE.Lo = cpu.BC.Hi
		case 0x59:
			cpu.DE.Lo = cpu.BC.Lo
		case 0x5a:
			cpu.DE.Lo = cpu.DE.Hi
		case 0x5b:
			//cpu.DE.Lo = cpu.DE.Lo
		case 0x5c:
			cpu.DE.Lo = uint8(cpu.IX >> 8)
		case 0x5d:
			cpu.DE.Lo = uint8(cpu.IX)
		case 0x5f:
			cpu.DE.Lo = cpu.AF.Hi
		case 0x60:
			cpu.IX = uint16(cpu.BC.Hi)<<8 | cpu.IX&0x00ff
		case 0x61:
			cpu.IX = uint16(cpu.BC.Lo)<<8 | cpu.IX&0x00ff
		case 0x62:
			cpu.IX = uint16(cpu.DE.Hi)<<8 | cpu.IX&0x00ff
		case 0x63:
			cpu.IX = uint16(cpu.DE.Lo)<<8 | cpu.IX&0x00ff
		case 0x64:
			//cpu.IX = uint16(uint8(cpu.IX >> 8))<<8 | cpu.IX&0x00ff
		case 0x65:
			cpu.IX = uint16(uint8(cpu.IX))<<8 | cpu.IX&0x00ff
		case 0x67:
			cpu.IX = uint16(cpu.AF.Hi)<<8 | cpu.IX&0x00ff
		case 0x68:
			cpu.IX = uint16(cpu.BC.Hi) | cpu.IX&0xff00
		case 0x69:
			cpu.IX = uint16(cpu.BC.Lo) | cpu.IX&0xff00
		case 0x6a:
			cpu.IX = uint16(cpu.DE.Hi) | cpu.IX&0xff00
		case 0x6b:
			cpu.IX = uint16(cpu.DE.Lo) | cpu.IX&0xff00
		case 0x6c:
			cpu.IX = uint16(uint8(cpu.IX>>8)) | cpu.IX&0xff00
		case 0x6d:
			//cpu.IX = uint16(uint8(cpu.IX)) | cpu.IX&0xff00
		case 0x6f:
			cpu.IX = uint16(cpu.AF.Hi) | cpu.IX&0xff00
		case 0x78:
			cpu.AF.Hi = cpu.BC.Hi
		case 0x79:
			cpu.AF.Hi = cpu.BC.Lo
		case 0x7a:
			cpu.AF.Hi = cpu.DE.Hi
		case 0x7b:
			cpu.AF.Hi = cpu.DE.Lo
		case 0x7c:
			cpu.AF.Hi = uint8(cpu.IX >> 8)
		case 0x7d:
			cpu.AF.Hi = uint8(cpu.IX)
		case 0x7f:
			//cpu.AF.Hi = cpu.AF.Hi

		// LD r, (IX+d)
		case 0x46:
			d := f.fetch()
			xopLDbIXdP(cpu, d)
		case 0x4e:
			d := f.fetch()
			xopLDcIXdP(cpu, d)
		case 0x56:
			d := f.fetch()
			xopLDdIXdP(cpu, d)
		case 0x5e:
			d := f.fetch()
			xopLDeIXdP(cpu, d)
		case 0x66:
			d := f.fetch()
			xopLDhIXdP(cpu, d)
		case 0x6e:
			d := f.fetch()
			xopLDlIXdP(cpu, d)
		case 0x7e:
			d := f.fetch()
			xopLDaIXdP(cpu, d)

		// LD (IX+d), r
		case 0x70:
			d := f.fetch()
			xopLDIXdPb(cpu, d)
		case 0x71:
			d := f.fetch()
			xopLDIXdPc(cpu, d)
		case 0x72:
			d := f.fetch()
			xopLDIXdPd(cpu, d)
		case 0x73:
			d := f.fetch()
			xopLDIXdPe(cpu, d)
		case 0x74:
			d := f.fetch()
			xopLDIXdPh(cpu, d)
		case 0x75:
			d := f.fetch()
			xopLDIXdPl(cpu, d)
		case 0x77:
			d := f.fetch()
			xopLDIXdPa(cpu, d)

		// ADD A, rx (undocumented)
		case 0x80:
			xopADDAb(cpu)
		case 0x81:
			xopADDAc(cpu)
		case 0x82:
			xopADDAd(cpu)
		case 0x83:
			xopADDAe(cpu)
		case 0x84:
			xopADDAixh(cpu)
		case 0x85:
			xopADDAixl(cpu)
		case 0x87:
			xopADDAa(cpu)

		case 0x86:
			d := f.fetch()
			oopADDAIXdP(cpu, d)

		// ADC A, rx
		case 0x88:
			xopADCAb(cpu)
		case 0x89:
			xopADCAc(cpu)
		case 0x8a:
			xopADCAd(cpu)
		case 0x8b:
			xopADCAe(cpu)
		case 0x8c:
			xopADCAixh(cpu)
		case 0x8d:
			xopADCAixl(cpu)
		case 0x8f:
			xopADCAa(cpu)

		case 0x8e:
			d := f.fetch()
			oopADCAIXdP(cpu, d)

		// SUB A, rx
		case 0x90:
			xopSUBAb(cpu)
		case 0x91:
			xopSUBAc(cpu)
		case 0x92:
			xopSUBAd(cpu)
		case 0x93:
			xopSUBAe(cpu)
		case 0x94:
			xopSUBAixh(cpu)
		case 0x95:
			xopSUBAixl(cpu)
		case 0x97:
			xopSUBAa(cpu)

		case 0x96:
			d := f.fetch()
			oopSUBAIXdP(cpu, d)

		// SBC A, rx
		case 0x98:
			xopSBCAb(cpu)
		case 0x99:
			xopSBCAc(cpu)
		case 0x9a:
			xopSBCAd(cpu)
		case 0x9b:
			xopSBCAe(cpu)
		case 0x9c:
			xopSBCAixh(cpu)
		case 0x9d:
			xopSBCAixl(cpu)
		case 0x9f:
			xopSBCAa(cpu)

		case 0x9e:
			d := f.fetch()
			oopSBCAIXdP(cpu, d)

		// ADD rx
		case 0xa0:
			xopANDAb(cpu)
		case 0xa1:
			xopANDAc(cpu)
		case 0xa2:
			xopANDAd(cpu)
		case 0xa3:
			xopANDAe(cpu)
		case 0xa4:
			xopANDixh(cpu)
		case 0xa5:
			xopANDixl(cpu)
		case 0xa7:
			xopANDAa(cpu)

		// AND (IX+d)
		case 0xa6:
			d := f.fetch()
			oopANDIXdP(cpu, d)

		// XOR rx
		case 0xa8:
			xopXORb(cpu)
		case 0xa9:
			xopXORc(cpu)
		case 0xaa:
			xopXORd(cpu)
		case 0xab:
			xopXORe(cpu)
		case 0xac:
			xopXORixh(cpu)
		case 0xad:
			xopXORixl(cpu)
		case 0xaf:
			xopXORa(cpu)

		case 0xae:
			d := f.fetch()
			oopXORIXdP(cpu, d)

		// OR rx
		case 0xb0:
			xopORb(cpu)
		case 0xb1:
			xopORc(cpu)
		case 0xb2:
			xopORd(cpu)
		case 0xb3:
			xopORe(cpu)
		case 0xb4:
			xopORixh(cpu)
		case 0xb5:
			xopORixl(cpu)
		case 0xb7:
			xopORa(cpu)

		case 0xb6:
			d := f.fetch()
			oopORIXdP(cpu, d)

		// CP rx
		case 0xb8:
			xopCPb(cpu)
		case 0xb9:
			xopCPc(cpu)
		case 0xba:
			xopCPd(cpu)
		case 0xbb:
			xopCPe(cpu)
		case 0xbc:
			xopCPixh(cpu)
		case 0xbd:
			xopCPixl(cpu)
		case 0xbf:
			xopCPa(cpu)

		case 0xbe:
			d := f.fetch()
			oopCPIXdP(cpu, d)

		case 0xe1:
			oopPOPIX(cpu)

		case 0xe3:
			oopEXSPPIX(cpu)

		case 0xe5:
			oopPUSHIX(cpu)

		case 0xe9:
			oopJPIXP(cpu)

		case 0xf9:
			oopLDSPIX(cpu)

		case 0xcb:
			d := f.fetch()
			c3 := f.fetch()
			switch c3 {

			case 0x06:
				oopRLCIXdP(cpu, d)

			case 0x0e:
				oopRRCIXdP(cpu, d)

			case 0x16:
				oopRLIXdP(cpu, d)

			case 0x1e:
				oopRRIXdP(cpu, d)

			case 0x26:
				oopSLAIXdP(cpu, d)

			case 0x2e:
				oopSRAIXdP(cpu, d)

			case 0x36:
				oopSL1IXdP(cpu, d)

			case 0x3e:
				oopSRLIXdP(cpu, d)

			// BIT b, (IX+d)
			case 0x46:
				xopBITbIXdP(cpu, 0, d)
			case 0x4e:
				xopBITbIXdP(cpu, 1, d)
			case 0x56:
				xopBITbIXdP(cpu, 2, d)
			case 0x5e:
				xopBITbIXdP(cpu, 3, d)
			case 0x66:
				xopBITbIXdP(cpu, 4, d)
			case 0x6e:
				xopBITbIXdP(cpu, 5, d)
			case 0x76:
				xopBITbIXdP(cpu, 6, d)
			case 0x7e:
				xopBITbIXdP(cpu, 7, d)

			// RES b, (IX+d)
			case 0x86:
				xopRESbIXdP(cpu, 0, d)
			case 0x8e:
				xopRESbIXdP(cpu, 1, d)
			case 0x96:
				xopRESbIXdP(cpu, 2, d)
			case 0x9e:
				xopRESbIXdP(cpu, 3, d)
			case 0xa6:
				xopRESbIXdP(cpu, 4, d)
			case 0xae:
				xopRESbIXdP(cpu, 5, d)
			case 0xb6:
				xopRESbIXdP(cpu, 6, d)
			case 0xbe:
				xopRESbIXdP(cpu, 7, d)

			// SET b, (IX+d)
			case 0xc6:
				xopSETbIXdP(cpu, 0, d)
			case 0xce:
				xopSETbIXdP(cpu, 1, d)
			case 0xd6:
				xopSETbIXdP(cpu, 2, d)
			case 0xde:
				xopSETbIXdP(cpu, 3, d)
			case 0xe6:
				xopSETbIXdP(cpu, 4, d)
			case 0xee:
				xopSETbIXdP(cpu, 5, d)
			case 0xf6:
				xopSETbIXdP(cpu, 6, d)
			case 0xfe:
				xopSETbIXdP(cpu, 7, d)

			default:
				cpu.invalidCode(c0, c1, d, c3)
			}
		default:
			cpu.invalidCode(c0, c1)
		}

	case 0xed:
		c1 := f.fetch()
		switch c1 {

		// IN r, (C)
		// FIXME: IN r[6], (C) to apply flags only.
		case 0x40:
			xopINbCP(cpu)
		case 0x48:
			xopINcCP(cpu)
		case 0x50:
			xopINdCP(cpu)
		case 0x58:
			xopINeCP(cpu)
		case 0x60:
			xopINhCP(cpu)
		case 0x68:
			xopINlCP(cpu)
		case 0x78:
			xopINaCP(cpu)

		// OUT (C), r
		case 0x41:
			xopOUTCPb(cpu)
		case 0x49:
			xopOUTCPc(cpu)
		case 0x51:
			xopOUTCPd(cpu)
		case 0x59:
			xopOUTCPe(cpu)
		case 0x61:
			xopOUTCPh(cpu)
		case 0x69:
			xopOUTCPl(cpu)
		case 0x79:
			xopOUTCPa(cpu)

		// SBC HL, ss
		case 0x42:
			xopSBCHLbc(cpu)
		case 0x52:
			xopSBCHLde(cpu)
		case 0x62:
			xopSBCHLhl(cpu)
		case 0x72:
			xopSBCHLsp(cpu)

		// LD (nn), dd
		case 0x43:
			l := f.fetch()
			h := f.fetch()
			xopLDnnPbc(cpu, l, h)
		case 0x53:
			l := f.fetch()
			h := f.fetch()
			xopLDnnPde(cpu, l, h)
		case 0x63:
			l := f.fetch()
			h := f.fetch()
			xopLDnnPhl(cpu, l, h)
		case 0x73:
			l := f.fetch()
			h := f.fetch()
			xopLDnnPsp(cpu, l, h)

		case 0x44:
			oopNEG(cpu)

		case 0x45:
			oopRETN(cpu)

		case 0x46:
			oopIM0(cpu)

		case 0x47:
			oopLDIA(cpu)

		// ADC HL, ss
		case 0x4a:
			xopADCHLbc(cpu)
		case 0x5a:
			xopADCHLde(cpu)
		case 0x6a:
			xopADCHLhl(cpu)
		case 0x7a:
			xopADCHLsp(cpu)

		// LD dd, (nn)
		case 0x4b:
			l := f.fetch()
			h := f.fetch()
			xopLDbcnnP(cpu, l, h)
		case 0x5b:
			l := f.fetch()
			h := f.fetch()
			xopLDdennP(cpu, l, h)
		case 0x6b:
			l := f.fetch()
			h := f.fetch()
			xopLDhlnnP(cpu, l, h)
		case 0x7b:
			l := f.fetch()
			h := f.fetch()
			xopLDspnnP(cpu, l, h)

		case 0x4d:
			oopRETI(cpu)

		case 0x4f:
			oopLDRA(cpu)

		case 0x56:
			oopIM1(cpu)

		case 0x57:
			oopLDAI(cpu)

		case 0x5e:
			oopIM2(cpu)

		case 0x5f:
			oopLDAR(cpu)

		case 0x67:
			oopRRD(cpu)

		case 0x6f:
			oopRLD(cpu)

		case 0xa0:
			oopLDI(cpu)

		case 0xa1:
			oopCPI(cpu)

		case 0xa2:
			oopINI(cpu)

		case 0xa3:
			oopOUTI(cpu)

		case 0xa8:
			oopLDD(cpu)

		case 0xa9:
			oopCPD(cpu)

		case 0xaa:
			oopIND(cpu)

		case 0xab:
			oopOUTD(cpu)

		case 0xb0:
			oopLDIR(cpu)

		case 0xb1:
			oopCPIR(cpu)

		case 0xb2:
			oopINIR(cpu)

		case 0xb3:
			oopOTIR(cpu)

		case 0xb8:
			oopLDDR(cpu)

		case 0xb9:
			oopCPDR(cpu)

		case 0xba:
			oopINDR(cpu)

		case 0xbb:
			oopOTDR(cpu)

		default:
			cpu.invalidCode(c0, c1)
		}

	case 0xfd:
		c1 := f.fetch()
		switch c1 {

		// ADD IY, pp
		case 0x09:
			xopADDIYbc(cpu)
		case 0x19:
			xopADDIYde(cpu)
		case 0x29:
			xopADDIYiy(cpu)
		case 0x39:
			xopADDIYsp(cpu)

		case 0x21:
			l := f.fetch()
			h := f.fetch()
			oopLDIYnn(cpu, l, h)

		case 0x22:
			l := f.fetch()
			h := f.fetch()
			oopLDnnPIY(cpu, l, h)

		case 0x23:
			oopINCIY(cpu)

		case 0x24:
			oopINCIYH(cpu)

		case 0x25:
			oopDECIYH(cpu)

		case 0x26:
			n := f.fetch()
			oopLDIYHn(cpu, n)

		case 0x2a:
			l := f.fetch()
			h := f.fetch()
			oopLDIYnnP(cpu, l, h)

		case 0x2b:
			oopDECIY(cpu)

		case 0x2c:
			oopINCIYL(cpu)

		case 0x2d:
			oopDECIYL(cpu)

		case 0x2e:
			n := f.fetch()
			oopLDIYLn(cpu, n)

		case 0x34:
			d := f.fetch()
			oopINCIYdP(cpu, d)

		case 0x35:
			d := f.fetch()
			oopDECIYdP(cpu, d)

		case 0x36:
			d := f.fetch()
			n := f.fetch()
			oopLDIYdPn(cpu, d, n)

		// LD ry1, ry2
		case 0x40:
			//cpu.BC.Hi = cpu.BC.Hi
		case 0x41:
			cpu.BC.Hi = cpu.BC.Lo
		case 0x42:
			cpu.BC.Hi = cpu.DE.Hi
		case 0x43:
			cpu.BC.Hi = cpu.DE.Lo
		case 0x44:
			cpu.BC.Hi = uint8(cpu.IY >> 8)
		case 0x45:
			cpu.BC.Hi = uint8(cpu.IY)
		case 0x47:
			cpu.BC.Hi = cpu.AF.Hi
		case 0x48:
			cpu.BC.Lo = cpu.BC.Hi
		case 0x49:
			//cpu.BC.Lo = cpu.BC.Lo
		case 0x4a:
			cpu.BC.Lo = cpu.DE.Hi
		case 0x4b:
			cpu.BC.Lo = cpu.DE.Lo
		case 0x4c:
			cpu.BC.Lo = uint8(cpu.IY >> 8)
		case 0x4d:
			cpu.BC.Lo = uint8(cpu.IY)
		case 0x4f:
			cpu.BC.Lo = cpu.AF.Hi
		case 0x50:
			cpu.DE.Hi = cpu.BC.Hi
		case 0x51:
			cpu.DE.Hi = cpu.BC.Lo
		case 0x52:
			//cpu.DE.Hi = cpu.DE.Hi
		case 0x53:
			cpu.DE.Hi = cpu.DE.Lo
		case 0x54:
			cpu.DE.Hi = uint8(cpu.IY >> 8)
		case 0x55:
			cpu.DE.Hi = uint8(cpu.IY)
		case 0x57:
			cpu.DE.Hi = cpu.AF.Hi
		case 0x58:
			cpu.DE.Lo = cpu.BC.Hi
		case 0x59:
			cpu.DE.Lo = cpu.BC.Lo
		case 0x5a:
			cpu.DE.Lo = cpu.DE.Hi
		case 0x5b:
			//cpu.DE.Lo = cpu.DE.Lo
		case 0x5c:
			cpu.DE.Lo = uint8(cpu.IY >> 8)
		case 0x5d:
			cpu.DE.Lo = uint8(cpu.IY)
		case 0x5f:
			cpu.DE.Lo = cpu.AF.Hi
		case 0x60:
			cpu.IY = uint16(cpu.BC.Hi)<<8 | cpu.IY&0x00ff
		case 0x61:
			cpu.IY = uint16(cpu.BC.Lo)<<8 | cpu.IY&0x00ff
		case 0x62:
			cpu.IY = uint16(cpu.DE.Hi)<<8 | cpu.IY&0x00ff
		case 0x63:
			cpu.IY = uint16(cpu.DE.Lo)<<8 | cpu.IY&0x00ff
		case 0x64:
			//cpu.IY = uint16(uint8(cpu.IY >> 8))<<8 | cpu.IY&0x00ff
		case 0x65:
			cpu.IY = uint16(uint8(cpu.IY))<<8 | cpu.IY&0x00ff
		case 0x67:
			cpu.IY = uint16(cpu.AF.Hi)<<8 | cpu.IY&0x00ff
		case 0x68:
			cpu.IY = uint16(cpu.BC.Hi) | cpu.IY&0xff00
		case 0x69:
			cpu.IY = uint16(cpu.BC.Lo) | cpu.IY&0xff00
		case 0x6a:
			cpu.IY = uint16(cpu.DE.Hi) | cpu.IY&0xff00
		case 0x6b:
			cpu.IY = uint16(cpu.DE.Lo) | cpu.IY&0xff00
		case 0x6c:
			cpu.IY = uint16(uint8(cpu.IY>>8)) | cpu.IY&0xff00
		case 0x6d:
			//cpu.IY = uint16(uint8(cpu.IY)) | cpu.IY&0xff00
		case 0x6f:
			cpu.IY = uint16(cpu.AF.Hi) | cpu.IY&0xff00
		case 0x78:
			cpu.AF.Hi = cpu.BC.Hi
		case 0x79:
			cpu.AF.Hi = cpu.BC.Lo
		case 0x7a:
			cpu.AF.Hi = cpu.DE.Hi
		case 0x7b:
			cpu.AF.Hi = cpu.DE.Lo
		case 0x7c:
			cpu.AF.Hi = uint8(cpu.IY >> 8)
		case 0x7d:
			cpu.AF.Hi = uint8(cpu.IY)
		case 0x7f:
			//cpu.AF.Hi = cpu.AF.Hi

		// LD r, (IY+d)
		case 0x46:
			d := f.fetch()
			xopLDbIYdP(cpu, d)
		case 0x4e:
			d := f.fetch()
			xopLDcIYdP(cpu, d)
		case 0x56:
			d := f.fetch()
			xopLDdIYdP(cpu, d)
		case 0x5e:
			d := f.fetch()
			xopLDeIYdP(cpu, d)
		case 0x66:
			d := f.fetch()
			xopLDhIYdP(cpu, d)
		case 0x6e:
			d := f.fetch()
			xopLDlIYdP(cpu, d)
		case 0x7e:
			d := f.fetch()
			xopLDaIYdP(cpu, d)

		// LD (IY+d), r
		case 0x70:
			d := f.fetch()
			xopLDIYdPb(cpu, d)
		case 0x71:
			d := f.fetch()
			xopLDIYdPc(cpu, d)
		case 0x72:
			d := f.fetch()
			xopLDIYdPd(cpu, d)
		case 0x73:
			d := f.fetch()
			xopLDIYdPe(cpu, d)
		case 0x74:
			d := f.fetch()
			xopLDIYdPh(cpu, d)
		case 0x75:
			d := f.fetch()
			xopLDIYdPl(cpu, d)
		case 0x77:
			d := f.fetch()
			xopLDIYdPa(cpu, d)

		// ADD A, ry (undocumented)
		case 0x80:
			xopADDAb(cpu)
		case 0x81:
			xopADDAc(cpu)
		case 0x82:
			xopADDAd(cpu)
		case 0x83:
			xopADDAe(cpu)
		case 0x84:
			xopADDAiyh(cpu)
		case 0x85:
			xopADDAiyl(cpu)
		case 0x87:
			xopADDAa(cpu)

		case 0x86:
			d := f.fetch()
			oopADDAIYdP(cpu, d)

		// ADC A, ry
		case 0x88:
			xopADCAb(cpu)
		case 0x89:
			xopADCAc(cpu)
		case 0x8a:
			xopADCAd(cpu)
		case 0x8b:
			xopADCAe(cpu)
		case 0x8c:
			xopADCAiyh(cpu)
		case 0x8d:
			xopADCAiyl(cpu)
		case 0x8f:
			xopADCAa(cpu)

		case 0x8e:
			d := f.fetch()
			oopADCAIYdP(cpu, d)

		// SUB A, ry
		case 0x90:
			xopSUBAb(cpu)
		case 0x91:
			xopSUBAc(cpu)
		case 0x92:
			xopSUBAd(cpu)
		case 0x93:
			xopSUBAe(cpu)
		case 0x94:
			xopSUBAiyh(cpu)
		case 0x95:
			xopSUBAiyl(cpu)
		case 0x97:
			xopSUBAa(cpu)

		case 0x96:
			d := f.fetch()
			oopSUBAIYdP(cpu, d)

		// SBC A, ry
		case 0x98:
			xopSBCAb(cpu)
		case 0x99:
			xopSBCAc(cpu)
		case 0x9a:
			xopSBCAd(cpu)
		case 0x9b:
			xopSBCAe(cpu)
		case 0x9c:
			xopSBCAiyh(cpu)
		case 0x9d:
			xopSBCAiyl(cpu)
		case 0x9f:
			xopSBCAa(cpu)

		case 0x9e:
			d := f.fetch()
			oopSBCAIYdP(cpu, d)

		// ADD rx
		case 0xa0:
			xopANDAb(cpu)
		case 0xa1:
			xopANDAc(cpu)
		case 0xa2:
			xopANDAd(cpu)
		case 0xa3:
			xopANDAe(cpu)
		case 0xa4:
			xopANDiyh(cpu)
		case 0xa5:
			xopANDiyl(cpu)
		case 0xa7:
			xopANDAa(cpu)

		// AND (IY+d)
		case 0xa6:
			d := f.fetch()
			oopANDIYdP(cpu, d)

		// XOR rx
		case 0xa8:
			xopXORb(cpu)
		case 0xa9:
			xopXORc(cpu)
		case 0xaa:
			xopXORd(cpu)
		case 0xab:
			xopXORe(cpu)
		case 0xac:
			xopXORiyh(cpu)
		case 0xad:
			xopXORiyl(cpu)
		case 0xaf:
			xopXORa(cpu)

		case 0xae:
			d := f.fetch()
			oopXORIYdP(cpu, d)

		// OR ry
		case 0xb0:
			xopORb(cpu)
		case 0xb1:
			xopORc(cpu)
		case 0xb2:
			xopORd(cpu)
		case 0xb3:
			xopORe(cpu)
		case 0xb4:
			xopORiyh(cpu)
		case 0xb5:
			xopORiyl(cpu)
		case 0xb7:
			xopORa(cpu)

		case 0xb6:
			d := f.fetch()
			oopORIYdP(cpu, d)

		// CP ry
		case 0xb8:
			xopCPb(cpu)
		case 0xb9:
			xopCPc(cpu)
		case 0xba:
			xopCPd(cpu)
		case 0xbb:
			xopCPe(cpu)
		case 0xbc:
			xopCPiyh(cpu)
		case 0xbd:
			xopCPiyl(cpu)
		case 0xbf:
			xopCPa(cpu)

		case 0xbe:
			d := f.fetch()
			oopCPIYdP(cpu, d)

		case 0xe1:
			oopPOPIY(cpu)

		case 0xe3:
			oopEXSPPIY(cpu)

		case 0xe5:
			oopPUSHIY(cpu)

		case 0xe9:
			oopJPIYP(cpu)

		case 0xf9:
			oopLDSPIY(cpu)

		case 0xcb:
			d := f.fetch()
			c3 := f.fetch()
			switch c3 {

			case 0x06:
				oopRLCIYdP(cpu, d)

			case 0x0e:
				oopRRCIYdP(cpu, d)

			case 0x16:
				oopRLIYdP(cpu, d)

			case 0x1e:
				oopRRIYdP(cpu, d)

			case 0x26:
				oopSLAIYdP(cpu, d)

			case 0x2e:
				oopSRAIYdP(cpu, d)

			case 0x36:
				oopSL1IYdP(cpu, d)

			case 0x3e:
				oopSRLIYdP(cpu, d)

			// BIT b, (IY+d)
			case 0x46:
				xopBITbIYdP(cpu, 0, d)
			case 0x4e:
				xopBITbIYdP(cpu, 1, d)
			case 0x56:
				xopBITbIYdP(cpu, 2, d)
			case 0x5e:
				xopBITbIYdP(cpu, 3, d)
			case 0x66:
				xopBITbIYdP(cpu, 4, d)
			case 0x6e:
				xopBITbIYdP(cpu, 5, d)
			case 0x76:
				xopBITbIYdP(cpu, 6, d)
			case 0x7e:
				xopBITbIYdP(cpu, 7, d)

			// RES b, (IY+d)
			case 0x86:
				xopRESbIYdP(cpu, 0, d)
			case 0x8e:
				xopRESbIYdP(cpu, 1, d)
			case 0x96:
				xopRESbIYdP(cpu, 2, d)
			case 0x9e:
				xopRESbIYdP(cpu, 3, d)
			case 0xa6:
				xopRESbIYdP(cpu, 4, d)
			case 0xae:
				xopRESbIYdP(cpu, 5, d)
			case 0xb6:
				xopRESbIYdP(cpu, 6, d)
			case 0xbe:
				xopRESbIYdP(cpu, 7, d)

			// SET b, (IY+d)
			case 0xc6:
				xopSETbIYdP(cpu, 0, d)
			case 0xce:
				xopSETbIYdP(cpu, 1, d)
			case 0xd6:
				xopSETbIYdP(cpu, 2, d)
			case 0xde:
				xopSETbIYdP(cpu, 3, d)
			case 0xe6:
				xopSETbIYdP(cpu, 4, d)
			case 0xee:
				xopSETbIYdP(cpu, 5, d)
			case 0xf6:
				xopSETbIYdP(cpu, 6, d)
			case 0xfe:
				xopSETbIYdP(cpu, 7, d)

			default:
				cpu.invalidCode(c0, c1, d, c3)
			}
		default:
			cpu.invalidCode(c0, c1)
		}

	default:
		cpu.invalidCode(c0)
	}
}
