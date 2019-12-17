package diffsquares

// SumOfSquares returns the sum of the first n natural number squares
func SumOfSquares(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

// SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// Difference returns the difference between sum of squares, and square of sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
