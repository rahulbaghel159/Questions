package others

// https://leetcode.com/explore/interview/card/google/66/others-4/3098/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict, reverse_dict := make(map[rune]rune, 0), make(map[rune]rune)
	for i := 0; i < len(s); i++ {
		m12, ok12 := dict[rune(s[i])]
		m21, ok21 := reverse_dict[rune(t[i])]

		if !ok12 && !ok21 {
			dict[rune(s[i])] = rune(t[i])
			reverse_dict[rune(t[i])] = rune(s[i])
		} else if ok12 && ok21 {
			if m12 != rune(t[i]) || m21 != rune(s[i]) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
