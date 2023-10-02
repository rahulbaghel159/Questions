package arrays

//https://leetcode.com/problems/h-index

func hIndex(citations []int) int {
	n := len(citations)
	papers := make([]int, n+1)

	for _, c := range citations {
		papers[min(n, c)]++
	}

	k := n
	for s := papers[n]; k > s; s += papers[k] {
		k--
	}

	return k
}
