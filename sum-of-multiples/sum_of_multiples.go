package summultiples

// SumMultiples calculates the sum of the multiples up to the limit
func SumMultiples(limit int, divisors ...int) int {
	set := make(map[int]bool)

	sum := 0

	for _, d := range divisors {
		if d == 0 {
			continue
		}

		for i := 1; i*d < limit; i++ {
			product := i * d

			if set[product] {
				continue
			}

			set[product] = true
			sum += product
		}
	}

	return sum
}
