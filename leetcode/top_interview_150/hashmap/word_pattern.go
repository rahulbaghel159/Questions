package hashmap

import "strings"

// https://leetcode.com/problems/word-pattern
func wordPattern(pattern string, s string) bool {
	str_arr := stringToArray(s)

	if len(pattern) != len(str_arr) {
		return false
	}

	dict_s, dict_p := make(map[string]rune), make(map[rune]string)

	for i := 0; i < len(pattern); i++ {
		mapping1, exist1 := dict_p[rune(pattern[i])]
		mapping2, exist2 := dict_s[str_arr[i]]

		if !exist1 && !exist2 {
			dict_p[rune(pattern[i])] = str_arr[i]
			dict_s[str_arr[i]] = rune(pattern[i])
		} else if mapping1 != str_arr[i] && mapping2 != rune(pattern[i]) {
			return false
		}
	}

	return true
}

func stringToArray(s string) []string {
	return strings.Split(s, " ")
}
