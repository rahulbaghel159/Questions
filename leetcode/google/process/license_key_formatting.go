package process

import "strings"

// https://leetcode.com/explore/interview/card/google/67/sql-2/472/
func licenseKeyFormatting(s string, k int) string {
	if len(s) == 0 {
		return ""
	}

	arr, temp, j := make([]string, 0), make([]byte, k), k-1

	for i := len(s) - 1; i >= 0; i-- {
		if j < 0 {
			temp, j = make([]byte, k), k-1
			arr = append(arr, strings.ToUpper(string(temp)))
		}
		temp[j] = s[i]
		j--
	}

	return strings.Join(arr, "-")
}
