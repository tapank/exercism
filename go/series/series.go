package series

// All returns a slice of all series that are n long
func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	if n == len(s) {
		return []string{s}
	}
	out := []string{}
	for i := 0; i <= len(s)-n; i++ {
		out = append(out, s[i:i+n])
	}
	return out
}

// UnsafFirst returns the first series that is n long.
// if the string is not long enough, it returns an empty string.
func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[:n]
}

// First returns the first series that is n long and a boolean
// indicating whether the series was sucessfully created.
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
