package pointers

import "strings"

// https://leetcode.com/problems/valid-palindrome
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)
	left, right := 0, len(s)-1

	for left < right {
		if !isValid(rune(s[left])) {
			left++
			continue
		}
		if !isValid(rune(s[right])) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isValid(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= '0' && r <= '9'
}
