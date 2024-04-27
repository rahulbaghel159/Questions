package adobe

import "sort"

// https://leetcode.com/problems/smallest-range-ii/
func smallestRangeII(nums []int, k int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := nums[len(nums)-1] - nums[0]
	for i := 0; i < len(nums)-1; i++ {
		a, b := nums[i], nums[i+1]
		high := max(nums[len(nums)-1]-k, a+k)
		low := min(nums[0]+k, b-k)

		ans = min(ans, high-low)
	}

	return ans
}
