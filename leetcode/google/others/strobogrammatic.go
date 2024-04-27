package others

// https://leetcode.com/explore/interview/card/google/66/others-4/3099/
func isStrobogrammatic(num string) bool {
	if len(num) == 0 {
		return true
	}

	if len(num) > 1 && num[len(num)-1] == '0' {
		return false
	}

	master := map[rune]string{
		'0': "0",
		'1': "1",
		'6': "9",
		'8': "8",
		'9': "6",
	}

	generated_num := ""
	for i := len(num) - 1; i >= 0; i-- {
		digit := num[i]
		if _, ok := master[rune(digit)]; !ok {
			return false
		}

		generated_num += master[rune(digit)]
	}

	return generated_num == num
}
