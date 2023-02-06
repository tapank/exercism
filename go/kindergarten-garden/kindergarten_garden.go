package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	rows     [2][]rune
	children map[string]int
}

var Plants = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	garden := &Garden{rows: [2][]rune{}}

	// The diagram argument starts each row with a '\n'.  This allows Go's
	// raw string literals to present diagrams in source code nicely as two
	// rows flush left, for example,
	//
	//	diagram := `
	//	VVCCGG
	//	VVCCGG`
	if len(diagram) == 0 || diagram[0] != '\n' {
		return nil, errors.New("invalid garden")
	}
	rows := strings.Split(strings.TrimSpace(diagram), "\n")
	if len(rows) != 2 || len(rows[0]) != len(rows[1]) {
		return nil, errors.New("invalid garden")
	}
	garden.rows[0] = []rune(rows[0])
	garden.rows[1] = []rune(rows[1])

	// sort a copy of names alphabetically, then index them
	names := append([]string{}, children...)
	sort.Strings(names)
	garden.children = make(map[string]int)
	for i, child := range names {
		if _, found := garden.children[child]; found {
			return nil, errors.New("duplicate name")
		}
		garden.children[child] = i
	}
	if len(garden.rows[0]) != 2*len(children) {
		return nil, errors.New("invalid garden")
	}

	// validate cup codes
	for _, row := range garden.rows {
		for _, cupcode := range row {
			switch cupcode {
			case 'V', 'C', 'R', 'G':
				// valid code, do nothing
			default:
				return nil, errors.New("invalid cup code")
			}
		}
	}
	return garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if i, ok := g.children[child]; ok {
		p := make([]string, 4)
		p[0] = Plants[g.rows[0][i*2]]
		p[1] = Plants[g.rows[0][i*2+1]]
		p[2] = Plants[g.rows[1][i*2]]
		p[3] = Plants[g.rows[1][i*2+1]]
		return p, true
	}
	return nil, false
}
