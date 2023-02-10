package wordsearch

import (
	"errors"
	"strings"
)

func Solve(words []string, puzzle []string) (result map[string][2][2]int, err error) {
	// create word map to track finds
	wordmap := map[string]bool{}
	for _, word := range words {
		wordmap[word] = true
	}
	result = make(map[string][2][2]int)

	// setup puzzle matrix and validate row lengths
	in := make([][]rune, 0, len(puzzle))
	lenr, lenc := 0, 0
	for i, line := range puzzle {
		l := []rune(line)
		if i == 0 {
			lenc = len(l)
		} else if len(l) != lenc {
			err = errors.New("unequal row lengths in puzzle")
			return
		}
		in = append(in, []rune(line))
	}
	lenr = len(in)

	var (
		lim                    = [2]int{lenr, lenc}
		topL, topR, botL       = [2]int{0, 0}, [2]int{0, lenc - 1}, [2]int{lenr - 1, 0}
		sout, nrth, east, west = [2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}
		nest, nwst, swst, sest = [2]int{-1, 1}, [2]int{-1, -1}, [2]int{1, -1}, [2]int{1, 1}
	)

	scanWords(topL, sout, east, lim, in, wordmap, result) // left to right
	scanWords(topR, sout, west, lim, in, wordmap, result) // right to left
	scanWords(topL, east, sout, lim, in, wordmap, result) // top to bottom
	scanWords(botL, east, nrth, lim, in, wordmap, result) // bottom to top
	scanWords(topL, sout, nest, lim, in, wordmap, result) // left to right up
	scanWords(topL, sout, sest, lim, in, wordmap, result) // left to right down
	scanWords(topR, sout, nwst, lim, in, wordmap, result) // right to left up
	scanWords(topR, sout, swst, lim, in, wordmap, result) // right to left down
	scanWords(topL, east, sest, lim, in, wordmap, result) // top to bottom right
	scanWords(topL, east, swst, lim, in, wordmap, result) // top to bottom left
	scanWords(botL, east, nest, lim, in, wordmap, result) // bottom to top right
	scanWords(botL, east, nwst, lim, in, wordmap, result) // bottom to top left

	// all words should be found
	for _, b := range wordmap {
		if b {
			err = errors.New("not all words found")
		}
	}
	return
}

func scanWords(start [2]int, startInc [2]int, incr [2]int, lim [2]int, in [][]rune, wm map[string]bool, res map[string][2][2]int) {
	valid := func(r, c int) bool {
		return r >= 0 && r < lim[0] && c >= 0 && c < lim[1]
	}
	for r, c := start[0], start[1]; valid(r, c); r, c = r+startInc[0], c+startInc[1] {
		for rs, cs := r, c; valid(rs, cs); rs, cs = rs+incr[0], cs+incr[1] {
			var wb strings.Builder
			for re, ce := rs, cs; valid(re, ce); re, ce = re+incr[0], ce+incr[1] {
				wb.WriteRune(in[re][ce])
				if word := wb.String(); wm[word] {
					res[word] = [2][2]int{{cs, rs}, {ce, re}}
					wm[word] = false
				}
			}
		}
	}
}
