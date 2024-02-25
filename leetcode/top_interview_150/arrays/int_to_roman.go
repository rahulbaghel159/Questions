package arrays

// https://leetcode.com/problems/integer-to-roman/
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for num > 0 {
		for i, v := range values {
			if v <= num {
				times := num / v
				num = num - (v * times)
				for times > 0 {
					res += symbol[i]
					times--
				}
			}
		}
	}

	return res
}
