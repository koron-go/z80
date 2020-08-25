package opcode

import (
	"bufio"
	"fmt"
	"io"

	"github.com/koron-go/z80/internal/opname"
)

// WriteSwitchDecoder writes codes for decoder "switch" statements.
func WriteSwitchDecoder(w io.Writer) error {
	bw := bufio.NewWriter(w)
	_, err := bw.WriteString(`package z80

func decodeExec(cpu *CPU, f fetcher) error {
	var b uint8
	buf := cpu.decodeBuf[:4]`)
	if err != nil {
		return err
	}
	l := defaultDecodeLayer()
	nr := 0
	err = writeLayerCode(bw, l, nr)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(bw, "\n}\n")
	if err != nil {
		return err
	}
	return bw.Flush()
}

type nextItem struct {
	d uint8
	l *decodeLayer
}

func writeLayerCode(w *bufio.Writer, l *decodeLayer, nr int) error {
	if l.anyNode != nil {
		fmt.Fprintf(w, "\nbuf[%d] = f.fetch()", nr)
		if l.anyNode.next == nil {
			return fmt.Errorf("invalid any node: %+v", l.anyNode)
		}
		return writeLayerCode(w, l.anyNode.next, nr+1)
	}

	fmt.Fprintf(w, `
b = f.fetch()
buf[%d] = b`, nr)
	nr++
	w.WriteString("\nswitch b {")

	var singleOps []*OPCode
	var codesMap = map[string][]uint8{}
	var nexts []*nextItem
	for i, n := range l.nodes {
		if n == nil {
			continue
		}
		if n.next != nil {
			nexts = append(nexts, &nextItem{d: uint8(i), l: n.next})
			continue
		}
		if n.opcode == nil {
			return fmt.Errorf("node without opcode: %+v", n)
		}
		// store opcode to group.
		op := n.opcode
		codes, ok := codesMap[op.N]
		if !ok {
			singleOps = append(singleOps, op)
		}
		codes = append(codes, uint8(i))
		codesMap[op.N] = codes
	}

	for _, op := range singleOps {
		codes, ok := codesMap[op.N]
		if !ok || len(codes) == 0 {
			panic("something wrong: failed to build codesMap")
		}
		w.WriteString("\ncase")
		for i, code := range codes {
			fmt.Fprintf(w, " 0x%02x", code)
			if i+1 < len(codes) {
				w.WriteRune(',')
			} else {
				w.WriteRune(':')
			}
		}
		if d := len(op.C) - nr; d > 0 {
			for i := 0; i < d; i++ {
				fmt.Fprintf(w, "\nbuf[%d] = f.fetch()", nr+i)
			}
		}
		name := opname.Mangle(op.N)
		fmt.Fprintf(w, "\n%s(cpu, buf[:%d])\nreturn nil", name, len(op.C))
	}

	for _, n := range nexts {
		fmt.Fprintf(w, "\ncase 0x%02x:", n.d)
		writeLayerCode(w, n.l, nr)
	}

	w.WriteString(`
default:
	return ErrInvalidCodes
}`)
	return nil
}
