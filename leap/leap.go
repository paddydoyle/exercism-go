// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// Exclude most years: non-divisible by 4
	if year % 4 != 0 {
		return false
	}

	// It is divisible by 4.
	// Next test: is it divisible by 100, but not by 400?
	if year % 100 == 0 && year % 400 != 0 {
		return false
	}

	// Otherwise, it is indeed a leap year.
	return true
}
