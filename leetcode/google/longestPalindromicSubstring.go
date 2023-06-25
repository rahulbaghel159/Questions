package google

//https://leetcode.com/problems/longest-palindromic-substring/solution/

func longestPalindrome(s string) string {
	start, end := 0, 0
	result := ""

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i > (end-start) && isPalindrome(s, i, j) {
				start, end = i, j
			}
		}
	}

	result = string([]rune(s)[start : end+1])

	if result == "" {
		result = string([]rune(s[0:1]))
	}

	return result
}

func isPalindrome(s string, start, end int) bool {
	mid := (start + end) / 2

	for i, j := start, end; i <= mid && j >= mid; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
