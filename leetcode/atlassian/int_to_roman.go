package atlassian

// https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	rep := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for i := 0; i < len(nums) && num > 0; i++ {
		if num >= nums[i] {
			times := num / nums[i]
			for j := 0; j < times; j++ {
				res += rep[i]
			}
			num -= (nums[i] * times)
		}
	}

	return res
}
