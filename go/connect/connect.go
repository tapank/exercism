package connect

import "errors"

// valid step increments
var steps = [][2]int{
	{0, 1},  // right
	{-1, 0}, // left
	{0, -1}, // up
	{1, 0},  // down
	{-1, 1}, // north east
	{1, -1}, // south west
}

type Connect struct {
	board      [][]rune
	rlen, clen int
}

func New(lines []string) (*Connect, error) {
	var rlen, clen int
	rlen = len(lines)
	board := make([][]rune, 0, rlen)
	for i, l := range lines {
		row := []rune(l)
		if i == 0 {
			clen = len(row)
		} else if clen != len(row) {
			return nil, errors.New("unequal rows")
		}
		board = append(board, row)
	}
	return &Connect{board: board, rlen: rlen, clen: clen}, nil
}

func (c *Connect) startAndEnd(player rune) ([][2]int, [2]int) {
	positions := [][2]int{}
	var endPos [2]int
	if player == 'O' {
		// O goes from top to bottom
		for i := 0; i < c.clen; i++ {
			positions = append(positions, [2]int{0, i})
		}
		endPos = [2]int{c.rlen - 1, -1}
	} else if player == 'X' {
		// X goes from left to right
		for i := 0; i < c.rlen; i++ {
			positions = append(positions, [2]int{i, 0})
		}
		endPos = [2]int{-1, c.clen - 1}
	} else {
		panic("unknown player")
	}
	return positions, endPos
}

func (c *Connect) nextSteps(row, col int) (nextSteps [][2]int) {
	nextSteps = make([][2]int, 0, len(steps))
	for _, step := range steps {
		nr, nc := row+step[0], col+step[1]
		if nr >= 0 && nr < c.rlen && nc >= 0 && nc < c.clen {
			nextSteps = append(nextSteps, [2]int{nr, nc})
		}
	}
	return nextSteps
}

func (c *Connect) traverse(pos [2]int, player rune, end [2]int) bool {
	if c.board[pos[0]][pos[1]] == player {
		c.board[pos[0]][pos[1]] = '.'
	} else {
		return false
	}
	if (end[0] != -1 && pos[0] == end[0]) || (end[1] != -1 && pos[1] == end[1]) {
		c.board[pos[0]][pos[1]] = player + 1
		return true
	}
	for _, step := range c.nextSteps(pos[0], pos[1]) {
		if c.traverse(step, player, end) {
			c.board[pos[0]][pos[1]] = player + 1
			return true
		}
	}
	return false
}

func ResultOf(lines []string) (string, error) {
	c, err := New(lines)
	if err != nil {
		return "", err
	}
	for _, player := range []rune{'O', 'X'} {
		start, end := c.startAndEnd(player)
		for _, pos := range start {
			if c.traverse(pos, player, end) {
				return string(player), nil
			}
		}
	}
	return "", nil
}
