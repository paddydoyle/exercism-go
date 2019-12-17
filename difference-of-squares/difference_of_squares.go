package diffsquares

// SumOfSquares returns the sum of the first n natural number squares
func SumOfSquares(n int) int {
	sum := n * (n + 1) * (2*n + 1) / 6

	return sum
}

// SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2

	return sum * sum
}

// Difference returns the difference between sum of squares, and square of sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
