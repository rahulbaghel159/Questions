package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3047/
// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}

// 	index, maxLen, currLen, str := make(map[rune]int, 0), 0, 0, ""
// 	for _, r := range s {
// 		if v, ok := index[r]; ok && v >= 1 {
// 			for i := 0; i < len(str); i++ {
// 				index[rune(str[i])]--
// 				currLen--
// 				if index[r] < 1 {
// 					str = string(str[i+1:])
// 					break
// 				}
// 			}
// 		}
// 		str += string(r)
// 		index[r] = 1
// 		currLen++
// 		maxLen = max(maxLen, currLen)
// 	}

// 	return maxLen
// }

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	index, start, maxLen := make(map[rune]int, 0), 0, 0
	for i := 0; i < len(s); i++ {
		if v, ok := index[rune(s[i])]; ok {
			if v >= start {
				start = v + 1
			}
		}
		index[rune(s[i])] = i
		maxLen = max(maxLen, i-start+1)
	}

	return maxLen
}
