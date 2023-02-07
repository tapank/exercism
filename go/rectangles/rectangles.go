package rectangles

func Count(diagram []string) int {
	grid := [][]rune{}
	for _, row := range diagram {
		grid = append(grid, []rune(row))
	}

	rects := 0
	// process rows top to bottom
	for row := 0; row < len(grid)-1; row++ {
		// locate the next top left corner of rect
		for left := 0; left < len(grid[row]); left++ {
			if grid[row][left] != '+' {
				continue
			}
			// locate the next top right corner of rect
			for right := left + 1; right < len(grid[row]); right++ {
				if grid[row][right] != '+' {
					continue
				}
				// check if the "left" to "right" makes a valid top
				if isTopBottom(grid, row, left, right) {
					// count rectangles by scanning down from "left" and "right"
					rects += countRects(grid, row, left, right)
				}
			}
		}
	}
	return rects
}

func isTopBottom(grid [][]rune, row, left, right int) bool {
	if left >= right {
		return false
	}
	if grid[row][left] != '+' || grid[row][right] != '+' {
		return false
	}
	for i := left + 1; i < right; i++ {
		switch grid[row][i] {
		case '-', '+':
			// all good
		default:
			return false
		}
	}
	return true
}

func countRects(grid [][]rune, row, left, right int) (count int) {
	if left >= right {
		return 0
	}
	for r := row + 1; r < len(grid); r++ {
		for _, c := range []int{left, right} {
			switch grid[r][c] {
			case '|', '+':
				// all good
			default:
				return
			}
		}
		if isTopBottom(grid, r, left, right) {
			count++
		}
	}
	return
}
