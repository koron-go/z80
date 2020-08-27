package z80

func decodeExec(cpu *CPU, f fetcher) error {
	buf := cpu.decodeBuf[:4]
	buf[0] = f.fetch()
	switch buf[0] {
	case 0x00:
		opNOP(cpu, buf[:1])
		return nil
	case 0x01, 0x11, 0x21, 0x31:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opLDddnn(cpu, buf[:3])
		return nil
	case 0x02:
		opLDBCPA(cpu, buf[:1])
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
	case 0x04, 0x0c, 0x14, 0x1c, 0x24, 0x2c, 0x3c:
		opINCr(cpu, buf[:1])
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

	case 0x06, 0x0e, 0x16, 0x1e, 0x26, 0x2e, 0x3e:
		buf[1] = f.fetch()
		opLDrn(cpu, buf[:2])
		return nil
	case 0x07:
		opRLCA(cpu, buf[:1])
		return nil
	case 0x08:
		opEXAFAF(cpu, buf[:1])
		return nil
	case 0x09, 0x19, 0x29, 0x39:
		opADDHLss(cpu, buf[:1])
		return nil
	case 0x0a:
		opLDABCP(cpu, buf[:1])
		return nil
	case 0x0b, 0x1b, 0x2b, 0x3b:
		opDECss(cpu, buf[:1])
		return nil
	case 0x0f:
		opRRCA(cpu, buf[:1])
		return nil
	case 0x10:
		buf[1] = f.fetch()
		opDJNZe(cpu, buf[:2])
		return nil
	case 0x12:
		opLDDEPA(cpu, buf[:1])
		return nil
	case 0x17:
		opRLA(cpu, buf[:1])
		return nil
	case 0x18:
		buf[1] = f.fetch()
		opJRe(cpu, buf[:2])
		return nil
	case 0x1a:
		opLDADEP(cpu, buf[:1])
		return nil
	case 0x1f:
		opRRA(cpu, buf[:1])
		return nil
	case 0x20:
		buf[1] = f.fetch()
		opJRNZe(cpu, buf[:2])
		return nil
	case 0x22:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opLDnnPHL(cpu, buf[:3])
		return nil
	case 0x27:
		opDAA(cpu, buf[:1])
		return nil
	case 0x28:
		buf[1] = f.fetch()
		opJRZe(cpu, buf[:2])
		return nil
	case 0x2a:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opLDHLnnP(cpu, buf[:3])
		return nil
	case 0x2f:
		opCPL(cpu, buf[:1])
		return nil
	case 0x30:
		buf[1] = f.fetch()
		opJRNCe(cpu, buf[:2])
		return nil
	case 0x32:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opLDnnPA(cpu, buf[:3])
		return nil
	case 0x34:
		opINCHLP(cpu, buf[:1])
		return nil
	case 0x35:
		opDECHLP(cpu, buf[:1])
		return nil
	case 0x36:
		buf[1] = f.fetch()
		opLDHLPn(cpu, buf[:2])
		return nil
	case 0x37:
		opSCF(cpu, buf[:1])
		return nil
	case 0x38:
		buf[1] = f.fetch()
		opJRCe(cpu, buf[:2])
		return nil
	case 0x3a:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opLDAnnP(cpu, buf[:3])
		return nil
	case 0x3f:
		opCCF(cpu, buf[:1])
		return nil
	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
		opLDr1r2(cpu, buf[:1])
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
	case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
		opLDHLPr(cpu, buf[:1])
		return nil
	case 0x76:
		opHALT(cpu, buf[:1])
		return nil
	case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87:
		opADDAr(cpu, buf[:1])
		return nil
	case 0x86:
		opADDAHLP(cpu, buf[:1])
		return nil
	case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f:
		opADCAr(cpu, buf[:1])
		return nil
	case 0x8e:
		opADCAHLP(cpu, buf[:1])
		return nil
	case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97:
		opSUBAr(cpu, buf[:1])
		return nil
	case 0x96:
		opSUBAHLP(cpu, buf[:1])
		return nil
	case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f:
		opSBCAr(cpu, buf[:1])
		return nil
	case 0x9e:
		opSBCAHLP(cpu, buf[:1])
		return nil
	case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7:
		opANDr(cpu, buf[:1])
		return nil
	case 0xa6:
		opANDHLP(cpu, buf[:1])
		return nil
	case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf:
		opXORr(cpu, buf[:1])
		return nil
	case 0xae:
		opXORHLP(cpu, buf[:1])
		return nil
	case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7:
		opORr(cpu, buf[:1])
		return nil
	case 0xb6:
		opORHLP(cpu, buf[:1])
		return nil
	case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
		opCPr(cpu, buf[:1])
		return nil
	case 0xbe:
		opCPHLP(cpu, buf[:1])
		return nil
	case 0xc0, 0xc8, 0xd0, 0xd8, 0xe0, 0xe8, 0xf0, 0xf8:
		opRETcc(cpu, buf[:1])
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
	case 0xc2, 0xca, 0xd2, 0xda, 0xe2, 0xea, 0xf2, 0xfa:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opJPccnn(cpu, buf[:3])
		return nil
	case 0xc3:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opJPnn(cpu, buf[:3])
		return nil
	case 0xc4, 0xcc, 0xd4, 0xdc, 0xe4, 0xec, 0xf4, 0xfc:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opCALLccnn(cpu, buf[:3])
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
		buf[1] = f.fetch()
		opADDAn(cpu, buf[:2])
		return nil
	case 0xc7, 0xcf, 0xd7, 0xdf, 0xe7, 0xef, 0xf7, 0xff:
		opRSTp(cpu, buf[:1])
		return nil
	case 0xc9:
		opRET(cpu, buf[:1])
		return nil
	case 0xcd:
		buf[1] = f.fetch()
		buf[2] = f.fetch()
		opCALLnn(cpu, buf[:3])
		return nil
	case 0xce:
		buf[1] = f.fetch()
		opADCAn(cpu, buf[:2])
		return nil
	case 0xd3:
		buf[1] = f.fetch()
		opOUTnPA(cpu, buf[:2])
		return nil
	case 0xd6:
		buf[1] = f.fetch()
		opSUBAn(cpu, buf[:2])
		return nil
	case 0xd9:
		opEXX(cpu, buf[:1])
		return nil
	case 0xdb:
		buf[1] = f.fetch()
		opINAnP(cpu, buf[:2])
		return nil
	case 0xde:
		buf[1] = f.fetch()
		opSBCAn(cpu, buf[:2])
		return nil
	case 0xe3:
		opEXSPPHL(cpu, buf[:1])
		return nil
	case 0xe6:
		buf[1] = f.fetch()
		opANDn(cpu, buf[:2])
		return nil
	case 0xe9:
		opJPHLP(cpu, buf[:1])
		return nil
	case 0xeb:
		opEXDEHL(cpu, buf[:1])
		return nil
	case 0xee:
		buf[1] = f.fetch()
		opXORn(cpu, buf[:2])
		return nil
	case 0xf3:
		opDI(cpu, buf[:1])
		return nil
	case 0xf6:
		buf[1] = f.fetch()
		opORn(cpu, buf[:2])
		return nil
	case 0xf9:
		opLDSPHL(cpu, buf[:1])
		return nil
	case 0xfb:
		opEI(cpu, buf[:1])
		return nil
	case 0xfe:
		buf[1] = f.fetch()
		opCPn(cpu, buf[:2])
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
