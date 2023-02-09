package transpose

import (
	"strings"
)

func Transpose(input []string) []string {
	if len(input) == 0 {
		return input
	}

	// convert strings to rune slices
	in := [][]rune{}
	for _, row := range input {
		in = append(in, []rune(row))
	}

	// pad spaces to ensure no row is longer than the previous row.
	var length int
	for i := len(input) - 1; i >= 0; i-- {
		if len(input[i]) < length {
			in[i] = append(in[i], []rune(strings.Repeat(" ", length-len(in[i])))...)
		} else {
			length = len(input[i])
		}
	}

	// construct output strings
	output := []string{}
	for c := 0; c < len(in[0]); c++ {
		var sb strings.Builder
		for r := 0; r < len(in) && c < len(in[r]); r++ {
			sb.WriteRune(in[r][c])
		}
		output = append(output, sb.String())
	}
	return output
}
