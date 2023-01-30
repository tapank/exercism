package pangram

func IsPangram(input string) bool {
	// create alphabet tracker
	unseen := map[rune]bool{}
	for r := 'a'; r <= 'z'; r++ {
		unseen[r] = true
	}

	ctr := 0 // alphabet counter
	for _, r := range input {
		// convert to lower case
		if r >= 'A' && r <= 'Z' {
			r |= 0b100000
		}

		// update tracker
		if unseen[r] {
			unseen[r] = false
			ctr++
		}

		// have we seen it all?
		if ctr == 26 {
			return true
		}
	}
	return false
}
