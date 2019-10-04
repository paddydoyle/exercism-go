// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
	"unicode"
)

// Hey returns a string according to the rules of Bob.
// 900-1000ns / op
func Hey3(remark string) string {
	//remark = strings.TrimSpace(remark)
	isEmpty := regexp.MustCompile(`^\s*$`).MatchString
	hasChars := regexp.MustCompile(`[[:alpha:]]`).MatchString
	//hasUpperChars := regexp.MustCompile(`[A-Z]`).MatchString
	isUpperQuestion := regexp.MustCompile(`^[A-Z ]+\?\s*$`).MatchString
	isUpperString := regexp.MustCompile(`^[A-Z0-9 !,]+\s*$`).MatchString
	isQuestion := regexp.MustCompile(`\?\s*$`).MatchString

	switch {
	case isEmpty(remark):
		return "Fine. Be that way!"
	case hasChars(remark) && isUpperQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case hasChars(remark) && isUpperString(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

// Hey returns a string according to the rules of Bob.
// 900-1000ns / op
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case IsUpperString(remark) && strings.HasSuffix(remark, "?"):
		return "Calm down, I know what I'm doing!"
	case IsUpperString(remark):
		return "Whoa, chill out!"
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}

// IsUpperString tests if the characters in the string are only uppercase
func IsUpperString(remark string) bool {
	containsAnyUppercase := false

	for _, ch := range remark {
		if unicode.IsLetter(ch) && unicode.IsLower(ch) {
			return false
		}
		if !containsAnyUppercase && unicode.IsLetter(ch) {
			containsAnyUppercase = true
		}
	}

	return containsAnyUppercase
}
