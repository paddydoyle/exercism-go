package raindrops

import (
	"strconv"
)

var divisorMappings = []struct {
	divisor int
	output  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert a number to a string, the contents of which depend on the
// number's factors.
func Convert(input int) (output string) {
	for _, mapping := range divisorMappings {
		if input%mapping.divisor == 0 {
			output += mapping.output
		}
	}

	if len(output) == 0 {
		return strconv.Itoa(input)
	}

	return output
}
