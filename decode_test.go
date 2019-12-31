package z80

import (
	"encoding/json"
	"os"
	"testing"
)

func TestDecodeLayer(t *testing.T) {
	l := defaultDecodeLayer()

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	err := e.Encode(l.mapTo())
	if err != nil {
		t.Fatal(err)
	}
}
