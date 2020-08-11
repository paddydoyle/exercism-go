package prime

import (
	//"fmt"
	"math"
)

// Nth returns the Nth prime number
func Nth(n int) (int, bool) {
	// Split off trivial edge cases
	if n == 0 {
		return 0, false
	} else if n == 1 {
		// In particular, avoid even numbers in the sieve.
		return 2, true
	}

	// Build the sieve array incrementally. Start with this size as a guess.
	sieveLength := 2 * n

	// Assume that all numbers are primes, and mark
	// the composites via the sieve. Initialised to false.
	composites := make([]bool, sieveLength, sieveLength)
	sliceStart := 0

	foundPrimes := 0

	//fmt.Println("\nNth: START n = ", n, "; foundPrimes = ", foundPrimes, "; cap = ", cap(composites))
	//fmt.Println("Nth: START n = ", n, "; foundPrimes = ", foundPrimes, "; len = ", len(composites))

	for foundPrimes < n {
		//foundPrimes = sieveOdd(composites, n)
		foundPrimes = sieveOddInRange(composites, n, sliceStart)

		// Increase the size of the sieve by n
		sliceStart = len(composites)
		extraComposites := make([]bool, n, n)
		//fmt.Println("Nth: APPEND BEFORE ; composites = ", composites)
		composites = append(composites, extraComposites...)
		//fmt.Println("Nth: APPEND AFTER ; composites = ", composites)

		//fmt.Println("Nth: OUTER n = ", n, "; foundPrimes = ", foundPrimes, "; len = ", len(composites))
	}

	//fmt.Println("Nth: composites = ", composites)

	return findNthPrime(composites, n), true
}

// sieveOdd runs the sieve on odd primes
func sieveOdd(composites []bool, n int) int {
	// Assume that 2 is already covered.
	foundPrimes := 1

	for p := 3; p < len(composites); p += 2 {
		//fmt.Println("sieve: >> p = ", p, "; foundPrimes = ", foundPrimes, "; comp[p] = ", composites[p])

		if composites[p] {
			continue
		}

		foundPrimes++

		// Sieve on multiples of primes, starting from p^2
		for i := p; i*p < len(composites); i++ {
			//fmt.Println("sieveOdd: >>>> INNER SETTING p = ", p, ";i = ", i*p)
			composites[i*p] = true
		}
	}

	//fmt.Println("sieveOdd: returning foundPrimes = ", foundPrimes)
	return foundPrimes
}

// sieveOdd runs the sieve on odd primes, starting from a given position
func sieveOddInRange(composites []bool, n int, sliceStart int) int {
	// Assume that 2 is already covered.
	foundPrimes := 1
	sieveStart := 0

	//fmt.Println("sieveOddInRange: >> n = ", n, "; sliceStart = ", sliceStart)

	for p := 3; p < len(composites); p += 2 {
		//fmt.Println("sieveOddInRange: >> p = ", p, "; foundPrimes = ", foundPrimes, "; comp[p] = ", composites[p])

		if composites[p] {
			continue
		}

		foundPrimes++

		if sliceStart == 0 {
			// Start the sieve from p^2, if this is the first run through
			sieveStart = p
		} else {
			// Else start from the beginning of the newly-appended section
			sieveStart = int(math.Max(float64(p), math.Ceil(float64(sliceStart)/float64(p))))
			//guessMax := input*int(math.Log(float64(input))) + input
		}
		//fmt.Println("sieveOddInRange: >> sieveStart = ", sieveStart, "; mul = ", sieveStart*p)

		// Sieve on multiples of primes
		for i := sieveStart; i*p < len(composites); i++ {
			//fmt.Println("sieveOddInRange: >>>> INNER SETTING p = ", p, ";i*p = ", i*p)
			composites[i*p] = true
		}
	}

	//fmt.Println("sieveOddInRange: END composites = ", composites)
	//fmt.Println("sieveOddInRange: returning foundPrimes = ", foundPrimes)
	return foundPrimes
}

// findNthPrime finds the nth prime in a slice of composites
func findNthPrime(composites []bool, n int) int {
	// Assume that 2 is already covered.
	count := 1

	//fmt.Println("findNthPrime: search for n = ", n, "; len = ", len(composites))

	// Skip indexes 0 and 1, as they are not primes. Odd numbers only.
	for i := 3; count < len(composites); i += 2 {
		//fmt.Println("findNthPrime: loop: for i = ", i, "; val = ", composites[i])
		if !composites[i] {
			count++
		}

		if count == n {
			//fmt.Println("findNthPrime: found for n = ", n, "; i = ", i)
			return i
		}
	}
	return 0
}

// Nth2 returns the Nth prime number
func Nth2(input int) (int, bool) {
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
	//fmt.Println("-----------------------\n-- input = ", input, "; initial guess = ", guessMax)

	// Except for small N, less than 10^6, that's way too small
	// FIXME: magic numbers alert!
	if input < 10e6 {
		guessMax *= 4
	}
	//fmt.Println("new guess = ", guessMax)
	sqrtGuessMax := int(math.Sqrt(float64(guessMax)))
	//fmt.Println("sqrt guess = ", sqrtGuessMax)

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
		//fmt.Println("inner loop j = ", j, "; j*prime = ", j*prime)
	}
	//fmt.Println(">> after even loop: composites = ", composites)

	// Now the odd numbers.
	prime = 3

	// Outer loop: the nth prime
	for n := 2; n < sqrtGuessMax; n++ {
		if composites[n] {
			continue
		}

		//fmt.Println("OUTER loop; n = ", n, "; prime = ", prime)
		// Mark all odd multiples of the current prime
		for j := prime; j*prime < guessMax; j += 2 {
			composites[j*prime] = true
			//fmt.Println("inner loop j = ", j, "; j*prime = ", j*prime)
		}
		//fmt.Println("after inner loop; sqrt(guessMax) = ", math.Sqrt(float64(guessMax)))
		//fmt.Println(">> after inner loop; composites = ", composites)

		//fmt.Println("starting with old prime = ", prime, "; next starting at ", prime+2)

		// Find next prime
		for j := prime + 2; j < guessMax; j += 2 {
			//fmt.Println(">> checking ", j, " which is: ", composites[j], " (false => prime)")
			if !composites[j] {
				prime = j
				//fmt.Println("FOUND ", n, "th prime = ", prime)
				break
			}
		}
	}

	// Now look through the sieve for the nth prime
	primeIndex := 1
	for i := 2; i < guessMax; i++ {
		//fmt.Println(">> checking ", j, " which is: ", composites[j], " (false => prime)")
		if !composites[i] {
			if primeIndex == input {
				prime = i
				break
			}
			primeIndex++
		}
	}

	return prime, true
}
