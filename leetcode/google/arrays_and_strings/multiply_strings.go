package google

import "fmt"

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/3051/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	if num1 == "1" {
		return num2
	}

	if num2 == "1" {
		return num1
	}

	strings, j := make([]string, len(num2)), 0
	for i := len(num2) - 1; i >= 0; i-- {
		strings[j] = multiplyBy(num1, rune(num2[i]), j)
		j++
	}

	maxLen := len(strings[0])
	for _, s := range strings {
		maxLen = max(maxLen, len(s))
	}

	for i := 0; i < len(strings); i++ {
		for j := maxLen - len(strings[i]); j > 0; j-- {
			strings[i] = "0" + strings[i]
		}
	}

	res, carry := "", 0
	for i := maxLen - 1; i >= 0; i-- {
		tmp := 0
		for j := 0; j < len(strings); j++ {
			tmp = tmp + int([]byte(strings[j])[i]-'0')
		}
		tmp += carry
		if tmp > 9 {
			carry = tmp / 10
			tmp = tmp % 10
		} else {
			carry = 0
		}
		res = fmt.Sprint(tmp) + res
	}

	if carry > 0 {
		res = fmt.Sprint(carry) + res
	}

	return res
}

func multiplyBy(num string, n rune, i int) string {
	res, arr, carry := "", []byte(num), 0

	for i := len(arr) - 1; i >= 0; i-- {
		r := (int(n-'0') * int(arr[i]-'0')) + carry
		if r > 9 {
			carry = r / 10
			r = r % 10
		} else {
			carry = 0
		}
		res = fmt.Sprint(r) + res
	}

	if carry > 0 {
		res = fmt.Sprint(carry) + res
	}

	for i > 0 {
		res += "0"
		i--
	}

	return res
}
