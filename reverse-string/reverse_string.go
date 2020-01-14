package reverse

// Reverse returns a reversal of the input string
func Reverse(input string) string {
	// solution copied from: https://yourbasic.org/golang/reverse-utf8-string/

	// Same length
	output := make([]byte, len(input))

	// Keep track of previous input position (increasing),
	// and output position (decreasing)
	prevPos, outputPos := 0, len(input)

	for pos := range input {
		// For non-UTF8 strings, this will subtract 1
		outputPos -= pos - prevPos

		// Copy the rune/byte from the start of input to
		// end of output
		copy(output[outputPos:], input[prevPos:pos])

		prevPos = pos
	}

	// Copy final rune to the start of the output
	copy(output[0:], input[prevPos:])

	return string(output)
}
