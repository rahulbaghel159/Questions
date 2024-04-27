package wonderfulsubstring

// https://leetcode.com/problems/number-of-wonderful-substrings/
func wonderfulSubstrings(word string) int64 {
	if len(word) == 0 {
		return int64(0)
	}

	count := int64(0)
	for i := 0; i < len(word); i++ {
		freq := make(map[rune]int)
		for j := i; j < len(word); j++ {
			freq[rune(word[j])]++

			sub_str := word[i:]
			if j < len(word)-1 {
				sub_str = word[i : j+1]
			}

			if isWonderful(sub_str, freq) {
				count++
			}
		}
	}

	return count
}

func isWonderful(s string, freq map[rune]int) bool {
	odd_found := false
	for _, v := range freq {
		if v%2 != 0 {
			if !odd_found {
				odd_found = true
			} else {
				return false
			}
		}
	}

	return true
}
