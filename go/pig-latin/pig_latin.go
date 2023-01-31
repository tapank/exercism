package piglatin

import "strings"

// words that start with vowel sounds start with one of these.
var VSounds = []string{"a", "e", "i", "o", "u", "xr", "yt"}

// Sentence converts a given sentence to pig latin. It splits the
// sentence by a single space and calls Word for translation.
// Limitations of this implementation: handles only lower case text
// and does not handle punctuation and numbers.
func Sentence(sentence string) string {
	pig := ""
	for i, word := range strings.Split(sentence, " ") {
		if i > 0 {
			pig += " "
		}
		pig += Word(word)
	}
	return pig
}

// Word converts a given word to pig latin. Limitations
// of this implementation: handles only lower case words
// and does not handle punctuation and numbers.
func Word(word string) string {
	// words beginning with vowel sounds
	for _, v := range VSounds {
		if strings.HasPrefix(word, v) {
			return word + "ay"
		}
	}
	// words beginning with "qu"
	if strings.HasPrefix(word, "qu") {
		return word[2:] + "quay"
	}
	// words beginning with a consonant followed by "qu"
	if strings.HasPrefix(word[1:], "qu") {
		return word[3:] + word[:3] + "ay"
	}
	// words that contains y after a consonant cluster
	if i := strings.IndexRune(word, 'y'); i > 0 {
		if !strings.ContainsAny(word[:i], "aeiou") {
			return word[i:] + word[:i] + "ay"
		}
	}

	// rest of the words that start with a consonant cluster
	var vIndex int
loop:
	for i, r := range word {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vIndex = i
			break loop
		}
	}
	return word[vIndex:] + word[:vIndex] + "ay"
}
