package etl

import (
	"strings"
)

// Transform converts the int-to-chars map into a char-to-int map
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)

	for i, chars := range input {
		for _, c := range chars {
			output[strings.ToLower(c)] = i
		}
	}

	return output
}
