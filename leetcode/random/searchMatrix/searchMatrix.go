package searchmatrix

//https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	left, right, row := 0, len(matrix)-1, 0

	for left <= right {
		mid := (left + right) / 2

		if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
			row = mid
			break
		} else if target < matrix[mid][0] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	left, right = 0, len(matrix[row])-1

	for left <= right {
		mid := (left + right) / 2
		if target == matrix[row][mid] {
			return true
		} else if target < matrix[row][mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
