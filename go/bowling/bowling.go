package bowling

import "errors"

type Game struct {
	score  int
	pins   int
	frame  int
	throw  int
	strike bool
	status int
}

const (
	NOT_STARTED = iota
	STARTED
	ENDED
)

func NewGame() *Game {
	return &Game{frame: 1}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 || pins > 10 {
		return errors.New("illegal number of pins")
	}
	if g.status == NOT_STARTED {
		g.status = STARTED
	} else if g.status == ENDED {
		return errors.New("cannot roll after game is over")
	}
	g.throw++
	switch g.throw {
	case 1:
		if pins == 10 {
			g.strike = true
		} else {
			g.pins = pins
		}
	case 2:
		if pins != 10 && pins+g.pins > 10 {
			return errors.New("cannot score more than 10 in two throws")
		}
		if pins == 10 {
			g.score += 10
		} else if g.strike {
			g.score += 10 + pins
			g.pins = pins
			g.pins = 0
			g.throw = 0
			g.frame++
		} else if g.pins+pins == 10 {
			g.score += 10
			g.pins = 0
			g.throw = 0
			g.frame++
		} else {
			g.score += g.pins + pins
			g.pins = 0
			g.throw = 0
			g.frame++
		}
	case 3:
		if pins+g.pins > 10 {
			return errors.New("cannot score more than 10 in two throws")
		}
		if g.strike {
			g.score += 10 + pins
		} else {
			g.score += pins
		}
		g.pins = 0
		g.throw = 0
		g.frame++
		g.strike = false
	}
	if g.frame > 10 {
		g.status = ENDED
	}
	return nil
}

func (g *Game) Score() (int, error) {
	if g.status != ENDED {
		return g.score, errors.New("game not ended yet")
	}
	return g.score, nil
}
