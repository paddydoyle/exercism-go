package hamming

import "errors"

func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return distance, errors.New("inputs of unequal lengths")
	}

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
