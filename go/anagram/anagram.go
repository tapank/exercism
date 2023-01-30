package anagram

import "strings"

// Detect returns anagrams of `subject` from `candidates`.
func Detect(subject string, candidates []string) (match []string) {
	for _, s := range candidates {
		if Anagram(subject, s) {
			match = append(match, s)
		}
	}
	return
}

// Anagram detects if the two given words are anagrams.
// Case does not matter.
func Anagram(s, t string) bool {
	s, t = strings.ToLower(s), strings.ToLower(t)
	if s == t {
		return false
	}
	ctr := map[rune]int{}
	// count alphabet frequency in s
	for _, r := range s {
		ctr[r]++
	}

	// count down alphabet frequency in t
	for _, r := range t {
		ctr[r]--
	}

	// a non zero frequency indicates non-anagram
	for _, v := range ctr {
		if v != 0 {
			return false
		}
	}
	return true
}
