package luhn

import (
	"strconv"
	"strings"
)

// Valid validates input according to the Luhn algorithm
func Valid(input string) bool {
	// Remove whitespace
	input = strings.ReplaceAll(input, " ", "")

	// Any non-digit chars?
	if _, err := strconv.Atoi(input); err != nil {
		return false
	}

	// Must be at least 2 chars in length
	if len(input) <= 1 {
		return false
	}

	// Simpler to loop from left to right, so prepend a zero if length
	// is even. This doesn't affect the result, as anything times 0 is 0.
	if len(input)%2 != 0 {
		input = "0" + input
	}

	// Last char is the checksum
	sum := int(input[len(input)-1] - '0')

	for i, c := range input[0 : len(input)-1] {
		digit := int(c) - '0'

		if i%2 == 0 {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
	}

	return (sum % 10) == 0
}
