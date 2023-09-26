package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	message string
	//details string
}

func (e *InvalidCowsError) Error() string {
	return e.message
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(fCalc FodderCalculator, cows int) (float64, error) {
	if err := ValidateNumberOfCows(cows); err != nil {
		return 0, err
	}
	fAmt, fErr := fCalc.FodderAmount(cows)
	if fErr != nil {
		return 0, fErr
	}

	fFactor, faErr := fCalc.FatteningFactor()
	if faErr != nil {
		return 0, faErr
	}

	if fAmt < 0 {
		return 0, errors.New("negative fodder")
	} else if fAmt == 0 {
		return 0, errors.New("amount could not be determined")
	}

	return fAmt * fFactor / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if err := ValidateNumberOfCows(cows); err != nil {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

func ValidateNumberOfCows(cows int) error {
	if cows > 0 {
		return nil
	} else if cows == 0 {
		return fmt.Errorf("%d cows are invalid: no cows don't need food", cows)
	}
	return fmt.Errorf("%d cows are invalid: there are no negative cows", cows)
}
