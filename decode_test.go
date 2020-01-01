package z80

import (
	"testing"
)

func TestDecodeLayer_CheckAllOPCodes(t *testing.T) {
	l := defaultDecodeLayer()

	m := map[string]int{}
	for _, cc := range allOPCodes {
		for _, c := range cc {
			k := c.String()
			m[k] = 0
		}
	}

	mark := func(c *OPCode) {
		t.Helper()
		k := c.String()
		v, ok := m[k]
		if !ok {
			t.Errorf("unknown OPCode: %s", c)
			return
		}
		m[k] = v + 1
	}

	var procNode func(*decodeNode)
	procLayer := func(l *decodeLayer) {
		t.Helper()
		for _, n := range l.nodes {
			if n == nil {
				continue
			}
			procNode(n)
		}
	}
	procNode = func(n *decodeNode) {
		t.Helper()
		if n.opcode != nil {
			mark(n.opcode)
		}
		for _, c := range n.codes {
			mark(c)
		}
		if n.next != nil {
			procLayer(n.next)
		}
	}
	procLayer(l)

	for c, n := range m {
		if n > 0 {
			continue
		}
		t.Errorf("unseen OPCode: %s %d", c, n)
	}
}
