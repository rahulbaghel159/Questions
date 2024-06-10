package others

// https://leetcode.com/explore/interview/card/google/66/others-4/477/
type NumMatrix struct {
	matrix [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	// m := make([][]int, len(matrix))
	// for i := 0; i < len(matrix); i++ {
	// 	m[i] = make([]int, len(matrix[i]))
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		m[i][j] = matrix[i][j]
	// 	}
	// }

	return NumMatrix{
		matrix: matrix,
	}
}

func (this *NumMatrix) Update(row int, col int, val int) {
	if len(this.matrix) == 0 {
		return
	}

	if row < len(this.matrix) && col < len(this.matrix[0]) {
		this.matrix[row][col] = val
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.matrix[i][j]
		}
	}

	return sum
}
