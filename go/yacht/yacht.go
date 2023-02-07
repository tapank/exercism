package yacht

import (
	"sort"
)

var categories map[string]func([]int) int

func init() {
	categories = make(map[string]func([]int) int)

	// singles solver
	singles := func(val int) func(dice []int) int {
		return func(dice []int) (score int) {
			for _, n := range dice {
				if n == val {
					score += val
				}
			}
			return
		}
	}
	categories["ones"] = singles(1)
	categories["twos"] = singles(2)
	categories["threes"] = singles(3)
	categories["fours"] = singles(4)
	categories["fives"] = singles(5)
	categories["sixes"] = singles(6)

	// full house solver
	categories["full house"] = func(dice []int) int {
		counter := make(map[int]int)
		for _, n := range dice {
			counter[n]++
		}
		var score int
		var two, three bool
		for val, count := range counter {
			if count == 2 {
				two = true
			} else if count == 3 {
				three = true
			}
			score += val * count
		}
		if two && three {
			return score
		}
		return 0
	}

	// four of a kind solver
	categories["four of a kind"] = func(dice []int) int {
		counter := make(map[int]int)
		for _, n := range dice {
			counter[n]++
		}
		for val, count := range counter {
			if count >= 4 {
				return val * 4
			}
		}
		return 0
	}

	// straights solver
	straights := func(delta int) func(dice []int) int {
		return func(dice []int) int {
			sort.Ints(dice)
			for i := 0; i < len(dice); i++ {
				if dice[i] != i+delta {
					return 0
				}
			}
			return 30
		}
	}
	categories["little straight"] = straights(1)
	categories["big straight"] = straights(2)

	// choice solver
	categories["choice"] = func(dice []int) (score int) {
		for _, n := range dice {
			score += n
		}
		return
	}

	// yatch solver
	categories["yacht"] = func(dice []int) int {
		val := dice[0]
		for _, n := range dice {
			if n != val {
				return 0
			}
		}
		return 50
	}
}

func Score(dice []int, category string) int {
	if len(dice) != 5 {
		return 0
	}
	if f, ok := categories[category]; ok {
		return f(dice)
	}
	return 0
}
