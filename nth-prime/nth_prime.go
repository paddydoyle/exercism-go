package prime

import (
	//"fmt"
	"math"
)

// Nth returns the Nth prime number
func Nth(input int) (int, bool) {
	// Split off trivial edge cases
	if input == 0 {
		return 0, false
	} else if input == 1 {
		return 2, true
	} else if input == 2 {
		return 3, true
	}

	// According to the Prime Number Theorem, the
	// Nth prime number roughly satisfies "n * log(n)"
	// So for the Sieve of Eratosthenes we'll allocate a bool
	// array up to that size.
	//guessMax := input*int(math.Log(float64(input))) + input
	guessMax := input * int(math.Log(float64(input)))
	//fmt.Println("input = ", input, "; initial guess = ", guessMax)

	// Except for small N, less than 10^6, that's way too small
	// FIXME: magic numbers alert!
	if input < 10e6 {
		guessMax *= 4
	}
	//fmt.Println("new guess = ", guessMax)

	// Assume that all numbers are primes, and mark
	// the composites. Initialised to false.
	composites := make([]bool, guessMax, guessMax)

	// Mark trivial cases. FIXME: remove
	//composites[0] = true
	//composites[1] = true
	//fmt.Println("composites = ", composites)

	prime := 2

	// Treat even numbers separately, as it allows small optimisations
	// in the odd number loop.
	// Mark all multiples of the current prime
	for j := prime; j*prime < guessMax; j++ {
		composites[j*prime] = true
	}
	//fmt.Println("after even loop: composites = ", composites)

	// Now the odd numbers.
	prime = 3

	// Outer loop: the nth prime
	for n := 2; n < input; n++ {
		// Mark all odd multiples of the current prime
		for j := prime; j*prime < guessMax; j += 2 {
			composites[j*prime] = true
			//fmt.Println("inner loop j = ", j, "; j*prime = ", j*prime)
		}
		//fmt.Println("after inner loop; composites = ", composites)

		//fmt.Println("starting with old prime = ", prime, "; next starting at ", prime+2)

		// Find next prime
		for j := prime + 2; j < guessMax; j += 2 {
			//fmt.Println(">> checking ", j, " which is: ", composites[j], " (false => prime)")
			if !composites[j] {
				prime = j
				break
			}
		}
	}

	return prime, true
}
