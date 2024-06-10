package facebook

// https://leetcode.com/explore/interview/card/facebook/5/array-and-strings/3008/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	freq, maxLen := make(map[rune]int), 0

	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := freq[rune(s[j])]; ok {
			for freq[rune(s[j])] > 0 && i <= j {
				freq[rune(s[i])]--
				i++
			}
		}
		freq[rune(s[j])]++
		maxLen = max(maxLen, j-i+1)
	}

	return maxLen
}
