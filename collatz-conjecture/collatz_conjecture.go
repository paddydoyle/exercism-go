package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps to reach 1
func CollatzConjecture(number int) (steps int, err error) {
	if number <= 0 {
		return 0, errors.New("invalid input: must be strictly positive")
	}

	for number != 1 {
		if number%2 == 0 {
			number = number / 2
		} else {
			number = 3*number + 1
		}

		steps++
	}

	return steps, err
}
