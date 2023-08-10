package mindeletesum

//https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/

func minimumDeleteSum(s1 string, s2 string) int {
	return computeCost(s1, s2, len(s1)-1, len(s2)-1, make(map[int]map[int]int, 0))
}

func computeCost(s1, s2 string, i, j int, memo map[int]map[int]int) int {
	if i < 0 && j < 0 {
		return 0
	}

	if _, o := memo[i]; o {
		if _, k := memo[i][j]; k {
			return memo[i][j]
		}
	} else {
		memo[i] = make(map[int]int)
	}

	if i < 0 {
		memo[i][j] = int(s2[j]) + computeCost(s1, s2, i, j-1, memo)
		return memo[i][j]
	}

	if j < 0 {
		memo[i][j] = int(s1[i]) + computeCost(s1, s2, i-1, j, memo)
		return memo[i][j]
	}

	if s1[i] == s2[j] {
		memo[i][j] = computeCost(s1, s2, i-1, j-1, memo)
	} else {
		memo[i][j] = min(
			int(s1[i])+computeCost(s1, s2, i-1, j, memo),
			int(s2[j])+computeCost(s1, s2, i, j-1, memo),
		)
	}

	return memo[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
