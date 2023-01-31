package prime

// Factors returns all factors of a given positive integer.
// Solution is unoptimized.
func Factors(n int64) []int64 {
	if n < 0 {
		return nil
	}
	factors := []int64{}
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n = n / i
		}
	}
	return factors
}
