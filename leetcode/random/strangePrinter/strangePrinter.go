package strangeprinter

//https://leetcode.com/problems/strange-printer/

func strangePrinter(s string) int {
	memo := make([][]int, len(s))

	for i := 0; i < len(s); i++ {
		memo[i] = make([]int, len(s))
		for j := 0; j < len(s); j++ {
			memo[i][j] = -1
		}
	}

	return solve(s, 0, len(s)-1, memo) + 1
}

func solve(s string, left, right int, memo [][]int) int {
	if memo[left][right] != -1 {
		return memo[left][right]
	}

	memo[left][right] = len(s)
	j := -1
	for i := left; i < right; i++ {
		if s[i] != s[right] && j == -1 {
			j = i
		}

		if j != -1 {
			memo[left][right] = min(memo[left][right], 1+solve(s, j, i, memo)+solve(s, i+1, right, memo))
		}
	}

	if j == -1 {
		memo[left][right] = 0
	}

	return memo[left][right]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
