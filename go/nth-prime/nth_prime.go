package prime

import "errors"

var primes = []int{2}

// Nth returns the nth prime number where n is a positive integer.
// An error is returned if the nth prime number can't be calculated.
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be a positive integer")
	}
	for i := 3; len(primes) < n; i += 2 {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes[len(primes)-1], nil
}
