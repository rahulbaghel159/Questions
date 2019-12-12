package longestsubstring

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	maxLength := 0
	characters := make(map[rune]bool)
	for _, ch := range s {
		if _, ok := characters[ch]; ok {
			for _, c := range s[start:end] {
				if _, okay := characters[ch]; !okay {
					break
				}
				delete(characters, c)
				start++
			}
		}
		if maxLength < (end - start + 1) {
			maxLength = (end - start + 1)
		}
		characters[ch] = true
		end++
	}
	return maxLength
}
