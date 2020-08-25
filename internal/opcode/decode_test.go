package opcode

import (
	"testing"
)

func TestDecodeLayer_CheckAllOPCodes(t *testing.T) {
	t.Parallel()
	l := defaultDecodeLayer()

	m := map[string]int{}
	for _, cc := range AllOPCodes {
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
