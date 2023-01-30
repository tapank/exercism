package armstrong

// IsNumber returns true if the given number is an Armstrong
// number, else false
func IsNumber(n int) bool {
	if n < 0 {
		return false
	}
	sum := 0
	for num, dcount := n, digitCount(n); num > 0; num /= 10 {
		sum += power(num%10, dcount)
	}
	return sum == n
}

// Counts the number of digits in a positive integer
func digitCount(n int) (cnt int) {
	for n > 0 {
		cnt++
		n /= 10
	}
	return
}

// computes the power of exp on base
func power(base, exp int) int {
	p := 1
	for i := 0; i < exp; i++ {
		p *= base
	}
	return p
}
