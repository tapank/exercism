package luhn

func Valid(id string) bool {
	n, ok := parse(id)
	if !ok || len(n) < 2 {
		return false
	}

	// double every other digit from the right,
	// starting with the second last digit
	for i := len(n) - 2; i >= 0; i -= 2 {
		n[i] = n[i] * 2
		if n[i] > 9 {
			n[i] -= 9
		}
	}

	// sum all digits
	sum := 0
	for _, v := range n {
		sum += v
	}

	// sum divisible by 10 is valid
	return sum%10 == 0
}

func parse(s string) ([]int, bool) {
	n := []int{}
	for _, r := range s {
		switch {
		// keep digits
		case r >= '0' && r <= '9':
			n = append(n, int(r-'0'))
		// ignore spaces
		case r == ' ':
			// no op
		// anything else is illegal
		default:
			return nil, false
		}
	}
	return n, true
}
