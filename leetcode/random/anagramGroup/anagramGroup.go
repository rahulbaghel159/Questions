package groupanagrams

func groupAnagrams(strs []string) [][]string {
	var (
		result    [][]string
		resultMap map[string][]string
		found     bool
	)

	resultMap = make(map[string][]string)

	for _, s := range strs {
		found = false
		for k := range resultMap {
			if isAnagram(k, s) {
				resultMap[k] = append(resultMap[k], s)
				found = true
				break
			}
		}
		if !found {
			resultMap[s] = []string{s}
		}
	}

	for _, v := range resultMap {
		result = append(result, v)
	}

	return result
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var (
		charCount map[rune]int
	)

	charCount = make(map[rune]int)

	for _, r := range s1 {
		if _, ok := charCount[r]; ok {
			charCount[r]++
		} else {
			charCount[r] = 1
		}
	}

	for _, r := range s2 {
		if _, ok := charCount[r]; ok {
			charCount[r]--
		} else {
			return false
		}
	}

	for _, v := range charCount {
		if v != 0 {
			return false
		}
	}

	return true
}
