package pascal

func Triangle(n int) [][]int {
	// validate argument
	if n < 1 {
		return nil
	}

	// seed the first row
	t := [][]int{{1}}
	if n == 1 {
		return t
	}

	// construct triangle larger than a single row
	for i := 1; i < n; i++ {
		row := make([]int, i+1)
		row[0] = 1
		for j := 1; j < len(t[i-1]); j++ {
			row[j] = t[i-1][j-1] + t[i-1][j]
		}
		row[i] = 1
		t = append(t, row)
	}
	return t
}
