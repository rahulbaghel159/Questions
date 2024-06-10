package dynamicprogramming

import "strconv"

// https://leetcode.com/explore/interview/card/facebook/55/dynamic-programming-3/264/
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	return recursiveCount(0, s, make(map[int]int))
}

func recursiveCount(index int, s string, memo map[int]int) int {
	if _, ok := memo[index]; ok {
		return memo[index]
	}

	if len(s) == 0 {
		return 0
	}

	if index == len(s) {
		return 1
	}

	if s[index] == '0' {
		return 0
	}

	if index == len(s)-1 {
		return 1
	}

	ans := recursiveCount(index+1, s, memo)
	sub, _ := strconv.Atoi(s[index : index+2])
	if sub <= 26 {
		ans += recursiveCount(index+2, s, memo)
	}

	memo[index] = ans

	return ans
}
