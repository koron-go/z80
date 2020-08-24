package z80

type fetcher interface {
	fetch() uint8
	fetchLabel() string
}

type memSrc []uint8

func (m *memSrc) fetch() uint8 {
	if len(*m) == 0 {
		return 0
	}
	var b uint8
	b, *m = (*m)[0], (*m)[1:]
	return b
}

func (m *memSrc) fetchLabel() string {
	return "IM0"
}
