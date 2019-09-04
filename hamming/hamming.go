package hamming

import "errors"

func Distance(a, b string) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return distance, errors.New("Inputs of unequal lengths.")
	}

	for pos, charA := range(a) {
		// Must convert to avoid rune/byte mismatch
		if byte(charA) != b[pos] {
			distance++
		}
	}

	return distance, nil
}
