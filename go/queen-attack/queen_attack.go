package queenattack

import "errors"

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	// parse coordinates
	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errors.New("invalid position")
	}
	// white coordinates
	xw, yw := int(whitePosition[0]-'a'), int(whitePosition[1]-'1')
	// black coordinates
	xb, yb := int(blackPosition[0]-'a'), int(blackPosition[1]-'1')

	// validate coordinates
	nok := func(n int) bool {
		return n < 0 || n > 7
	}
	if nok(xw) || nok(yw) || nok(xb) || nok(yb) || whitePosition == blackPosition {
		return false, errors.New("invalid position")
	}

	// check attack position
	if x, y := xw-xb, yw-yb; x == 0 || y == 0 || x*x == y*y {
		// queens are on the same line or diagonal
		return true, nil
	}

	// queens are not on the same line or diagonal
	return false, nil
}
