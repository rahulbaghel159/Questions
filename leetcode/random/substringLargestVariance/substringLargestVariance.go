package substringlargestvariance

//https://leetcode.com/problems/substring-with-largest-variance/

func largestVariance(s string) int {
	counter := make([]int, 26)

	for _, r := range s {
		counter[r-'a']++
	}

	global_max := 0
	for i := 0; i < len(counter); i++ {
		for j := 0; j < len(counter); j++ {
			if i == j || counter[i] == 0 || counter[j] == 0 {
				continue
			}
			major, minor := i, j

			major_count, minor_count, rest_minor := 0, 0, counter[j]

			for i := 0; i < len(s); i++ {
				if int(s[i]-'a') == major {
					major_count++
				}
				if int(s[i]-'a') == minor {
					minor_count++
					rest_minor--
				}
				if minor_count > 0 {
					global_max = max(global_max, major_count-minor_count)
				}
				if major_count < minor_count && rest_minor > 0 {
					major_count, minor_count = 0, 0
				}
			}
		}
	}

	return global_max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
