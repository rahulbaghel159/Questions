package hashmap

import "sort"

// https://leetcode.com/problems/group-anagrams
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	res, dict := [][]string{}, make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		r := []rune(strs[i])
		sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
		if _, exists := dict[string(r)]; exists {
			dict[string(r)] = append(dict[string(r)], strs[i])
		} else {
			dict[string(r)] = []string{strs[i]}
		}
	}

	for _, v := range dict {
		res = append(res, v)
	}

	return res
}
