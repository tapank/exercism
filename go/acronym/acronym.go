package acronym

import "strings"

// Abbreviate the given term
func Abbreviate(s string) string {
	// sanitize input
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	// abbreviate
	var accr strings.Builder
	for _, word := range strings.Split(s, " ") {
		if len(word) > 0 {
			letter := rune(word[0])
			switch {
			case letter >= 'a' && letter <= 'z':
				accr.WriteRune(letter - 32)
			case letter >= 'A' && letter <= 'Z':
				accr.WriteRune(letter)
			}
		}
	}
	return accr.String()
}
