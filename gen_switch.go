package z80

import (
	"bufio"
	"fmt"
	"io"

	"github.com/koron-go/z80/internal/opname"
)

func GenerateSwitchDecoder(w io.Writer) error {
	bw := bufio.NewWriter(w)
	_, err := bw.WriteString(`package z80

func decodeExec(cpu *CPU, f fetcher) error {
	var err error
	var b uint8
	buf := cpu.decodeBuf[:0]`)
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

func writeLayerCode(w *bufio.Writer, l *decodeLayer, nr int) error {
	_, err := w.WriteString(`
b, err = f.fetch()
if err != nil {
	return err
}
buf = append(buf, b)`)
	if err != nil {
		return err
	}
	nr++
	if l.anyNode != nil {
		if l.anyNode.next == nil {
			return fmt.Errorf("invalid any node: %+v", l.anyNode)
		}
		return writeLayerCode(w, l.anyNode.next, nr)
	}
	_, err = w.WriteString("\nswitch b {")
	if err != nil {
		return err
	}
	for i, n := range l.nodes {
		if n == nil {
			continue
		}
		_, err := fmt.Fprintf(w, "\ncase 0x%02x:", i)
		if err != nil {
			return err
		}
		if n.next != nil {
			err := writeLayerCode(w, n.next, nr)
			if err != nil {
				return err
			}
			continue
		}
		if n.opcode == nil {
			return fmt.Errorf("node without opcode: %+v", n)
		}
		op := n.opcode
		if d := len(op.C) - nr; d > 0 {
			for i := 0; i < d; i++ {
				_, err := w.WriteString(`
b, err = f.fetch()
if err != nil {
	return err
}
buf = append(buf, b)`)
				if err != nil {
					return err
				}
			}
		}
		name := opname.Mangle(op.N)
		_, err = fmt.Fprintf(w, "\n%s(cpu, buf)\nreturn nil", name)
		if err != nil {
			return err
		}
	}
	_, err = w.WriteString(`
default:
	return ErrInvalidCodes
}
return nil`)
	if err != nil {
		return err
	}
	return nil
}
