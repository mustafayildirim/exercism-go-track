package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type error interface {
	Error() string
}

type SillyNephewError struct {
	cows int
}

func NewSillyNephewError(cows int) *SillyNephewError {
	return &SillyNephewError{cows: cows}
}

func (m *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", m.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	if cows < 0 {
		return 0, NewSillyNephewError(cows)
	} else if cows == 0 {
		return 0, errors.New("division by zero")
	}
	fodder, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction {
		if fodder > 0 {
			return 2 * fodder / float64(cows), nil
		}
	} else if err != nil {
		return 0, err
	}
	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}

	return fodder / float64(cows), nil
}
