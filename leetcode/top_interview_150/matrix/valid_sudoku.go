package matrix

// https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	row, col, box := make([]int, 10), make([]int, 10), make([]int, 10)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {

			if board[i][j] == '.' {
				continue
			}

			val := board[i][j] - '0'
			pos := 1 << (val - 1)

			if row[i]&pos > 0 {
				return false
			}
			row[i] |= pos

			if col[j]&pos > 0 {
				return false
			}
			col[j] |= pos

			idx := (i/3)*3 + j/3
			if box[idx]&pos > 0 {
				return false
			}
			box[idx] |= pos
		}
	}

	return true
}
