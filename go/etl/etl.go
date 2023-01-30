package etl

// Transform transforms old format of scores to the new format.
// This implementation does not handle bad input.
func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for score, alphabets := range in {
		for _, char := range alphabets {
			out[string(rune(char[0])|0b100000)] = score
		}
	}
	return out
}
