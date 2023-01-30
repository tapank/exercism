package wordcount

import "strings"

type Frequency map[string]int

// WordCount returns frequencies of words in a phrase.
// Case is ignored and special characters are removed.
func WordCount(phrase string) Frequency {
	ctr := Frequency{}
	phrase = strings.ReplaceAll(phrase, "\n", " ")
	phrase = strings.ReplaceAll(phrase, "\t", " ")
	phrase = strings.ReplaceAll(phrase, ",", " ")
	phrase = strings.ToLower(phrase)
	for _, word := range strings.Split(phrase, " ") {
		word = strings.Trim(word, "!.;:&'@$%^\"")
		if word != "" {
			ctr[word]++
		}
	}
	return ctr
}
