package ocr

import (
	"strings"
)

var digits = map[[4][3]rune]string{
	{{' ', '_', ' '}, {'|', ' ', '|'}, {'|', '_', '|'}, {' ', ' ', ' '}}: "0",
	{{' ', ' ', ' '}, {' ', ' ', '|'}, {' ', ' ', '|'}, {' ', ' ', ' '}}: "1",
	{{' ', '_', ' '}, {' ', '_', '|'}, {'|', '_', ' '}, {' ', ' ', ' '}}: "2",
	{{' ', '_', ' '}, {' ', '_', '|'}, {' ', '_', '|'}, {' ', ' ', ' '}}: "3",
	{{' ', ' ', ' '}, {'|', '_', '|'}, {' ', ' ', '|'}, {' ', ' ', ' '}}: "4",
	{{' ', '_', ' '}, {'|', '_', ' '}, {' ', '_', '|'}, {' ', ' ', ' '}}: "5",
	{{' ', '_', ' '}, {'|', '_', ' '}, {'|', '_', '|'}, {' ', ' ', ' '}}: "6",
	{{' ', '_', ' '}, {' ', ' ', '|'}, {' ', ' ', '|'}, {' ', ' ', ' '}}: "7",
	{{' ', '_', ' '}, {'|', '_', '|'}, {'|', '_', '|'}, {' ', ' ', ' '}}: "8",
	{{' ', '_', ' '}, {'|', '_', '|'}, {' ', '_', '|'}, {' ', ' ', ' '}}: "9",
}

func recognizeDigit(ds [4][3]rune) string {
	if d, ok := digits[ds]; ok {
		return d
	}
	return "?"
}

func readLine(s [4]string) (digits string) {
	// create rune matrix
	chars := make([][]rune, 4)
	var linelen int
	for i, line := range s {
		if i == 0 {
			linelen = len(line)
			if linelen%3 != 0 {
				panic("line length not a multiple of three")
			}
		} else if linelen != len(line) {
			panic("unequal lines!")
		}
		chars[i] = []rune(s[i])
	}

	// chop off one digit at a time and convert
	for len(chars[0]) > 0 {
		rawdigit := [4][3]rune{}
		for i := range chars {
			rawdigit[i] = [3]rune(chars[i][:3])
			chars[i] = chars[i][3:]
		}
		digits += recognizeDigit(rawdigit)
	}
	return digits
}

func Recognize(s string) []string {
	if len(s) == 0 {
		return nil
	}

	lines := strings.Split(s[1:], "\n") // ignore the leading newline
	if len(lines)%4 != 0 {
		panic("lines not multiples of four")
	}

	result := []string{}
	for len(lines) > 0 {
		result = append(result, readLine([4]string(lines[:4])))
		lines = lines[4:]
	}
	return result
}
