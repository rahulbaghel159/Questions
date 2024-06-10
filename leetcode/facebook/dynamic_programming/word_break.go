package dynamicprogramming

// https://leetcode.com/explore/interview/card/facebook/55/dynamic-programming-3/3036/
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for _, w := range wordDict {
		words[w] = struct{}{}
	}

	return recursiveWordBreak(s, words, make(map[string]bool))
}

func recursiveWordBreak(s string, wordDict map[string]struct{}, memo map[string]bool) bool {
	if _, ok := memo[s]; ok {
		return memo[s]
	}

	if len(s) == 0 {
		return true
	}

	ans := false
	for i := 0; i < len(s); i++ {
		if _, ok := wordDict[s[:i+1]]; ok {
			ans = ans || recursiveWordBreak(s[i+1:], wordDict, memo)
		}
	}

	memo[s] = ans

	return memo[s]
}
