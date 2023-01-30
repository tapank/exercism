package summultiples

// SumMultiples returns sum of numbers from 1 to `limit`
// that are divisible by any of the `divisors`.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, div := range divisors {
			if div != 0 && i%div == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
