package house

import (
	"strings"
)

var firstLine = map[int]string{
	1:  "This is the house that Jack built.",
	2:  "This is the malt",
	3:  "This is the rat",
	4:  "This is the cat",
	5:  "This is the dog",
	6:  "This is the cow with the crumpled horn",
	7:  "This is the maiden all forlorn",
	8:  "This is the man all tattered and torn",
	9:  "This is the priest all shaven and shorn",
	10: "This is the rooster that crowed in the morn",
	11: "This is the farmer sowing his corn",
	12: "This is the horse and the hound and the horn",
}

var otherLines = map[int]string{
	2:  "that belonged to the farmer sowing his corn",
	3:  "that kept the rooster that crowed in the morn",
	4:  "that woke the priest all shaven and shorn",
	5:  "that married the man all tattered and torn",
	6:  "that kissed the maiden all forlorn",
	7:  "that milked the cow with the crumpled horn",
	8:  "that tossed the dog",
	9:  "that worried the cat",
	10: "that killed the rat",
	11: "that ate the malt",
	12: "that lay in the house that Jack built.",
}

func Verse(v int) string {
	if v < 1 || v > 12 {
		return ""
	}
	var verse strings.Builder
	verse.WriteString(firstLine[v])
	for i := 14 - v; i <= 12; i++ {
		verse.WriteRune('\n')
		verse.WriteString(otherLines[i])
	}
	return verse.String()
}

func Song() string {
	var song strings.Builder
	song.WriteString(Verse(1))
	for i := 2; i <= 12; i++ {
		song.WriteString("\n\n")
		song.WriteString(Verse(i))
	}
	return song.String()
}
