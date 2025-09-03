package utils

import (
	"regexp"
)

// EscapeCharsForPerlSQL escapes the known characters in an SQL that needs to be escaped in Perl.
//
// Current known characters: '(', ')', and '?'.
//
// An example usage is with DATA-DOG go-sqlmock expected SQL parameter.
func EscapeCharsForPerlSQL(s string) string {
	return regexp.MustCompile(`([()?])`).ReplaceAllString(s, `\$1`)
}
