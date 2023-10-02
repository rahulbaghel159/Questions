package arrays

//https://leetcode.com/problems/longest-common-prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	index := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][0:index]
			}
		}
		index++
	}

	return strs[0][0:index]
}
