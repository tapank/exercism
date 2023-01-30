package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	cnt := 0
	for _, b := range cb[file] {
		if b {
			cnt++
		}
	}
	return cnt
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	cnt := 0
	for _, file := range cb {
		if file[rank-1] {
			cnt++
		}
	}
	return cnt
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	cnt := 0
	for _, file := range cb {
		cnt += len(file)
	}
	return cnt
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	cnt := 0
	for i := 'A'; i <= 'H'; i++ {
		cnt += CountInFile(cb, string(i))
	}
	return cnt
}
