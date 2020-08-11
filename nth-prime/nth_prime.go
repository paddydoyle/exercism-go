package prime

import (
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
	sieveLength := estimateInitialSize(n)

	// Assume that all numbers are primes, and mark the composites via
	// the sieve. Initialised to false.
	composites := make([]bool, sieveLength, sieveLength)

	foundPrimes := 0

	for foundPrimes < n {
		foundPrimes = sieveOdd(composites, n)

		// Increase the size of the sieve
		sieveLength = len(composites)
		extraComposites := make([]bool, sieveLength, sieveLength)
		composites = append(composites, extraComposites...)
	}

	return findNthPrime(composites, n), true
}

func estimateInitialSize(n int) int {
	// According to the Prime Number Theorem, the
	// Nth prime number roughly satisfies "n * log(n)"
	// So for the Sieve of Eratosthenes we'll allocate a bool
	// array up to that size.
	// Except that doesn't hold for small "n", so in that
	// case conservatively use "n * 2"

	return n * int(math.Max(2.0, math.Log(float64(n))))
}

// sieveOdd runs the sieve on odd primes
func sieveOdd(composites []bool, n int) int {
	// Assume that 2 is already covered.
	foundPrimes := 1

	for p := 3; p < len(composites); p += 2 {
		if composites[p] {
			continue
		}

		foundPrimes++

		// Sieve on multiples of primes, starting from p^2
		for i := p; i*p < len(composites); i++ {
			composites[i*p] = true
		}
	}

	return foundPrimes
}

// findNthPrime finds the nth prime in a slice of composites
func findNthPrime(composites []bool, n int) int {
	// Assume that 2 is already covered.
	count := 1

	// Skip indexes 0 and 1, as they are not primes. Odd numbers only.
	for i := 3; i < len(composites); i += 2 {
		if !composites[i] {
			count++
		}

		if count == n {
			return i
		}
	}
	return 0
}
