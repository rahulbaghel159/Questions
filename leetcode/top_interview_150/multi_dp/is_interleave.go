package multidp

// https://leetcode.com/problems/interleaving-string
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	memo := make([][]int, len(s1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(s2))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}

	return is_interleave(s1, 0, s2, 0, 0, memo, s3)
}

func is_interleave(s1 string, i int, s2 string, j int, k int, memo [][]int, s3 string) bool {
	if i == len(s1) {
		return s3[k:] == s2[j:]
	}

	if j == len(s2) {
		return s3[k:] == s1[i:]
	}

	if memo[i][j] != -1 {
		return memo[i][j] == 1
	}

	ans := false

	if (s3[k] == s1[i] && is_interleave(s1, i+1, s2, j, k+1, memo, s3)) ||
		(s2[j] == s3[k] && is_interleave(s1, i, s2, j+1, k+1, memo, s3)) {
		ans = true
	}

	if ans {
		memo[i][j] = 1
	} else {
		memo[i][j] = 0
	}

	return ans
}
