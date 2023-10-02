package arrays

//https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func strStr(haystack string, needle string) int {
	ans, needleIndex, flag := -1, 0, false

	for i := 0; i < len(haystack); i++ {
		if flag {
			if haystack[i] != needle[needleIndex] {
				flag = false
				needleIndex = 0
			} else {
				needleIndex++
			}
		} else {
			if haystack[i] == needle[needleIndex] {
				flag = true
				ans = i
				needleIndex++
			}
		}

		if needleIndex == len(needle) && flag {
			return ans
		}
	}

	return -1
}
