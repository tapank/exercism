package minesweeper

var deltas = [][2]int{
	{0, 1},   // right
	{0, -1},  // left
	{-1, 0},  // up
	{1, 0},   // down
	{-1, 1},  // north east
	{-1, -1}, // north west
	{1, 1},   // south east
	{1, -1},  // south west
}

func neighbors(r, c, rlen, clen int) [][2]int {
	next := make([][2]int, 0, 8)
	for _, step := range deltas {
		nr, nc := r+step[0], c+step[1]
		if nr >= 0 && nc >= 0 && nr < rlen && nc < clen {
			next = append(next, [2]int{nr, nc})
		}
	}
	return next
}

// Annotate returns an annotated board
func Annotate(board []string) []string {
	var rlen, clen int
	rlen = len(board)

	// create a rune board for annotation
	b := make([][]rune, 0, rlen)
	for i, line := range board {
		l := []rune(line)
		if i == 0 {
			clen = len(l)
		} else if len(l) != clen {
			panic("unequal row lengths")
		}
		b = append(b, l)
	}

	// annotate the board
	for r := 0; r < rlen; r++ {
		for c := 0; c < clen; c++ {
			if b[r][c] != ' ' {
				continue
			}
			var sum rune
			for _, pos := range neighbors(r, c, rlen, clen) {
				if b[pos[0]][pos[1]] == '*' {
					sum++
				}
			}
			if sum > 0 {
				b[r][c] = sum + '0'
			}
		}
	}

	// recreate the string board
	for i := range board {
		board[i] = string(b[i])
	}
	return board
}
