package facebook

import "sort"

// https://leetcode.com/explore/interview/card/facebook/5/array-and-strings/283/
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	if len(nums) == 3 && nums[0]+nums[1]+nums[2] == 0 {
		return [][]int{{nums[0], nums[1], nums[2]}}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			start, end := i+1, len(nums)-1
			for start < end {
				sum := nums[i] + nums[start] + nums[end]
				if sum > 0 {
					end--
				} else if sum < 0 {
					start++
				} else {
					res = append(res, []int{nums[start], nums[i], nums[end]})
					start++
					end--
					for start < end && nums[start] == nums[start-1] {
						start++
					}
				}
			}
		}
	}

	return res
}
