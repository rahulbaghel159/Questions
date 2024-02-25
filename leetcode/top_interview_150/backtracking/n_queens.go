package backtracking

// https://leetcode.com/problems/n-queens-ii/
func totalNQueens(n int) int {
	return backtrackQueens(n, 0, make(map[int]struct{}), make(map[int]struct{}), make(map[int]struct{}))
}

func backtrackQueens(size, row int, cols, diagnoal, antidiagonal map[int]struct{}) int {
	if row == size {
		return 1
	}

	solutions := 0
	for col := 0; col < size; col++ {
		currDiagonal, currAntiDiagonal := row-col, row+col
		_, ok_col := cols[col]
		_, ok_diagonal := diagnoal[currDiagonal]
		_, ok_anti_diagonal := antidiagonal[currAntiDiagonal]

		if ok_col || ok_diagonal || ok_anti_diagonal {
			continue
		}

		cols[col] = struct{}{}
		diagnoal[currDiagonal] = struct{}{}
		antidiagonal[currAntiDiagonal] = struct{}{}

		solutions += backtrackQueens(size, row+1, cols, diagnoal, antidiagonal)

		delete(cols, col)
		delete(diagnoal, currDiagonal)
		delete(antidiagonal, currAntiDiagonal)
	}

	return solutions
}
