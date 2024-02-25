package binarysearch

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	leftRow, rightRow := 0, m-1

	for leftRow <= rightRow {
		midRow := leftRow + (rightRow-leftRow)/2
		//target in midrow
		if matrix[midRow][0] <= target && matrix[midRow][n-1] >= target {
			leftCol, rightCol := 0, n-1
			for leftCol <= rightCol {
				midCol := leftCol + (rightCol-leftCol)/2
				if matrix[midRow][midCol] == target {
					return true
				} else if matrix[midRow][midCol] > target {
					rightCol = midCol - 1
				} else {
					leftCol = midCol + 1
				}
			}
			return false
		} else if matrix[midRow][0] > target {
			rightRow = midRow - 1
		} else {
			leftRow = midRow + 1
		}
	}

	return false
}
