// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb: Given a list of inputs, generate the relevant proverb.
package proverb

import (
	"fmt"
)

var proverbLine = "For want of a %s the %s was lost."
var proverbLast = "And all for the want of a %s."

// Given a list of inputs, generate the relevant proverb.
func Proverb(rhyme []string) []string {
	results := make([]string, 0, len(rhyme))

	if len(rhyme) == 0 {
		return []string{}
	}

	for i, _ := range rhyme[0 : len(rhyme)-1] {
		results = append(results, fmt.Sprintf(proverbLine, rhyme[i], rhyme[i+1]))
	}

	results = append(results, fmt.Sprintf(proverbLast, rhyme[0]))

	return results
}
