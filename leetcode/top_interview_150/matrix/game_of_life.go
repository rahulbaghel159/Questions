package matrix

// https://leetcode.com/problems/game-of-life
func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			n := countNeighbour(board, i, j)

			if board[i][j] == 1 {
				if n < 2 {
					//rule 1
					board[i][j] = -1
				} else if n == 2 || n == 3 {
					//rule 2
					board[i][j] = 1
				} else if n > 3 {
					//rule 3
					board[i][j] = -1
				}
			} else {
				if n == 3 {
					//rule 4
					board[i][j] = 2
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func countNeighbour(board [][]int, r, c int) int {
	if len(board) == 0 {
		return 0
	}

	count := 0

	if r-1 >= 0 {
		if c-1 >= 0 && (board[r-1][c-1] == 1 || board[r-1][c-1] == -1) {
			count++
		}

		if board[r-1][c] == 1 || board[r-1][c] == -1 {
			count++
		}

		if c+1 < len(board[0]) && (board[r-1][c+1] == 1 || board[r-1][c+1] == -1) {
			count++
		}
	}

	if c-1 >= 0 && (board[r][c-1] == 1 || board[r][c-1] == -1) {
		count++
	}

	if c+1 < len(board[0]) && (board[r][c+1] == 1 || board[r][c+1] == -1) {
		count++
	}

	if r+1 < len(board) {
		if c-1 >= 0 && (board[r+1][c-1] == 1 || board[r+1][c-1] == -1) {
			count++
		}

		if board[r+1][c] == 1 || board[r+1][c] == -1 {
			count++
		}

		if c+1 < len(board[0]) && (board[r+1][c+1] == 1 || board[r+1][c+1] == -1) {
			count++
		}
	}

	return count
}
