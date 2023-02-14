package stateoftictactoe

import (
	"errors"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (state State, err error) {
	// create board
	b := make([][]rune, 0, len(board))
	for i, line := range board {
		b = append(b, []rune(line))
		if i > 2 || len(b[i]) != 3 {
			err = errors.New("illegal board dimensions")
			return
		}
	}

	// validate board
	xcount, ocount, bcount := count('X', b), count('O', b), count(' ', b)
	if xcount+ocount+bcount != 9 {
		err = errors.New("illegal moves")
		return
	}
	if diff := xcount - ocount; diff > 1 || diff < -1 {
		err = errors.New("players kept playing")
		return
	}
	if ocount > xcount {
		err = errors.New("o started")
		return
	}

	// trackers
	wins := make(map[rune]int)

	// lines to check (rows, columns, and diagonals)
	lines := [][3]rune{
		// rows
		{b[0][0], b[0][1], b[0][2]},
		{b[1][0], b[1][1], b[1][2]},
		{b[2][0], b[2][1], b[2][2]},
		// columns
		{b[0][0], b[1][0], b[2][0]},
		{b[0][1], b[1][1], b[2][1]},
		{b[0][2], b[1][2], b[2][2]},
		// diagonals
		{b[0][0], b[1][1], b[2][2]},
		{b[0][2], b[1][1], b[2][0]},
	}
	for _, l := range lines {
		countWins(l, wins)
	}

	// determine outcome
	switch {
	case wins['O'] > 0 && wins['X'] > 0:
		err = errors.New("players kept playing")
	case wins['O'] > 0 || wins['X'] > 0:
		state = Win
	case bcount == 0:
		state = Draw
	default:
		state = Ongoing
	}
	return state, err
}

func countWins(marks [3]rune, wins map[rune]int) {
	var xcount, ocount int
	for _, v := range marks {
		switch v {
		case 'X':
			xcount++
		case 'O':
			ocount++
		}
	}
	if xcount == 3 {
		wins['X']++
	} else if ocount == 3 {
		wins['O']++
	}
}

func count(r rune, board [][]rune) (cnt int) {
	for _, row := range board {
		for _, v := range row {
			if v == r {
				cnt++
			}
		}
	}
	return
}
