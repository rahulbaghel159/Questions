package arrays

//https://leetcode.com/problems/product-of-array-except-self

func productExceptSelf(nums []int) []int {
	suffix, output := make([]int, len(nums)), make([]int, len(nums))

	product := 1
	for i := len(nums) - 1; i >= 0; i-- {
		product = product * nums[i]
		suffix[i] = product
	}

	product = 1
	for i := 0; i < len(nums); i++ {
		s := 1

		if i < len(suffix)-1 {
			s = suffix[i+1]
		}

		output[i] = product * s

		product = product * nums[i]
	}

	return output
}
