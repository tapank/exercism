package ocr

import (
	"fmt"
	"strings"
)

func Recognize(s string) []string {
	if len(s) > 0 && s[0] == '\n' {
		s = s[1:]
	}
	result := []string{}
	fmt.Println("*****input:", s, "***")
	lines := strings.Split(s, "\n")
	fmt.Println("***number of lines:", len(lines))
	if len(lines)%4 != 0 {
		panic("lines not multiples of four")
	}
	for i := 0; len(lines) > 0; i += 4 {
		rowofdigits := [4]string{lines[i], lines[i+1], lines[i+2], lines[i+3]}
		lines = lines[4:]
		result = append(result, readLine(rowofdigits))
	}
	return result
}

func readLine(s [4]string) string {
	digits := ""

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

func recognizeDigit(ds [4][3]rune) string {
	if d, ok := digits[ds]; ok {
		return d
	}
	return "?"
}

var digits = map[[4][3]rune]string{
	{{' ', '_', ' '}, {'|', ' ', '|'}, {'|', '_', '|'}, {' ', ' ', ' '}}: "0",
}
