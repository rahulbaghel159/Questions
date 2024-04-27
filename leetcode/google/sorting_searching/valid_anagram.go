package sortingsearching

// https://leetcode.com/explore/interview/card/google/63/sorting-and-searching-4/3082/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]int, 26)
	for _, r := range s {
		count[r-'a']++
	}

	for _, r := range t {
		count[r-'a']--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
