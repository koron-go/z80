package z80

func decodeExec(cpu *CPU, f fetcher) error {
	var err error
	var b uint8
	buf := cpu.decodeBuf[:0]
	b, err = f.fetch()
	if err != nil {
		return err
	}
	buf = append(buf, b)
	switch b {
	case 0x00:
		opNOP(cpu, buf)
		return nil
	case 0x01:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDddnn(cpu, buf)
		return nil
	case 0x02:
		opLDBCPA(cpu, buf)
		return nil
	case 0x03:
		opINCss(cpu, buf)
		return nil
	case 0x04:
		opINCr(cpu, buf)
		return nil
	case 0x05:
		opDECr(cpu, buf)
		return nil
	case 0x06:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDrn(cpu, buf)
		return nil
	case 0x07:
		opRLCA(cpu, buf)
		return nil
	case 0x08:
		opEXAFAF(cpu, buf)
		return nil
	case 0x09:
		opADDHLss(cpu, buf)
		return nil
	case 0x0a:
		opLDABCP(cpu, buf)
		return nil
	case 0x0b:
		opDECss(cpu, buf)
		return nil
	case 0x0c:
		opINCr(cpu, buf)
		return nil
	case 0x0d:
		opDECr(cpu, buf)
		return nil
	case 0x0e:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDrn(cpu, buf)
		return nil
	case 0x0f:
		opRRCA(cpu, buf)
		return nil
	case 0x10:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opDJNZe(cpu, buf)
		return nil
	case 0x11:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDddnn(cpu, buf)
		return nil
	case 0x12:
		opLDDEPA(cpu, buf)
		return nil
	case 0x13:
		opINCss(cpu, buf)
		return nil
	case 0x14:
		opINCr(cpu, buf)
		return nil
	case 0x15:
		opDECr(cpu, buf)
		return nil
	case 0x16:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDrn(cpu, buf)
		return nil
	case 0x17:
		opRLA(cpu, buf)
		return nil
	case 0x18:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJRe(cpu, buf)
		return nil
	case 0x19:
		opADDHLss(cpu, buf)
		return nil
	case 0x1a:
		opLDADEP(cpu, buf)
		return nil
	case 0x1b:
		opDECss(cpu, buf)
		return nil
	case 0x1c:
		opINCr(cpu, buf)
		return nil
	case 0x1d:
		opDECr(cpu, buf)
		return nil
	case 0x1e:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDrn(cpu, buf)
		return nil
	case 0x1f:
		opRRA(cpu, buf)
		return nil
	case 0x20:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJRNZe(cpu, buf)
		return nil
	case 0x21:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDddnn(cpu, buf)
		return nil
	case 0x22:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDnnPHL(cpu, buf)
		return nil
	case 0x23:
		opINCss(cpu, buf)
		return nil
	case 0x24:
		opINCr(cpu, buf)
		return nil
	case 0x25:
		opDECr(cpu, buf)
		return nil
	case 0x26:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDrn(cpu, buf)
		return nil
	case 0x27:
		opDAA(cpu, buf)
		return nil
	case 0x28:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJRZe(cpu, buf)
		return nil
	case 0x29:
		opADDHLss(cpu, buf)
		return nil
	case 0x2a:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDHLnnP(cpu, buf)
		return nil
	case 0x2b:
		opDECss(cpu, buf)
		return nil
	case 0x2c:
		opINCr(cpu, buf)
		return nil
	case 0x2d:
		opDECr(cpu, buf)
		return nil
	case 0x2e:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDrn(cpu, buf)
		return nil
	case 0x2f:
		opCPL(cpu, buf)
		return nil
	case 0x30:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJRNCe(cpu, buf)
		return nil
	case 0x31:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDddnn(cpu, buf)
		return nil
	case 0x32:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDnnPA(cpu, buf)
		return nil
	case 0x33:
		opINCss(cpu, buf)
		return nil
	case 0x34:
		opINCHLP(cpu, buf)
		return nil
	case 0x35:
		opDECHLP(cpu, buf)
		return nil
	case 0x36:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDHLPn(cpu, buf)
		return nil
	case 0x37:
		opSCF(cpu, buf)
		return nil
	case 0x38:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJRCe(cpu, buf)
		return nil
	case 0x39:
		opADDHLss(cpu, buf)
		return nil
	case 0x3a:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDAnnP(cpu, buf)
		return nil
	case 0x3b:
		opDECss(cpu, buf)
		return nil
	case 0x3c:
		opINCr(cpu, buf)
		return nil
	case 0x3d:
		opDECr(cpu, buf)
		return nil
	case 0x3e:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opLDrn(cpu, buf)
		return nil
	case 0x3f:
		opCCF(cpu, buf)
		return nil
	case 0x40:
		opLDr1r2(cpu, buf)
		return nil
	case 0x41:
		opLDr1r2(cpu, buf)
		return nil
	case 0x42:
		opLDr1r2(cpu, buf)
		return nil
	case 0x43:
		opLDr1r2(cpu, buf)
		return nil
	case 0x44:
		opLDr1r2(cpu, buf)
		return nil
	case 0x45:
		opLDr1r2(cpu, buf)
		return nil
	case 0x46:
		opLDrHLP(cpu, buf)
		return nil
	case 0x47:
		opLDr1r2(cpu, buf)
		return nil
	case 0x48:
		opLDr1r2(cpu, buf)
		return nil
	case 0x49:
		opLDr1r2(cpu, buf)
		return nil
	case 0x4a:
		opLDr1r2(cpu, buf)
		return nil
	case 0x4b:
		opLDr1r2(cpu, buf)
		return nil
	case 0x4c:
		opLDr1r2(cpu, buf)
		return nil
	case 0x4d:
		opLDr1r2(cpu, buf)
		return nil
	case 0x4e:
		opLDrHLP(cpu, buf)
		return nil
	case 0x4f:
		opLDr1r2(cpu, buf)
		return nil
	case 0x50:
		opLDr1r2(cpu, buf)
		return nil
	case 0x51:
		opLDr1r2(cpu, buf)
		return nil
	case 0x52:
		opLDr1r2(cpu, buf)
		return nil
	case 0x53:
		opLDr1r2(cpu, buf)
		return nil
	case 0x54:
		opLDr1r2(cpu, buf)
		return nil
	case 0x55:
		opLDr1r2(cpu, buf)
		return nil
	case 0x56:
		opLDrHLP(cpu, buf)
		return nil
	case 0x57:
		opLDr1r2(cpu, buf)
		return nil
	case 0x58:
		opLDr1r2(cpu, buf)
		return nil
	case 0x59:
		opLDr1r2(cpu, buf)
		return nil
	case 0x5a:
		opLDr1r2(cpu, buf)
		return nil
	case 0x5b:
		opLDr1r2(cpu, buf)
		return nil
	case 0x5c:
		opLDr1r2(cpu, buf)
		return nil
	case 0x5d:
		opLDr1r2(cpu, buf)
		return nil
	case 0x5e:
		opLDrHLP(cpu, buf)
		return nil
	case 0x5f:
		opLDr1r2(cpu, buf)
		return nil
	case 0x60:
		opLDr1r2(cpu, buf)
		return nil
	case 0x61:
		opLDr1r2(cpu, buf)
		return nil
	case 0x62:
		opLDr1r2(cpu, buf)
		return nil
	case 0x63:
		opLDr1r2(cpu, buf)
		return nil
	case 0x64:
		opLDr1r2(cpu, buf)
		return nil
	case 0x65:
		opLDr1r2(cpu, buf)
		return nil
	case 0x66:
		opLDrHLP(cpu, buf)
		return nil
	case 0x67:
		opLDr1r2(cpu, buf)
		return nil
	case 0x68:
		opLDr1r2(cpu, buf)
		return nil
	case 0x69:
		opLDr1r2(cpu, buf)
		return nil
	case 0x6a:
		opLDr1r2(cpu, buf)
		return nil
	case 0x6b:
		opLDr1r2(cpu, buf)
		return nil
	case 0x6c:
		opLDr1r2(cpu, buf)
		return nil
	case 0x6d:
		opLDr1r2(cpu, buf)
		return nil
	case 0x6e:
		opLDrHLP(cpu, buf)
		return nil
	case 0x6f:
		opLDr1r2(cpu, buf)
		return nil
	case 0x70:
		opLDHLPr(cpu, buf)
		return nil
	case 0x71:
		opLDHLPr(cpu, buf)
		return nil
	case 0x72:
		opLDHLPr(cpu, buf)
		return nil
	case 0x73:
		opLDHLPr(cpu, buf)
		return nil
	case 0x74:
		opLDHLPr(cpu, buf)
		return nil
	case 0x75:
		opLDHLPr(cpu, buf)
		return nil
	case 0x76:
		opHALT(cpu, buf)
		return nil
	case 0x77:
		opLDHLPr(cpu, buf)
		return nil
	case 0x78:
		opLDr1r2(cpu, buf)
		return nil
	case 0x79:
		opLDr1r2(cpu, buf)
		return nil
	case 0x7a:
		opLDr1r2(cpu, buf)
		return nil
	case 0x7b:
		opLDr1r2(cpu, buf)
		return nil
	case 0x7c:
		opLDr1r2(cpu, buf)
		return nil
	case 0x7d:
		opLDr1r2(cpu, buf)
		return nil
	case 0x7e:
		opLDrHLP(cpu, buf)
		return nil
	case 0x7f:
		opLDr1r2(cpu, buf)
		return nil
	case 0x80:
		opADDAr(cpu, buf)
		return nil
	case 0x81:
		opADDAr(cpu, buf)
		return nil
	case 0x82:
		opADDAr(cpu, buf)
		return nil
	case 0x83:
		opADDAr(cpu, buf)
		return nil
	case 0x84:
		opADDAr(cpu, buf)
		return nil
	case 0x85:
		opADDAr(cpu, buf)
		return nil
	case 0x86:
		opADDAHLP(cpu, buf)
		return nil
	case 0x87:
		opADDAr(cpu, buf)
		return nil
	case 0x88:
		opADCAr(cpu, buf)
		return nil
	case 0x89:
		opADCAr(cpu, buf)
		return nil
	case 0x8a:
		opADCAr(cpu, buf)
		return nil
	case 0x8b:
		opADCAr(cpu, buf)
		return nil
	case 0x8c:
		opADCAr(cpu, buf)
		return nil
	case 0x8d:
		opADCAr(cpu, buf)
		return nil
	case 0x8e:
		opADCAHLP(cpu, buf)
		return nil
	case 0x8f:
		opADCAr(cpu, buf)
		return nil
	case 0x90:
		opSUBAr(cpu, buf)
		return nil
	case 0x91:
		opSUBAr(cpu, buf)
		return nil
	case 0x92:
		opSUBAr(cpu, buf)
		return nil
	case 0x93:
		opSUBAr(cpu, buf)
		return nil
	case 0x94:
		opSUBAr(cpu, buf)
		return nil
	case 0x95:
		opSUBAr(cpu, buf)
		return nil
	case 0x96:
		opSUBAHLP(cpu, buf)
		return nil
	case 0x97:
		opSUBAr(cpu, buf)
		return nil
	case 0x98:
		opSBCAr(cpu, buf)
		return nil
	case 0x99:
		opSBCAr(cpu, buf)
		return nil
	case 0x9a:
		opSBCAr(cpu, buf)
		return nil
	case 0x9b:
		opSBCAr(cpu, buf)
		return nil
	case 0x9c:
		opSBCAr(cpu, buf)
		return nil
	case 0x9d:
		opSBCAr(cpu, buf)
		return nil
	case 0x9e:
		opSBCAHLP(cpu, buf)
		return nil
	case 0x9f:
		opSBCAr(cpu, buf)
		return nil
	case 0xa0:
		opANDr(cpu, buf)
		return nil
	case 0xa1:
		opANDr(cpu, buf)
		return nil
	case 0xa2:
		opANDr(cpu, buf)
		return nil
	case 0xa3:
		opANDr(cpu, buf)
		return nil
	case 0xa4:
		opANDr(cpu, buf)
		return nil
	case 0xa5:
		opANDr(cpu, buf)
		return nil
	case 0xa6:
		opANDHLP(cpu, buf)
		return nil
	case 0xa7:
		opANDr(cpu, buf)
		return nil
	case 0xa8:
		opXORr(cpu, buf)
		return nil
	case 0xa9:
		opXORr(cpu, buf)
		return nil
	case 0xaa:
		opXORr(cpu, buf)
		return nil
	case 0xab:
		opXORr(cpu, buf)
		return nil
	case 0xac:
		opXORr(cpu, buf)
		return nil
	case 0xad:
		opXORr(cpu, buf)
		return nil
	case 0xae:
		opXORHLP(cpu, buf)
		return nil
	case 0xaf:
		opXORr(cpu, buf)
		return nil
	case 0xb0:
		opORr(cpu, buf)
		return nil
	case 0xb1:
		opORr(cpu, buf)
		return nil
	case 0xb2:
		opORr(cpu, buf)
		return nil
	case 0xb3:
		opORr(cpu, buf)
		return nil
	case 0xb4:
		opORr(cpu, buf)
		return nil
	case 0xb5:
		opORr(cpu, buf)
		return nil
	case 0xb6:
		opORHLP(cpu, buf)
		return nil
	case 0xb7:
		opORr(cpu, buf)
		return nil
	case 0xb8:
		opCPr(cpu, buf)
		return nil
	case 0xb9:
		opCPr(cpu, buf)
		return nil
	case 0xba:
		opCPr(cpu, buf)
		return nil
	case 0xbb:
		opCPr(cpu, buf)
		return nil
	case 0xbc:
		opCPr(cpu, buf)
		return nil
	case 0xbd:
		opCPr(cpu, buf)
		return nil
	case 0xbe:
		opCPHLP(cpu, buf)
		return nil
	case 0xbf:
		opCPr(cpu, buf)
		return nil
	case 0xc0:
		opRETcc(cpu, buf)
		return nil
	case 0xc1:
		opPOPqq(cpu, buf)
		return nil
	case 0xc2:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xc3:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPnn(cpu, buf)
		return nil
	case 0xc4:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xc5:
		opPUSHqq(cpu, buf)
		return nil
	case 0xc6:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opADDAn(cpu, buf)
		return nil
	case 0xc7:
		opRSTp(cpu, buf)
		return nil
	case 0xc8:
		opRETcc(cpu, buf)
		return nil
	case 0xc9:
		opRET(cpu, buf)
		return nil
	case 0xca:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xcb:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		switch b {
		case 0x00:
			opRLCr(cpu, buf)
			return nil
		case 0x01:
			opRLCr(cpu, buf)
			return nil
		case 0x02:
			opRLCr(cpu, buf)
			return nil
		case 0x03:
			opRLCr(cpu, buf)
			return nil
		case 0x04:
			opRLCr(cpu, buf)
			return nil
		case 0x05:
			opRLCr(cpu, buf)
			return nil
		case 0x06:
			opRLCHLP(cpu, buf)
			return nil
		case 0x07:
			opRLCr(cpu, buf)
			return nil
		case 0x08:
			opRRCr(cpu, buf)
			return nil
		case 0x09:
			opRRCr(cpu, buf)
			return nil
		case 0x0a:
			opRRCr(cpu, buf)
			return nil
		case 0x0b:
			opRRCr(cpu, buf)
			return nil
		case 0x0c:
			opRRCr(cpu, buf)
			return nil
		case 0x0d:
			opRRCr(cpu, buf)
			return nil
		case 0x0e:
			opRRCHLP(cpu, buf)
			return nil
		case 0x0f:
			opRRCr(cpu, buf)
			return nil
		case 0x10:
			opRLr(cpu, buf)
			return nil
		case 0x11:
			opRLr(cpu, buf)
			return nil
		case 0x12:
			opRLr(cpu, buf)
			return nil
		case 0x13:
			opRLr(cpu, buf)
			return nil
		case 0x14:
			opRLr(cpu, buf)
			return nil
		case 0x15:
			opRLr(cpu, buf)
			return nil
		case 0x16:
			opRLHLP(cpu, buf)
			return nil
		case 0x17:
			opRLr(cpu, buf)
			return nil
		case 0x18:
			opRRr(cpu, buf)
			return nil
		case 0x19:
			opRRr(cpu, buf)
			return nil
		case 0x1a:
			opRRr(cpu, buf)
			return nil
		case 0x1b:
			opRRr(cpu, buf)
			return nil
		case 0x1c:
			opRRr(cpu, buf)
			return nil
		case 0x1d:
			opRRr(cpu, buf)
			return nil
		case 0x1e:
			opRRHLP(cpu, buf)
			return nil
		case 0x1f:
			opRRr(cpu, buf)
			return nil
		case 0x20:
			opSLAr(cpu, buf)
			return nil
		case 0x21:
			opSLAr(cpu, buf)
			return nil
		case 0x22:
			opSLAr(cpu, buf)
			return nil
		case 0x23:
			opSLAr(cpu, buf)
			return nil
		case 0x24:
			opSLAr(cpu, buf)
			return nil
		case 0x25:
			opSLAr(cpu, buf)
			return nil
		case 0x26:
			opSLAHLP(cpu, buf)
			return nil
		case 0x27:
			opSLAr(cpu, buf)
			return nil
		case 0x28:
			opSRAr(cpu, buf)
			return nil
		case 0x29:
			opSRAr(cpu, buf)
			return nil
		case 0x2a:
			opSRAr(cpu, buf)
			return nil
		case 0x2b:
			opSRAr(cpu, buf)
			return nil
		case 0x2c:
			opSRAr(cpu, buf)
			return nil
		case 0x2d:
			opSRAr(cpu, buf)
			return nil
		case 0x2e:
			opSRAHLP(cpu, buf)
			return nil
		case 0x2f:
			opSRAr(cpu, buf)
			return nil
		case 0x30:
			opSL1r(cpu, buf)
			return nil
		case 0x31:
			opSL1r(cpu, buf)
			return nil
		case 0x32:
			opSL1r(cpu, buf)
			return nil
		case 0x33:
			opSL1r(cpu, buf)
			return nil
		case 0x34:
			opSL1r(cpu, buf)
			return nil
		case 0x35:
			opSL1r(cpu, buf)
			return nil
		case 0x36:
			opSL1HLP(cpu, buf)
			return nil
		case 0x37:
			opSL1r(cpu, buf)
			return nil
		case 0x38:
			opSRLr(cpu, buf)
			return nil
		case 0x39:
			opSRLr(cpu, buf)
			return nil
		case 0x3a:
			opSRLr(cpu, buf)
			return nil
		case 0x3b:
			opSRLr(cpu, buf)
			return nil
		case 0x3c:
			opSRLr(cpu, buf)
			return nil
		case 0x3d:
			opSRLr(cpu, buf)
			return nil
		case 0x3e:
			opSRLHLP(cpu, buf)
			return nil
		case 0x3f:
			opSRLr(cpu, buf)
			return nil
		case 0x40:
			opBITbr(cpu, buf)
			return nil
		case 0x41:
			opBITbr(cpu, buf)
			return nil
		case 0x42:
			opBITbr(cpu, buf)
			return nil
		case 0x43:
			opBITbr(cpu, buf)
			return nil
		case 0x44:
			opBITbr(cpu, buf)
			return nil
		case 0x45:
			opBITbr(cpu, buf)
			return nil
		case 0x46:
			opBITbHLP(cpu, buf)
			return nil
		case 0x47:
			opBITbr(cpu, buf)
			return nil
		case 0x48:
			opBITbr(cpu, buf)
			return nil
		case 0x49:
			opBITbr(cpu, buf)
			return nil
		case 0x4a:
			opBITbr(cpu, buf)
			return nil
		case 0x4b:
			opBITbr(cpu, buf)
			return nil
		case 0x4c:
			opBITbr(cpu, buf)
			return nil
		case 0x4d:
			opBITbr(cpu, buf)
			return nil
		case 0x4e:
			opBITbHLP(cpu, buf)
			return nil
		case 0x4f:
			opBITbr(cpu, buf)
			return nil
		case 0x50:
			opBITbr(cpu, buf)
			return nil
		case 0x51:
			opBITbr(cpu, buf)
			return nil
		case 0x52:
			opBITbr(cpu, buf)
			return nil
		case 0x53:
			opBITbr(cpu, buf)
			return nil
		case 0x54:
			opBITbr(cpu, buf)
			return nil
		case 0x55:
			opBITbr(cpu, buf)
			return nil
		case 0x56:
			opBITbHLP(cpu, buf)
			return nil
		case 0x57:
			opBITbr(cpu, buf)
			return nil
		case 0x58:
			opBITbr(cpu, buf)
			return nil
		case 0x59:
			opBITbr(cpu, buf)
			return nil
		case 0x5a:
			opBITbr(cpu, buf)
			return nil
		case 0x5b:
			opBITbr(cpu, buf)
			return nil
		case 0x5c:
			opBITbr(cpu, buf)
			return nil
		case 0x5d:
			opBITbr(cpu, buf)
			return nil
		case 0x5e:
			opBITbHLP(cpu, buf)
			return nil
		case 0x5f:
			opBITbr(cpu, buf)
			return nil
		case 0x60:
			opBITbr(cpu, buf)
			return nil
		case 0x61:
			opBITbr(cpu, buf)
			return nil
		case 0x62:
			opBITbr(cpu, buf)
			return nil
		case 0x63:
			opBITbr(cpu, buf)
			return nil
		case 0x64:
			opBITbr(cpu, buf)
			return nil
		case 0x65:
			opBITbr(cpu, buf)
			return nil
		case 0x66:
			opBITbHLP(cpu, buf)
			return nil
		case 0x67:
			opBITbr(cpu, buf)
			return nil
		case 0x68:
			opBITbr(cpu, buf)
			return nil
		case 0x69:
			opBITbr(cpu, buf)
			return nil
		case 0x6a:
			opBITbr(cpu, buf)
			return nil
		case 0x6b:
			opBITbr(cpu, buf)
			return nil
		case 0x6c:
			opBITbr(cpu, buf)
			return nil
		case 0x6d:
			opBITbr(cpu, buf)
			return nil
		case 0x6e:
			opBITbHLP(cpu, buf)
			return nil
		case 0x6f:
			opBITbr(cpu, buf)
			return nil
		case 0x70:
			opBITbr(cpu, buf)
			return nil
		case 0x71:
			opBITbr(cpu, buf)
			return nil
		case 0x72:
			opBITbr(cpu, buf)
			return nil
		case 0x73:
			opBITbr(cpu, buf)
			return nil
		case 0x74:
			opBITbr(cpu, buf)
			return nil
		case 0x75:
			opBITbr(cpu, buf)
			return nil
		case 0x76:
			opBITbHLP(cpu, buf)
			return nil
		case 0x77:
			opBITbr(cpu, buf)
			return nil
		case 0x78:
			opBITbr(cpu, buf)
			return nil
		case 0x79:
			opBITbr(cpu, buf)
			return nil
		case 0x7a:
			opBITbr(cpu, buf)
			return nil
		case 0x7b:
			opBITbr(cpu, buf)
			return nil
		case 0x7c:
			opBITbr(cpu, buf)
			return nil
		case 0x7d:
			opBITbr(cpu, buf)
			return nil
		case 0x7e:
			opBITbHLP(cpu, buf)
			return nil
		case 0x7f:
			opBITbr(cpu, buf)
			return nil
		case 0x80:
			opRESbr(cpu, buf)
			return nil
		case 0x81:
			opRESbr(cpu, buf)
			return nil
		case 0x82:
			opRESbr(cpu, buf)
			return nil
		case 0x83:
			opRESbr(cpu, buf)
			return nil
		case 0x84:
			opRESbr(cpu, buf)
			return nil
		case 0x85:
			opRESbr(cpu, buf)
			return nil
		case 0x86:
			opRESbHLP(cpu, buf)
			return nil
		case 0x87:
			opRESbr(cpu, buf)
			return nil
		case 0x88:
			opRESbr(cpu, buf)
			return nil
		case 0x89:
			opRESbr(cpu, buf)
			return nil
		case 0x8a:
			opRESbr(cpu, buf)
			return nil
		case 0x8b:
			opRESbr(cpu, buf)
			return nil
		case 0x8c:
			opRESbr(cpu, buf)
			return nil
		case 0x8d:
			opRESbr(cpu, buf)
			return nil
		case 0x8e:
			opRESbHLP(cpu, buf)
			return nil
		case 0x8f:
			opRESbr(cpu, buf)
			return nil
		case 0x90:
			opRESbr(cpu, buf)
			return nil
		case 0x91:
			opRESbr(cpu, buf)
			return nil
		case 0x92:
			opRESbr(cpu, buf)
			return nil
		case 0x93:
			opRESbr(cpu, buf)
			return nil
		case 0x94:
			opRESbr(cpu, buf)
			return nil
		case 0x95:
			opRESbr(cpu, buf)
			return nil
		case 0x96:
			opRESbHLP(cpu, buf)
			return nil
		case 0x97:
			opRESbr(cpu, buf)
			return nil
		case 0x98:
			opRESbr(cpu, buf)
			return nil
		case 0x99:
			opRESbr(cpu, buf)
			return nil
		case 0x9a:
			opRESbr(cpu, buf)
			return nil
		case 0x9b:
			opRESbr(cpu, buf)
			return nil
		case 0x9c:
			opRESbr(cpu, buf)
			return nil
		case 0x9d:
			opRESbr(cpu, buf)
			return nil
		case 0x9e:
			opRESbHLP(cpu, buf)
			return nil
		case 0x9f:
			opRESbr(cpu, buf)
			return nil
		case 0xa0:
			opRESbr(cpu, buf)
			return nil
		case 0xa1:
			opRESbr(cpu, buf)
			return nil
		case 0xa2:
			opRESbr(cpu, buf)
			return nil
		case 0xa3:
			opRESbr(cpu, buf)
			return nil
		case 0xa4:
			opRESbr(cpu, buf)
			return nil
		case 0xa5:
			opRESbr(cpu, buf)
			return nil
		case 0xa6:
			opRESbHLP(cpu, buf)
			return nil
		case 0xa7:
			opRESbr(cpu, buf)
			return nil
		case 0xa8:
			opRESbr(cpu, buf)
			return nil
		case 0xa9:
			opRESbr(cpu, buf)
			return nil
		case 0xaa:
			opRESbr(cpu, buf)
			return nil
		case 0xab:
			opRESbr(cpu, buf)
			return nil
		case 0xac:
			opRESbr(cpu, buf)
			return nil
		case 0xad:
			opRESbr(cpu, buf)
			return nil
		case 0xae:
			opRESbHLP(cpu, buf)
			return nil
		case 0xaf:
			opRESbr(cpu, buf)
			return nil
		case 0xb0:
			opRESbr(cpu, buf)
			return nil
		case 0xb1:
			opRESbr(cpu, buf)
			return nil
		case 0xb2:
			opRESbr(cpu, buf)
			return nil
		case 0xb3:
			opRESbr(cpu, buf)
			return nil
		case 0xb4:
			opRESbr(cpu, buf)
			return nil
		case 0xb5:
			opRESbr(cpu, buf)
			return nil
		case 0xb6:
			opRESbHLP(cpu, buf)
			return nil
		case 0xb7:
			opRESbr(cpu, buf)
			return nil
		case 0xb8:
			opRESbr(cpu, buf)
			return nil
		case 0xb9:
			opRESbr(cpu, buf)
			return nil
		case 0xba:
			opRESbr(cpu, buf)
			return nil
		case 0xbb:
			opRESbr(cpu, buf)
			return nil
		case 0xbc:
			opRESbr(cpu, buf)
			return nil
		case 0xbd:
			opRESbr(cpu, buf)
			return nil
		case 0xbe:
			opRESbHLP(cpu, buf)
			return nil
		case 0xbf:
			opRESbr(cpu, buf)
			return nil
		case 0xc0:
			opSETbr(cpu, buf)
			return nil
		case 0xc1:
			opSETbr(cpu, buf)
			return nil
		case 0xc2:
			opSETbr(cpu, buf)
			return nil
		case 0xc3:
			opSETbr(cpu, buf)
			return nil
		case 0xc4:
			opSETbr(cpu, buf)
			return nil
		case 0xc5:
			opSETbr(cpu, buf)
			return nil
		case 0xc6:
			opSETbHLP(cpu, buf)
			return nil
		case 0xc7:
			opSETbr(cpu, buf)
			return nil
		case 0xc8:
			opSETbr(cpu, buf)
			return nil
		case 0xc9:
			opSETbr(cpu, buf)
			return nil
		case 0xca:
			opSETbr(cpu, buf)
			return nil
		case 0xcb:
			opSETbr(cpu, buf)
			return nil
		case 0xcc:
			opSETbr(cpu, buf)
			return nil
		case 0xcd:
			opSETbr(cpu, buf)
			return nil
		case 0xce:
			opSETbHLP(cpu, buf)
			return nil
		case 0xcf:
			opSETbr(cpu, buf)
			return nil
		case 0xd0:
			opSETbr(cpu, buf)
			return nil
		case 0xd1:
			opSETbr(cpu, buf)
			return nil
		case 0xd2:
			opSETbr(cpu, buf)
			return nil
		case 0xd3:
			opSETbr(cpu, buf)
			return nil
		case 0xd4:
			opSETbr(cpu, buf)
			return nil
		case 0xd5:
			opSETbr(cpu, buf)
			return nil
		case 0xd6:
			opSETbHLP(cpu, buf)
			return nil
		case 0xd7:
			opSETbr(cpu, buf)
			return nil
		case 0xd8:
			opSETbr(cpu, buf)
			return nil
		case 0xd9:
			opSETbr(cpu, buf)
			return nil
		case 0xda:
			opSETbr(cpu, buf)
			return nil
		case 0xdb:
			opSETbr(cpu, buf)
			return nil
		case 0xdc:
			opSETbr(cpu, buf)
			return nil
		case 0xdd:
			opSETbr(cpu, buf)
			return nil
		case 0xde:
			opSETbHLP(cpu, buf)
			return nil
		case 0xdf:
			opSETbr(cpu, buf)
			return nil
		case 0xe0:
			opSETbr(cpu, buf)
			return nil
		case 0xe1:
			opSETbr(cpu, buf)
			return nil
		case 0xe2:
			opSETbr(cpu, buf)
			return nil
		case 0xe3:
			opSETbr(cpu, buf)
			return nil
		case 0xe4:
			opSETbr(cpu, buf)
			return nil
		case 0xe5:
			opSETbr(cpu, buf)
			return nil
		case 0xe6:
			opSETbHLP(cpu, buf)
			return nil
		case 0xe7:
			opSETbr(cpu, buf)
			return nil
		case 0xe8:
			opSETbr(cpu, buf)
			return nil
		case 0xe9:
			opSETbr(cpu, buf)
			return nil
		case 0xea:
			opSETbr(cpu, buf)
			return nil
		case 0xeb:
			opSETbr(cpu, buf)
			return nil
		case 0xec:
			opSETbr(cpu, buf)
			return nil
		case 0xed:
			opSETbr(cpu, buf)
			return nil
		case 0xee:
			opSETbHLP(cpu, buf)
			return nil
		case 0xef:
			opSETbr(cpu, buf)
			return nil
		case 0xf0:
			opSETbr(cpu, buf)
			return nil
		case 0xf1:
			opSETbr(cpu, buf)
			return nil
		case 0xf2:
			opSETbr(cpu, buf)
			return nil
		case 0xf3:
			opSETbr(cpu, buf)
			return nil
		case 0xf4:
			opSETbr(cpu, buf)
			return nil
		case 0xf5:
			opSETbr(cpu, buf)
			return nil
		case 0xf6:
			opSETbHLP(cpu, buf)
			return nil
		case 0xf7:
			opSETbr(cpu, buf)
			return nil
		case 0xf8:
			opSETbr(cpu, buf)
			return nil
		case 0xf9:
			opSETbr(cpu, buf)
			return nil
		case 0xfa:
			opSETbr(cpu, buf)
			return nil
		case 0xfb:
			opSETbr(cpu, buf)
			return nil
		case 0xfc:
			opSETbr(cpu, buf)
			return nil
		case 0xfd:
			opSETbr(cpu, buf)
			return nil
		case 0xfe:
			opSETbHLP(cpu, buf)
			return nil
		case 0xff:
			opSETbr(cpu, buf)
			return nil
		default:
			return ErrInvalidCodes
		}
		return nil
	case 0xcc:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xcd:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLnn(cpu, buf)
		return nil
	case 0xce:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opADCAn(cpu, buf)
		return nil
	case 0xcf:
		opRSTp(cpu, buf)
		return nil
	case 0xd0:
		opRETcc(cpu, buf)
		return nil
	case 0xd1:
		opPOPqq(cpu, buf)
		return nil
	case 0xd2:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xd3:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opOUTnPA(cpu, buf)
		return nil
	case 0xd4:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xd5:
		opPUSHqq(cpu, buf)
		return nil
	case 0xd6:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opSUBAn(cpu, buf)
		return nil
	case 0xd7:
		opRSTp(cpu, buf)
		return nil
	case 0xd8:
		opRETcc(cpu, buf)
		return nil
	case 0xd9:
		opEXX(cpu, buf)
		return nil
	case 0xda:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xdb:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opINAnP(cpu, buf)
		return nil
	case 0xdc:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xdd:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		switch b {
		case 0x09:
			opADDIXpp(cpu, buf)
			return nil
		case 0x19:
			opADDIXpp(cpu, buf)
			return nil
		case 0x21:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXnn(cpu, buf)
			return nil
		case 0x22:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
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
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXHn(cpu, buf)
			return nil
		case 0x29:
			opADDIXpp(cpu, buf)
			return nil
		case 0x2a:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
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
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXLn(cpu, buf)
			return nil
		case 0x34:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opINCIXdP(cpu, buf)
			return nil
		case 0x35:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opDECIXdP(cpu, buf)
			return nil
		case 0x36:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPn(cpu, buf)
			return nil
		case 0x39:
			opADDIXpp(cpu, buf)
			return nil
		case 0x40:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x41:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x42:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x43:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x44:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x45:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x46:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIXdP(cpu, buf)
			return nil
		case 0x47:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x48:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x49:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x4a:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x4b:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x4c:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x4d:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x4e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIXdP(cpu, buf)
			return nil
		case 0x4f:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x50:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x51:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x52:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x53:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x54:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x55:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x56:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIXdP(cpu, buf)
			return nil
		case 0x57:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x58:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x59:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x5a:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x5b:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x5c:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x5d:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x5e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIXdP(cpu, buf)
			return nil
		case 0x5f:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x60:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x61:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x62:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x63:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x64:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x65:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x66:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIXdP(cpu, buf)
			return nil
		case 0x67:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x68:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x69:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x6a:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x6b:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x6c:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x6d:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x6e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIXdP(cpu, buf)
			return nil
		case 0x6f:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x70:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPr(cpu, buf)
			return nil
		case 0x71:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPr(cpu, buf)
			return nil
		case 0x72:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPr(cpu, buf)
			return nil
		case 0x73:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPr(cpu, buf)
			return nil
		case 0x74:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPr(cpu, buf)
			return nil
		case 0x75:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPr(cpu, buf)
			return nil
		case 0x77:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIXdPr(cpu, buf)
			return nil
		case 0x78:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x79:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x7a:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x7b:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x7c:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x7d:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x7e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIXdP(cpu, buf)
			return nil
		case 0x7f:
			opLDrx1rx2(cpu, buf)
			return nil
		case 0x80:
			opADDArx(cpu, buf)
			return nil
		case 0x81:
			opADDArx(cpu, buf)
			return nil
		case 0x82:
			opADDArx(cpu, buf)
			return nil
		case 0x83:
			opADDArx(cpu, buf)
			return nil
		case 0x84:
			opADDArx(cpu, buf)
			return nil
		case 0x85:
			opADDArx(cpu, buf)
			return nil
		case 0x86:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opADDAIXdP(cpu, buf)
			return nil
		case 0x87:
			opADDArx(cpu, buf)
			return nil
		case 0x88:
			opADCArx(cpu, buf)
			return nil
		case 0x89:
			opADCArx(cpu, buf)
			return nil
		case 0x8a:
			opADCArx(cpu, buf)
			return nil
		case 0x8b:
			opADCArx(cpu, buf)
			return nil
		case 0x8c:
			opADCArx(cpu, buf)
			return nil
		case 0x8d:
			opADCArx(cpu, buf)
			return nil
		case 0x8e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opADCAIXdP(cpu, buf)
			return nil
		case 0x8f:
			opADCArx(cpu, buf)
			return nil
		case 0x90:
			opSUBArx(cpu, buf)
			return nil
		case 0x91:
			opSUBArx(cpu, buf)
			return nil
		case 0x92:
			opSUBArx(cpu, buf)
			return nil
		case 0x93:
			opSUBArx(cpu, buf)
			return nil
		case 0x94:
			opSUBArx(cpu, buf)
			return nil
		case 0x95:
			opSUBArx(cpu, buf)
			return nil
		case 0x96:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opSUBAIXdP(cpu, buf)
			return nil
		case 0x97:
			opSUBArx(cpu, buf)
			return nil
		case 0x98:
			opSBCArx(cpu, buf)
			return nil
		case 0x99:
			opSBCArx(cpu, buf)
			return nil
		case 0x9a:
			opSBCArx(cpu, buf)
			return nil
		case 0x9b:
			opSBCArx(cpu, buf)
			return nil
		case 0x9c:
			opSBCArx(cpu, buf)
			return nil
		case 0x9d:
			opSBCArx(cpu, buf)
			return nil
		case 0x9e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opSBCAIXdP(cpu, buf)
			return nil
		case 0x9f:
			opSBCArx(cpu, buf)
			return nil
		case 0xa0:
			opANDrx(cpu, buf)
			return nil
		case 0xa1:
			opANDrx(cpu, buf)
			return nil
		case 0xa2:
			opANDrx(cpu, buf)
			return nil
		case 0xa3:
			opANDrx(cpu, buf)
			return nil
		case 0xa4:
			opANDrx(cpu, buf)
			return nil
		case 0xa5:
			opANDrx(cpu, buf)
			return nil
		case 0xa6:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opANDIXdP(cpu, buf)
			return nil
		case 0xa7:
			opANDrx(cpu, buf)
			return nil
		case 0xa8:
			opXORrx(cpu, buf)
			return nil
		case 0xa9:
			opXORrx(cpu, buf)
			return nil
		case 0xaa:
			opXORrx(cpu, buf)
			return nil
		case 0xab:
			opXORrx(cpu, buf)
			return nil
		case 0xac:
			opXORrx(cpu, buf)
			return nil
		case 0xad:
			opXORrx(cpu, buf)
			return nil
		case 0xae:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opXORIXdP(cpu, buf)
			return nil
		case 0xaf:
			opXORrx(cpu, buf)
			return nil
		case 0xb0:
			opORrx(cpu, buf)
			return nil
		case 0xb1:
			opORrx(cpu, buf)
			return nil
		case 0xb2:
			opORrx(cpu, buf)
			return nil
		case 0xb3:
			opORrx(cpu, buf)
			return nil
		case 0xb4:
			opORrx(cpu, buf)
			return nil
		case 0xb5:
			opORrx(cpu, buf)
			return nil
		case 0xb6:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opORIXdP(cpu, buf)
			return nil
		case 0xb7:
			opORrx(cpu, buf)
			return nil
		case 0xb8:
			opCPrx(cpu, buf)
			return nil
		case 0xb9:
			opCPrx(cpu, buf)
			return nil
		case 0xba:
			opCPrx(cpu, buf)
			return nil
		case 0xbb:
			opCPrx(cpu, buf)
			return nil
		case 0xbc:
			opCPrx(cpu, buf)
			return nil
		case 0xbd:
			opCPrx(cpu, buf)
			return nil
		case 0xbe:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opCPIXdP(cpu, buf)
			return nil
		case 0xbf:
			opCPrx(cpu, buf)
			return nil
		case 0xcb:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
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
			case 0x46:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x4e:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x56:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x5e:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x66:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x6e:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x76:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x7e:
				opBITbIXdP(cpu, buf)
				return nil
			case 0x86:
				opRESbIXdP(cpu, buf)
				return nil
			case 0x8e:
				opRESbIXdP(cpu, buf)
				return nil
			case 0x96:
				opRESbIXdP(cpu, buf)
				return nil
			case 0x9e:
				opRESbIXdP(cpu, buf)
				return nil
			case 0xa6:
				opRESbIXdP(cpu, buf)
				return nil
			case 0xae:
				opRESbIXdP(cpu, buf)
				return nil
			case 0xb6:
				opRESbIXdP(cpu, buf)
				return nil
			case 0xbe:
				opRESbIXdP(cpu, buf)
				return nil
			case 0xc6:
				opSETbIXdP(cpu, buf)
				return nil
			case 0xce:
				opSETbIXdP(cpu, buf)
				return nil
			case 0xd6:
				opSETbIXdP(cpu, buf)
				return nil
			case 0xde:
				opSETbIXdP(cpu, buf)
				return nil
			case 0xe6:
				opSETbIXdP(cpu, buf)
				return nil
			case 0xee:
				opSETbIXdP(cpu, buf)
				return nil
			case 0xf6:
				opSETbIXdP(cpu, buf)
				return nil
			case 0xfe:
				opSETbIXdP(cpu, buf)
				return nil
			default:
				return ErrInvalidCodes
			}
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
		default:
			return ErrInvalidCodes
		}
		return nil
	case 0xde:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opSBCAn(cpu, buf)
		return nil
	case 0xdf:
		opRSTp(cpu, buf)
		return nil
	case 0xe0:
		opRETcc(cpu, buf)
		return nil
	case 0xe1:
		opPOPqq(cpu, buf)
		return nil
	case 0xe2:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xe3:
		opEXSPPHL(cpu, buf)
		return nil
	case 0xe4:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xe5:
		opPUSHqq(cpu, buf)
		return nil
	case 0xe6:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opANDn(cpu, buf)
		return nil
	case 0xe7:
		opRSTp(cpu, buf)
		return nil
	case 0xe8:
		opRETcc(cpu, buf)
		return nil
	case 0xe9:
		opJPHLP(cpu, buf)
		return nil
	case 0xea:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xeb:
		opEXDEHL(cpu, buf)
		return nil
	case 0xec:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xed:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		switch b {
		case 0x40:
			opINrCP(cpu, buf)
			return nil
		case 0x41:
			opOUTCPr(cpu, buf)
			return nil
		case 0x42:
			opSBCHLss(cpu, buf)
			return nil
		case 0x43:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
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
		case 0x48:
			opINrCP(cpu, buf)
			return nil
		case 0x49:
			opOUTCPr(cpu, buf)
			return nil
		case 0x4a:
			opADCHLss(cpu, buf)
			return nil
		case 0x4b:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDddnnP(cpu, buf)
			return nil
		case 0x4d:
			opRETI(cpu, buf)
			return nil
		case 0x4f:
			opLDRA(cpu, buf)
			return nil
		case 0x50:
			opINrCP(cpu, buf)
			return nil
		case 0x51:
			opOUTCPr(cpu, buf)
			return nil
		case 0x52:
			opSBCHLss(cpu, buf)
			return nil
		case 0x53:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDnnPdd(cpu, buf)
			return nil
		case 0x56:
			opIM1(cpu, buf)
			return nil
		case 0x57:
			opLDAI(cpu, buf)
			return nil
		case 0x58:
			opINrCP(cpu, buf)
			return nil
		case 0x59:
			opOUTCPr(cpu, buf)
			return nil
		case 0x5a:
			opADCHLss(cpu, buf)
			return nil
		case 0x5b:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDddnnP(cpu, buf)
			return nil
		case 0x5e:
			opIM2(cpu, buf)
			return nil
		case 0x5f:
			opLDAR(cpu, buf)
			return nil
		case 0x60:
			opINrCP(cpu, buf)
			return nil
		case 0x61:
			opOUTCPr(cpu, buf)
			return nil
		case 0x62:
			opSBCHLss(cpu, buf)
			return nil
		case 0x63:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDnnPdd(cpu, buf)
			return nil
		case 0x67:
			opRRD(cpu, buf)
			return nil
		case 0x68:
			opINrCP(cpu, buf)
			return nil
		case 0x69:
			opOUTCPr(cpu, buf)
			return nil
		case 0x6a:
			opADCHLss(cpu, buf)
			return nil
		case 0x6b:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDddnnP(cpu, buf)
			return nil
		case 0x6f:
			opRLD(cpu, buf)
			return nil
		case 0x72:
			opSBCHLss(cpu, buf)
			return nil
		case 0x73:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDnnPdd(cpu, buf)
			return nil
		case 0x78:
			opINrCP(cpu, buf)
			return nil
		case 0x79:
			opOUTCPr(cpu, buf)
			return nil
		case 0x7a:
			opADCHLss(cpu, buf)
			return nil
		case 0x7b:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDddnnP(cpu, buf)
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
		return nil
	case 0xee:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opXORn(cpu, buf)
		return nil
	case 0xef:
		opRSTp(cpu, buf)
		return nil
	case 0xf0:
		opRETcc(cpu, buf)
		return nil
	case 0xf1:
		opPOPqq(cpu, buf)
		return nil
	case 0xf2:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xf3:
		opDI(cpu, buf)
		return nil
	case 0xf4:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xf5:
		opPUSHqq(cpu, buf)
		return nil
	case 0xf6:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opORn(cpu, buf)
		return nil
	case 0xf7:
		opRSTp(cpu, buf)
		return nil
	case 0xf8:
		opRETcc(cpu, buf)
		return nil
	case 0xf9:
		opLDSPHL(cpu, buf)
		return nil
	case 0xfa:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opJPccnn(cpu, buf)
		return nil
	case 0xfb:
		opEI(cpu, buf)
		return nil
	case 0xfc:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCALLccnn(cpu, buf)
		return nil
	case 0xfd:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		switch b {
		case 0x09:
			opADDIYrr(cpu, buf)
			return nil
		case 0x19:
			opADDIYrr(cpu, buf)
			return nil
		case 0x21:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYnn(cpu, buf)
			return nil
		case 0x22:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
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
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYHn(cpu, buf)
			return nil
		case 0x29:
			opADDIYrr(cpu, buf)
			return nil
		case 0x2a:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
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
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYLn(cpu, buf)
			return nil
		case 0x34:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opINCIYdP(cpu, buf)
			return nil
		case 0x35:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opDECIYdP(cpu, buf)
			return nil
		case 0x36:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPn(cpu, buf)
			return nil
		case 0x39:
			opADDIYrr(cpu, buf)
			return nil
		case 0x40:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x41:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x42:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x43:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x44:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x45:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x46:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIYdP(cpu, buf)
			return nil
		case 0x47:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x48:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x49:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x4a:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x4b:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x4c:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x4d:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x4e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIYdP(cpu, buf)
			return nil
		case 0x4f:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x50:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x51:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x52:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x53:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x54:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x55:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x56:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIYdP(cpu, buf)
			return nil
		case 0x57:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x58:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x59:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x5a:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x5b:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x5c:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x5d:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x5e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIYdP(cpu, buf)
			return nil
		case 0x5f:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x60:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x61:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x62:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x63:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x64:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x65:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x66:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIYdP(cpu, buf)
			return nil
		case 0x67:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x68:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x69:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x6a:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x6b:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x6c:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x6d:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x6e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIYdP(cpu, buf)
			return nil
		case 0x6f:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x70:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPr(cpu, buf)
			return nil
		case 0x71:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPr(cpu, buf)
			return nil
		case 0x72:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPr(cpu, buf)
			return nil
		case 0x73:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPr(cpu, buf)
			return nil
		case 0x74:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPr(cpu, buf)
			return nil
		case 0x75:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPr(cpu, buf)
			return nil
		case 0x77:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDIYdPr(cpu, buf)
			return nil
		case 0x78:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x79:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x7a:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x7b:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x7c:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x7d:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x7e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opLDrIYdP(cpu, buf)
			return nil
		case 0x7f:
			opLDry1ry2(cpu, buf)
			return nil
		case 0x80:
			opADDAry(cpu, buf)
			return nil
		case 0x81:
			opADDAry(cpu, buf)
			return nil
		case 0x82:
			opADDAry(cpu, buf)
			return nil
		case 0x83:
			opADDAry(cpu, buf)
			return nil
		case 0x84:
			opADDAry(cpu, buf)
			return nil
		case 0x85:
			opADDAry(cpu, buf)
			return nil
		case 0x86:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opADDAIYdP(cpu, buf)
			return nil
		case 0x87:
			opADDAry(cpu, buf)
			return nil
		case 0x88:
			opADCAry(cpu, buf)
			return nil
		case 0x89:
			opADCAry(cpu, buf)
			return nil
		case 0x8a:
			opADCAry(cpu, buf)
			return nil
		case 0x8b:
			opADCAry(cpu, buf)
			return nil
		case 0x8c:
			opADCAry(cpu, buf)
			return nil
		case 0x8d:
			opADCAry(cpu, buf)
			return nil
		case 0x8e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opADCAIYdP(cpu, buf)
			return nil
		case 0x8f:
			opADCAry(cpu, buf)
			return nil
		case 0x90:
			opSUBAry(cpu, buf)
			return nil
		case 0x91:
			opSUBAry(cpu, buf)
			return nil
		case 0x92:
			opSUBAry(cpu, buf)
			return nil
		case 0x93:
			opSUBAry(cpu, buf)
			return nil
		case 0x94:
			opSUBAry(cpu, buf)
			return nil
		case 0x95:
			opSUBAry(cpu, buf)
			return nil
		case 0x96:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opSUBAIYdP(cpu, buf)
			return nil
		case 0x97:
			opSUBAry(cpu, buf)
			return nil
		case 0x98:
			opSBCAry(cpu, buf)
			return nil
		case 0x99:
			opSBCAry(cpu, buf)
			return nil
		case 0x9a:
			opSBCAry(cpu, buf)
			return nil
		case 0x9b:
			opSBCAry(cpu, buf)
			return nil
		case 0x9c:
			opSBCAry(cpu, buf)
			return nil
		case 0x9d:
			opSBCAry(cpu, buf)
			return nil
		case 0x9e:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opSBCAIYdP(cpu, buf)
			return nil
		case 0x9f:
			opSBCAry(cpu, buf)
			return nil
		case 0xa0:
			opANDry(cpu, buf)
			return nil
		case 0xa1:
			opANDry(cpu, buf)
			return nil
		case 0xa2:
			opANDry(cpu, buf)
			return nil
		case 0xa3:
			opANDry(cpu, buf)
			return nil
		case 0xa4:
			opANDry(cpu, buf)
			return nil
		case 0xa5:
			opANDry(cpu, buf)
			return nil
		case 0xa6:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opANDIYdP(cpu, buf)
			return nil
		case 0xa7:
			opANDry(cpu, buf)
			return nil
		case 0xa8:
			opXORry(cpu, buf)
			return nil
		case 0xa9:
			opXORry(cpu, buf)
			return nil
		case 0xaa:
			opXORry(cpu, buf)
			return nil
		case 0xab:
			opXORry(cpu, buf)
			return nil
		case 0xac:
			opXORry(cpu, buf)
			return nil
		case 0xad:
			opXORry(cpu, buf)
			return nil
		case 0xae:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opXORIYdP(cpu, buf)
			return nil
		case 0xaf:
			opXORry(cpu, buf)
			return nil
		case 0xb0:
			opORry(cpu, buf)
			return nil
		case 0xb1:
			opORry(cpu, buf)
			return nil
		case 0xb2:
			opORry(cpu, buf)
			return nil
		case 0xb3:
			opORry(cpu, buf)
			return nil
		case 0xb4:
			opORry(cpu, buf)
			return nil
		case 0xb5:
			opORry(cpu, buf)
			return nil
		case 0xb6:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opORIYdP(cpu, buf)
			return nil
		case 0xb7:
			opORry(cpu, buf)
			return nil
		case 0xb8:
			opCPry(cpu, buf)
			return nil
		case 0xb9:
			opCPry(cpu, buf)
			return nil
		case 0xba:
			opCPry(cpu, buf)
			return nil
		case 0xbb:
			opCPry(cpu, buf)
			return nil
		case 0xbc:
			opCPry(cpu, buf)
			return nil
		case 0xbd:
			opCPry(cpu, buf)
			return nil
		case 0xbe:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			opCPIYdP(cpu, buf)
			return nil
		case 0xbf:
			opCPry(cpu, buf)
			return nil
		case 0xcb:
			b, err = f.fetch()
			if err != nil {
				return err
			}
			buf = append(buf, b)
			b, err = f.fetch()
			if err != nil {
				return err
			}
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
			case 0x46:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x4e:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x56:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x5e:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x66:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x6e:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x76:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x7e:
				opBITbIYdP(cpu, buf)
				return nil
			case 0x86:
				opRESbIYdP(cpu, buf)
				return nil
			case 0x8e:
				opRESbIYdP(cpu, buf)
				return nil
			case 0x96:
				opRESbIYdP(cpu, buf)
				return nil
			case 0x9e:
				opRESbIYdP(cpu, buf)
				return nil
			case 0xa6:
				opRESbIYdP(cpu, buf)
				return nil
			case 0xae:
				opRESbIYdP(cpu, buf)
				return nil
			case 0xb6:
				opRESbIYdP(cpu, buf)
				return nil
			case 0xbe:
				opRESbIYdP(cpu, buf)
				return nil
			case 0xc6:
				opSETbIYdP(cpu, buf)
				return nil
			case 0xce:
				opSETbIYdP(cpu, buf)
				return nil
			case 0xd6:
				opSETbIYdP(cpu, buf)
				return nil
			case 0xde:
				opSETbIYdP(cpu, buf)
				return nil
			case 0xe6:
				opSETbIYdP(cpu, buf)
				return nil
			case 0xee:
				opSETbIYdP(cpu, buf)
				return nil
			case 0xf6:
				opSETbIYdP(cpu, buf)
				return nil
			case 0xfe:
				opSETbIYdP(cpu, buf)
				return nil
			default:
				return ErrInvalidCodes
			}
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
		default:
			return ErrInvalidCodes
		}
		return nil
	case 0xfe:
		b, err = f.fetch()
		if err != nil {
			return err
		}
		buf = append(buf, b)
		opCPn(cpu, buf)
		return nil
	case 0xff:
		opRSTp(cpu, buf)
		return nil
	default:
		return ErrInvalidCodes
	}
	return nil
}
