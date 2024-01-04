package slidingwindow

//https://leetcode.com/problems/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right, maxLen, char_map, tempLen := 0, 0, 0, make(map[rune]struct{}), 0

	for left <= right && right < len(s) {
		if _, ok := char_map[rune(s[right])]; !ok {
			char_map[rune(s[right])] = struct{}{}
			tempLen++
			right++
		} else {
			for {
				delete(char_map, rune(s[left]))
				left++
				tempLen--
				if _, ok := char_map[rune(s[right])]; !ok {
					char_map[rune(s[right])] = struct{}{}
					right++
					tempLen++
					break
				}
			}
		}
		maxLen = max(maxLen, tempLen)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
