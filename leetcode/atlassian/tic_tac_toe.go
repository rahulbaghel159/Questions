package atlassian

// https://leetcode.com/problems/design-tic-tac-toe/
// type TicTacToe struct {
// 	board [][]int
// }

// func Constructor4(n int) TicTacToe {
// 	board := make([][]int, n)
// 	for i := 0; i < len(board); i++ {
// 		board[i] = make([]int, n)
// 	}

// 	return TicTacToe{
// 		board: board,
// 	}
// }

// func (this *TicTacToe) Move(row int, col int, player int) int {
// 	this.board[row][col] = player
// 	if this.isWinner(row, col, player) {
// 		return player
// 	}

// 	return 0
// }

// func (this *TicTacToe) isWinner(row, col, player int) bool {
// 	//check row
// 	row_flag := true
// 	for j := 0; j < len(this.board[row]); j++ {
// 		if this.board[row][j] != player {
// 			row_flag = false
// 		}

// 		if !row_flag {
// 			break
// 		}
// 	}

// 	if row_flag {
// 		return true
// 	}

// 	//check col
// 	col_flag := true
// 	for i := 0; i < len(this.board); i++ {
// 		if this.board[i][col] != player {
// 			col_flag = false
// 		}

// 		if !col_flag {
// 			break
// 		}
// 	}

// 	if col_flag {
// 		return true
// 	}

// 	//check diagonal
// 	diagonal_flag := true
// 	for i := 0; i < len(this.board); i++ {
// 		if this.board[i][i] != player {
// 			diagonal_flag = false
// 		}
// 		if !diagonal_flag {
// 			break
// 		}
// 	}

// 	if diagonal_flag {
// 		return true
// 	}

// 	//check anti-diagonal
// 	anti_diagonal_flag := true
// 	for i, j := len(this.board)-1, 0; i >= 0 && j < len(this.board); i, j = i-1, j+1 {
// 		if this.board[i][j] != player {
// 			anti_diagonal_flag = false
// 		}

// 		if !anti_diagonal_flag {
// 			break
// 		}
// 	}

// 	if anti_diagonal_flag {
// 		return true
// 	}

// 	return false
// }

type TicTacToe struct {
	rows, cols              []int
	diagonal, anti_diagonal int
}

func Constructor4(n int) TicTacToe {
	return TicTacToe{
		rows: make([]int, n),
		cols: make([]int, n),
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	current_player := player
	if current_player != 1 {
		current_player = -1
	}

	this.rows[row] += current_player
	this.cols[col] += current_player

	if row == col {
		this.diagonal += current_player
	}

	if col == len(this.cols)-row-1 {
		this.anti_diagonal += current_player
	}

	n := len(this.rows)
	if abs(this.rows[row]) == n || abs(this.cols[col]) == n || abs(this.diagonal) == n || abs(this.anti_diagonal) == n {
		return player
	}

	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
