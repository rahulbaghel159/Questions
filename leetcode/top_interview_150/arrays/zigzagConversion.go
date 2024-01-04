package arrays

// https://leetcode.com/problems/zigzag-conversion
// func convert(s string, numRows int) string {
// 	if numRows == 1 {
// 		return s
// 	}

// 	ans, n, numCols, matrix, currRow, currCol := "", len(s), 0, make([][]string, numRows), 0, 0

// 	sections := int(math.Ceil(float64(n) / float64(2*numRows-2.0)))
// 	numCols = sections * (numRows - 1)

// 	for i := 0; i < numRows; i++ {
// 		matrix[i] = make([]string, numCols)
// 	}

// 	for i := 0; i < len(s); {
// 		for currRow < numRows && i < len(s) {
// 			matrix[currRow][currCol] = string(s[i])
// 			currRow++
// 			i++
// 		}
// 		currRow -= 2
// 		currCol++
// 		for currRow > 0 && currCol < numCols && i < len(s) {
// 			matrix[currRow][currCol] = string(s[i])
// 			currRow--
// 			currCol++
// 			i++
// 		}
// 	}

// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j < numCols; j++ {
// 			if matrix[i][j] != "" {
// 				ans += matrix[i][j]
// 			}
// 		}
// 	}

// 	return ans
// }

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	ans, n, charsInSection := "", len(s), 2*(numRows-1)

	for currRow := 0; currRow < numRows; currRow++ {
		index := currRow

		for index < n {
			ans += string(s[index])

			if currRow != 0 && currRow != numRows-1 {
				charsInBetween := charsInSection - 2*currRow
				secondIndex := index + charsInBetween

				if secondIndex < n {
					ans += string(s[secondIndex])
				}
			}
			index += charsInSection
		}
	}

	return ans
}
