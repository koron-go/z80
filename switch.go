package z80

func decodeExec(cpu *CPU, f fetcher) error {
	var b uint8
	buf := cpu.decodeBuf[:0]
	b = f.fetch()
	buf = append(buf, b)
	switch b {
	case 0x00:
		opNOP(cpu, buf)
		return nil
	case 0x01, 0x11, 0x21, 0x31:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opLDddnn(cpu, buf)
		return nil
	case 0x02:
		opLDBCPA(cpu, buf)
		return nil
	case 0x03, 0x13, 0x23, 0x33:
		opINCss(cpu, buf)
		return nil
	case 0x04, 0x0c, 0x14, 0x1c, 0x24, 0x2c, 0x3c:
		opINCr(cpu, buf)
		return nil
	case 0x05, 0x0d, 0x15, 0x1d, 0x25, 0x2d, 0x3d:
		opDECr(cpu, buf)
		return nil
	case 0x06, 0x0e, 0x16, 0x1e, 0x26, 0x2e, 0x3e:
		buf = append(buf, f.fetch())
		opLDrn(cpu, buf)
		return nil
	case 0x07:
		opRLCA(cpu, buf)
		return nil
	case 0x08:
		opEXAFAF(cpu, buf)
		return nil
	case 0x09, 0x19, 0x29, 0x39:
		opADDHLss(cpu, buf)
		return nil
	case 0x0a:
		opLDABCP(cpu, buf)
		return nil
	case 0x0b, 0x1b, 0x2b, 0x3b:
		opDECss(cpu, buf)
		return nil
	case 0x0f:
		opRRCA(cpu, buf)
		return nil
	case 0x10:
		buf = append(buf, f.fetch())
		opDJNZe(cpu, buf)
		return nil
	case 0x12:
		opLDDEPA(cpu, buf)
		return nil
	case 0x17:
		opRLA(cpu, buf)
		return nil
	case 0x18:
		buf = append(buf, f.fetch())
		opJRe(cpu, buf)
		return nil
	case 0x1a:
		opLDADEP(cpu, buf)
		return nil
	case 0x1f:
		opRRA(cpu, buf)
		return nil
	case 0x20:
		buf = append(buf, f.fetch())
		opJRNZe(cpu, buf)
		return nil
	case 0x22:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opLDnnPHL(cpu, buf)
		return nil
	case 0x27:
		opDAA(cpu, buf)
		return nil
	case 0x28:
		buf = append(buf, f.fetch())
		opJRZe(cpu, buf)
		return nil
	case 0x2a:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opLDHLnnP(cpu, buf)
		return nil
	case 0x2f:
		opCPL(cpu, buf)
		return nil
	case 0x30:
		buf = append(buf, f.fetch())
		opJRNCe(cpu, buf)
		return nil
	case 0x32:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opLDnnPA(cpu, buf)
		return nil
	case 0x34:
		opINCHLP(cpu, buf)
		return nil
	case 0x35:
		opDECHLP(cpu, buf)
		return nil
	case 0x36:
		buf = append(buf, f.fetch())
		opLDHLPn(cpu, buf)
		return nil
	case 0x37:
		opSCF(cpu, buf)
		return nil
	case 0x38:
		buf = append(buf, f.fetch())
		opJRCe(cpu, buf)
		return nil
	case 0x3a:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opLDAnnP(cpu, buf)
		return nil
	case 0x3f:
		opCCF(cpu, buf)
		return nil
	case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
		opLDr1r2(cpu, buf)
		return nil
	case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x7e:
		opLDrHLP(cpu, buf)
		return nil
	case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
		opLDHLPr(cpu, buf)
		return nil
	case 0x76:
		opHALT(cpu, buf)
		return nil
	case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87:
		opADDAr(cpu, buf)
		return nil
	case 0x86:
		opADDAHLP(cpu, buf)
		return nil
	case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f:
		opADCAr(cpu, buf)
		return nil
	case 0x8e:
		opADCAHLP(cpu, buf)
		return nil
	case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97:
		opSUBAr(cpu, buf)
		return nil
	case 0x96:
		opSUBAHLP(cpu, buf)
		return nil
	case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f:
		opSBCAr(cpu, buf)
		return nil
	case 0x9e:
		opSBCAHLP(cpu, buf)
		return nil
	case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7:
		opANDr(cpu, buf)
		return nil
	case 0xa6:
		opANDHLP(cpu, buf)
		return nil
	case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf:
		opXORr(cpu, buf)
		return nil
	case 0xae:
		opXORHLP(cpu, buf)
		return nil
	case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7:
		opORr(cpu, buf)
		return nil
	case 0xb6:
		opORHLP(cpu, buf)
		return nil
	case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
		opCPr(cpu, buf)
		return nil
	case 0xbe:
		opCPHLP(cpu, buf)
		return nil
	case 0xc0, 0xc8, 0xd0, 0xd8, 0xe0, 0xe8, 0xf0, 0xf8:
		opRETcc(cpu, buf)
		return nil
	case 0xc1, 0xd1, 0xe1, 0xf1:
		opPOPqq(cpu, buf)
		return nil
	case 0xc2, 0xca, 0xd2, 0xda, 0xe2, 0xea, 0xf2, 0xfa:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opJPccnn(cpu, buf)
		return nil
	case 0xc3:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opJPnn(cpu, buf)
		return nil
	case 0xc4, 0xcc, 0xd4, 0xdc, 0xe4, 0xec, 0xf4, 0xfc:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opCALLccnn(cpu, buf)
		return nil
	case 0xc5, 0xd5, 0xe5, 0xf5:
		opPUSHqq(cpu, buf)
		return nil
	case 0xc6:
		buf = append(buf, f.fetch())
		opADDAn(cpu, buf)
		return nil
	case 0xc7, 0xcf, 0xd7, 0xdf, 0xe7, 0xef, 0xf7, 0xff:
		opRSTp(cpu, buf)
		return nil
	case 0xc9:
		opRET(cpu, buf)
		return nil
	case 0xcd:
		buf = append(buf, f.fetch())
		buf = append(buf, f.fetch())
		opCALLnn(cpu, buf)
		return nil
	case 0xce:
		buf = append(buf, f.fetch())
		opADCAn(cpu, buf)
		return nil
	case 0xd3:
		buf = append(buf, f.fetch())
		opOUTnPA(cpu, buf)
		return nil
	case 0xd6:
		buf = append(buf, f.fetch())
		opSUBAn(cpu, buf)
		return nil
	case 0xd9:
		opEXX(cpu, buf)
		return nil
	case 0xdb:
		buf = append(buf, f.fetch())
		opINAnP(cpu, buf)
		return nil
	case 0xde:
		buf = append(buf, f.fetch())
		opSBCAn(cpu, buf)
		return nil
	case 0xe3:
		opEXSPPHL(cpu, buf)
		return nil
	case 0xe6:
		buf = append(buf, f.fetch())
		opANDn(cpu, buf)
		return nil
	case 0xe9:
		opJPHLP(cpu, buf)
		return nil
	case 0xeb:
		opEXDEHL(cpu, buf)
		return nil
	case 0xee:
		buf = append(buf, f.fetch())
		opXORn(cpu, buf)
		return nil
	case 0xf3:
		opDI(cpu, buf)
		return nil
	case 0xf6:
		buf = append(buf, f.fetch())
		opORn(cpu, buf)
		return nil
	case 0xf9:
		opLDSPHL(cpu, buf)
		return nil
	case 0xfb:
		opEI(cpu, buf)
		return nil
	case 0xfe:
		buf = append(buf, f.fetch())
		opCPn(cpu, buf)
		return nil
	case 0xcb:
		b = f.fetch()
		buf = append(buf, b)
		switch b {
		case 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x07:
			opRLCr(cpu, buf)
			return nil
		case 0x06:
			opRLCHLP(cpu, buf)
			return nil
		case 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0f:
			opRRCr(cpu, buf)
			return nil
		case 0x0e:
			opRRCHLP(cpu, buf)
			return nil
		case 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x17:
			opRLr(cpu, buf)
			return nil
		case 0x16:
			opRLHLP(cpu, buf)
			return nil
		case 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1f:
			opRRr(cpu, buf)
			return nil
		case 0x1e:
			opRRHLP(cpu, buf)
			return nil
		case 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x27:
			opSLAr(cpu, buf)
			return nil
		case 0x26:
			opSLAHLP(cpu, buf)
			return nil
		case 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2f:
			opSRAr(cpu, buf)
			return nil
		case 0x2e:
			opSRAHLP(cpu, buf)
			return nil
		case 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x37:
			opSL1r(cpu, buf)
			return nil
		case 0x36:
			opSL1HLP(cpu, buf)
			return nil
		case 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3f:
			opSRLr(cpu, buf)
			return nil
		case 0x3e:
			opSRLHLP(cpu, buf)
			return nil
		case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
			opBITbr(cpu, buf)
			return nil
		case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x76, 0x7e:
			opBITbHLP(cpu, buf)
			return nil
		case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87, 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97, 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f, 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7, 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf, 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7, 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
			opRESbr(cpu, buf)
			return nil
		case 0x86, 0x8e, 0x96, 0x9e, 0xa6, 0xae, 0xb6, 0xbe:
			opRESbHLP(cpu, buf)
			return nil
		case 0xc0, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc7, 0xc8, 0xc9, 0xca, 0xcb, 0xcc, 0xcd, 0xcf, 0xd0, 0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd7, 0xd8, 0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xdf, 0xe0, 0xe1, 0xe2, 0xe3, 0xe4, 0xe5, 0xe7, 0xe8, 0xe9, 0xea, 0xeb, 0xec, 0xed, 0xef, 0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf7, 0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xff:
			opSETbr(cpu, buf)
			return nil
		case 0xc6, 0xce, 0xd6, 0xde, 0xe6, 0xee, 0xf6, 0xfe:
			opSETbHLP(cpu, buf)
			return nil
		default:
			return ErrInvalidCodes
		}
	case 0xdd:
		b = f.fetch()
		buf = append(buf, b)
		switch b {
		case 0x09, 0x19, 0x29, 0x39:
			opADDIXpp(cpu, buf)
			return nil
		case 0x21:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDIXnn(cpu, buf)
			return nil
		case 0x22:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDnnPIX(cpu, buf)
			return nil
		case 0x23:
			opINCIX(cpu, buf)
			return nil
		case 0x24:
			opINCIXH(cpu, buf)
			return nil
		case 0x25:
			opDECIXH(cpu, buf)
			return nil
		case 0x26:
			buf = append(buf, f.fetch())
			opLDIXHn(cpu, buf)
			return nil
		case 0x2a:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDIXnnP(cpu, buf)
			return nil
		case 0x2b:
			opDECIX(cpu, buf)
			return nil
		case 0x2c:
			opINCIXL(cpu, buf)
			return nil
		case 0x2d:
			opDECIXL(cpu, buf)
			return nil
		case 0x2e:
			buf = append(buf, f.fetch())
			opLDIXLn(cpu, buf)
			return nil
		case 0x34:
			buf = append(buf, f.fetch())
			opINCIXdP(cpu, buf)
			return nil
		case 0x35:
			buf = append(buf, f.fetch())
			opDECIXdP(cpu, buf)
			return nil
		case 0x36:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDIXdPn(cpu, buf)
			return nil
		case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x7e:
			buf = append(buf, f.fetch())
			opLDrIXdP(cpu, buf)
			return nil
		case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
			buf = append(buf, f.fetch())
			opLDIXdPr(cpu, buf)
			return nil
		case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87:
			opADDArx(cpu, buf)
			return nil
		case 0x86:
			buf = append(buf, f.fetch())
			opADDAIXdP(cpu, buf)
			return nil
		case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f:
			opADCArx(cpu, buf)
			return nil
		case 0x8e:
			buf = append(buf, f.fetch())
			opADCAIXdP(cpu, buf)
			return nil
		case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97:
			opSUBArx(cpu, buf)
			return nil
		case 0x96:
			buf = append(buf, f.fetch())
			opSUBAIXdP(cpu, buf)
			return nil
		case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f:
			opSBCArx(cpu, buf)
			return nil
		case 0x9e:
			buf = append(buf, f.fetch())
			opSBCAIXdP(cpu, buf)
			return nil
		case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7:
			opANDrx(cpu, buf)
			return nil
		case 0xa6:
			buf = append(buf, f.fetch())
			opANDIXdP(cpu, buf)
			return nil
		case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf:
			opXORrx(cpu, buf)
			return nil
		case 0xae:
			buf = append(buf, f.fetch())
			opXORIXdP(cpu, buf)
			return nil
		case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7:
			opORrx(cpu, buf)
			return nil
		case 0xb6:
			buf = append(buf, f.fetch())
			opORIXdP(cpu, buf)
			return nil
		case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
			opCPrx(cpu, buf)
			return nil
		case 0xbe:
			buf = append(buf, f.fetch())
			opCPIXdP(cpu, buf)
			return nil
		case 0xe1:
			opPOPIX(cpu, buf)
			return nil
		case 0xe3:
			opEXSPPIX(cpu, buf)
			return nil
		case 0xe5:
			opPUSHIX(cpu, buf)
			return nil
		case 0xe9:
			opJPIXP(cpu, buf)
			return nil
		case 0xf9:
			opLDSPIX(cpu, buf)
			return nil
		case 0xcb:
			buf = append(buf, f.fetch())
			b = f.fetch()
			buf = append(buf, b)
			switch b {
			case 0x06:
				opRLCIXdP(cpu, buf)
				return nil
			case 0x0e:
				opRRCIXdP(cpu, buf)
				return nil
			case 0x16:
				opRLIXdP(cpu, buf)
				return nil
			case 0x1e:
				opRRIXdP(cpu, buf)
				return nil
			case 0x26:
				opSLAIXdP(cpu, buf)
				return nil
			case 0x2e:
				opSRAIXdP(cpu, buf)
				return nil
			case 0x36:
				opSL1IXdP(cpu, buf)
				return nil
			case 0x3e:
				opSRLIXdP(cpu, buf)
				return nil
			case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x76, 0x7e:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x86, 0x8e, 0x96, 0x9e, 0xa6, 0xae, 0xb6, 0xbe:
				opRESbIXdP(cpu, buf)
				return nil
			case 0xc6, 0xce, 0xd6, 0xde, 0xe6, 0xee, 0xf6, 0xfe:
				opSETbIXdP(cpu, buf)
				return nil
			default:
				return ErrInvalidCodes
			}
		default:
			return ErrInvalidCodes
		}
	case 0xed:
		b = f.fetch()
		buf = append(buf, b)
		switch b {
		case 0x40, 0x48, 0x50, 0x58, 0x60, 0x68, 0x78:
			opINrCP(cpu, buf)
			return nil
		case 0x41, 0x49, 0x51, 0x59, 0x61, 0x69, 0x79:
			opOUTCPr(cpu, buf)
			return nil
		case 0x42, 0x52, 0x62, 0x72:
			opSBCHLss(cpu, buf)
			return nil
		case 0x43, 0x53, 0x63, 0x73:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDnnPdd(cpu, buf)
			return nil
		case 0x44:
			opNEG(cpu, buf)
			return nil
		case 0x45:
			opRETN(cpu, buf)
			return nil
		case 0x46:
			opIM0(cpu, buf)
			return nil
		case 0x47:
			opLDIA(cpu, buf)
			return nil
		case 0x4a, 0x5a, 0x6a, 0x7a:
			opADCHLss(cpu, buf)
			return nil
		case 0x4b, 0x5b, 0x6b, 0x7b:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDddnnP(cpu, buf)
			return nil
		case 0x4d:
			opRETI(cpu, buf)
			return nil
		case 0x4f:
			opLDRA(cpu, buf)
			return nil
		case 0x56:
			opIM1(cpu, buf)
			return nil
		case 0x57:
			opLDAI(cpu, buf)
			return nil
		case 0x5e:
			opIM2(cpu, buf)
			return nil
		case 0x5f:
			opLDAR(cpu, buf)
			return nil
		case 0x67:
			opRRD(cpu, buf)
			return nil
		case 0x6f:
			opRLD(cpu, buf)
			return nil
		case 0xa0:
			opLDI(cpu, buf)
			return nil
		case 0xa1:
			opCPI(cpu, buf)
			return nil
		case 0xa2:
			opINI(cpu, buf)
			return nil
		case 0xa3:
			opOUTI(cpu, buf)
			return nil
		case 0xa8:
			opLDD(cpu, buf)
			return nil
		case 0xa9:
			opCPD(cpu, buf)
			return nil
		case 0xaa:
			opIND(cpu, buf)
			return nil
		case 0xab:
			opOUTD(cpu, buf)
			return nil
		case 0xb0:
			opLDIR(cpu, buf)
			return nil
		case 0xb1:
			opCPIR(cpu, buf)
			return nil
		case 0xb2:
			opINIR(cpu, buf)
			return nil
		case 0xb3:
			opOTIR(cpu, buf)
			return nil
		case 0xb8:
			opLDDR(cpu, buf)
			return nil
		case 0xb9:
			opCPDR(cpu, buf)
			return nil
		case 0xba:
			opINDR(cpu, buf)
			return nil
		case 0xbb:
			opOTDR(cpu, buf)
			return nil
		default:
			return ErrInvalidCodes
		}
	case 0xfd:
		b = f.fetch()
		buf = append(buf, b)
		switch b {
		case 0x09, 0x19, 0x29, 0x39:
			opADDIYrr(cpu, buf)
			return nil
		case 0x21:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDIYnn(cpu, buf)
			return nil
		case 0x22:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDnnPIY(cpu, buf)
			return nil
		case 0x23:
			opINCIY(cpu, buf)
			return nil
		case 0x24:
			opINCIYH(cpu, buf)
			return nil
		case 0x25:
			opDECIYH(cpu, buf)
			return nil
		case 0x26:
			buf = append(buf, f.fetch())
			opLDIYHn(cpu, buf)
			return nil
		case 0x2a:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDIYnnP(cpu, buf)
			return nil
		case 0x2b:
			opDECIY(cpu, buf)
			return nil
		case 0x2c:
			opINCIYL(cpu, buf)
			return nil
		case 0x2d:
			opDECIYL(cpu, buf)
			return nil
		case 0x2e:
			buf = append(buf, f.fetch())
			opLDIYLn(cpu, buf)
			return nil
		case 0x34:
			buf = append(buf, f.fetch())
			opINCIYdP(cpu, buf)
			return nil
		case 0x35:
			buf = append(buf, f.fetch())
			opDECIYdP(cpu, buf)
			return nil
		case 0x36:
			buf = append(buf, f.fetch())
			buf = append(buf, f.fetch())
			opLDIYdPn(cpu, buf)
			return nil
		case 0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4f, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5f, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6f, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7f:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x7e:
			buf = append(buf, f.fetch())
			opLDrIYdP(cpu, buf)
			return nil
		case 0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x77:
			buf = append(buf, f.fetch())
			opLDIYdPr(cpu, buf)
			return nil
		case 0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x87:
			opADDAry(cpu, buf)
			return nil
		case 0x86:
			buf = append(buf, f.fetch())
			opADDAIYdP(cpu, buf)
			return nil
		case 0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8f:
			opADCAry(cpu, buf)
			return nil
		case 0x8e:
			buf = append(buf, f.fetch())
			opADCAIYdP(cpu, buf)
			return nil
		case 0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x97:
			opSUBAry(cpu, buf)
			return nil
		case 0x96:
			buf = append(buf, f.fetch())
			opSUBAIYdP(cpu, buf)
			return nil
		case 0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9f:
			opSBCAry(cpu, buf)
			return nil
		case 0x9e:
			buf = append(buf, f.fetch())
			opSBCAIYdP(cpu, buf)
			return nil
		case 0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa7:
			opANDry(cpu, buf)
			return nil
		case 0xa6:
			buf = append(buf, f.fetch())
			opANDIYdP(cpu, buf)
			return nil
		case 0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xaf:
			opXORry(cpu, buf)
			return nil
		case 0xae:
			buf = append(buf, f.fetch())
			opXORIYdP(cpu, buf)
			return nil
		case 0xb0, 0xb1, 0xb2, 0xb3, 0xb4, 0xb5, 0xb7:
			opORry(cpu, buf)
			return nil
		case 0xb6:
			buf = append(buf, f.fetch())
			opORIYdP(cpu, buf)
			return nil
		case 0xb8, 0xb9, 0xba, 0xbb, 0xbc, 0xbd, 0xbf:
			opCPry(cpu, buf)
			return nil
		case 0xbe:
			buf = append(buf, f.fetch())
			opCPIYdP(cpu, buf)
			return nil
		case 0xe1:
			opPOPIY(cpu, buf)
			return nil
		case 0xe3:
			opEXSPPIY(cpu, buf)
			return nil
		case 0xe5:
			opPUSHIY(cpu, buf)
			return nil
		case 0xe9:
			opJPIYP(cpu, buf)
			return nil
		case 0xf9:
			opLDSPIY(cpu, buf)
			return nil
		case 0xcb:
			buf = append(buf, f.fetch())
			b = f.fetch()
			buf = append(buf, b)
			switch b {
			case 0x06:
				opRLCIYdP(cpu, buf)
				return nil
			case 0x0e:
				opRRCIYdP(cpu, buf)
				return nil
			case 0x16:
				opRLIYdP(cpu, buf)
				return nil
			case 0x1e:
				opRRIYdP(cpu, buf)
				return nil
			case 0x26:
				opSLAIYdP(cpu, buf)
				return nil
			case 0x2e:
				opSRAIYdP(cpu, buf)
				return nil
			case 0x36:
				opSL1IYdP(cpu, buf)
				return nil
			case 0x3e:
				opSRLIYdP(cpu, buf)
				return nil
			case 0x46, 0x4e, 0x56, 0x5e, 0x66, 0x6e, 0x76, 0x7e:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x86, 0x8e, 0x96, 0x9e, 0xa6, 0xae, 0xb6, 0xbe:
				opRESbIYdP(cpu, buf)
				return nil
			case 0xc6, 0xce, 0xd6, 0xde, 0xe6, 0xee, 0xf6, 0xfe:
				opSETbIYdP(cpu, buf)
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
