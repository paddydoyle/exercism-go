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

	// Only perform the transformation on every second char.
	parity := len(input)%2 == 0

	// Last char is the checksum
	sum := int(input[len(input)-1] - '0')

	for _, c := range input[0 : len(input)-1] {
		digit := int(c) - '0'

		if parity {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}

		parity = !parity

		sum += digit
	}

	return (sum % 10) == 0
}
