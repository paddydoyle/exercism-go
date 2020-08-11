package summultiples

// SumMultiples calculates the sum of the multiples up to the limit
func SumMultiples(limit int, divisors ...int) int {
	set := make(map[int]bool)

	for _, d := range divisors {
		if d == 0 {
			continue
		}

		for i := 1; i*d < limit; i++ {
			set[i*d] = true
		}
	}

	sum := 0

	for key, _ := range set {
		sum += key
	}

	return sum
}
