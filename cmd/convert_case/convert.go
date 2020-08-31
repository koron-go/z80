package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	err := convert(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

func convert(w io.Writer, r io.Reader) error {
	fmt.Fprintf(w, "package zex\n")
	br := bufio.NewReader(r)
	var all []*testCase
	for {
		tc, err := readTestCase(br)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		err = writeTestCase(w, tc)
		if err != nil {
			return err
		}
		all = append(all, tc)
	}
	if len(all) == 0{
		return nil
	}
	fmt.Fprintf(w, "\n// DocCases has all test cases in zexdoc.asm\n")
	fmt.Fprintf(w, "var DocCases = []Case{\n")
	for _, tc := range all {
		fmt.Fprintf(w, "\tDoc%s,\n", strings.ToUpper(tc.label))
	}
	fmt.Fprintf(w, "}\n")
	// TODO:
	return nil
}

type testCase struct {
	comment string
	label   string
	mask    uint8
	base    *state
	inc     *state
	shift   *state
	crc     uint32
	desc    string
}

var rxComment = regexp.MustCompile(`^; (.* \([0-9,]+ cycles\))`)
var rxLabelMask = regexp.MustCompile(`^([^:]+):\s*db\s+(\S+)\s*; flag mask`)

func readTestCase(r *bufio.Reader) (*testCase, error) {
	for {

		s1, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		m1 := rxComment.FindStringSubmatch(s1)
		if len(m1) == 0 {
			continue
		}
		tc := &testCase{comment: m1[1]}

		s2, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		m2 := rxLabelMask.FindStringSubmatch(s2)
		if len(m2) == 0 {
			return nil, fmt.Errorf("can't find label and mask for: %s", tc.comment)
		}
		tc.label = m2[1]
		tc.mask, err = parseU8(m2[2])
		if err != nil {
			return nil, fmt.Errorf("failed to parse mask for %s: %s", tc.label, err)
		}

		base, err := readState(r)
		if err != nil {
			return nil, fmt.Errorf("failed to parse basecase for %s: %s", tc.label, err)
		}
		tc.base = base

		inc, err := readState(r)
		if err != nil {
			return nil, fmt.Errorf("failed to parse incvec for %s: %s", tc.label, err)
		}
		tc.inc = inc

		shift, err := readState(r)
		if err != nil {
			return nil, fmt.Errorf("failed to parse shiftvec for %s: %s", tc.label, err)
		}
		tc.shift = shift

		crc, err := readCrc(r)
		if err != nil {
			return nil, fmt.Errorf("failed to parse CRC for %s: %s", tc.label, err)
		}
		tc.crc = crc

		desc, err := readDesc(r)
		if err != nil {
			return nil, fmt.Errorf("failed to parse desc for %s: %s", tc.label, err)
		}
		tc.desc = desc
		// TODO:

		return tc, nil
	}
}

type state struct {
	values []string
}

var rxTstr = regexp.MustCompile(`^\s*tstr\s+((?:[^,]+,){12}[^,;]+)(?:\s*;.*)?\n$`)

func readState(r *bufio.Reader) (*state, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	m := rxTstr.FindStringSubmatch(s)
	if len(m) == 0 {
		return nil, fmt.Errorf("not found \"tstr\": %q", s)
	}
	v, err := commaSplit(strings.TrimSpace(m[1]), 13)
	if err != nil {
		return nil, fmt.Errorf("unexpected \"tstr\" in %q: %s", s, err)
	}
	return &state{values: v}, nil
}

var rxCommaSplit = regexp.MustCompile(`\s*,\s*`)

func commaSplit(s string, n int) ([]string, error) {
	v := rxCommaSplit.Split(s, n+1)
	if len(v) != n {
		return nil, fmt.Errorf("wrong length: want=%d got=%d", n, len(v))
	}
	return v, nil
}

var rxCrc = regexp.MustCompile(`^\s*db\s+((?:[^,]+,){3}[^,;]+)(?:\s*;.*)?\n$`)

func readCrc(r *bufio.Reader) (uint32, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	m := rxCrc.FindStringSubmatch(s)
	if m == nil {
		return 0, fmt.Errorf("not found CRC: %q", s)
	}
	v, err := commaSplit(strings.TrimSpace(m[1]), 4)
	if err != nil {
		return 0, fmt.Errorf("unexpected CRC in %q: %s", s, err)
	}
	var crc uint32
	for _, s := range v {
		d, err := parseU8(s)
		if err != nil {
			return 0, fmt.Errorf("invalid CRC in %q: %s", s, err)
		}
		crc <<= 8
		crc |= uint32(d)
	}
	return crc, nil
}

var rxDesc = regexp.MustCompile(`^\s*tmsg\s+'([^']*)'`)

func readDesc(r *bufio.Reader) (string, error) {
	s, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	m := rxDesc.FindStringSubmatch(s)
	if m == nil {
		return "", fmt.Errorf("not found tmsg(desc): %q", s)
	}
	return strings.Trim(m[1], "."), nil
}

func writeTestCase(w io.Writer, tc *testCase) error {
	uplabel := strings.ToUpper(tc.label)
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "// Doc%s: %s\n", uplabel, tc.comment)
	fmt.Fprintf(w, "var Doc%s = Case{\n", uplabel)
	fmt.Fprintf(w, "\t0x%02x,\n", tc.mask)
	if err := writeState(w, tc.base); err != nil {
		return err
	}
	if err := writeState(w, tc.inc); err != nil {
		return err
	}
	if err := writeState(w, tc.shift); err != nil {
		return err
	}
	fmt.Fprintf(w, "\tCRC32(0x%08x),\n", tc.crc)
	fmt.Fprintf(w, "\t%q,\n", tc.desc)
	fmt.Fprintf(w, "}\n")
	return nil
}

func writeState(w io.Writer, s *state) error {
	fmt.Fprintf(w, "\tStatus{\n\t\t")

	if err := writeAsU8(w, s.values[0], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 0, err)
	}
	if err := writeAsU8(w, s.values[1], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 1, err)
	}
	if err := writeAsU8(w, s.values[2], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 2, err)
	}
	if err := writeAsU8(w, s.values[3], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 3, err)
	}

	if err := writeAsU16(w, s.values[4], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 4, err)
	}
	if err := writeAsU16(w, s.values[5], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 5, err)
	}
	if err := writeAsU16(w, s.values[6], ",\n\t\t"); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 6, err)
	}
	if err := writeAsU16(w, s.values[7], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 7, err)
	}
	if err := writeAsU16(w, s.values[8], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 8, err)
	}
	if err := writeAsU16(w, s.values[9], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 9, err)
	}

	if err := writeAsU8(w, s.values[10], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 10, err)
	}
	if err := writeAsU8(w, s.values[11], ", "); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 11, err)
	}

	if err := writeAsU16(w, s.values[12], ",\n"); err != nil {
		return fmt.Errorf("failed to write state#%d: %w", 12, err)
	}

	fmt.Fprintf(w, "\t},\n")
	return nil
}

var exceptions = map[string]string{
	"msbt":    "Msbt",
	"msbt-1":  "Msbt - 1",
	"msbt+17": "Msbt + 17",
	"msbtlo":  "MsbtLo",
	"msbthi":  "MsbtHi",
	"msbt+3":  "Msbt + 3",
	"msbt+1":  "Msbt + 1",
	"msbt+2":  "Msbt + 2",
}

func writeAsU8(w io.Writer, s, suffix string) error {
	if rxInt.MatchString(s) {
		n, err := strconv.ParseInt(s, 10, 9)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "0x%02x%s", uint8(n), suffix)
		return nil
	}
	if m := rxHex.FindStringSubmatch(s); len(m) != 0 {
		if m[1] == "" {
			m[1] = "0"
		}
		n, err := strconv.ParseInt(m[1], 16, 9)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "0x%02x%s", uint8(n), suffix)
		return nil
	}
	if r, ok := exceptions[s]; ok {
		fmt.Fprintf(w, "%s%s", r, suffix)
		return nil
	}
	return fmt.Errorf("unknown U8 format: %q", s)
}

func writeAsU16(w io.Writer, s, suffix string) error {
	if rxInt.MatchString(s) {
		n, err := strconv.ParseInt(s, 10, 17)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "0x%04x%s", uint16(n), suffix)
		return nil
	}
	if m := rxHex.FindStringSubmatch(s); len(m) != 0 {
		if m[1] == "" {
			m[1] = "0"
		}
		n, err := strconv.ParseInt(m[1], 16, 17)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "0x%04x%s", uint16(n), suffix)
		return nil
	}
	if r, ok := exceptions[s]; ok {
		fmt.Fprintf(w, "%s%s", r, suffix)
		return nil
	}
	return fmt.Errorf("unknown U8 format: %q", s)
}

var rxInt = regexp.MustCompile(`^-?(?:0|[1-9][0-9]*)$`)
var rxHex = regexp.MustCompile(`^0([0-9a-fA-F]*)h$`)

func parseU8(s string) (uint8, error) {
	if rxInt.MatchString(s) {
		n, err := strconv.ParseInt(s, 10, 9)
		if err != nil {
			return 0, err
		}
		return uint8(n), nil
	}
	if m := rxHex.FindStringSubmatch(s); len(m) != 0 {
		if m[1] == "" {
			m[1] = "0"
		}
		n, err := strconv.ParseInt(m[1], 16, 9)
		if err != nil {
			return 0, err
		}
		return uint8(n), nil
	}
	return 0, fmt.Errorf("unknown U8 format: %q", s)
}
