package z80

import (
	"encoding/json"
	"os"
	"testing"
)

func TestDecodeLayer(t *testing.T) {
	l := newDecodeLayer(0, load8)
	m := l.mapTo()

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	err := e.Encode(m)
	if err != nil {
		t.Fatal(err)
	}
}
