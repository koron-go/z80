package z80

func decodeExec(cpu *CPU, f fetcher) error {
	buf := cpu.decodeBuf[:4]
	switch f.fetch() {
	case 0x00:
		oopNOP(cpu)
		return nil

	case 0x01:
		l := f.fetch()
		h := f.fetch()
		xopLDbcnn(cpu, l, h)
		return nil
	case 0x11:
		l := f.fetch()
		h := f.fetch()
		xopLDdenn(cpu, l, h)
		return nil
	case 0x21:
		l := f.fetch()
		h := f.fetch()
		xopLDhlnn(cpu, l, h)
		return nil
	case 0x31:
		l := f.fetch()
		h := f.fetch()
		xopLDspnn(cpu, l, h)
		return nil

	case 0x02:
		oopLDBCPA(cpu)
		return nil

	case 0x03:
		xopINCbc(cpu)
		return nil
	case 0x13:
		xopINCde(cpu)
		return nil
	case 0x23:
		xopINChl(cpu)
		return nil
	case 0x33:
		xopINCsp(cpu)
		return nil

	case 0x04:
		xopINCb(cpu)
		return nil
	case 0x0c:
		xopINCc(cpu)
		return nil
	case 0x14:
		xopINCd(cpu)
		return nil
	case 0x1c:
		xopINCe(cpu)
		return nil
	case 0x24:
		xopINCh(cpu)
		return nil
	case 0x2c:
		xopINCl(cpu)
		return nil
	case 0x3c:
		xopINCa(cpu)
		return nil

	case 0x05:
		xopDECb(cpu)
		return nil
	case 0x0d:
		xopDECc(cpu)
		return nil
	case 0x15:
		xopDECd(cpu)
		return nil
	case 0x1d:
		xopDECe(cpu)
		return nil
	case 0x25:
		xopDECh(cpu)
		return nil
	case 0x2d:
		xopDECl(cpu)
		return nil
	case 0x3d:
		xopDECa(cpu)
		return nil

	case 0x06:
		n := f.fetch()
		xopLDbn(cpu, n)
		return nil
	case 0x0e:
		n := f.fetch()
		xopLDcn(cpu, n)
		return nil
	case 0x16:
		n := f.fetch()
		xopLDdn(cpu, n)
		return nil
	case 0x1e:
		n := f.fetch()
		xopLDen(cpu, n)
		return nil
	case 0x26:
		n := f.fetch()
		xopLDhn(cpu, n)
		return nil
	case 0x2e:
		n := f.fetch()
		xopLDln(cpu, n)
		return nil
	case 0x3e:
		n := f.fetch()
		xopLDan(cpu, n)
		return nil

	case 0x07:
		oopRLCA(cpu)
		return nil

	case 0x08:
		oopEXAFAF(cpu)
		return nil

	case 0x09:
		xopADDHLbc(cpu)
		return nil
	case 0x19:
		xopADDHLde(cpu)
		return nil
	case 0x29:
		xopADDHLhl(cpu)
		return nil
	case 0x39:
		xopADDHLsp(cpu)
		return nil

	case 0x0a:
		oopLDABCP(cpu)
		return nil

	case 0x0b:
		xopDECbc(cpu)
		return nil
	case 0x1b:
		xopDECde(cpu)
		return nil
	case 0x2b:
		xopDEChl(cpu)
		return nil
	case 0x3b:
		xopDECsp(cpu)
		return nil

	case 0x0f:
		oopRRCA(cpu)
		return nil

	case 0x10:
		off := f.fetch()
		oopDJNZe(cpu, off)
		return nil

	case 0x12:
		oopLDDEPA(cpu)
		return nil

	case 0x17:
		oopRLA(cpu)
		return nil

	case 0x18:
		off := f.fetch()
		oopJRe(cpu, off)
		return nil

	case 0x1a:
		oopLDADEP(cpu)
		return nil

	case 0x1f:
		oopRRA(cpu)
		return nil

	case 0x20:
		off := f.fetch()
		oopJRNZe(cpu, off)
		return nil

	case 0x22:
		l := f.fetch()
		h := f.fetch()
		oopLDnnPHL(cpu, l, h)
		return nil

	case 0x27:
		oopDAA(cpu)
		return nil

	case 0x28:
		off := f.fetch()
		oopJRZe(cpu, off)
		return nil

	case 0x2a:
		l := f.fetch()
		h := f.fetch()
		oopLDHLnnP(cpu, l, h)
		return nil

	case 0x2f:
		oopCPL(cpu)
		return nil

	case 0x30:
		off := f.fetch()
		oopJRNCe(cpu, off)
		return nil

	case 0x32:
		l := f.fetch()
		h := f.fetch()
		oopLDnnPA(cpu, l, h)
		return nil

	case 0x34:
		oopINCHLP(cpu)
		return nil

	case 0x35:
		oopDECHLP(cpu)
		return nil

	case 0x36:
		n := f.fetch()
		oopLDHLPn(cpu, n)
		return nil

	case 0x37:
		oopSCF(cpu)
		return nil

	case 0x38:
		off := f.fetch()
		oopJRCe(cpu, off)
		return nil

	case 0x3a:
		l := f.fetch()
		h := f.fetch()
		oopLDAnnP(cpu, l, h)
		return nil

	case 0x3f:
		oopCCF(cpu)
		return nil

	// LD r1, r2
	case 0x40:
		//cpu.BC.Hi = cpu.BC.Hi
		return nil
	case 0x41:
		cpu.BC.Hi = cpu.BC.Lo
		return nil
	case 0x42:
		cpu.BC.Hi = cpu.DE.Hi
		return nil
	case 0x43:
		cpu.BC.Hi = cpu.DE.Lo
		return nil
	case 0x44:
		cpu.BC.Hi = cpu.HL.Hi
		return nil
	case 0x45:
		cpu.BC.Hi = cpu.HL.Lo
		return nil
	case 0x47:
		cpu.BC.Hi = cpu.AF.Hi
		return nil
	case 0x48:
		cpu.BC.Lo = cpu.BC.Hi
		return nil
	case 0x49:
		//cpu.BC.Lo = cpu.BC.Lo
		return nil
	case 0x4a:
		cpu.BC.Lo = cpu.DE.Hi
		return nil
	case 0x4b:
		cpu.BC.Lo = cpu.DE.Lo
		return nil
	case 0x4c:
		cpu.BC.Lo = cpu.HL.Hi
		return nil
	case 0x4d:
		cpu.BC.Lo = cpu.HL.Lo
		return nil
	case 0x4f:
		cpu.BC.Lo = cpu.AF.Hi
		return nil
	case 0x50:
		cpu.DE.Hi = cpu.BC.Hi
		return nil
	case 0x51:
		cpu.DE.Hi = cpu.BC.Lo
		return nil
	case 0x52:
		//cpu.DE.Hi = cpu.DE.Hi
		return nil
	case 0x53:
		cpu.DE.Hi = cpu.DE.Lo
		return nil
	case 0x54:
		cpu.DE.Hi = cpu.HL.Hi
		return nil
	case 0x55:
		cpu.DE.Hi = cpu.HL.Lo
		return nil
	case 0x57:
		cpu.DE.Hi = cpu.AF.Hi
		return nil
	case 0x58:
		cpu.DE.Lo = cpu.BC.Hi
		return nil
	case 0x59:
		cpu.DE.Lo = cpu.BC.Lo
		return nil
	case 0x5a:
		cpu.DE.Lo = cpu.DE.Hi
		return nil
	case 0x5b:
		//cpu.DE.Lo = cpu.DE.Lo
		return nil
	case 0x5c:
		cpu.DE.Lo = cpu.HL.Hi
		return nil
	case 0x5d:
		cpu.DE.Lo = cpu.HL.Lo
		return nil
	case 0x5f:
		cpu.DE.Lo = cpu.AF.Hi
		return nil
	case 0x60:
		cpu.HL.Hi = cpu.BC.Hi
		return nil
	case 0x61:
		cpu.HL.Hi = cpu.BC.Lo
		return nil
	case 0x62:
		cpu.HL.Hi = cpu.DE.Hi
		return nil
	case 0x63:
		cpu.HL.Hi = cpu.DE.Lo
		return nil
	case 0x64:
		//cpu.HL.Hi = cpu.HL.Hi
		return nil
	case 0x65:
		cpu.HL.Hi = cpu.HL.Lo
		return nil
	case 0x67:
		cpu.HL.Hi = cpu.AF.Hi
		return nil
	case 0x68:
		cpu.HL.Lo = cpu.BC.Hi
		return nil
	case 0x69:
		cpu.HL.Lo = cpu.BC.Lo
		return nil
	case 0x6a:
		cpu.HL.Lo = cpu.DE.Hi
		return nil
	case 0x6b:
		cpu.HL.Lo = cpu.DE.Lo
		return nil
	case 0x6c:
		cpu.HL.Lo = cpu.HL.Hi
		return nil
	case 0x6d:
		//cpu.HL.Lo = cpu.HL.Lo
		return nil
	case 0x6f:
		cpu.HL.Lo = cpu.AF.Hi
		return nil
	case 0x78:
		cpu.AF.Hi = cpu.BC.Hi
		return nil
	case 0x79:
		cpu.AF.Hi = cpu.BC.Lo
		return nil
	case 0x7a:
		cpu.AF.Hi = cpu.DE.Hi
		return nil
	case 0x7b:
		cpu.AF.Hi = cpu.DE.Lo
		return nil
	case 0x7c:
		cpu.AF.Hi = cpu.HL.Hi
		return nil
	case 0x7d:
		cpu.AF.Hi = cpu.HL.Lo
		return nil
	case 0x7f:
		//cpu.AF.Hi = cpu.AF.Hi
		return nil

	case 0x46:
		xopLDbchHLP(cpu)
		return nil
	case 0x4e:
		xopLDbclHLP(cpu)
		return nil
	case 0x56:
		xopLDdehHLP(cpu)
		return nil
	case 0x5e:
		xopLDdelHLP(cpu)
		return nil
	case 0x66:
		xopLDhlhHLP(cpu)
		return nil
	case 0x6e:
		xopLDhllHLP(cpu)
		return nil
	case 0x7e:
		xopLDafhHLP(cpu)
		return nil

	case 0x70:
		xopLDHLPb(cpu)
		return nil
	case 0x71:
		xopLDHLPc(cpu)
		return nil
	case 0x72:
		xopLDHLPd(cpu)
		return nil
	case 0x73:
		xopLDHLPe(cpu)
		return nil
	case 0x74:
		xopLDHLPh(cpu)
		return nil
	case 0x75:
		xopLDHLPl(cpu)
		return nil
	case 0x77:
		xopLDHLPa(cpu)
		return nil

	case 0x76:
		oopHALT(cpu)
		return nil

	// ADD A, r
	case 0x80:
		xopADDAb(cpu)
		return nil
	case 0x81:
		xopADDAc(cpu)
		return nil
	case 0x82:
		xopADDAd(cpu)
		return nil
	case 0x83:
		xopADDAe(cpu)
		return nil
	case 0x84:
		xopADDAh(cpu)
		return nil
	case 0x85:
		xopADDAl(cpu)
		return nil
	case 0x87:
		xopADDAa(cpu)
		return nil

	// ADD A, (HL)
	case 0x86:
		oopADDAHLP(cpu)
		return nil

	// ADC A, r
	case 0x88:
		xopADCAb(cpu)
		return nil
	case 0x89:
		xopADCAc(cpu)
		return nil
	case 0x8a:
		xopADCAd(cpu)
		return nil
	case 0x8b:
		xopADCAe(cpu)
		return nil
	case 0x8c:
		xopADCAh(cpu)
		return nil
	case 0x8d:
		xopADCAl(cpu)
		return nil
	case 0x8f:
		xopADCAa(cpu)
		return nil

	case 0x8e:
		oopADCAHLP(cpu)
		return nil

	// SUB A, r
	case 0x90:
		xopSUBAb(cpu)
		return nil
	case 0x91:
		xopSUBAc(cpu)
		return nil
	case 0x92:
		xopSUBAd(cpu)
		return nil
	case 0x93:
		xopSUBAe(cpu)
		return nil
	case 0x94:
		xopSUBAh(cpu)
		return nil
	case 0x95:
		xopSUBAl(cpu)
		return nil
	case 0x97:
		xopSUBAa(cpu)
		return nil

	case 0x96:
		oopSUBAHLP(cpu)
		return nil

	// SBC A, r
	case 0x98:
		xopSBCAb(cpu)
		return nil
	case 0x99:
		xopSBCAc(cpu)
		return nil
	case 0x9a:
		xopSBCAd(cpu)
		return nil
	case 0x9b:
		xopSBCAe(cpu)
		return nil
	case 0x9c:
		xopSBCAh(cpu)
		return nil
	case 0x9d:
		xopSBCAl(cpu)
		return nil
	case 0x9f:
		xopSBCAa(cpu)
		return nil

	case 0x9e:
		oopSBCAHLP(cpu)
		return nil

	// ADD r
	case 0xa0:
		xopANDAb(cpu)
		return nil
	case 0xa1:
		xopANDAc(cpu)
		return nil
	case 0xa2:
		xopANDAd(cpu)
		return nil
	case 0xa3:
		xopANDAe(cpu)
		return nil
	case 0xa4:
		xopANDAh(cpu)
		return nil
	case 0xa5:
		xopANDAl(cpu)
		return nil
	case 0xa7:
		xopANDAa(cpu)
		return nil

	case 0xa6:
		oopANDHLP(cpu)
		return nil

	// XOR r
	case 0xa8:
		xopXORb(cpu)
		return nil
	case 0xa9:
		xopXORc(cpu)
		return nil
	case 0xaa:
		xopXORd(cpu)
		return nil
	case 0xab:
		xopXORe(cpu)
		return nil
	case 0xac:
		xopXORh(cpu)
		return nil
	case 0xad:
		xopXORl(cpu)
		return nil
	case 0xaf:
		xopXORa(cpu)
		return nil

	case 0xae:
		oopXORHLP(cpu)
		return nil

	// OR r
	case 0xb0:
		xopORb(cpu)
		return nil
	case 0xb1:
		xopORc(cpu)
		return nil
	case 0xb2:
		xopORd(cpu)
		return nil
	case 0xb3:
		xopORe(cpu)
		return nil
	case 0xb4:
		xopORh(cpu)
		return nil
	case 0xb5:
		xopORl(cpu)
		return nil
	case 0xb7:
		xopORa(cpu)
		return nil

	case 0xb6:
		oopORHLP(cpu)
		return nil

	// CP r
	case 0xb8:
		xopCPb(cpu)
		return nil
	case 0xb9:
		xopCPc(cpu)
		return nil
	case 0xba:
		xopCPd(cpu)
		return nil
	case 0xbb:
		xopCPe(cpu)
		return nil
	case 0xbc:
		xopCPh(cpu)
		return nil
	case 0xbd:
		xopCPl(cpu)
		return nil
	case 0xbf:
		xopCPa(cpu)
		return nil

	case 0xbe:
		oopCPHLP(cpu)
		return nil

	case 0xc0:
		xopRETnZ(cpu)
		return nil
	case 0xc8:
		xopRETfZ(cpu)
		return nil
	case 0xd0:
		xopRETnC(cpu)
		return nil
	case 0xd8:
		xopRETfC(cpu)
		return nil
	case 0xe0:
		xopRETnPV(cpu)
		return nil
	case 0xe8:
		xopRETfPV(cpu)
		return nil
	case 0xf0:
		xopRETnS(cpu)
		return nil
	case 0xf8:
		xopRETfS(cpu)
		return nil

	case 0xc1:
		xopPOPbc(cpu)
		return nil
	case 0xd1:
		xopPOPde(cpu)
		return nil
	case 0xe1:
		xopPOPhl(cpu)
		return nil
	case 0xf1:
		xopPOPaf(cpu)
		return nil

	case 0xc2:
		l := f.fetch()
		h := f.fetch()
		xopJPnZnn(cpu, l, h)
		return nil
	case 0xca:
		l := f.fetch()
		h := f.fetch()
		xopJPfZnn(cpu, l, h)
		return nil
	case 0xd2:
		l := f.fetch()
		h := f.fetch()
		xopJPnCnn(cpu, l, h)
		return nil
	case 0xda:
		l := f.fetch()
		h := f.fetch()
		xopJPfCnn(cpu, l, h)
		return nil
	case 0xe2:
		l := f.fetch()
		h := f.fetch()
		xopJPnPVnn(cpu, l, h)
		return nil
	case 0xea:
		l := f.fetch()
		h := f.fetch()
		xopJPfPVnn(cpu, l, h)
		return nil
	case 0xf2:
		l := f.fetch()
		h := f.fetch()
		xopJPnSnn(cpu, l, h)
		return nil
	case 0xfa:
		l := f.fetch()
		h := f.fetch()
		xopJPfSnn(cpu, l, h)
		return nil

	case 0xc3:
		l := f.fetch()
		h := f.fetch()
		oopJPnn(cpu, l, h)
		return nil

	case 0xc4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnZnn(cpu, l, h)
		return nil
	case 0xcc:
		l := f.fetch()
		h := f.fetch()
		xopCALLfZnn(cpu, l, h)
		return nil
	case 0xd4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnCnn(cpu, l, h)
		return nil
	case 0xdc:
		l := f.fetch()
		h := f.fetch()
		xopCALLfCnn(cpu, l, h)
		return nil
	case 0xe4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnPVnn(cpu, l, h)
		return nil
	case 0xec:
		l := f.fetch()
		h := f.fetch()
		xopCALLfPVnn(cpu, l, h)
		return nil
	case 0xf4:
		l := f.fetch()
		h := f.fetch()
		xopCALLnSnn(cpu, l, h)
		return nil
	case 0xfc:
		l := f.fetch()
		h := f.fetch()
		xopCALLfSnn(cpu, l, h)
		return nil

	case 0xc5:
		xopPUSHbc(cpu)
		return nil
	case 0xd5:
		xopPUSHde(cpu)
		return nil
	case 0xe5:
		xopPUSHhl(cpu)
		return nil
	case 0xf5:
		xopPUSHaf(cpu)
		return nil

	case 0xc6:
		n := f.fetch()
		oopADDAn(cpu, n)
		return nil

	case 0xc7:
		xopRST00(cpu)
		return nil
	case 0xcf:
		xopRST08(cpu)
		return nil
	case 0xd7:
		xopRST10(cpu)
		return nil
	case 0xdf:
		xopRST18(cpu)
		return nil
	case 0xe7:
		xopRST20(cpu)
		return nil
	case 0xef:
		xopRST28(cpu)
		return nil
	case 0xf7:
		xopRST30(cpu)
		return nil
	case 0xff:
		xopRST38(cpu)
		return nil

	case 0xc9:
		oopRET(cpu)
		return nil

	case 0xcd:
		l := f.fetch()
		h := f.fetch()
		xopCALLnn(cpu, l, h)
		return nil

	case 0xce:
		n := f.fetch()
		oopADCAn(cpu, n)
		return nil

	case 0xd3:
		n := f.fetch()
		oopOUTnPA(cpu, n)
		return nil

	case 0xd6:
		n := f.fetch()
		oopSUBAn(cpu, n)
		return nil

	case 0xd9:
		oopEXX(cpu)
		return nil

	case 0xdb:
		n := f.fetch()
		oopINAnP(cpu, n)
		return nil

	case 0xde:
		n := f.fetch()
		oopSBCAn(cpu, n)
		return nil

	case 0xe3:
		oopEXSPPHL(cpu)
		return nil

	case 0xe6:
		n := f.fetch()
		oopANDn(cpu, n)
		return nil

	case 0xe9:
		oopJPHLP(cpu)
		return nil

	case 0xeb:
		oopEXDEHL(cpu)
		return nil

	case 0xee:
		n := f.fetch()
		oopXORn(cpu, n)
		return nil

	case 0xf3:
		oopDI(cpu)
		return nil

	case 0xf6:
		n := f.fetch()
		oopORn(cpu, n)
		return nil

	case 0xf9:
		oopLDSPHL(cpu)
		return nil

	case 0xfb:
		oopEI(cpu)
		return nil

	case 0xfe:
		n := f.fetch()
		oopCPn(cpu, n)
		return nil

	case 0xcb:
		switch f.fetch() {

		// RLC r / RLC (HL)
		case 0x00:
			xopRLCb(cpu)
			return nil
		case 0x01:
			xopRLCc(cpu)
			return nil
		case 0x02:
			xopRLCd(cpu)
			return nil
		case 0x03:
			xopRLCe(cpu)
			return nil
		case 0x04:
			xopRLCh(cpu)
			return nil
		case 0x05:
			xopRLCl(cpu)
			return nil
		case 0x06:
			xopRLCHLP(cpu)
			return nil
		case 0x07:
			xopRLCa(cpu)
			return nil

		// RRC r / RRC (HL)
		case 0x08:
			xopRRCb(cpu)
			return nil
		case 0x09:
			xopRRCc(cpu)
			return nil
		case 0x0a:
			xopRRCd(cpu)
			return nil
		case 0x0b:
			xopRRCe(cpu)
			return nil
		case 0x0c:
			xopRRCh(cpu)
			return nil
		case 0x0d:
			xopRRCl(cpu)
			return nil
		case 0x0e:
			xopRRCHLP(cpu)
			return nil
		case 0x0f:
			xopRRCa(cpu)
			return nil

		// RL r / RL (HL)
		case 0x10:
			xopRLb(cpu)
			return nil
		case 0x11:
			xopRLc(cpu)
			return nil
		case 0x12:
			xopRLd(cpu)
			return nil
		case 0x13:
			xopRLe(cpu)
			return nil
		case 0x14:
			xopRLh(cpu)
			return nil
		case 0x15:
			xopRLl(cpu)
			return nil
		case 0x16:
			xopRLHLP(cpu)
			return nil
		case 0x17:
			xopRLa(cpu)
			return nil

		// RR r / RR (HL)
		case 0x18:
			xopRRb(cpu)
			return nil
		case 0x19:
			xopRRc(cpu)
			return nil
		case 0x1a:
			xopRRd(cpu)
			return nil
		case 0x1b:
			xopRRe(cpu)
			return nil
		case 0x1c:
			xopRRh(cpu)
			return nil
		case 0x1d:
			xopRRl(cpu)
			return nil
		case 0x1e:
			xopRRHLP(cpu)
			return nil
		case 0x1f:
			xopRRa(cpu)
			return nil

		// SLA r / SLA (HL)
		case 0x20:
			xopSLAb(cpu)
			return nil
		case 0x21:
			xopSLAc(cpu)
			return nil
		case 0x22:
			xopSLAd(cpu)
			return nil
		case 0x23:
			xopSLAe(cpu)
			return nil
		case 0x24:
			xopSLAh(cpu)
			return nil
		case 0x25:
			xopSLAl(cpu)
			return nil
		case 0x26:
			xopSLAHLP(cpu)
			return nil
		case 0x27:
			xopSLAa(cpu)
			return nil

		// SRA r / SRA (HL)
		case 0x28:
			xopSRAb(cpu)
			return nil
		case 0x29:
			xopSRAc(cpu)
			return nil
		case 0x2a:
			xopSRAd(cpu)
			return nil
		case 0x2b:
			xopSRAe(cpu)
			return nil
		case 0x2c:
			xopSRAh(cpu)
			return nil
		case 0x2d:
			xopSRAl(cpu)
			return nil
		case 0x2e:
			xopSRAHLP(cpu)
			return nil
		case 0x2f:
			xopSRAa(cpu)
			return nil

		// SL1 r / SL1 (HL) (undocumented)
		case 0x30:
			xopSL1b(cpu)
			return nil
		case 0x31:
			xopSL1c(cpu)
			return nil
		case 0x32:
			xopSL1d(cpu)
			return nil
		case 0x33:
			xopSL1e(cpu)
			return nil
		case 0x34:
			xopSL1h(cpu)
			return nil
		case 0x35:
			xopSL1l(cpu)
			return nil
		case 0x36:
			xopSL1HLP(cpu)
			return nil
		case 0x37:
			xopSL1a(cpu)
			return nil

		// SRL r / SRL (HL)
		case 0x38:
			xopSRLb(cpu)
			return nil
		case 0x39:
			xopSRLc(cpu)
			return nil
		case 0x3a:
			xopSRLd(cpu)
			return nil
		case 0x3b:
			xopSRLe(cpu)
			return nil
		case 0x3c:
			xopSRLh(cpu)
			return nil
		case 0x3d:
			xopSRLl(cpu)
			return nil
		case 0x3e:
			xopSRLHLP(cpu)
			return nil
		case 0x3f:
			xopSRLa(cpu)
			return nil

		// BIT 0, r|(HL)
		case 0x40:
			cpu.bitchk8(0, cpu.BC.Hi)
			return nil
		case 0x41:
			cpu.bitchk8(0, cpu.BC.Lo)
			return nil
		case 0x42:
			cpu.bitchk8(0, cpu.DE.Hi)
			return nil
		case 0x43:
			cpu.bitchk8(0, cpu.DE.Lo)
			return nil
		case 0x44:
			cpu.bitchk8(0, cpu.HL.Hi)
			return nil
		case 0x45:
			cpu.bitchk8(0, cpu.HL.Lo)
			return nil
		case 0x46:
			cpu.bitchk8(0, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x47:
			cpu.bitchk8(0, cpu.AF.Hi)
			return nil

		// BIT 1, r|(HL)
		case 0x48:
			cpu.bitchk8(1, cpu.BC.Hi)
			return nil
		case 0x49:
			cpu.bitchk8(1, cpu.BC.Lo)
			return nil
		case 0x4a:
			cpu.bitchk8(1, cpu.DE.Hi)
			return nil
		case 0x4b:
			cpu.bitchk8(1, cpu.DE.Lo)
			return nil
		case 0x4c:
			cpu.bitchk8(1, cpu.HL.Hi)
			return nil
		case 0x4d:
			cpu.bitchk8(1, cpu.HL.Lo)
			return nil
		case 0x4e:
			cpu.bitchk8(1, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x4f:
			cpu.bitchk8(1, cpu.AF.Hi)
			return nil

		// BIT 2, r|(HL)
		case 0x50:
			cpu.bitchk8(2, cpu.BC.Hi)
			return nil
		case 0x51:
			cpu.bitchk8(2, cpu.BC.Lo)
			return nil
		case 0x52:
			cpu.bitchk8(2, cpu.DE.Hi)
			return nil
		case 0x53:
			cpu.bitchk8(2, cpu.DE.Lo)
			return nil
		case 0x54:
			cpu.bitchk8(2, cpu.HL.Hi)
			return nil
		case 0x55:
			cpu.bitchk8(2, cpu.HL.Lo)
			return nil
		case 0x56:
			cpu.bitchk8(2, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x57:
			cpu.bitchk8(2, cpu.AF.Hi)
			return nil

		// BIT 3, r|(HL)
		case 0x58:
			cpu.bitchk8(3, cpu.BC.Hi)
			return nil
		case 0x59:
			cpu.bitchk8(3, cpu.BC.Lo)
			return nil
		case 0x5a:
			cpu.bitchk8(3, cpu.DE.Hi)
			return nil
		case 0x5b:
			cpu.bitchk8(3, cpu.DE.Lo)
			return nil
		case 0x5c:
			cpu.bitchk8(3, cpu.HL.Hi)
			return nil
		case 0x5d:
			cpu.bitchk8(3, cpu.HL.Lo)
			return nil
		case 0x5e:
			cpu.bitchk8(3, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x5f:
			cpu.bitchk8(3, cpu.AF.Hi)
			return nil

		// BIT 4, r|(HL)
		case 0x60:
			cpu.bitchk8(4, cpu.BC.Hi)
			return nil
		case 0x61:
			cpu.bitchk8(4, cpu.BC.Lo)
			return nil
		case 0x62:
			cpu.bitchk8(4, cpu.DE.Hi)
			return nil
		case 0x63:
			cpu.bitchk8(4, cpu.DE.Lo)
			return nil
		case 0x64:
			cpu.bitchk8(4, cpu.HL.Hi)
			return nil
		case 0x65:
			cpu.bitchk8(4, cpu.HL.Lo)
			return nil
		case 0x66:
			cpu.bitchk8(4, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x67:
			cpu.bitchk8(4, cpu.AF.Hi)
			return nil

		// BIT 5, r|(HL)
		case 0x68:
			cpu.bitchk8(5, cpu.BC.Hi)
			return nil
		case 0x69:
			cpu.bitchk8(5, cpu.BC.Lo)
			return nil
		case 0x6a:
			cpu.bitchk8(5, cpu.DE.Hi)
			return nil
		case 0x6b:
			cpu.bitchk8(5, cpu.DE.Lo)
			return nil
		case 0x6c:
			cpu.bitchk8(5, cpu.HL.Hi)
			return nil
		case 0x6d:
			cpu.bitchk8(5, cpu.HL.Lo)
			return nil
		case 0x6e:
			cpu.bitchk8(5, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x6f:
			cpu.bitchk8(5, cpu.AF.Hi)
			return nil

		// BIT 6, r|(HL)
		case 0x70:
			cpu.bitchk8(6, cpu.BC.Hi)
			return nil
		case 0x71:
			cpu.bitchk8(6, cpu.BC.Lo)
			return nil
		case 0x72:
			cpu.bitchk8(6, cpu.DE.Hi)
			return nil
		case 0x73:
			cpu.bitchk8(6, cpu.DE.Lo)
			return nil
		case 0x74:
			cpu.bitchk8(6, cpu.HL.Hi)
			return nil
		case 0x75:
			cpu.bitchk8(6, cpu.HL.Lo)
			return nil
		case 0x76:
			cpu.bitchk8(6, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x77:
			cpu.bitchk8(6, cpu.AF.Hi)
			return nil

		// BIT 7, r|(HL)
		case 0x78:
			cpu.bitchk8(7, cpu.BC.Hi)
			return nil
		case 0x79:
			cpu.bitchk8(7, cpu.BC.Lo)
			return nil
		case 0x7a:
			cpu.bitchk8(7, cpu.DE.Hi)
			return nil
		case 0x7b:
			cpu.bitchk8(7, cpu.DE.Lo)
			return nil
		case 0x7c:
			cpu.bitchk8(7, cpu.HL.Hi)
			return nil
		case 0x7d:
			cpu.bitchk8(7, cpu.HL.Lo)
			return nil
		case 0x7e:
			cpu.bitchk8(7, cpu.Memory.Get(cpu.HL.U16()))
			return nil
		case 0x7f:
			cpu.bitchk8(7, cpu.AF.Hi)
			return nil

		// RES 0, r|(HL)
		case 0x80:
			cpu.BC.Hi = cpu.bitres8(0, cpu.BC.Hi)
			return nil
		case 0x81:
			cpu.BC.Lo = cpu.bitres8(0, cpu.BC.Lo)
			return nil
		case 0x82:
			cpu.DE.Hi = cpu.bitres8(0, cpu.DE.Hi)
			return nil
		case 0x83:
			cpu.DE.Lo = cpu.bitres8(0, cpu.DE.Lo)
			return nil
		case 0x84:
			cpu.HL.Hi = cpu.bitres8(0, cpu.HL.Hi)
			return nil
		case 0x85:
			cpu.HL.Lo = cpu.bitres8(0, cpu.HL.Lo)
			return nil
		case 0x86:
			xopBITbHLP(cpu, 0)
			return nil
		case 0x87:
			cpu.AF.Hi = cpu.bitres8(0, cpu.AF.Hi)
			return nil

		// RES 1, r|(HL)
		case 0x88:
			cpu.BC.Hi = cpu.bitres8(1, cpu.BC.Hi)
			return nil
		case 0x89:
			cpu.BC.Lo = cpu.bitres8(1, cpu.BC.Lo)
			return nil
		case 0x8a:
			cpu.DE.Hi = cpu.bitres8(1, cpu.DE.Hi)
			return nil
		case 0x8b:
			cpu.DE.Lo = cpu.bitres8(1, cpu.DE.Lo)
			return nil
		case 0x8c:
			cpu.HL.Hi = cpu.bitres8(1, cpu.HL.Hi)
			return nil
		case 0x8d:
			cpu.HL.Lo = cpu.bitres8(1, cpu.HL.Lo)
			return nil
		case 0x8e:
			xopBITbHLP(cpu, 1)
			return nil
		case 0x8f:
			cpu.AF.Hi = cpu.bitres8(1, cpu.AF.Hi)
			return nil

		// RES 2, r|(HL)
		case 0x90:
			cpu.BC.Hi = cpu.bitres8(2, cpu.BC.Hi)
			return nil
		case 0x91:
			cpu.BC.Lo = cpu.bitres8(2, cpu.BC.Lo)
			return nil
		case 0x92:
			cpu.DE.Hi = cpu.bitres8(2, cpu.DE.Hi)
			return nil
		case 0x93:
			cpu.DE.Lo = cpu.bitres8(2, cpu.DE.Lo)
			return nil
		case 0x94:
			cpu.HL.Hi = cpu.bitres8(2, cpu.HL.Hi)
			return nil
		case 0x95:
			cpu.HL.Lo = cpu.bitres8(2, cpu.HL.Lo)
			return nil
		case 0x96:
			xopBITbHLP(cpu, 2)
			return nil
		case 0x97:
			cpu.AF.Hi = cpu.bitres8(2, cpu.AF.Hi)
			return nil

		// RES 3, r|(HL)
		case 0x98:
			cpu.BC.Hi = cpu.bitres8(3, cpu.BC.Hi)
			return nil
		case 0x99:
			cpu.BC.Lo = cpu.bitres8(3, cpu.BC.Lo)
			return nil
		case 0x9a:
			cpu.DE.Hi = cpu.bitres8(3, cpu.DE.Hi)
			return nil
		case 0x9b:
			cpu.DE.Lo = cpu.bitres8(3, cpu.DE.Lo)
			return nil
		case 0x9c:
			cpu.HL.Hi = cpu.bitres8(3, cpu.HL.Hi)
			return nil
		case 0x9d:
			cpu.HL.Lo = cpu.bitres8(3, cpu.HL.Lo)
			return nil
		case 0x9e:
			xopBITbHLP(cpu, 3)
			return nil
		case 0x9f:
			cpu.AF.Hi = cpu.bitres8(3, cpu.AF.Hi)
			return nil

		// RES 4, r|(HL)
		case 0xa0:
			cpu.BC.Hi = cpu.bitres8(4, cpu.BC.Hi)
			return nil
		case 0xa1:
			cpu.BC.Lo = cpu.bitres8(4, cpu.BC.Lo)
			return nil
		case 0xa2:
			cpu.DE.Hi = cpu.bitres8(4, cpu.DE.Hi)
			return nil
		case 0xa3:
			cpu.DE.Lo = cpu.bitres8(4, cpu.DE.Lo)
			return nil
		case 0xa4:
			cpu.HL.Hi = cpu.bitres8(4, cpu.HL.Hi)
			return nil
		case 0xa5:
			cpu.HL.Lo = cpu.bitres8(4, cpu.HL.Lo)
			return nil
		case 0xa6:
			xopBITbHLP(cpu, 4)
			return nil
		case 0xa7:
			cpu.AF.Hi = cpu.bitres8(4, cpu.AF.Hi)
			return nil

		// RES 5, r|(HL)
		case 0xa8:
			cpu.BC.Hi = cpu.bitres8(5, cpu.BC.Hi)
			return nil
		case 0xa9:
			cpu.BC.Lo = cpu.bitres8(5, cpu.BC.Lo)
			return nil
		case 0xaa:
			cpu.DE.Hi = cpu.bitres8(5, cpu.DE.Hi)
			return nil
		case 0xab:
			cpu.DE.Lo = cpu.bitres8(5, cpu.DE.Lo)
			return nil
		case 0xac:
			cpu.HL.Hi = cpu.bitres8(5, cpu.HL.Hi)
			return nil
		case 0xad:
			cpu.HL.Lo = cpu.bitres8(5, cpu.HL.Lo)
			return nil
		case 0xae:
			xopBITbHLP(cpu, 5)
			return nil
		case 0xaf:
			cpu.AF.Hi = cpu.bitres8(5, cpu.AF.Hi)
			return nil

		// RES 6, r|(HL)
		case 0xb0:
			cpu.BC.Hi = cpu.bitres8(6, cpu.BC.Hi)
			return nil
		case 0xb1:
			cpu.BC.Lo = cpu.bitres8(6, cpu.BC.Lo)
			return nil
		case 0xb2:
			cpu.DE.Hi = cpu.bitres8(6, cpu.DE.Hi)
			return nil
		case 0xb3:
			cpu.DE.Lo = cpu.bitres8(6, cpu.DE.Lo)
			return nil
		case 0xb4:
			cpu.HL.Hi = cpu.bitres8(6, cpu.HL.Hi)
			return nil
		case 0xb5:
			cpu.HL.Lo = cpu.bitres8(6, cpu.HL.Lo)
			return nil
		case 0xb6:
			xopBITbHLP(cpu, 6)
			return nil
		case 0xb7:
			cpu.AF.Hi = cpu.bitres8(6, cpu.AF.Hi)
			return nil

		// RES 7, r|(HL)
		case 0xb8:
			cpu.BC.Hi = cpu.bitres8(7, cpu.BC.Hi)
			return nil
		case 0xb9:
			cpu.BC.Lo = cpu.bitres8(7, cpu.BC.Lo)
			return nil
		case 0xba:
			cpu.DE.Hi = cpu.bitres8(7, cpu.DE.Hi)
			return nil
		case 0xbb:
			cpu.DE.Lo = cpu.bitres8(7, cpu.DE.Lo)
			return nil
		case 0xbc:
			cpu.HL.Hi = cpu.bitres8(7, cpu.HL.Hi)
			return nil
		case 0xbd:
			cpu.HL.Lo = cpu.bitres8(7, cpu.HL.Lo)
			return nil
		case 0xbe:
			xopBITbHLP(cpu, 7)
			return nil
		case 0xbf:
			cpu.AF.Hi = cpu.bitres8(7, cpu.AF.Hi)
			return nil

		// SET 0, r|(HL)
		case 0xc0:
			cpu.BC.Hi = cpu.bitset8(0, cpu.BC.Hi)
			return nil
		case 0xc1:
			cpu.BC.Lo = cpu.bitset8(0, cpu.BC.Lo)
			return nil
		case 0xc2:
			cpu.DE.Hi = cpu.bitset8(0, cpu.DE.Hi)
			return nil
		case 0xc3:
			cpu.DE.Lo = cpu.bitset8(0, cpu.DE.Lo)
			return nil
		case 0xc4:
			cpu.HL.Hi = cpu.bitset8(0, cpu.HL.Hi)
			return nil
		case 0xc5:
			cpu.HL.Lo = cpu.bitset8(0, cpu.HL.Lo)
			return nil
		case 0xc6:
			xopSETbHLP(cpu, 0)
			return nil
		case 0xc7:
			cpu.AF.Hi = cpu.bitset8(0, cpu.AF.Hi)
			return nil

		// SET 1, r|(HL)
		case 0xc8:
			cpu.BC.Hi = cpu.bitset8(1, cpu.BC.Hi)
			return nil
		case 0xc9:
			cpu.BC.Lo = cpu.bitset8(1, cpu.BC.Lo)
			return nil
		case 0xca:
			cpu.DE.Hi = cpu.bitset8(1, cpu.DE.Hi)
			return nil
		case 0xcb:
			cpu.DE.Lo = cpu.bitset8(1, cpu.DE.Lo)
			return nil
		case 0xcc:
			cpu.HL.Hi = cpu.bitset8(1, cpu.HL.Hi)
			return nil
		case 0xcd:
			cpu.HL.Lo = cpu.bitset8(1, cpu.HL.Lo)
			return nil
		case 0xce:
			xopSETbHLP(cpu, 1)
			return nil
		case 0xcf:
			cpu.AF.Hi = cpu.bitset8(1, cpu.AF.Hi)
			return nil

		// SET 2, r|(HL)
		case 0xd0:
			cpu.BC.Hi = cpu.bitset8(2, cpu.BC.Hi)
			return nil
		case 0xd1:
			cpu.BC.Lo = cpu.bitset8(2, cpu.BC.Lo)
			return nil
		case 0xd2:
			cpu.DE.Hi = cpu.bitset8(2, cpu.DE.Hi)
			return nil
		case 0xd3:
			cpu.DE.Lo = cpu.bitset8(2, cpu.DE.Lo)
			return nil
		case 0xd4:
			cpu.HL.Hi = cpu.bitset8(2, cpu.HL.Hi)
			return nil
		case 0xd5:
			cpu.HL.Lo = cpu.bitset8(2, cpu.HL.Lo)
			return nil
		case 0xd6:
			xopSETbHLP(cpu, 2)
			return nil
		case 0xd7:
			cpu.AF.Hi = cpu.bitset8(2, cpu.AF.Hi)
			return nil

		// SET 3, r|(HL)
		case 0xd8:
			cpu.BC.Hi = cpu.bitset8(3, cpu.BC.Hi)
			return nil
		case 0xd9:
			cpu.BC.Lo = cpu.bitset8(3, cpu.BC.Lo)
			return nil
		case 0xda:
			cpu.DE.Hi = cpu.bitset8(3, cpu.DE.Hi)
			return nil
		case 0xdb:
			cpu.DE.Lo = cpu.bitset8(3, cpu.DE.Lo)
			return nil
		case 0xdc:
			cpu.HL.Hi = cpu.bitset8(3, cpu.HL.Hi)
			return nil
		case 0xdd:
			cpu.HL.Lo = cpu.bitset8(3, cpu.HL.Lo)
			return nil
		case 0xde:
			xopSETbHLP(cpu, 3)
			return nil
		case 0xdf:
			cpu.AF.Hi = cpu.bitset8(3, cpu.AF.Hi)
			return nil

		// SET 4, r|(HL)
		case 0xe0:
			cpu.BC.Hi = cpu.bitset8(4, cpu.BC.Hi)
			return nil
		case 0xe1:
			cpu.BC.Lo = cpu.bitset8(4, cpu.BC.Lo)
			return nil
		case 0xe2:
			cpu.DE.Hi = cpu.bitset8(4, cpu.DE.Hi)
			return nil
		case 0xe3:
			cpu.DE.Lo = cpu.bitset8(4, cpu.DE.Lo)
			return nil
		case 0xe4:
			cpu.HL.Hi = cpu.bitset8(4, cpu.HL.Hi)
			return nil
		case 0xe5:
			cpu.HL.Lo = cpu.bitset8(4, cpu.HL.Lo)
			return nil
		case 0xe6:
			xopSETbHLP(cpu, 4)
			return nil
		case 0xe7:
			cpu.AF.Hi = cpu.bitset8(4, cpu.AF.Hi)
			return nil

		// SET 5, r|(HL)
		case 0xe8:
			cpu.BC.Hi = cpu.bitset8(5, cpu.BC.Hi)
			return nil
		case 0xe9:
			cpu.BC.Lo = cpu.bitset8(5, cpu.BC.Lo)
			return nil
		case 0xea:
			cpu.DE.Hi = cpu.bitset8(5, cpu.DE.Hi)
			return nil
		case 0xeb:
			cpu.DE.Lo = cpu.bitset8(5, cpu.DE.Lo)
			return nil
		case 0xec:
			cpu.HL.Hi = cpu.bitset8(5, cpu.HL.Hi)
			return nil
		case 0xed:
			cpu.HL.Lo = cpu.bitset8(5, cpu.HL.Lo)
			return nil
		case 0xee:
			xopSETbHLP(cpu, 5)
			return nil
		case 0xef:
			cpu.AF.Hi = cpu.bitset8(5, cpu.AF.Hi)
			return nil

		// SET 6, r|(HL)
		case 0xf0:
			cpu.BC.Hi = cpu.bitset8(6, cpu.BC.Hi)
			return nil
		case 0xf1:
			cpu.BC.Lo = cpu.bitset8(6, cpu.BC.Lo)
			return nil
		case 0xf2:
			cpu.DE.Hi = cpu.bitset8(6, cpu.DE.Hi)
			return nil
		case 0xf3:
			cpu.DE.Lo = cpu.bitset8(6, cpu.DE.Lo)
			return nil
		case 0xf4:
			cpu.HL.Hi = cpu.bitset8(6, cpu.HL.Hi)
			return nil
		case 0xf5:
			cpu.HL.Lo = cpu.bitset8(6, cpu.HL.Lo)
			return nil
		case 0xf6:
			xopSETbHLP(cpu, 6)
			return nil
		case 0xf7:
			cpu.AF.Hi = cpu.bitset8(6, cpu.AF.Hi)
			return nil

		// SET 7, r|(HL)
		case 0xf8:
			cpu.BC.Hi = cpu.bitset8(7, cpu.BC.Hi)
			return nil
		case 0xf9:
			cpu.BC.Lo = cpu.bitset8(7, cpu.BC.Lo)
			return nil
		case 0xfa:
			cpu.DE.Hi = cpu.bitset8(7, cpu.DE.Hi)
			return nil
		case 0xfb:
			cpu.DE.Lo = cpu.bitset8(7, cpu.DE.Lo)
			return nil
		case 0xfc:
			cpu.HL.Hi = cpu.bitset8(7, cpu.HL.Hi)
			return nil
		case 0xfd:
			cpu.HL.Lo = cpu.bitset8(7, cpu.HL.Lo)
			return nil
		case 0xfe:
			xopSETbHLP(cpu, 7)
			return nil
		case 0xff:
			cpu.AF.Hi = cpu.bitset8(7, cpu.AF.Hi)
			return nil

		default:
			return ErrInvalidCodes
		}

	case 0xdd:
		buf[1] = f.fetch()
		switch buf[1] {

		// ADD IX, pp
		case 0x09:
			xopADDIXbc(cpu)
			return nil
		case 0x19:
			xopADDIXde(cpu)
			return nil
		case 0x29:
			xopADDIXix(cpu)
			return nil
		case 0x39:
			xopADDIXsp(cpu)
			return nil

		case 0x21:
			l := f.fetch()
			h := f.fetch()
			oopLDIXnn(cpu, l, h)
			return nil

		case 0x22:
			l := f.fetch()
			h := f.fetch()
			oopLDnnPIX(cpu, l, h)
			return nil

		case 0x23:
			oopINCIX(cpu)
			return nil

		case 0x24:
			oopINCIXH(cpu)
			return nil

		case 0x25:
			oopDECIXH(cpu)
			return nil

		case 0x26:
			n := f.fetch()
			oopLDIXHn(cpu, n)
			return nil

		case 0x2a:
			l := f.fetch()
			h := f.fetch()
			oopLDIXnnP(cpu, l, h)
			return nil

		case 0x2b:
			oopDECIX(cpu)
			return nil

		case 0x2c:
			oopINCIXL(cpu)
			return nil

		case 0x2d:
			oopDECIXL(cpu)
			return nil

		case 0x2e:
			n := f.fetch()
			oopLDIXLn(cpu, n)
			return nil

		case 0x34:
			d := f.fetch()
			oopINCIXdP(cpu, d)
			return nil

		case 0x35:
			d := f.fetch()
			oopDECIXdP(cpu, d)
			return nil

		case 0x36:
			d := f.fetch()
			n := f.fetch()
			oopLDIXdPn(cpu, d, n)
			return nil

		// LD rx1, rx2
		case 0x40:
			//cpu.BC.Hi = cpu.BC.Hi
			return nil
		case 0x41:
			cpu.BC.Hi = cpu.BC.Lo
			return nil
		case 0x42:
			cpu.BC.Hi = cpu.DE.Hi
			return nil
		case 0x43:
			cpu.BC.Hi = cpu.DE.Lo
			return nil
		case 0x44:
			cpu.BC.Hi = uint8(cpu.IX >> 8)
			return nil
		case 0x45:
			cpu.BC.Hi = uint8(cpu.IX)
			return nil
		case 0x47:
			cpu.BC.Hi = cpu.AF.Hi
			return nil
		case 0x48:
			cpu.BC.Lo = cpu.BC.Hi
			return nil
		case 0x49:
			//cpu.BC.Lo = cpu.BC.Lo
			return nil
		case 0x4a:
			cpu.BC.Lo = cpu.DE.Hi
			return nil
		case 0x4b:
			cpu.BC.Lo = cpu.DE.Lo
			return nil
		case 0x4c:
			cpu.BC.Lo = uint8(cpu.IX >> 8)
			return nil
		case 0x4d:
			cpu.BC.Lo = uint8(cpu.IX)
			return nil
		case 0x4f:
			cpu.BC.Lo = cpu.AF.Hi
			return nil
		case 0x50:
			cpu.DE.Hi = cpu.BC.Hi
			return nil
		case 0x51:
			cpu.DE.Hi = cpu.BC.Lo
			return nil
		case 0x52:
			//cpu.DE.Hi = cpu.DE.Hi
			return nil
		case 0x53:
			cpu.DE.Hi = cpu.DE.Lo
			return nil
		case 0x54:
			cpu.DE.Hi = uint8(cpu.IX >> 8)
			return nil
		case 0x55:
			cpu.DE.Hi = uint8(cpu.IX)
			return nil
		case 0x57:
			cpu.DE.Hi = cpu.AF.Hi
			return nil
		case 0x58:
			cpu.DE.Lo = cpu.BC.Hi
			return nil
		case 0x59:
			cpu.DE.Lo = cpu.BC.Lo
			return nil
		case 0x5a:
			cpu.DE.Lo = cpu.DE.Hi
			return nil
		case 0x5b:
			//cpu.DE.Lo = cpu.DE.Lo
			return nil
		case 0x5c:
			cpu.DE.Lo = uint8(cpu.IX >> 8)
			return nil
		case 0x5d:
			cpu.DE.Lo = uint8(cpu.IX)
			return nil
		case 0x5f:
			cpu.DE.Lo = cpu.AF.Hi
			return nil
		case 0x60:
			cpu.IX = uint16(cpu.BC.Hi)<<8 | cpu.IX&0x00ff
			return nil
		case 0x61:
			cpu.IX = uint16(cpu.BC.Lo)<<8 | cpu.IX&0x00ff
			return nil
		case 0x62:
			cpu.IX = uint16(cpu.DE.Hi)<<8 | cpu.IX&0x00ff
			return nil
		case 0x63:
			cpu.IX = uint16(cpu.DE.Lo)<<8 | cpu.IX&0x00ff
			return nil
		case 0x64:
			//cpu.IX = uint16(uint8(cpu.IX >> 8))<<8 | cpu.IX&0x00ff
			return nil
		case 0x65:
			cpu.IX = uint16(uint8(cpu.IX))<<8 | cpu.IX&0x00ff
			return nil
		case 0x67:
			cpu.IX = uint16(cpu.AF.Hi)<<8 | cpu.IX&0x00ff
			return nil
		case 0x68:
			cpu.IX = uint16(cpu.BC.Hi) | cpu.IX&0xff00
			return nil
		case 0x69:
			cpu.IX = uint16(cpu.BC.Lo) | cpu.IX&0xff00
			return nil
		case 0x6a:
			cpu.IX = uint16(cpu.DE.Hi) | cpu.IX&0xff00
			return nil
		case 0x6b:
			cpu.IX = uint16(cpu.DE.Lo) | cpu.IX&0xff00
			return nil
		case 0x6c:
			cpu.IX = uint16(uint8(cpu.IX>>8)) | cpu.IX&0xff00
			return nil
		case 0x6d:
			//cpu.IX = uint16(uint8(cpu.IX)) | cpu.IX&0xff00
			return nil
		case 0x6f:
			cpu.IX = uint16(cpu.AF.Hi) | cpu.IX&0xff00
			return nil
		case 0x78:
			cpu.AF.Hi = cpu.BC.Hi
			return nil
		case 0x79:
			cpu.AF.Hi = cpu.BC.Lo
			return nil
		case 0x7a:
			cpu.AF.Hi = cpu.DE.Hi
			return nil
		case 0x7b:
			cpu.AF.Hi = cpu.DE.Lo
			return nil
		case 0x7c:
			cpu.AF.Hi = uint8(cpu.IX >> 8)
			return nil
		case 0x7d:
			cpu.AF.Hi = uint8(cpu.IX)
			return nil
		case 0x7f:
			//cpu.AF.Hi = cpu.AF.Hi
			return nil

		// LD r, (IX+d)
		case 0x46:
			d := f.fetch()
			xopLDbIXdP(cpu, d)
			return nil
		case 0x4e:
			d := f.fetch()
			xopLDcIXdP(cpu, d)
			return nil
		case 0x56:
			d := f.fetch()
			xopLDdIXdP(cpu, d)
			return nil
		case 0x5e:
			d := f.fetch()
			xopLDeIXdP(cpu, d)
			return nil
		case 0x66:
			d := f.fetch()
			xopLDhIXdP(cpu, d)
			return nil
		case 0x6e:
			d := f.fetch()
			xopLDlIXdP(cpu, d)
			return nil
		case 0x7e:
			d := f.fetch()
			xopLDaIXdP(cpu, d)
			return nil

		// LD (IX+d), r
		case 0x70:
			d := f.fetch()
			xopLDIXdPb(cpu, d)
			return nil
		case 0x71:
			d := f.fetch()
			xopLDIXdPc(cpu, d)
			return nil
		case 0x72:
			d := f.fetch()
			xopLDIXdPd(cpu, d)
			return nil
		case 0x73:
			d := f.fetch()
			xopLDIXdPe(cpu, d)
			return nil
		case 0x74:
			d := f.fetch()
			xopLDIXdPh(cpu, d)
			return nil
		case 0x75:
			d := f.fetch()
			xopLDIXdPl(cpu, d)
			return nil
		case 0x77:
			d := f.fetch()
			xopLDIXdPa(cpu, d)
			return nil

		// ADD A, rx (undocumented)
		case 0x80:
			xopADDAb(cpu)
			return nil
		case 0x81:
			xopADDAc(cpu)
			return nil
		case 0x82:
			xopADDAd(cpu)
			return nil
		case 0x83:
			xopADDAe(cpu)
			return nil
		case 0x84:
			xopADDAixh(cpu)
			return nil
		case 0x85:
			xopADDAixl(cpu)
			return nil
		case 0x87:
			xopADDAa(cpu)
			return nil

		case 0x86:
			d := f.fetch()
			oopADDAIXdP(cpu, d)
			return nil

		// ADC A, rx
		case 0x88:
			xopADCAb(cpu)
			return nil
		case 0x89:
			xopADCAc(cpu)
			return nil
		case 0x8a:
			xopADCAd(cpu)
			return nil
		case 0x8b:
			xopADCAe(cpu)
			return nil
		case 0x8c:
			xopADCAixh(cpu)
			return nil
		case 0x8d:
			xopADCAixl(cpu)
			return nil
		case 0x8f:
			xopADCAa(cpu)
			return nil

		case 0x8e:
			d := f.fetch()
			oopADCAIXdP(cpu, d)
			return nil

		// SUB A, rx
		case 0x90:
			xopSUBAb(cpu)
			return nil
		case 0x91:
			xopSUBAc(cpu)
			return nil
		case 0x92:
			xopSUBAd(cpu)
			return nil
		case 0x93:
			xopSUBAe(cpu)
			return nil
		case 0x94:
			xopSUBAixh(cpu)
			return nil
		case 0x95:
			xopSUBAixl(cpu)
			return nil
		case 0x97:
			xopSUBAa(cpu)
			return nil

		case 0x96:
			d := f.fetch()
			oopSUBAIXdP(cpu, d)
			return nil

		// SBC A, rx
		case 0x98:
			xopSBCAb(cpu)
			return nil
		case 0x99:
			xopSBCAc(cpu)
			return nil
		case 0x9a:
			xopSBCAd(cpu)
			return nil
		case 0x9b:
			xopSBCAe(cpu)
			return nil
		case 0x9c:
			xopSBCAixh(cpu)
			return nil
		case 0x9d:
			xopSBCAixl(cpu)
			return nil
		case 0x9f:
			xopSBCAa(cpu)
			return nil

		case 0x9e:
			d := f.fetch()
			oopSBCAIXdP(cpu, d)
			return nil

		// ADD rx
		case 0xa0:
			xopANDAb(cpu)
			return nil
		case 0xa1:
			xopANDAc(cpu)
			return nil
		case 0xa2:
			xopANDAd(cpu)
			return nil
		case 0xa3:
			xopANDAe(cpu)
			return nil
		case 0xa4:
			xopANDixh(cpu)
			return nil
		case 0xa5:
			xopANDixl(cpu)
			return nil
		case 0xa7:
			xopANDAa(cpu)
			return nil

		// AND (IX+d)
		case 0xa6:
			d := f.fetch()
			oopANDIXdP(cpu, d)
			return nil

		// XOR rx
		case 0xa8:
			xopXORb(cpu)
			return nil
		case 0xa9:
			xopXORc(cpu)
			return nil
		case 0xaa:
			xopXORd(cpu)
			return nil
		case 0xab:
			xopXORe(cpu)
			return nil
		case 0xac:
			xopXORixh(cpu)
			return nil
		case 0xad:
			xopXORixl(cpu)
			return nil
		case 0xaf:
			xopXORa(cpu)
			return nil

		case 0xae:
			d := f.fetch()
			oopXORIXdP(cpu, d)
			return nil

		// OR rx
		case 0xb0:
			xopORb(cpu)
			return nil
		case 0xb1:
			xopORc(cpu)
			return nil
		case 0xb2:
			xopORd(cpu)
			return nil
		case 0xb3:
			xopORe(cpu)
			return nil
		case 0xb4:
			xopORixh(cpu)
			return nil
		case 0xb5:
			xopORixl(cpu)
			return nil
		case 0xb7:
			xopORa(cpu)
			return nil

		case 0xb6:
			d := f.fetch()
			oopORIXdP(cpu, d)
			return nil

		// CP rx
		case 0xb8:
			xopCPb(cpu)
			return nil
		case 0xb9:
			xopCPc(cpu)
			return nil
		case 0xba:
			xopCPd(cpu)
			return nil
		case 0xbb:
			xopCPe(cpu)
			return nil
		case 0xbc:
			xopCPixh(cpu)
			return nil
		case 0xbd:
			xopCPixl(cpu)
			return nil
		case 0xbf:
			xopCPa(cpu)
			return nil

		case 0xbe:
			d := f.fetch()
			oopCPIXdP(cpu, d)
			return nil

		case 0xe1:
			oopPOPIX(cpu)
			return nil

		case 0xe3:
			oopEXSPPIX(cpu)
			return nil

		case 0xe5:
			oopPUSHIX(cpu)
			return nil

		case 0xe9:
			oopJPIXP(cpu)
			return nil

		case 0xf9:
			oopLDSPIX(cpu)
			return nil

		case 0xcb:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			switch buf[3] {
			case 0x06:
				opRLCIXdP(cpu, buf[:4])
				return nil
			case 0x0e:
				opRRCIXdP(cpu, buf[:4])
				return nil
			case 0x16:
				opRLIXdP(cpu, buf[:4])
				return nil
			case 0x1e:
				opRRIXdP(cpu, buf[:4])
				return nil
			case 0x26:
				opSLAIXdP(cpu, buf[:4])
				return nil
			case 0x2e:
				opSRAIXdP(cpu, buf[:4])
				return nil
			case 0x36:
				opSL1IXdP(cpu, buf[:4])
				return nil
			case 0x3e:
				opSRLIXdP(cpu, buf[:4])
				return nil
			case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x76, 0x7e:
				opBITbIXdP(cpu, buf[:4])
				return nil
			case 0x86, 0x8e, 0x96, 0x9e, 0xa6, 0xae, 0xb6, 0xbe:
				opRESbIXdP(cpu, buf[:4])
				return nil
			case 0xc6, 0xce, 0xd6, 0xde, 0xe6, 0xee, 0xf6, 0xfe:
				opSETbIXdP(cpu, buf[:4])
				return nil
			default:
				return ErrInvalidCodes
			}
		default:
			return ErrInvalidCodes
		}

	case 0xed:
		buf[1] = f.fetch()
		switch buf[1] {
		case 0x40, 0x48, 0x50, 0x58, 0x60, 0x68, 0x78:
			opINrCP(cpu, buf[:2])
			return nil
		case 0x41, 0x49, 0x51, 0x59, 0x61, 0x69, 0x79:
			opOUTCPr(cpu, buf[:2])
			return nil
		case 0x42, 0x52, 0x62, 0x72:
			opSBCHLss(cpu, buf[:2])
			return nil
		case 0x43, 0x53, 0x63, 0x73:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDnnPdd(cpu, buf[:4])
			return nil
		case 0x44:
			opNEG(cpu, buf[:2])
			return nil

		case 0x45:
			oopRETN(cpu)
			return nil

		case 0x46:
			opIM0(cpu, buf[:2])
			return nil
		case 0x47:
			opLDIA(cpu, buf[:2])
			return nil
		case 0x4a, 0x5a, 0x6a, 0x7a:
			opADCHLss(cpu, buf[:2])
			return nil
		case 0x4b, 0x5b, 0x6b, 0x7b:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDddnnP(cpu, buf[:4])
			return nil

		case 0x4d:
			oopRETI(cpu)
			return nil

		case 0x4f:
			opLDRA(cpu, buf[:2])
			return nil
		case 0x56:
			opIM1(cpu, buf[:2])
			return nil
		case 0x57:
			opLDAI(cpu, buf[:2])
			return nil
		case 0x5e:
			opIM2(cpu, buf[:2])
			return nil
		case 0x5f:
			opLDAR(cpu, buf[:2])
			return nil
		case 0x67:
			opRRD(cpu, buf[:2])
			return nil
		case 0x6f:
			opRLD(cpu, buf[:2])
			return nil
		case 0xa0:
			opLDI(cpu, buf[:2])
			return nil
		case 0xa1:
			opCPI(cpu, buf[:2])
			return nil
		case 0xa2:
			opINI(cpu, buf[:2])
			return nil
		case 0xa3:
			opOUTI(cpu, buf[:2])
			return nil
		case 0xa8:
			opLDD(cpu, buf[:2])
			return nil
		case 0xa9:
			opCPD(cpu, buf[:2])
			return nil
		case 0xaa:
			opIND(cpu, buf[:2])
			return nil
		case 0xab:
			opOUTD(cpu, buf[:2])
			return nil
		case 0xb0:
			opLDIR(cpu, buf[:2])
			return nil
		case 0xb1:
			opCPIR(cpu, buf[:2])
			return nil
		case 0xb2:
			opINIR(cpu, buf[:2])
			return nil
		case 0xb3:
			opOTIR(cpu, buf[:2])
			return nil
		case 0xb8:
			opLDDR(cpu, buf[:2])
			return nil
		case 0xb9:
			opCPDR(cpu, buf[:2])
			return nil
		case 0xba:
			opINDR(cpu, buf[:2])
			return nil
		case 0xbb:
			opOTDR(cpu, buf[:2])
			return nil
		default:
			return ErrInvalidCodes
		}

	case 0xfd:
		buf[1] = f.fetch()
		switch buf[1] {

		// ADD IY, pp
		case 0x09:
			xopADDIYbc(cpu)
			return nil
		case 0x19:
			xopADDIYde(cpu)
			return nil
		case 0x29:
			xopADDIYiy(cpu)
			return nil
		case 0x39:
			xopADDIYsp(cpu)
			return nil

		case 0x21:
			l := f.fetch()
			h := f.fetch()
			oopLDIYnn(cpu, l, h)
			return nil

		case 0x22:
			l := f.fetch()
			h := f.fetch()
			oopLDnnPIY(cpu, l, h)
			return nil

		case 0x23:
			oopINCIY(cpu)
			return nil

		case 0x24:
			oopINCIYH(cpu)
			return nil

		case 0x25:
			oopDECIYH(cpu)
			return nil

		case 0x26:
			n := f.fetch()
			oopLDIYHn(cpu, n)
			return nil

		case 0x2a:
			l := f.fetch()
			h := f.fetch()
			oopLDIYnnP(cpu, l, h)
			return nil

		case 0x2b:
			oopDECIY(cpu)
			return nil

		case 0x2c:
			oopINCIYL(cpu)
			return nil

		case 0x2d:
			oopDECIYL(cpu)
			return nil

		case 0x2e:
			n := f.fetch()
			oopLDIYLn(cpu, n)
			return nil

		case 0x34:
			d := f.fetch()
			oopINCIYdP(cpu, d)
			return nil

		case 0x35:
			d := f.fetch()
			oopDECIYdP(cpu, d)
			return nil

		case 0x36:
			d := f.fetch()
			n := f.fetch()
			oopLDIYdPn(cpu, d, n)
			return nil

		// LD ry1, ry2
		case 0x40:
			//cpu.BC.Hi = cpu.BC.Hi
			return nil
		case 0x41:
			cpu.BC.Hi = cpu.BC.Lo
			return nil
		case 0x42:
			cpu.BC.Hi = cpu.DE.Hi
			return nil
		case 0x43:
			cpu.BC.Hi = cpu.DE.Lo
			return nil
		case 0x44:
			cpu.BC.Hi = uint8(cpu.IY >> 8)
			return nil
		case 0x45:
			cpu.BC.Hi = uint8(cpu.IY)
			return nil
		case 0x47:
			cpu.BC.Hi = cpu.AF.Hi
			return nil
		case 0x48:
			cpu.BC.Lo = cpu.BC.Hi
			return nil
		case 0x49:
			//cpu.BC.Lo = cpu.BC.Lo
			return nil
		case 0x4a:
			cpu.BC.Lo = cpu.DE.Hi
			return nil
		case 0x4b:
			cpu.BC.Lo = cpu.DE.Lo
			return nil
		case 0x4c:
			cpu.BC.Lo = uint8(cpu.IY >> 8)
			return nil
		case 0x4d:
			cpu.BC.Lo = uint8(cpu.IY)
			return nil
		case 0x4f:
			cpu.BC.Lo = cpu.AF.Hi
			return nil
		case 0x50:
			cpu.DE.Hi = cpu.BC.Hi
			return nil
		case 0x51:
			cpu.DE.Hi = cpu.BC.Lo
			return nil
		case 0x52:
			//cpu.DE.Hi = cpu.DE.Hi
			return nil
		case 0x53:
			cpu.DE.Hi = cpu.DE.Lo
			return nil
		case 0x54:
			cpu.DE.Hi = uint8(cpu.IY >> 8)
			return nil
		case 0x55:
			cpu.DE.Hi = uint8(cpu.IY)
			return nil
		case 0x57:
			cpu.DE.Hi = cpu.AF.Hi
			return nil
		case 0x58:
			cpu.DE.Lo = cpu.BC.Hi
			return nil
		case 0x59:
			cpu.DE.Lo = cpu.BC.Lo
			return nil
		case 0x5a:
			cpu.DE.Lo = cpu.DE.Hi
			return nil
		case 0x5b:
			//cpu.DE.Lo = cpu.DE.Lo
			return nil
		case 0x5c:
			cpu.DE.Lo = uint8(cpu.IY >> 8)
			return nil
		case 0x5d:
			cpu.DE.Lo = uint8(cpu.IY)
			return nil
		case 0x5f:
			cpu.DE.Lo = cpu.AF.Hi
			return nil
		case 0x60:
			cpu.IY = uint16(cpu.BC.Hi)<<8 | cpu.IY&0x00ff
			return nil
		case 0x61:
			cpu.IY = uint16(cpu.BC.Lo)<<8 | cpu.IY&0x00ff
			return nil
		case 0x62:
			cpu.IY = uint16(cpu.DE.Hi)<<8 | cpu.IY&0x00ff
			return nil
		case 0x63:
			cpu.IY = uint16(cpu.DE.Lo)<<8 | cpu.IY&0x00ff
			return nil
		case 0x64:
			//cpu.IY = uint16(uint8(cpu.IY >> 8))<<8 | cpu.IY&0x00ff
			return nil
		case 0x65:
			cpu.IY = uint16(uint8(cpu.IY))<<8 | cpu.IY&0x00ff
			return nil
		case 0x67:
			cpu.IY = uint16(cpu.AF.Hi)<<8 | cpu.IY&0x00ff
			return nil
		case 0x68:
			cpu.IY = uint16(cpu.BC.Hi) | cpu.IY&0xff00
			return nil
		case 0x69:
			cpu.IY = uint16(cpu.BC.Lo) | cpu.IY&0xff00
			return nil
		case 0x6a:
			cpu.IY = uint16(cpu.DE.Hi) | cpu.IY&0xff00
			return nil
		case 0x6b:
			cpu.IY = uint16(cpu.DE.Lo) | cpu.IY&0xff00
			return nil
		case 0x6c:
			cpu.IY = uint16(uint8(cpu.IY>>8)) | cpu.IY&0xff00
			return nil
		case 0x6d:
			//cpu.IY = uint16(uint8(cpu.IY)) | cpu.IY&0xff00
			return nil
		case 0x6f:
			cpu.IY = uint16(cpu.AF.Hi) | cpu.IY&0xff00
			return nil
		case 0x78:
			cpu.AF.Hi = cpu.BC.Hi
			return nil
		case 0x79:
			cpu.AF.Hi = cpu.BC.Lo
			return nil
		case 0x7a:
			cpu.AF.Hi = cpu.DE.Hi
			return nil
		case 0x7b:
			cpu.AF.Hi = cpu.DE.Lo
			return nil
		case 0x7c:
			cpu.AF.Hi = uint8(cpu.IY >> 8)
			return nil
		case 0x7d:
			cpu.AF.Hi = uint8(cpu.IY)
			return nil
		case 0x7f:
			//cpu.AF.Hi = cpu.AF.Hi
			return nil

		// LD r, (IY+d)
		case 0x46:
			d := f.fetch()
			xopLDbIYdP(cpu, d)
			return nil
		case 0x4e:
			d := f.fetch()
			xopLDcIYdP(cpu, d)
			return nil
		case 0x56:
			d := f.fetch()
			xopLDdIYdP(cpu, d)
			return nil
		case 0x5e:
			d := f.fetch()
			xopLDeIYdP(cpu, d)
			return nil
		case 0x66:
			d := f.fetch()
			xopLDhIYdP(cpu, d)
			return nil
		case 0x6e:
			d := f.fetch()
			xopLDlIYdP(cpu, d)
			return nil
		case 0x7e:
			d := f.fetch()
			xopLDaIYdP(cpu, d)
			return nil

		// LD (IY+d), r
		case 0x70:
			d := f.fetch()
			xopLDIYdPb(cpu, d)
			return nil
		case 0x71:
			d := f.fetch()
			xopLDIYdPc(cpu, d)
			return nil
		case 0x72:
			d := f.fetch()
			xopLDIYdPd(cpu, d)
			return nil
		case 0x73:
			d := f.fetch()
			xopLDIYdPe(cpu, d)
			return nil
		case 0x74:
			d := f.fetch()
			xopLDIYdPh(cpu, d)
			return nil
		case 0x75:
			d := f.fetch()
			xopLDIYdPl(cpu, d)
			return nil
		case 0x77:
			d := f.fetch()
			xopLDIYdPa(cpu, d)
			return nil

		// ADD A, ry (undocumented)
		case 0x80:
			xopADDAb(cpu)
			return nil
		case 0x81:
			xopADDAc(cpu)
			return nil
		case 0x82:
			xopADDAd(cpu)
			return nil
		case 0x83:
			xopADDAe(cpu)
			return nil
		case 0x84:
			xopADDAiyh(cpu)
			return nil
		case 0x85:
			xopADDAiyl(cpu)
			return nil
		case 0x87:
			xopADDAa(cpu)
			return nil

		case 0x86:
			d := f.fetch()
			oopADDAIYdP(cpu, d)
			return nil

		// ADC A, ry
		case 0x88:
			xopADCAb(cpu)
			return nil
		case 0x89:
			xopADCAc(cpu)
			return nil
		case 0x8a:
			xopADCAd(cpu)
			return nil
		case 0x8b:
			xopADCAe(cpu)
			return nil
		case 0x8c:
			xopADCAiyh(cpu)
			return nil
		case 0x8d:
			xopADCAiyl(cpu)
			return nil
		case 0x8f:
			xopADCAa(cpu)
			return nil

		case 0x8e:
			d := f.fetch()
			oopADCAIYdP(cpu, d)
			return nil

		// SUB A, ry
		case 0x90:
			xopSUBAb(cpu)
			return nil
		case 0x91:
			xopSUBAc(cpu)
			return nil
		case 0x92:
			xopSUBAd(cpu)
			return nil
		case 0x93:
			xopSUBAe(cpu)
			return nil
		case 0x94:
			xopSUBAiyh(cpu)
			return nil
		case 0x95:
			xopSUBAiyl(cpu)
			return nil
		case 0x97:
			xopSUBAa(cpu)
			return nil

		case 0x96:
			d := f.fetch()
			oopSUBAIYdP(cpu, d)
			return nil

		// SBC A, ry
		case 0x98:
			xopSBCAb(cpu)
			return nil
		case 0x99:
			xopSBCAc(cpu)
			return nil
		case 0x9a:
			xopSBCAd(cpu)
			return nil
		case 0x9b:
			xopSBCAe(cpu)
			return nil
		case 0x9c:
			xopSBCAiyh(cpu)
			return nil
		case 0x9d:
			xopSBCAiyl(cpu)
			return nil
		case 0x9f:
			xopSBCAa(cpu)
			return nil

		case 0x9e:
			d := f.fetch()
			oopSBCAIYdP(cpu, d)
			return nil

		// ADD rx
		case 0xa0:
			xopANDAb(cpu)
			return nil
		case 0xa1:
			xopANDAc(cpu)
			return nil
		case 0xa2:
			xopANDAd(cpu)
			return nil
		case 0xa3:
			xopANDAe(cpu)
			return nil
		case 0xa4:
			xopANDiyh(cpu)
			return nil
		case 0xa5:
			xopANDiyl(cpu)
			return nil
		case 0xa7:
			xopANDAa(cpu)
			return nil

		// AND (IY+d)
		case 0xa6:
			d := f.fetch()
			oopANDIYdP(cpu, d)
			return nil

		// XOR rx
		case 0xa8:
			xopXORb(cpu)
			return nil
		case 0xa9:
			xopXORc(cpu)
			return nil
		case 0xaa:
			xopXORd(cpu)
			return nil
		case 0xab:
			xopXORe(cpu)
			return nil
		case 0xac:
			xopXORiyh(cpu)
			return nil
		case 0xad:
			xopXORiyl(cpu)
			return nil
		case 0xaf:
			xopXORa(cpu)
			return nil

		case 0xae:
			d := f.fetch()
			oopXORIYdP(cpu, d)
			return nil

		// OR ry
		case 0xb0:
			xopORb(cpu)
			return nil
		case 0xb1:
			xopORc(cpu)
			return nil
		case 0xb2:
			xopORd(cpu)
			return nil
		case 0xb3:
			xopORe(cpu)
			return nil
		case 0xb4:
			xopORiyh(cpu)
			return nil
		case 0xb5:
			xopORiyl(cpu)
			return nil
		case 0xb7:
			xopORa(cpu)
			return nil

		case 0xb6:
			d := f.fetch()
			oopORIYdP(cpu, d)
			return nil

		// CP ry
		case 0xb8:
			xopCPb(cpu)
			return nil
		case 0xb9:
			xopCPc(cpu)
			return nil
		case 0xba:
			xopCPd(cpu)
			return nil
		case 0xbb:
			xopCPe(cpu)
			return nil
		case 0xbc:
			xopCPiyh(cpu)
			return nil
		case 0xbd:
			xopCPiyl(cpu)
			return nil
		case 0xbf:
			xopCPa(cpu)
			return nil

		case 0xbe:
			d := f.fetch()
			oopCPIYdP(cpu, d)
			return nil

		case 0xe1:
			oopPOPIY(cpu)
			return nil

		case 0xe3:
			oopEXSPPIY(cpu)
			return nil

		case 0xe5:
			oopPUSHIY(cpu)
			return nil

		case 0xe9:
			oopJPIYP(cpu)
			return nil

		case 0xf9:
			oopLDSPIY(cpu)
			return nil

		case 0xcb:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			switch buf[3] {
			case 0x06:
				opRLCIYdP(cpu, buf[:4])
				return nil
			case 0x0e:
				opRRCIYdP(cpu, buf[:4])
				return nil
			case 0x16:
				opRLIYdP(cpu, buf[:4])
				return nil
			case 0x1e:
				opRRIYdP(cpu, buf[:4])
				return nil
			case 0x26:
				opSLAIYdP(cpu, buf[:4])
				return nil
			case 0x2e:
				opSRAIYdP(cpu, buf[:4])
				return nil
			case 0x36:
				opSL1IYdP(cpu, buf[:4])
				return nil
			case 0x3e:
				opSRLIYdP(cpu, buf[:4])
				return nil
			case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x76, 0x7e:
				opBITbIYdP(cpu, buf[:4])
				return nil
			case 0x86, 0x8e, 0x96, 0x9e, 0xa6, 0xae, 0xb6, 0xbe:
				opRESbIYdP(cpu, buf[:4])
				return nil
			case 0xc6, 0xce, 0xd6, 0xde, 0xe6, 0xee, 0xf6, 0xfe:
				opSETbIYdP(cpu, buf[:4])
				return nil
			default:
				return ErrInvalidCodes
			}
		default:
			return ErrInvalidCodes
		}

	default:
		return ErrInvalidCodes
	}
}
