package google

//https://leetcode.com/explore/interview/card/google/59/array-and-strings/3047/

//LengthOfLongestSubstring finds maximum length of substring with non-repeated characters
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		start, end, max int
		chars           []int
	)

	chars = make([]int, 128)

	for end < len(s) {
		chars[rune(s[end])]++

		for chars[rune(s[end])] > 1 {
			chars[rune(s[start])]--
			start++
		}
		end++

		if end-start > max {
			max = end - start
		}
	}

	return max
}
