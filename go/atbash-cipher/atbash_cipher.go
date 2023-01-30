package atbash

var atbash = map[rune]rune{}

// populate `atbash` map
func init() {
	// translate lower case alphabet.
	for i, j := 'a', 'z'; i <= 'z'; i, j = i+1, j-1 {
		atbash[i] = j
	}

	// translate upper case alphabet.
	for i, j := 'A', 'z'; i <= 'Z'; i, j = i+1, j-1 {
		atbash[i] = j
	}

	// fill in digits as is. no need for translation.
	for i := '0'; i <= '9'; i++ {
		atbash[i] = i
	}
}

// Atbash replaces alphabets with its transposed equivalent.
// Digits are left unchanged and result is groups of 5 lower
// case alphanumeric strings separated by space.
func Atbash(s string) string {
	out := []rune{}
	pos := 0
	for _, r := range s {
		if ch, ok := atbash[r]; ok {
			if pos == 5 {
				out = append(out, ' ')
				pos = 0
			}
			out = append(out, ch)
			pos++
		}
	}
	return string(out)
}
