package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	switch {
	case distance <= 1.0:
		return 10.0
	case distance <= 5.0:
		return 5.0
	case distance <= 10.0:
		return 1.0
	}
	return 0.0
}
