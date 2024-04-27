package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3052/
func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	//transpose
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			if i != j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}

	//reverse each row
	for i := 0; i < len(matrix); i++ {
		for l, r := 0, len(matrix[i])-1; l < r; {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}
