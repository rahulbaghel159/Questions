package atlassian

// https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	if len(s) == 0 {
		return 0
	}

	count, unique, sub := make([]byte, 26), 0, make(map[string]int)
	for i := 0; i < len(s); i++ {
		if i >= minSize {
			ch := s[i-minSize] - 'a'
			count[ch]--
			if count[ch] == 0 {
				unique--
			}
		}

		ch := s[i] - 'a'
		count[ch]++
		if count[ch] == 1 {
			unique++
		}

		if i+1 >= minSize && unique <= maxLetters {
			sub[string(s[i+1-minSize:i+1])]++
		}
	}

	ans := 0
	for _, v := range sub {
		ans = max(ans, v)
	}

	return ans
}
