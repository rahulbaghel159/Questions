package hashmap

// https://leetcode.com/problems/isomorphic-strings
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict_s, dict_t := make(map[rune]rune), make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		mapping1, exist1 := dict_s[rune(s[i])]
		mapping2, exist2 := dict_t[rune(t[i])]

		if !exist1 && !exist2 {
			dict_s[rune(s[i])] = rune(t[i])
			dict_t[rune(t[i])] = rune(s[i])
		} else if mapping1 != rune(t[i]) && mapping2 != rune(s[i]) {
			return false
		}
	}

	return true
}
