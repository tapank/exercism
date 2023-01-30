package isbn

// IsValidISBN returns true if `isbn` is valid ISBN10 id, else false.
func IsValidISBN(isbn string) bool {
	sum := 0
	i := 10
	for _, r := range isbn {
		if r >= '0' && r <= '9' {
			sum += int(r-'0') * i
			i--
		} else if i == 1 && r == 'X' {
			sum += 10
			i--
		} else if r != '-' {
			return false
		}
	}
	return i == 0 && sum%11 == 0
}
