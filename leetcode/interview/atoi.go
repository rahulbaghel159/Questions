package interview

const (
	intMin = -2147483648
	intMax = 2147483647
)

func myAtoi(s string) int {
	//empty string
	if len(s) == 0 {
		return 0
	}

	var (
		sign   byte
		result int
		index  int
	)

	//discard whitespace characters
	for i, ch := range s {
		if !isWhiteSpace(ch) {
			index = i
			break
		}
	}

	//consider optional plus or minus sign
	if s[index] == '+' || s[index] == '-' {
		sign = s[index]
		index++
	} else {
		//first non-white space character non numeric or string consist of only whitespace characters
		if !isNumeric(rune(s[index])) {
			return 0
		}
	}

	for _, ch := range s[index:] {
		if isNumeric(ch) {
			//check overflow/underflow condition
			if (result == intMax/10 && ch-'0' > 7) || (result > intMax/10) {
				if sign == '-' {
					return intMin
				}
				return intMax
			}
			result = result*10 + int(ch-'0')
		} else {
			break
		}
	}

	if sign == '-' {
		result = -1 * result
	}

	return int(result)
}

func isWhiteSpace(ch rune) bool {
	if ch == ' ' {
		return true
	}
	return false
}

func isNumeric(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
