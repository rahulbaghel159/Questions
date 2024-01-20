package jan2024

import "sort"

func missingInteger(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	seqential_prefix_sum, prev := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if prev+1 == nums[i] {
			seqential_prefix_sum += nums[i]
			prev = nums[i]
		} else {
			break
		}
	}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] == seqential_prefix_sum {
			seqential_prefix_sum++
		}
	}

	return seqential_prefix_sum
}
