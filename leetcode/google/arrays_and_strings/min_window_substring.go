package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/345/
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	if len(t) > len(s) {
		return ""
	}

	dict_t := make(map[rune]int)
	for _, r := range t {
		dict_t[r] = dict_t[r] + 1
	}

	required, l, r, formed, dict_window, ans := len(dict_t), 0, 0, 0, make(map[rune]int), []int{-1, 0, 0}

	for r < len(s) {
		dict_window[rune(s[r])] = dict_window[rune(s[r])] + 1
		if v, ok := dict_t[rune(s[r])]; ok && v == dict_window[rune(s[r])] {
			formed++
		}

		for l <= r && formed == required {
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}

			dict_window[rune(s[l])]--
			if v, ok := dict_t[rune(s[l])]; ok && dict_window[rune(s[l])] < v {
				formed--
			}
			l++
		}
		r++
	}

	if ans[0] == -1 {
		return ""
	}

	return s[ans[1] : ans[2]+1]
}
