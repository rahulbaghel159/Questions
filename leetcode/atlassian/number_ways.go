package atlassian

// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary
func numWays(words []string, target string) int {
	if len(words) == 0 && len(target) > 0 {
		return 0
	}

	ALPHABET_SIZE, MOD, targetLength, wordLength := 26, 1_000_000_007, len(target), len(words[0])

	charOccurences := make([][]int, ALPHABET_SIZE)
	for i := 0; i < len(charOccurences); i++ {
		charOccurences[i] = make([]int, wordLength)
	}

	for col := 0; col < wordLength; col++ {
		for _, word := range words {
			charOccurences[word[col]-'a'][col]++
		}
	}

	dp := make([][]int64, targetLength+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int64, wordLength+1)
	}
	dp[0][0] = 1

	for length := 0; length <= targetLength; length++ {
		for col := 0; col < wordLength; col++ {
			if length < targetLength {
				dp[length+1][col+1] += int64(charOccurences[target[length]-'a'][col]) * dp[length][col]
				dp[length+1][col+1] %= int64(MOD)
			}
			dp[length][col+1] += dp[length][col]
			dp[length][col+1] %= int64(MOD)
		}
	}

	return int(dp[targetLength][wordLength])
}
