package arrays

//https://leetcode.com/problems/reverse-words-in-a-string

func reverseWords(s string) string {
	ans, str, found := "", "", false
	for _, r := range s {
		if !found {
			if !isWhiteSpace(r) {
				found = true
				str += string(r)
			}
		} else {
			if isWhiteSpace(r) {
				found = false
				ans = str + " " + ans
				str = ""
			} else {
				str += string(r)
			}
		}
	}

	if len(str) != 0 {
		ans = str + " " + ans
	}

	if isWhiteSpace(rune(ans[0])) {
		ans = ans[1:]
	}

	if isWhiteSpace(rune(ans[len(ans)-1])) {
		ans = ans[:len(ans)-1]
	}

	return ans
}
