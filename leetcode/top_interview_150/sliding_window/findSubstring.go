package slidingwindow

//https://leetcode.com/problems/substring-with-concatenation-of-all-words

func findSubstring(s string, words []string) []int {
	dict := make(map[string]int)
	for _, w := range words {
		dict[w] += 1
	}

	wordLen, wordsCount := len(words[0]), len(words)
	substringSize := wordsCount * wordLen
	result := []int{}

	sliding_window := func(left int) {
		wordsFound, wordsUsed, excessWord := make(map[string]int), 0, false

		for right := left; right+wordLen <= len(s); right += wordLen {
			sub := s[right : right+len(words[0])]
			if v, ok := dict[sub]; !ok {
				wordsFound, wordsUsed, excessWord, left = make(map[string]int), 0, false, right+wordLen
			} else {
				for right-left == substringSize || excessWord {
					leftmostWord := s[left : left+wordLen]
					left += wordLen
					wordsFound[leftmostWord]--
					if wordsFound[leftmostWord] >= dict[leftmostWord] {
						excessWord = false
					} else {
						wordsUsed--
					}
				}

				wordsFound[sub]++

				if wordsFound[sub] <= v {
					wordsUsed++
				} else {
					excessWord = true
				}

				if wordsUsed == wordsCount && !excessWord {
					result = append(result, left)
				}
			}
		}
	}

	for i := 0; i < wordLen; i++ {
		sliding_window(i)
	}

	return result
}
