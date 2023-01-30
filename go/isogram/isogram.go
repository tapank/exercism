package isogram

func IsIsogram(word string) bool {
	ctr := map[rune]bool{}
	for _, r := range word {
		r |= 0b100000
		if r >= 'a' && r <= 'z' {
			if ctr[r] {
				return false
			}
			ctr[r] = true
		}
	}
	return true
}
