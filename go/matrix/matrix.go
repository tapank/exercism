package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	m := Matrix{}
	rows := strings.Split(s, "\n")
	for i, row := range rows {
		mrow := []int{}
		for _, v := range strings.Split(row, " ") {
			if v == "" {
				continue
			}
			if d, err := strconv.Atoi(v); err != nil {
				return nil, err
			} else {
				mrow = append(mrow, d)
			}
		}
		if len(mrow) == 0 {
			return nil, errors.New("empty row")
		}
		if i > 0 && len(mrow) != len(m[0]) {
			return nil, errors.New("uneven row lengths")
		}
		m = append(m, mrow)
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := [][]int{}
	for c := 0; c < len(m[0]); c++ {
		col := []int{}
		for r := 0; r < len(m); r++ {
			col = append(col, m[r][c])
		}
		cols = append(cols, col)
	}
	return cols
}

func (m Matrix) Rows() [][]int {
	rows := [][]int{}
	for _, mrow := range m {
		row := []int{}
		for _, v := range mrow {
			row = append(row, v)
		}
		rows = append(rows, row)
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
