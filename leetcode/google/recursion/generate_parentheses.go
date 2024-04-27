package recursion

// https://leetcode.com/explore/interview/card/google/62/recursion-4/3079/
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	result := make([]string, 0)
	backtrackParentheses(n*2, "", &result)

	return result
}

func backtrackParentheses(n int, curr string, result *[]string) {
	if len(curr) == n {
		if isValidParentheses(curr) {
			*result = append(*result, curr)
		}
		return
	}

	backtrackParentheses(n, curr+"(", result)
	backtrackParentheses(n, curr+")", result)
}

func isValidParentheses(s string) bool {
	if len(s) == 0 {
		return true
	}

	open := 0
	for _, r := range s {
		if r == '(' {
			open++
		} else {
			open--
		}

		if open < 0 {
			return false
		}
	}

	return open == 0
}
