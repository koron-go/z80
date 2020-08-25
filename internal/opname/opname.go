/*
Package opname provides argorithms for name of Z80 operations.
*/
package opname

import "strings"

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

// Mangle mangles OPCode name
func Mangle(s string) string {
	return "op" + strings.Map(mapping, s)
}
