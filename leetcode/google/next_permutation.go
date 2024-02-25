package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3050/
func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

}

func backtrack(nums []int, start, prev int) {
	if start == len(nums)-2 && nums[start] < nums[start+1] {
		nums[start], nums[start+1] = nums[start+1], nums[start]
	}

}
