package arrays

//https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	ans, startFound := 0, false

	for i := len(s) - 1; i >= 0; i-- {
		if !startFound {
			if !isWhiteSpace(rune(s[i])) {
				ans++
				startFound = true
			}
		} else {
			if isWhiteSpace(rune(s[i])) {
				break
			}
			ans++
		}
	}

	return ans
}

func isWhiteSpace(r rune) bool {
	return r == ' '
}
