package google

import "sort"

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3057/
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	if len(s) == 0 || len(indices) == 0 {
		return s
	}

	sorted := make([][]int, len(indices))
	for i := 0; i < len(indices); i++ {
		sorted[i] = []int{indices[i], i}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	res := []byte(s)
	for i := len(sorted) - 1; i >= 0; i-- {
		source, target := sources[sorted[i][1]], targets[sorted[i][1]]
		start, end := sorted[i][0], sorted[i][0]+len(source)
		if end <= len(s) && s[start:end] == source {
			res = []byte(string(res[:start]) + target + string(res[end:]))
		}
	}

	return string(res)
}
