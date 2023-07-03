package fairDistributionCookies

//https://leetcode.com/problems/fair-distribution-of-cookies/

func distributeCookies(cookies []int, k int) int {
	if k == 0 {
		return 0
	}

	if k == 1 {
		sum := 0
		for _, val := range cookies {
			sum += val
		}
		return sum
	}

	return dfs(0, k, cookies, make([]int, k))
}

func dfs(i, zero_count int, cookies, distribute []int) int {
	if len(cookies)-i < zero_count {
		return 999999999999999999
	}

	if i == len(cookies) {
		max := distribute[0]
		for _, val := range distribute {
			if max < val {
				max = val
			}
		}
		return max
	}

	answer := 999999999999999999

	for j := 0; j < len(distribute); j++ {
		if distribute[j] == 0 {
			zero_count--
		}
		distribute[j] += cookies[i]

		answer = min(answer, dfs(i+1, zero_count, cookies, distribute))

		distribute[j] -= cookies[i]
		if distribute[j] == 0 {
			zero_count++
		}
	}

	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
