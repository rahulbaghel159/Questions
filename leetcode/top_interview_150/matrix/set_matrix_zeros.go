package matrix

// https://leetcode.com/problems/set-matrix-zeroes
// func setZeroes(matrix [][]int) {
// 	if len(matrix) == 0 {
// 		return
// 	}

// 	indexes := make([][]int, 0)

// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[0]); j++ {
// 			if matrix[i][j] == 0 {
// 				indexes = append(indexes, []int{i, j})
// 			}
// 		}
// 	}

// 	for _, idx := range indexes {
// 		r, c := idx[0], idx[1]

// 		for i := 0; i < len(matrix); i++ {
// 			matrix[i][c] = 0
// 		}

// 		for j := 0; j < len(matrix[0]); j++ {
// 			matrix[r][j] = 0
// 		}
// 	}
// }

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	isCol := false

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if isCol {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

}
