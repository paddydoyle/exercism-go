package prime

import (
	"math"
)

// Nth returns the Nth prime number
func Nth(input int) (int, bool) {
	// Split off edge cases
	if input == 0 {
		return 0, false
	} else if input == 1 {
		return 2, true
	} else if input == 2 {
		return 3, true
	}

	// Nth prime number roughly satisfies "n * log(n)"
	// So for the Sieve of Eratosthenes we'll allocate a bool
	// array up to that size.
	guessMax := input * int(math.Log(float64(input)))
	// Except for small N, that's probably way too small
	guessMax = int(math.Max(float64(guessMax), float64(input*input)))

	// Assume that all numbers are primes, and mark
	// the composites
	composites := make([]bool, guessMax, guessMax)

	// Mark trivial cases
	composites[0] = true
	composites[1] = true

	prime := 2

	for i := 1; i < input; i++ {
		for j := 1; j*prime < guessMax; j++ {
			// Mark all multiples of the current prime
			composites[j*prime] = true
		}

		// Find next prime
		for j := prime + 1; j < guessMax; j++ {
			if !composites[j] {
				prime = j
				break
			}
		}
	}

	return prime, true
}
