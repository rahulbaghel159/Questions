package process

import (
	"strings"
)

// https://leetcode.com/explore/interview/card/google/67/sql-2/472/
// func licenseKeyFormatting(s string, k int) string {
// 	if len(s) == 0 {
// 		return ""
// 	}

// 	arr, temp := make([]string, 0), make([]byte, 0)

// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] != '-' {
// 			if len(temp) == k {
// 				arr = append([]string{strings.ToUpper(string(temp))}, arr...)
// 				temp = make([]byte, 0)
// 			}
// 			temp = append([]byte{s[i]}, temp...)
// 		}
// 	}

// 	if len(temp) > 0 {
// 		arr = append([]string{strings.ToUpper(string(temp))}, arr...)
// 	}

// 	return strings.Join(arr, "-")
// }

func licenseKeyFormatting(s string, k int) string {
	if len(s) == 0 {
		return ""
	}

	ans, curr := "", 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			if curr == k {
				ans = "-" + ans
				curr = 0
			}

			ans = strings.ToUpper(string(s[i])) + ans
			curr++
		}
	}

	return ans
}
