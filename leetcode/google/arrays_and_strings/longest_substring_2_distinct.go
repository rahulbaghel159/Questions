package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3054/
func lengthOfLongestSubstringTwoDistinct(s string) int {
	if len(s) == 0 {
		return 0
	}

	start, end, maxLen, dict, distinct_count := 0, 0, 0, make(map[rune]int), 0
	for end < len(s) {
		if _, ok := dict[rune(s[end])]; !ok {
			distinct_count++
		}
		dict[rune(s[end])]++
		if distinct_count <= 2 {
			maxLen = max(maxLen, end-start+1)
		}
		for distinct_count > 2 {
			dict[rune(s[start])]--
			if dict[rune(s[start])] == 0 {
				delete(dict, rune(s[start]))
				distinct_count--
			}
			start++
		}
		end++
	}

	return maxLen
}
