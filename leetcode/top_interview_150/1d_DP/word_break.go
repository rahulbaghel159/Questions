package ddp

// https://leetcode.com/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))

	for i := 0; i < len(s); i++ {
		for _, word := range wordDict {
			if i < len(word)-1 {
				continue
			}

			if i == len(word)-1 || dp[i-len(word)] {
				if s[i-len(word)+1:i+1] == word {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(s)-1]
}
