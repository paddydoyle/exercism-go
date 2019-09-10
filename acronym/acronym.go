// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

var stopChars = []string{
	"-",
	"_",
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) (output string) {
	for _, stopChar := range stopChars {
		s = strings.Replace(s, stopChar, " ", -1)
	}

	words := strings.Fields(s)

	buf := new(strings.Builder)

	for _, word := range words {
		buf.WriteByte(word[0])
	}

	return strings.ToUpper(buf.String())
}
