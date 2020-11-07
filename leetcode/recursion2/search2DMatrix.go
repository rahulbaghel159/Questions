package recursion2

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])

	i, j := chosePivot(matrix, target)
	pivot := matrix[i][j]

	if pivot == target {
		return true
	}

	return searchMatrix(copyMatrix(matrix, i, m-1, 0, j-1), target) || searchMatrix(copyMatrix(matrix, 0, i, j+1, n-1), target)
}

func copyMatrix(matrix [][]int, m1, m2, n1, n2 int) [][]int {
	result := make([][]int, 0)
	for i := m1; i <= m2; i++ {
		temp := make([]int, 0)
		for j := n1; j <= n2; j++ {
			temp = append(temp, matrix[i][j])
		}
		result = append(result, temp)
	}
	return result
}

func chosePivot(matrix [][]int, target int) (i, j int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0, 0
	}

	//method 1: chose middle element as pivot, this would reduce search in 3 quadrants
	/*
		m := len(matrix)
		n := len(matrix[0])
		return m / 2, n / 2
	*/

	//method 2: chose middle column and then iterate over all elements to find boundary pivot
	// this would reduce search in 2 quadrants
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		//pivot same as target, search success, return this as pivot
		if matrix[i][n/2] == target {
			return i, n / 2
		}
		//if element at index is greater than target, pivot found
		if matrix[i][n/2] > target {
			return i, n / 2
		}
	}
	return m - 1, n / 2
}
