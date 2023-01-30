package perfect

import (
	"errors"
)

type Classification int

const (
	ClassificationAbundant = iota
	ClassificationDeficient
	ClassificationPerfect
)

var ErrOnlyPositive = errors.New("number must be a positive values")

// Classify categorizes a given positive integer based on Nicomachus'
// classification scheme. This solution is unoptimized.
func Classify(n int64) (Classification, error) {
	if n < 1 {
		return -1, ErrOnlyPositive
	}
	var sum int64
	for i := int64(1); i <= n/2 && sum <= n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	if sum < n {
		return ClassificationDeficient, nil
	}
	if sum > n {
		return ClassificationAbundant, nil
	}
	return ClassificationPerfect, nil
}
