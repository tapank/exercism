package rotationalcipher

// RotationalCipher rotates each english alphabet in the by the shiftkey.
// Other characters in the string remain unchanged.
func RotationalCipher(plain string, shiftKey int) string {
	in := []rune(plain)
	out := make([]rune, len(in))
	for i, r := range in {
		if r >= 'a' && r <= 'z' {
			out[i] = 'a' + (r-'a'+rune(shiftKey))%26
		} else if r >= 'A' && r <= 'Z' {
			out[i] = 'A' + (r-'A'+rune(shiftKey))%26
		} else {
			out[i] = r
		}
	}
	return string(out)
}
