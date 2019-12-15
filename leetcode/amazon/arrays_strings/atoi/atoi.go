package atoi

//https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/2962/
const (
	maxInt = 2147483647
	minInt = -2147483648
)

func atoi(str string) int {
	//handle blank string
	if len(str) == 0 {
		return 0
	}

	pos := 0
	sign := 1
	var num, result int32
	//ignore whitespace characters until non-whitespace is not found
	for p, ch := range str {
		if !isWhiteSpace(ch) {
			pos = p
			break
		}
	}

	//check for sign at start position
	if str[pos] == '-' {
		sign = -1
		pos++
	} else if str[pos] == '+' {
		sign = 1
		pos++
	}

	//start conversion
	for pos < len(str) && isValidNum(rune(str[pos])) {
		//handle overflow
		if num > maxInt/10 || (num == maxInt/10 && str[pos]-'0' > 7) {
			if sign < 0 {
				return minInt
			}
			return maxInt
		}
		num = num*10 + int32(str[pos]-'0')
		pos++
	}

	result = int32(sign) * num

	return int(result)
}

func isWhiteSpace(ch rune) bool {
	return (ch == ' ')
}

func isValidNum(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}
