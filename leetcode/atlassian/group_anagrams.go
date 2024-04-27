package atlassian

import "strconv"

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	index, res := make(map[string][]string), make([][]string, 0)
	for _, s := range strs {
		count := make([]int, 26)
		for _, r := range s {
			count[r-'a']++
		}
		key := ""
		for _, v := range count {
			key += strconv.Itoa(v) + "#"
		}

		if _, ok := index[key]; ok {
			index[key] = append(index[key], s)
		} else {
			index[key] = []string{s}
		}
	}

	for _, v := range index {
		res = append(res, v)
	}

	return res
}
