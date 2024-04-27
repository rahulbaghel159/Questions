package google

import "sort"

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3049/
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	ans := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			start, end := i+1, len(nums)-1
			for start < end {
				sum := nums[start] + nums[end] + nums[i]
				if sum < 0 {
					start++
				} else if sum > 0 {
					end--
				} else {
					ans = append(ans, []int{nums[i], nums[start], nums[end]})
					start++
					end--
					for start < end && nums[start] == nums[start-1] {
						start++
					}
				}
			}
		}
	}

	return ans
}
