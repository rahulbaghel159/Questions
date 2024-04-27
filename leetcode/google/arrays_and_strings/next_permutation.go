package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3050/
func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}

	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i+1] > nums[i] {
			break
		}
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	for l, r := i+1, len(nums)-1; l < r; {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
