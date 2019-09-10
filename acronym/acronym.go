// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	words := regexp.MustCompile(`[ \-_]`).Split(s, -1)

	var buf = new(strings.Builder)

	for _, word := range words {
		// Split includes empty matches
		if len(word) > 0 {
			buf.WriteByte(word[0])
		}
	}

	return strings.ToUpper(buf.String())
}
