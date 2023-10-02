package arrays

//https://leetcode.com/problems/majority-element

func majorityElement(nums []int) int {
	count, candidate := 0, 0

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}

		if candidate == v {
			count++
		} else {
			count--
		}
	}

	return candidate
}
