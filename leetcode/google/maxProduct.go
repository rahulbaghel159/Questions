package google

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max_so_far, min_so_far := nums[0], nums[0]
	result := max_so_far

	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		temp_max := max(curr, max(curr*max_so_far, curr*min_so_far))
		min_so_far = min(curr, min(curr, min(curr*min_so_far, curr*max_so_far)))

		max_so_far = temp_max

		result = max(result, max_so_far)
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
