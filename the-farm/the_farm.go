package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(cal FodderCalculator, cows int) (float64, error) {
	fodderAmount, err := cal.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	flatteningFactor, err := cal.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodderAmount * flatteningFactor / float64(cows), nil
}

func ValidateInputAndDivideFood(cal FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(cal, cows)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			message: "there are no negative cows",
			cows:    cows,
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			message: "no cows don't need food",
			cows:    0,
		}
	}

	return nil
}
