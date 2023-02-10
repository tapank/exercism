package bottlesong

import (
	"fmt"
)

// number words lower case
var numLwords = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

// number words title case
var numTwords = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

func Recite(startBottles, takeDown int) []string {
	// validate arguments
	if startBottles < 1 || startBottles > 10 || startBottles-takeDown < 0 {
		return nil
	}

	// construct song
	song := make([]string, 0, takeDown*5)
	for i := startBottles; takeDown > 0; i, takeDown = i-1, takeDown-1 {
		// insert empty string before every verse but the first
		if i != startBottles {
			song = append(song, "")
		}

		// compose verse line by line
		l1and2 := fmt.Sprintf("%s green %s hanging on the wall,", numTwords[i], bword(i))
		l3 := "And if one green bottle should accidentally fall,"
		l4 := fmt.Sprintf("There'll be %s green %s hanging on the wall.", numLwords[i-1], bword(i-1))
		song = append(song, l1and2, l1and2, l3, l4)
	}
	return song
}

func bword(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}
