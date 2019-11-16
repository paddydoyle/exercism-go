package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if the input word is an Isogram
func IsIsogram(word string) bool {
	charMap := make(map[rune]bool, len(word))

	for _, ch := range strings.ToLower(word) {
		if charMap[ch] {
			return false
		}

		if unicode.IsLetter(ch) {
			charMap[ch] = true
		}
	}

	return true
}
