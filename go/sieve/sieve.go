package sieve

func Sieve(limit int) []int {
	// setup sieve of Eratosthenes (soe)
	soe := make([]bool, limit+1)
	for i := 0; i < len(soe); i++ {
		soe[i] = true
	}
	soe[0], soe[1] = false, false

	// mark non primes as false
	for i := 0; i < len(soe); i++ {
		if soe[i] {
			for j := i * 2; j < len(soe); j += i {
				soe[j] = false
			}
		}
	}

	// create slice of primes
	primes := []int{}
	for i := 0; i < len(soe); i++ {
		if soe[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
