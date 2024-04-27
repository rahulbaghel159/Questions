package atlassian

// https://leetcode.com/problems/longest-string-chain/
func longestStrChain(words []string) int {
	memo, wordsPresent := make(map[string]int), make(map[string]struct{})

	for _, w := range words {
		wordsPresent[w] = struct{}{}
	}

	ans := 0
	for _, w := range words {
		ans = max(ans, dfsChain(wordsPresent, memo, w))
	}

	return ans
}

func dfsChain(wordsPresent map[string]struct{}, memo map[string]int, word string) int {
	if _, ok := memo[word]; ok {
		return memo[word]
	}

	maxLen, word_arr := 1, []byte(word)
	for i := 0; i < len(word_arr); i++ {
		s := ""
		if i > 0 && i+1 < len(word_arr) {
			s = string(word_arr[:i]) + string(word_arr[i+1:])
		} else if i == 0 {
			s = string(word_arr[1:])
		} else if i == len(word_arr)-1 {
			s = string(word[:len(word_arr)-1])
		}

		if _, ok := wordsPresent[s]; ok {
			currLen := 1 + dfsChain(wordsPresent, memo, s)
			maxLen = max(maxLen, currLen)
		}
	}

	memo[word] = maxLen

	return maxLen
}
