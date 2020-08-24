package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func invalidRune(r rune) bool {
	// valid: 0-9 A-Z a-z
	// invalid: others
	return r < '0' ||
		(r > '9' && r < 'A') ||
		(r > 'Z' && r < 'a') ||
		r > 'z'
}

func mapping(r rune) rune {
	if r == ')' {
		return 'P'
	}
	if invalidRune(r) {
		return -1
	}
	return r
}

func mangleN(s string) string {
	return "op" + strings.Map(mapping, s)
}

type item struct {
	n string
	f []string
}

func extractF(w *bufio.Writer, r *bufio.Reader) error {
	var items []*item

L:
	for {
		var (
			n      string
			f      []string
			indent string
			fEnd   string
		)
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break L
				}
				return err
			}
			_, err = w.WriteString(s)
			if err != nil {
				return err
			}
			m := rxN.FindStringSubmatch(s)
			if len(m) > 0 {
				n = mangleN(m[1])
				break
			}
		}
		for {
			s, err := r.ReadString('\n')
			if err != nil {
				return err
			}
			m := rxF.FindStringSubmatch(s)
			if len(m) > 0 {
				indent = m[1]
				fEnd = indent + "},\n"
				break
			}
			_, err = w.WriteString(s)
			if err != nil {
				return err
			}
		}
		_, err := fmt.Fprintf(w, "%sF: %s,\n", indent, n)
		if err != nil {
			return err
		}

		for {
			s, err := r.ReadString('\n')
			if err != nil {
				return err
			}
			if s == fEnd {
				break
			}
			f = append(f, s[len(indent):])
		}
		items = append(items, &item{n: n, f: f})
	}

	for _, i := range items {
		_, err := fmt.Fprintf(w, "\nfunc %s(cpu *CPU, codes []uint8) {\n", i.n)
		if err != nil {
			return err
		}
		for _, f := range i.f {
			_, err := fmt.Fprint(w, f)
			if err != nil {
				return err
			}
		}
		_, err = w.WriteString("}\n")
		if err != nil {
			return err
		}
	}

	return w.Flush()
}

func (i item) name() string {
	return strings.Map(mapping, i.n)
}

var rxN = regexp.MustCompile(`^\s*N: "([^"]*)",`)

var rxF = regexp.MustCompile(`^(\t+)F: func\(.*\) \{\n$`)

func rewrite(name string) error {
	in, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	r := bufio.NewReader(bytes.NewBuffer(in))

	out := &bytes.Buffer{}
	w := bufio.NewWriter(out)
	err = extractF(w, r)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(name, out.Bytes(), 0666)
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		r := bufio.NewReader(os.Stdin)
		w := bufio.NewWriter(os.Stdout)
		err := extractF(w, r)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, name := range flag.Args() {
		err := rewrite(name)
		if err != nil {
			log.Fatal(err)
		}
	}
}
