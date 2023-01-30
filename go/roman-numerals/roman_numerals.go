package romannumerals

import (
	"fmt"
)

var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	units     = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", fmt.Errorf("bad number: %d", input)
	}
	roman := thousands[input/1000]
	input %= 1000
	roman += hundreds[input/100]
	input %= 100
	roman += tens[input/10]
	input %= 10
	roman += units[input]
	return roman, nil
}
