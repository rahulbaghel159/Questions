package stack

import (
	"math"
	"reflect"
	"strconv"
)

// https://leetcode.com/problems/basic-calculator
func evaluateExpr(st str_stack) (int, str_stack) {
	if st.isEmpty() || !(reflect.TypeOf(st.peek()).Name() == "int") {
		_, err := strconv.Atoi(st.peek())
		if err != nil {
			st.push("0")
		}
	}

	res, _ := strconv.Atoi(st.pop())

	for !st.isEmpty() && st.peek() != ")" {
		sign := st.pop()
		val, _ := strconv.Atoi(st.pop())

		if sign == "+" {
			res += val
		} else {
			res -= val
		}
	}

	return res, st
}

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	operand, n, st := 0, 0, str_stack{arr: make([]string, 0)}

	for i := len(s) - 1; i >= 0; i-- {
		if isDigit(rune(s[i])) {
			operand = int(math.Pow10(n))*int(s[i]-'0') + operand
			n++
		} else if s[i] != ' ' {
			if n != 0 {
				st.push(strconv.Itoa(operand))
				n, operand = 0, 0
			}
			if s[i] == '(' {
				var res int
				res, st = evaluateExpr(st)
				st.pop()
				st.push(strconv.Itoa(res))
			} else {
				st.push(string(s[i]))
			}
		}
	}

	if n != 0 {
		st.push(strconv.Itoa(operand))
	}

	res, _ := evaluateExpr(st)

	return res
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
