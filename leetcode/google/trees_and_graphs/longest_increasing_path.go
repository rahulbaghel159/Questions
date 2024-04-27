package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3072/
var cache [][]int

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	cache = make([][]int, len(matrix))
	for i := 0; i < len(cache); i++ {
		cache[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(cache[i]); j++ {
			cache[i][j] = -1
		}
	}

	maxLen := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if cache[i][j] == -1 {
				cache[i][j] = triggerDFS(matrix, i, j)
			}
			maxLen = max(maxLen, cache[i][j])
		}
	}

	return maxLen
}

func triggerDFS(matrix [][]int, r, c int) int {
	if len(matrix) == 0 || r < 0 || r > len(matrix) || c < 0 || c > len(matrix[0]) {
		return 0
	}

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	maxLen := 0
	for _, d := range directions {
		new_r, new_c := r+d[0], c+d[1]
		if new_r >= 0 && new_r < len(matrix) && new_c >= 0 && new_c < len(matrix[0]) && matrix[new_r][new_c] > matrix[r][c] {
			if cache[new_r][new_c] == -1 {
				cache[new_r][new_c] = triggerDFS(matrix, new_r, new_c)
			}
			maxLen = max(maxLen, cache[new_r][new_c])
		}
	}

	return maxLen + 1
}
