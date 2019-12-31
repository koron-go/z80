package z80

import (
	"fmt"
	"math/bits"
)

type decodeLayer struct {
	nodes [256]*decodeNode
}

func defaultDecodeLayer() *decodeLayer {
	return newDecodeLayer(0, load8, load16, exbtsg, arith8)
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
	for _, n := range l.nodes {
		if n == nil {
			continue
		}
		switch len(n.codes) {
		case 0:
			continue
		case 1:
			n.opcode = n.codes[0]
			continue
		}
		n.next = newDecodeLayer(level+1, n.codes)
		n.codes = nil
	}
	return l
}

func (l *decodeLayer) put(level int, opcode *OPCode) {
	c := opcode.C[level]
	s := bits.OnesCount8(^c.M)
	for i, e := c.beginEnd(); i <= e; i++ {
		if !c.match(uint8(i)) {
			continue
		}
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
	if len(n.codes) >= 2 {
		m["codes"] = len(n.codes)
	}
	return m
}
