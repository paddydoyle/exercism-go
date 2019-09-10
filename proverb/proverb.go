// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb: Given a list of inputs, generate the relevant proverb.
package proverb

import (
	"fmt"
)

const (
	proverbLine = "For want of a %s the %s was lost."
	proverbLast = "And all for the want of a %s."
)

// Given a list of inputs, generate the relevant proverb.
func Proverb(rhyme []string) []string {
	rhyme_len := len(rhyme)
	results := make([]string, rhyme_len)

	if rhyme_len == 0 {
		return []string{}
	}

	for i, _ := range rhyme[:rhyme_len-1] {
		results[i] = fmt.Sprintf(proverbLine, rhyme[i], rhyme[i+1])
	}

	results[rhyme_len-1] = fmt.Sprintf(proverbLast, rhyme[0])

	return results
}
