package google

//https://leetcode.com/problems/longest-palindromic-substring/solution/

func longestPalindrome(s string) string {
	maxLength := -1
	i, j := 0, 0
	result := ""

	for i = 0; i < len(s); i++ {
		for j = i; j < len(s); j++ {
			if isPalindrome([]rune(s), i, j) {
				if j-i > maxLength {
					maxLength = j - i
					result = string([]rune(s)[i : j+1])
				}
			}
		}
	}

	return result
}

func isPalindrome(s []rune, i, j int) bool {
	if i == j {
		return true
	}

	if i >= 0 && i < len(s) && j >= 0 && j < len(s) && s[i] == s[j] {
		return j-i == 1 || isPalindrome([]rune(s), i+1, j-1)
	}

	return false
}
