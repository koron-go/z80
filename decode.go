package z80

import (
	"fmt"
	"math/bits"
)

type decodeLayer struct {
	nodes [256]*decodeNode
}

func newDecodeLayer(level int, opcodes ...[]*OPCode) *decodeLayer {
	l := &decodeLayer{}
	for _, cc := range opcodes {
		for _, opcode := range cc {
			if level < len(opcode.C) {
				l.put(level, opcode)
			}
		}
	}
	for i, n := range l.nodes {
		if n == nil {
			continue
		}
		switch len(n.codes) {
		case 0:
		case 1:
			n.opcode = n.codes[0]
			n.codes = nil
			fmt.Printf("HERE_A0: %02X %+v\n", i, n.opcode.C)
		default:
			//n.next = newDecodeLayer(level+1, n.codes)
			//n.codes = nil
			fmt.Printf("HERE_A1: %02X %+v\n", i, n.codes)
		}
	}
	return l
}

func (l *decodeLayer) put(level int, opcode *OPCode) {
	c := opcode.C[level]
	s := bits.OnesCount8(^c.M)
	for i, e := int(c.V), int(c.V|c.M); i <= e; i++ {
		n := l.nodes[i]
		if n == nil || s > n.score {
			n = &decodeNode{score: s}
			l.nodes[i] = n
		}
		n.codes = append(n.codes, opcode)
	}
}

func (l *decodeLayer) mapTo() map[string]interface{} {
	m := map[string]interface{}{}
	for i, n := range l.nodes {
		if n == nil {
			continue
		}
		m[fmt.Sprintf("%02X", i)] = n.mapTo()
	}
	return m
}

type decodeNode struct {
	score  int
	opcode *OPCode
	next   *decodeLayer
	codes  []*OPCode
}

func (n *decodeNode) mapTo() map[string]interface{} {
	m := map[string]interface{}{}
	m["score"] = n.score
	if n.opcode != nil {
		m["opcode"] = n.opcode.mapTo()
	}
	if n.next != nil {
		m["next"] = n.next.mapTo()
	}
	if len(n.codes) > 0 {
		m["codes"] = len(n.codes)
	}
	return m
}
