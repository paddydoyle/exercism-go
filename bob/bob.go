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
// 280000ns / op
//=== RUN   TestHey
//--- PASS: TestHey (0.00s)
//goos: linux
//goarch: amd64
//BenchmarkHey-8              4274            280134 ns/op          247985 B/op       3475 allocs/op
//PASS
//ok      _/home/paddy/exercism/go/bob    2.148s
func Hey(remark string) string {
	//remark = strings.TrimSpace(remark)
	isEmpty := regexp.MustCompile(`^\s*$`).MatchString
	hasChars := regexp.MustCompile(`[[:alpha:]]`).MatchString
	//hasUpperChars := regexp.MustCompile(`[A-Z]`).MatchString
	hasLowerChars := regexp.MustCompile(`[a-z]`).MatchString
	isUpperQuestion := regexp.MustCompile(`^[A-Z ]+\?\s*$`).MatchString
	//isUpperString := regexp.MustCompile(`^[A-Z0-9 !,]+\s*$`).MatchString
	isQuestion := regexp.MustCompile(`\?\s*$`).MatchString

	switch {
	case isEmpty(remark):
		return "Fine. Be that way!"
	case hasChars(remark) && isUpperQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case hasChars(remark) && !hasLowerChars(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

// Hey returns a string according to the rules of Bob.
// 900-1000ns / op
//=== RUN   TestHey
//--- PASS: TestHey (0.00s)
//goos: linux
//goarch: amd64
//BenchmarkHey-8           1306461               912 ns/op               0 B/op          0 allocs/op
//PASS
//ok      _/home/paddy/exercism/go/bob    1.976s
func Hey2(remark string) string {
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
