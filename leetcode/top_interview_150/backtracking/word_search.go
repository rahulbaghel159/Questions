package backtracking

// https://leetcode.com/problems/word-search/
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	rows, cols := len(board), len(board[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backtrack(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func backtrack(board [][]byte, row, col int, word string, index int) bool {
	if len(board) == 0 {
		return false
	}

	if index >= len(word) {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) || word[index] != board[row][col] {
		return false
	}

	board[row][col] = '#'
	ret, rowOffset, colOffset := false, []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	for d := 0; d < 4; d++ {
		ret = backtrack(board, row+rowOffset[d], col+colOffset[d], word, index+1)
		if ret {
			break
		}
	}

	board[row][col] = word[index]
	return ret
}
