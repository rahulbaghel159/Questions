package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3053/
func canJump(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	stepsRemaining := nums[0]
	for i := 1; i < len(nums); i++ {
		stepsRemaining--
		if stepsRemaining < 0 {
			return false
		}
		stepsRemaining = max(stepsRemaining, nums[i])
	}

	return true
}
