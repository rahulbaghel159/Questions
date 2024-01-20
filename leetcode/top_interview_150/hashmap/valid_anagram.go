package hashmap

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict_s := make(map[rune]int)
	for _, r := range s {
		dict_s[r]++
	}

	for _, r := range t {
		if count, exists := dict_s[r]; !exists || count == 0 {
			return false
		}
		dict_s[r]--
	}

	return true
}
