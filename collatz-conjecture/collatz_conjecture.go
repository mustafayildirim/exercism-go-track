package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {

	if n < 1 {
		return 0, errors.New("should be positive number")
	} else if n == 1 {
		return 0, nil
	}

	var counter int = 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		counter++
	}

	return counter, nil
}
