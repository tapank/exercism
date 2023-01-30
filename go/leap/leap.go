package leap

// IsLeapYear returns a true if the given year is a leap year, else false
func IsLeapYear(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%100 == 0 {
		return false
	}
	return y%4 == 0
}
