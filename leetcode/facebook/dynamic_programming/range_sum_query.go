package dynamicprogramming

// https://leetcode.com/explore/interview/card/facebook/55/dynamic-programming-3/3037/
// type NumMatrix struct {
// 	matrix [][]int
// 	memo   map[[2]int]map[[2]int]int
// }

// func Constructor(matrix [][]int) NumMatrix {
// 	return NumMatrix{
// 		matrix: matrix,
// 		memo:   make(map[[2]int]map[[2]int]int, 0),
// 	}
// }

// func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
// 	if _, ok := this.memo[[2]int{row1, col1}][[2]int{row2, col2}]; ok {
// 		return this.memo[[2]int{row1, col1}][[2]int{row2, col2}]
// 	}

// 	sum := 0
// 	for i := row1; i <= row2; i++ {
// 		for j := col1; j <= col2; j++ {
// 			sum += this.matrix[i][j]
// 		}
// 	}

// 	if _, ok := this.memo[[2]int{row1, col1}]; !ok {
// 		this.memo[[2]int{row1, col1}] = make(map[[2]int]int, 0)
// 	}

// 	this.memo[[2]int{row1, col1}][[2]int{row2, col2}] = sum

// 	return sum
// }

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	dp := make([][]int, len(matrix))

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[i])+1)
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j+1] = dp[i][j] + matrix[i][j]
		}
	}

	return NumMatrix{
		dp: dp,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.dp[i][col2+1] - this.dp[i][col1]
	}

	return sum
}
