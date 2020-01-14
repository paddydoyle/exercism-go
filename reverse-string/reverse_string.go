package reverse

// Reverse returns a reversal of the input string
func Reverse(input string) string {
	// solution copied from user kxnes

	reversed := make([]rune, 0)

	for _, r := range input {
		reversed = append([]rune{r}, reversed...)
	}

	return string(reversed)
}
