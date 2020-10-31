package recursion2

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	fmt.Print("matrix")
	fmt.Println(matrix)

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	n := len(matrix)
	m := len(matrix[0])

	pivot := matrix[m/2][n/2]

	fmt.Println("pivot", pivot)

	if pivot == target {
		return true
	}

	if pivot > target {
		return searchMatrix(copyMatrix(matrix, 0, m/2, 0, n/2), target) || searchMatrix(copyMatrix(matrix, 0, m/2, n/2, n), target) || searchMatrix(copyMatrix(matrix, m/2, m, 0, n/2), target)
	}
	return searchMatrix(copyMatrix(matrix, m/2+1, m, n/2+1, n), target) || searchMatrix(copyMatrix(matrix, m/2+1, m, 0, n/2+1), target) || searchMatrix(copyMatrix(matrix, 0, m/2+1, n/2+1, n), target)
}

func copyMatrix(matrix [][]int, m1, m2, n1, n2 int) [][]int {
	result := make([][]int, 0)
	for i := m1; i < m2; i++ {
		temp := make([]int, 0)
		for j := n1; j < n2; j++ {
			temp = append(temp, matrix[i][j])
		}
		result = append(result, temp)
	}
	return result
}
