package matrix

// https://leetcode.com/problems/rotate-image
func rotate(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
}

func transpose(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

func reverse(matrix [][]int) {
	mid := len(matrix) / 2
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < mid; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][len(matrix)-j-1]
			matrix[i][len(matrix)-j-1] = tmp
		}
	}
}

// func rotate(matrix [][]int) {
// 	n := len(matrix)
// 	for i := 0; i < (n+1)/2; i++ {
// 		for j := 0; j < n/2; j++ {
// 			temp := matrix[n-1-j][i]
// 			matrix[n-1-j][i] = matrix[n-1-i][n-j-1]
// 			matrix[n-1-i][n-j-1] = matrix[j][n-1-i]
// 			matrix[j][n-1-i] = matrix[i][j]
// 			matrix[i][j] = temp
// 		}
// 	}
// }
