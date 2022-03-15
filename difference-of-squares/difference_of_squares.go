package diffsquares

import "math"

func SquareOfSum(n int) int {
	v := n * (n + 1) / 2
	return int(math.Pow(float64(v), 2.0))
}

func SumOfSquares(n int) int {

	return n * (n + 1) * (2*n + 1) / 6

}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
