package z80

func decodeExec(cpu *CPU, f fetcher) error {
	buf := cpu.decodeBuf[:4]
	buf[0] = f.fetch()
	switch buf[0] {
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

	case 0x40:
		cpu.BC.Hi = cpu.BC.Hi
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
		cpu.BC.Lo = cpu.BC.Lo
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
		cpu.DE.Hi = cpu.DE.Hi
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
		cpu.DE.Lo = cpu.DE.Lo
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
		cpu.HL.Hi = cpu.HL.Hi
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
		cpu.HL.Lo = cpu.HL.Lo
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
		cpu.AF.Hi = cpu.AF.Hi
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

	case 0x86:
		oopADDAHLP(cpu)
		return nil

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
		buf[1] = f.fetch()
		switch buf[1] {
		case 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x07:
			opRLCr(cpu, buf[:2])
			return nil
		case 0x06:
			opRLCHLP(cpu, buf[:2])
			return nil
		case 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0f:
			opRRCr(cpu, buf[:2])
			return nil
		case 0x0e:
			opRRCHLP(cpu, buf[:2])
			return nil
		case 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x17:
			opRLr(cpu, buf[:2])
			return nil
		case 0x16:
			opRLHLP(cpu, buf[:2])
			return nil
		case 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1f:
			opRRr(cpu, buf[:2])
			return nil
		case 0x1e:
			opRRHLP(cpu, buf[:2])
			return nil
		case 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x27:
			opSLAr(cpu, buf[:2])
			return nil
		case 0x26:
			opSLAHLP(cpu, buf[:2])
			return nil
		case 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2f:
			opSRAr(cpu, buf[:2])
			return nil
		case 0x2e:
			opSRAHLP(cpu, buf[:2])
			return nil
		case 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x37:
			opSL1r(cpu, buf[:2])
			return nil
		case 0x36:
			opSL1HLP(cpu, buf[:2])
			return nil
		case 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3f:
			opSRLr(cpu, buf[:2])
			return nil
		case 0x3e:
			opSRLHLP(cpu, buf[:2])
			return nil
		case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
			opBITbr(cpu, buf[:2])
			return nil
		case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x76, 0x7e:
			opBITbHLP(cpu, buf[:2])
			return nil
		case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97, 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f, 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7, 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf, 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7, 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
			opRESbr(cpu, buf[:2])
			return nil
		case 0x86, 0x8e, 0x96, 0x9e, 0xa6, 0xae, 0xb6, 0xbe:
			opRESbHLP(cpu, buf[:2])
			return nil
		case 0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc7, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xcf, 0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd7, 0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xdf, 0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe7, 0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xef, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf7, 0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xff:
			opSETbr(cpu, buf[:2])
			return nil
		case 0xc6, 0xce, 0xd6, 0xde, 0xe6, 0xee, 0xf6, 0xfe:
			opSETbHLP(cpu, buf[:2])
			return nil
		default:
			return ErrInvalidCodes
		}
	case 0xdd:
		buf[1] = f.fetch()
		switch buf[1] {
		case 0x09, 0x19, 0x29, 0x39:
			opADDIXpp(cpu, buf[:2])
			return nil
		case 0x21:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDIXnn(cpu, buf[:4])
			return nil
		case 0x22:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDnnPIX(cpu, buf[:4])
			return nil
		case 0x23:
			opINCIX(cpu, buf[:2])
			return nil
		case 0x24:
			opINCIXH(cpu, buf[:2])
			return nil
		case 0x25:
			opDECIXH(cpu, buf[:2])
			return nil
		case 0x26:
			buf[2] = f.fetch()
			opLDIXHn(cpu, buf[:3])
			return nil
		case 0x2a:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDIXnnP(cpu, buf[:4])
			return nil
		case 0x2b:
			opDECIX(cpu, buf[:2])
			return nil
		case 0x2c:
			opINCIXL(cpu, buf[:2])
			return nil
		case 0x2d:
			opDECIXL(cpu, buf[:2])
			return nil
		case 0x2e:
			buf[2] = f.fetch()
			opLDIXLn(cpu, buf[:3])
			return nil
		case 0x34:
			buf[2] = f.fetch()
			opINCIXdP(cpu, buf[:3])
			return nil
		case 0x35:
			buf[2] = f.fetch()
			opDECIXdP(cpu, buf[:3])
			return nil
		case 0x36:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDIXdPn(cpu, buf[:4])
			return nil
		case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
			opLDrx1rx2(cpu, buf[:2])
			return nil
		case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x7e:
			buf[2] = f.fetch()
			opLDrIXdP(cpu, buf[:3])
			return nil
		case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
			buf[2] = f.fetch()
			opLDIXdPr(cpu, buf[:3])
			return nil
		case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87:
			opADDArx(cpu, buf[:2])
			return nil
		case 0x86:
			buf[2] = f.fetch()
			opADDAIXdP(cpu, buf[:3])
			return nil
		case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f:
			opADCArx(cpu, buf[:2])
			return nil
		case 0x8e:
			buf[2] = f.fetch()
			opADCAIXdP(cpu, buf[:3])
			return nil
		case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97:
			opSUBArx(cpu, buf[:2])
			return nil
		case 0x96:
			buf[2] = f.fetch()
			opSUBAIXdP(cpu, buf[:3])
			return nil
		case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f:
			opSBCArx(cpu, buf[:2])
			return nil
		case 0x9e:
			buf[2] = f.fetch()
			opSBCAIXdP(cpu, buf[:3])
			return nil
		case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7:
			opANDrx(cpu, buf[:2])
			return nil
		case 0xa6:
			buf[2] = f.fetch()
			opANDIXdP(cpu, buf[:3])
			return nil
		case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf:
			opXORrx(cpu, buf[:2])
			return nil
		case 0xae:
			buf[2] = f.fetch()
			opXORIXdP(cpu, buf[:3])
			return nil
		case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7:
			opORrx(cpu, buf[:2])
			return nil
		case 0xb6:
			buf[2] = f.fetch()
			opORIXdP(cpu, buf[:3])
			return nil
		case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
			opCPrx(cpu, buf[:2])
			return nil
		case 0xbe:
			buf[2] = f.fetch()
			opCPIXdP(cpu, buf[:3])
			return nil
		case 0xe1:
			opPOPIX(cpu, buf[:2])
			return nil
		case 0xe3:
			opEXSPPIX(cpu, buf[:2])
			return nil
		case 0xe5:
			opPUSHIX(cpu, buf[:2])
			return nil
		case 0xe9:
			opJPIXP(cpu, buf[:2])
			return nil
		case 0xf9:
			opLDSPIX(cpu, buf[:2])
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
			opRETN(cpu, buf[:2])
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
			opRETI(cpu, buf[:2])
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
		case 0x09, 0x19, 0x29, 0x39:
			opADDIYrr(cpu, buf[:2])
			return nil
		case 0x21:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDIYnn(cpu, buf[:4])
			return nil
		case 0x22:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDnnPIY(cpu, buf[:4])
			return nil
		case 0x23:
			opINCIY(cpu, buf[:2])
			return nil
		case 0x24:
			opINCIYH(cpu, buf[:2])
			return nil
		case 0x25:
			opDECIYH(cpu, buf[:2])
			return nil
		case 0x26:
			buf[2] = f.fetch()
			opLDIYHn(cpu, buf[:3])
			return nil
		case 0x2a:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDIYnnP(cpu, buf[:4])
			return nil
		case 0x2b:
			opDECIY(cpu, buf[:2])
			return nil
		case 0x2c:
			opINCIYL(cpu, buf[:2])
			return nil
		case 0x2d:
			opDECIYL(cpu, buf[:2])
			return nil
		case 0x2e:
			buf[2] = f.fetch()
			opLDIYLn(cpu, buf[:3])
			return nil
		case 0x34:
			buf[2] = f.fetch()
			opINCIYdP(cpu, buf[:3])
			return nil
		case 0x35:
			buf[2] = f.fetch()
			opDECIYdP(cpu, buf[:3])
			return nil
		case 0x36:
			buf[2] = f.fetch()
			buf[3] = f.fetch()
			opLDIYdPn(cpu, buf[:4])
			return nil
		case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
			opLDry1ry2(cpu, buf[:2])
			return nil
		case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x7e:
			buf[2] = f.fetch()
			opLDrIYdP(cpu, buf[:3])
			return nil
		case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
			buf[2] = f.fetch()
			opLDIYdPr(cpu, buf[:3])
			return nil
		case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87:
			opADDAry(cpu, buf[:2])
			return nil
		case 0x86:
			buf[2] = f.fetch()
			opADDAIYdP(cpu, buf[:3])
			return nil
		case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f:
			opADCAry(cpu, buf[:2])
			return nil
		case 0x8e:
			buf[2] = f.fetch()
			opADCAIYdP(cpu, buf[:3])
			return nil
		case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97:
			opSUBAry(cpu, buf[:2])
			return nil
		case 0x96:
			buf[2] = f.fetch()
			opSUBAIYdP(cpu, buf[:3])
			return nil
		case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f:
			opSBCAry(cpu, buf[:2])
			return nil
		case 0x9e:
			buf[2] = f.fetch()
			opSBCAIYdP(cpu, buf[:3])
			return nil
		case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7:
			opANDry(cpu, buf[:2])
			return nil
		case 0xa6:
			buf[2] = f.fetch()
			opANDIYdP(cpu, buf[:3])
			return nil
		case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf:
			opXORry(cpu, buf[:2])
			return nil
		case 0xae:
			buf[2] = f.fetch()
			opXORIYdP(cpu, buf[:3])
			return nil
		case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7:
			opORry(cpu, buf[:2])
			return nil
		case 0xb6:
			buf[2] = f.fetch()
			opORIYdP(cpu, buf[:3])
			return nil
		case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
			opCPry(cpu, buf[:2])
			return nil
		case 0xbe:
			buf[2] = f.fetch()
			opCPIYdP(cpu, buf[:3])
			return nil
		case 0xe1:
			opPOPIY(cpu, buf[:2])
			return nil
		case 0xe3:
			opEXSPPIY(cpu, buf[:2])
			return nil
		case 0xe5:
			opPUSHIY(cpu, buf[:2])
			return nil
		case 0xe9:
			opJPIYP(cpu, buf[:2])
			return nil
		case 0xf9:
			opLDSPIY(cpu, buf[:2])
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
