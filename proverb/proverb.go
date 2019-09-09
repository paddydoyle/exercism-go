// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import (
	"fmt"
)

var proverbLine = "For want of a %s the %s was lost."
var proverbLast = "And all for the want of a %s."

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	// Take last element, if there is one. For the final line.

	var results []string

	if len(rhyme) == 0 {
		return []string{}
	}

	for i, _ := range rhyme[0 : len(rhyme)-1] {
		results = append(results, fmt.Sprintf(proverbLine, rhyme[i], rhyme[i+1]))
	}

	results = append(results, fmt.Sprintf(proverbLast, rhyme[0]))

	return results
}
