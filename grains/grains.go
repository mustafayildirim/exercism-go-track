package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {

	if number < 1 || number > 64 {
		return 0, errors.New("square number should be between 1 and 64")
	}

	total := math.Pow(2.0, float64(number-1))

	return uint64(total), nil
}

func Total() uint64 {

	var total uint64 = 0
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		total += square
	}

	return total
}


