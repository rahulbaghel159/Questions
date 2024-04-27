package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3055/
func findMissingRanges(nums []int, lower int, upper int) [][]int {
	if len(nums) == 0 {
		return [][]int{{lower, upper}}
	}

	res := make([][]int, 0)
	if nums[0] != lower {
		res = append(res, []int{lower, nums[0] - 1})
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			res = append(res, []int{nums[i-1] + 1, nums[i] - 1})
		}
	}

	if nums[len(nums)-1] != upper {
		res = append(res, []int{nums[len(nums)-1] + 1, upper})
	}

	return res
}
