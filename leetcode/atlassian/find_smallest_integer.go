package atlassian

// https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/
func findSmallestInteger(nums []int, value int) int {
	if len(nums) == 0 || len(nums) == 1 && nums[0] != 0 {
		return 0
	}

	count := make([]int, value)

	for _, num := range nums {
		count[((num%value)+value)%value]++
	}

	for i := 0; ; i++ {
		count[i%value]--
		if count[i%value] < 0 {
			return i
		}
	}
}
