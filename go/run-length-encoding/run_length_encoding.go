package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	if len(input) <= 1 {
		return input
	}
	count := 0
	sb := strings.Builder{}
	prev := rune(input[0])
	for _, r := range input {
		if r >= '0' && r <= '9' {
			panic("encountered digit")
		} else if r == prev {
			count++
		} else {
			if count > 1 {
				sb.WriteString(strconv.Itoa(count))
			}
			sb.WriteRune(prev)
			prev = r
			count = 1
		}
	}
	if count > 1 {
		sb.WriteString(strconv.Itoa(count))
	}
	sb.WriteRune(prev)
	return sb.String()
}

func RunLengthDecode(input string) string {
	if len(input) <= 1 {
		return input
	}
	count := 0
	sb := strings.Builder{}
	for _, r := range input {
		if r >= '0' && r <= '9' {
			count = count*10 + int(r-'0')
		} else if count > 0 {
			for count > 0 {
				sb.WriteRune(r)
				count--
			}
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
