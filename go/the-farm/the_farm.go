package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}
	fAmt, fErr := weightFodder.FodderAmount()
	if fErr != nil {
		if fErr == ErrScaleMalfunction {
			if fAmt < 0 {
				return 0, errors.New("negative fodder")
			} else if cows < 0 {
				return 0, fmt.Errorf("silly nephew, there cannot be %d cows", cows)
			}
			return fAmt * 2 / float64(cows), nil
		} else {
			if fAmt < 0 {
				return 0.0, errors.New("non-scale error")
			}
			return 0, errors.New("non-scale error")
		}
	}
	if fAmt < 0 {
		return 0, errors.New("negative fodder")
	} else if cows < 0 {
		return 0, fmt.Errorf("silly nephew, there cannot be %d cows", cows)
	}
	return fAmt / float64(cows), nil
}
