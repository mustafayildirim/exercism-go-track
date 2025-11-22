package darts

import "math"

func Score(x, y float64) int {

	result := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if result <= 1 {
		return 10
	} else if result <= 5 {
		return 5
	} else if result <= 10 {
		return 1
	}

	return 0
}
