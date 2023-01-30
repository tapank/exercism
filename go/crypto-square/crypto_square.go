package cryptosquare

func Encode(pt string) string {
	if len(pt) < 2 {
		return pt
	}

	normal := normalize(pt)
	c, normal := sizeAndPad(normal)

	// create matrix
	matrix := [][]rune{}
	for i := 0; i < len(normal); i += c {
		matrix = append(matrix, normal[i:i+c])
	}
	return String(transpose(matrix))
}

func String(matrix [][]rune) string {
	s := ""
	for i, row := range matrix {
		if i > 0 {
			s += " "
		}
		s += string(row)
	}
	return s
}

// transpose two dimensional slice of runes
func transpose(matrix [][]rune) [][]rune {
	tmatrix := [][]rune{}
	for col := 0; col < len(matrix[0]); col++ {
		newrow := []rune{}
		for row := 0; row < len(matrix); row++ {
			newrow = append(newrow, matrix[row][col])
		}
		tmatrix = append(tmatrix, newrow)
	}
	return tmatrix
}

// returns column width and adds padding to slice as necessary
func sizeAndPad(normal []rune) (int, []rune) {
	// determine dimensions
	var r, c int
	for r*c < len(normal) {
		if r == c {
			c++
		} else {
			r++
		}
	}

	// pad spaces
	for i := len(normal); i < r*c; i++ {
		normal = append(normal, ' ')
	}
	return c, normal
}

// normalizes given string
func normalize(pt string) []rune {
	normal := []rune{}
	for _, r := range pt {
		if r >= 'A' && r <= 'Z' {
			r |= 0b100000
		}
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			normal = append(normal, r)
		}
	}
	return normal
}
