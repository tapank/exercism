package railfence

import "strings"

func Encode(message string, rails int) string {
	// bad argument
	if rails < 1 {
		return ""
	}

	// message remains unchanged
	if rails == 1 {
		return message
	}

	// track current row and the direction of increment
	pos, delta := 0, 1
	canvas := make([][]rune, rails)
	for _, r := range message {
		// adjust the current row position
		if pos == -1 {
			pos, delta = 1, 1
		} else if pos == rails {
			pos, delta = rails-2, -1
		}

		// create new row if not present
		if canvas[pos] == nil {
			canvas[pos] = make([]rune, 0, len(message)/rails+1)
		}

		// add the character to current row and increment position
		canvas[pos] = append(canvas[pos], r)
		pos += delta
	}

	// encrypted message is each row converted to string and joined together
	out := ""
	for _, row := range canvas {
		out += string(row)
	}
	return out
}

func Decode(message string, rails int) string {
	// illegal argument
	if rails < 1 {
		return ""
	}

	// in this case, message remains unchanged
	if len(message) < 3 || rails == 1 {
		return message
	}

	// track current row and the direction of increment
	pos, delta := 0, 1
	canvas := make([][]rune, rails)

	// create an empty canvas to determine the length of each row
	for range message {
		if pos == -1 {
			pos, delta = 1, 1
		} else if pos == rails {
			pos, delta = rails-2, -1
		}
		if canvas[pos] == nil {
			canvas[pos] = make([]rune, 0, len(message)/rails+1)
		}
		canvas[pos] = append(canvas[pos], 0)
		pos += delta
	}

	// relace each row with the actual message
	for i, row := range canvas {
		canvas[i] = []rune(message[:len(row)])
		message = message[len(row):]
	}

	// construct the decrypted message from the canvas
	var out strings.Builder
	pos, delta = 0, 1
	for {
		if pos == -1 {
			pos, delta = 1, 1
		} else if pos == rails {
			pos, delta = rails-2, -1
		}
		if len(canvas[pos]) == 0 {
			break
		}
		out.WriteRune(canvas[pos][0])
		canvas[pos] = canvas[pos][1:]
		pos += delta
	}
	return out.String()
}
