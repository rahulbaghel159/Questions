package backtracking

// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	return generateParenthesisHelper(0, n*2, 0, "")
}

func generateParenthesisHelper(length, n, val int, curr string) []string {
	if n <= 1 {
		return []string{}
	}

	if n == 2 {
		return []string{"()"}
	}

	if length == n {
		if val == 0 {
			return []string{curr}
		} else {
			return []string{}
		}
	}

	if val < 0 {
		return []string{}
	}

	tmp1, tmp2 := curr+"(", curr+")"
	ans1 := generateParenthesisHelper(length+1, n, val+1, tmp1)
	ans2 := generateParenthesisHelper(length+1, n, val-1, tmp2)

	res, k := make([]string, len(ans1)+len(ans2)), 0
	for _, s := range ans1 {
		res[k] = s
		k++
	}

	for _, s := range ans2 {
		res[k] = s
		k++
	}

	return res
}
