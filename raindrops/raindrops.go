package raindrops

import (
	"fmt"
)

// Easier to add new mappings
var divisorMappings = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert a number to a string, the contents of which depend on the
// number's factors.
func Convert(input int) (output string) {
	for div, str := range divisorMappings {
		if input%div == 0 {
			output += str
		}
	}

	if len(output) == 0 {
		return fmt.Sprintf("%d", input)
	}

	return output
}
