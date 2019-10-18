package scrabble

import (
	"strings"
)

func scoreLookup(ch rune) (score int) {
	switch ch {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

// Score returns the scrabble score for input string.
func Score(input string) (score int) {
	input = strings.ToUpper(input)

	for _, ch := range input {
		score += scoreLookup(ch)
	}

	return score
}
