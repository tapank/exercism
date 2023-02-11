package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int
type Pair [2]int

const (
	MINMASK = 1
	MAXMASK = 2
)

func New(s string) (*Matrix, error) {
	m := make(Matrix, 0, 10)
	if len(s) == 0 {
		return &m, nil
	}
	for i, line := range strings.Split(s, "\n") {
		gridRow := make([]int, 0, 10)
		for _, v := range strings.Split(line, " ") {
			if num, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				gridRow = append(gridRow, num)
			}
		}
		if i > 0 && len(gridRow) != len(m[0]) {
			return nil, errors.New("unequal row lengths")
		}
		m = append(m, gridRow)
	}
	return &m, nil
}

func (m *Matrix) MarkMaxInRows(tracker [][]int) {
	for r, row := range *m {
		var max int
		// find max
		for _, val := range row {
			if val > max {
				max = val
			}
		}
		// mark tracker
		for c, val := range row {
			if val == max {
				tracker[r][c] |= MAXMASK
			}
		}
	}
}

func (m *Matrix) MarkMinInCols(tracker [][]int) {
	for c := 0; c < len((*m)[0]); c++ {
		var min int
		// find min
		for r := 0; r < len(*m); r++ {
			if r == 0 {
				min = (*m)[r][c]
			} else if (*m)[r][c] < min {
				min = (*m)[r][c]
			}
		}
		// mark tracker
		for r := 0; r < len(*m); r++ {
			if (*m)[r][c] == min {
				tracker[r][c] |= MINMASK
			}
		}
	}
}

func (m *Matrix) NewTracker() [][]int {
	var tracker [][]int
	tracker = make([][]int, 0, len(*m))
	for range *m {
		tracker = append(tracker, make([]int, len((*m)[0])))
	}
	return tracker
}

func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}
	if len(*m) == 0 {
		return pairs
	}
	tracker := m.NewTracker()
	m.MarkMaxInRows(tracker)
	m.MarkMinInCols(tracker)
	for r, row := range tracker {
		for c, val := range row {
			if val == 3 {
				pairs = append(pairs, Pair{r + 1, c + 1})
			}
		}
	}
	return pairs
}
