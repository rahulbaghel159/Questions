package adobe

// https://leetcode.com/problems/smallest-range-i/
func smallestRangeI(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}

	max_num, min_num := nums[0], nums[0]
	for _, num := range nums {
		max_num = max(max_num, num)
		min_num = min(min_num, num)
	}

	res := max_num - min_num - 2*k

	if res < 0 {
		return 0
	}

	return res
}
