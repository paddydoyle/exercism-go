// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"bytes"
	//"fmt"
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	buf := bytes.Buffer{}

	strs := regexp.MustCompile(`[ \-_]`).Split(s, -1)

	//fmt.Printf("strs = %s\n", strs)

	//strs := re.Split(s, -1).FindAllString(`\w`)


        for _, word := range strs {
                if len(word) > 0 {
			buf.WriteByte(word[0])
		}
        }

        return strings.ToUpper(buf.String())
}
