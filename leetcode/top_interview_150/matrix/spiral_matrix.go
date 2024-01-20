package matrix

// https://leetcode.com/problems/spiral-matrix/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	up, down, right, left, res, index := 0, rows-1, cols-1, 0, make([]int, rows*cols), 0

	for index < rows*cols {
		for col := left; col <= right; col++ {
			res[index] = matrix[up][col]
			index++
		}

		for row := up + 1; row <= down; row++ {
			res[index] = matrix[row][right]
			index++
		}

		if up != down {
			for col := right - 1; col >= left; col-- {
				res[index] = matrix[down][col]
				index++
			}
		}

		if left != right {
			for row := down - 1; row > up; row-- {
				res[index] = matrix[row][left]
				index++
			}
		}

		left++
		right--
		up++
		down--
	}

	return res
}
