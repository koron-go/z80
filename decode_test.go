package z80

import (
	"testing"
)

func TestDecodeLayer_CheckAllOPCodes(t *testing.T) {
	t.Parallel()
	l := defaultDecodeLayer()

	m := map[string]int{}
	for _, cc := range allOPCodes {
		for _, c := range cc {
			k := c.String()
			if _, ok := m[k]; ok {
				t.Fatalf("duplicated: %+v", c)
			}
			m[k] = 0
		}
	}

	l.forNodes(func(n *decodeNode) {
		t.Helper()
		if c := n.opcode; c != nil {
			k := c.String()
			v, ok := m[k]
			if !ok {
				t.Errorf("unknown OPCode: %s", c)
				return
			}
			m[k] = v + 1
		}
	})

	for c, n := range m {
		if n > 0 {
			continue
		}
		t.Errorf("unseen OPCode: %s", c)
	}
}

func TestDecodeLayer_DD7E00(t *testing.T) {
	f := memSrc{0xdd, 0x7e, 0x00}
	l := defaultDecodeLayer()
	c, _, err := decode(l, nil, &f)
	if err != nil {
		t.Fatal(err)
	}
	if c.N != "LD r, (IX+d)" {
		t.Fatalf("unexpected opcode: %s", c.N)
	}
}
