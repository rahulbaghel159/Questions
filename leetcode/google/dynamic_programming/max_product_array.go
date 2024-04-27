package dynamicprogramming

// https://leetcode.com/explore/interview/card/google/64/dynamic-programming-4/3087/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max_so_far, min_so_far, result := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		curr := nums[i]

		temp_max := max(curr, max(curr*max_so_far, curr*min_so_far))
		min_so_far = min(curr, min(max_so_far*curr, min_so_far*curr))

		max_so_far = temp_max

		result = max(result, max_so_far)
	}

	return result
}
