package math

// https://leetcode.com/problems/palindrome-number
// func isPalindrome(x int) bool {
// 	s := fmt.Sprint(x)
// 	left, right := 0, len(s)-1
// 	for left < right {
// 		if s[left] != s[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}

// 	return true
// }

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	return x == reverted || x == reverted/10
}
