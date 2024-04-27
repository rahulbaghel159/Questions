package adobe

// https://leetcode.com/problems/concatenation-of-array/description/
func getConcatenation(nums []int) []int {
	res := make([]int, len(nums)*2)

	for i := 0; i < len(nums); i++ {
		res[i], res[i+len(nums)] = nums[i], nums[i]
	}

	return res
}
