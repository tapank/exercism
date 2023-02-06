package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	// quick answer if char is A
	if char == 'A' {
		return "A", nil
	}

	// validate char range
	if char < 'A' || char > 'Z' {
		return "", errors.New("invalid argument")
	}

	// build the lines to create diamond
	l := int(char - 'A')
	lines := []string{}
	for i := 0; i <= l; i++ {
		s := strings.Repeat(" ", l-i)
		s += string('A' + rune(i))
		if i > 0 {
			s += strings.Repeat(" ", 2*i-1)
			s += string('A' + rune(i))
		}
		s += strings.Repeat(" ", l-i)
		lines = append(lines, s)
	}

	// now create the diamond by stitching together the lines
	var diamond strings.Builder
	// upper half of diamond
	for _, line := range lines {
		diamond.WriteString(line)
		diamond.WriteRune('\n')
	}
	// lower half of diamond
	for i := len(lines) - 2; i >= 0; i-- {
		diamond.WriteString(lines[i])
		if i > 0 {
			diamond.WriteRune('\n')
		}
	}
	return diamond.String(), nil
}
