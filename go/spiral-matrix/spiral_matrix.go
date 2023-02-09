package spiralmatrix

// Position remembers the current position
// and direction of movement within the box.
type Position struct {
	size      int
	r, c      int
	steps     [4][2]int
	stepindex int
}

// New returns a new Position such that the first
// call to Peek or Next returns the 0, 0 position.
func New(size int) *Position {
	steps := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	return &Position{size: size, steps: steps, c: -1}
}

// Next returns the new position and also updates it.
func (p *Position) Next(turn bool) (int, int) {
	if turn {
		p.stepindex++
		if p.stepindex == len(p.steps) {
			p.stepindex = 0
		}
	}
	p.r, p.c = p.Peek()
	return p.r, p.c
}

// Peek returns the next position in the current direction.
// It does not alter the current position
func (p *Position) Peek() (int, int) {
	nr := p.r + p.steps[p.stepindex][0]
	nc := p.c + p.steps[p.stepindex][1]
	return nr, nc
}

func SpiralMatrix(size int) [][]int {
	if size < 1 {
		return [][]int{}
	}

	// create empty box
	box := make([][]int, size)
	for i := range box {
		box[i] = make([]int, size)
	}

	// fill the box
	p := New(size)
	cnt, lim := 0, size*size
	for {
		// check if the next position is valid
		r, c := p.Peek()
		if r >= 0 && r < size && c >= 0 && c < size && box[r][c] == 0 {
			r, c = p.Next(false)
		} else {
			r, c = p.Next(true)
		}
		cnt++
		box[r][c] = cnt
		if cnt == lim {
			break
		}
	}
	return box
}
