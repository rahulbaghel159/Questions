package stack

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	st := str_stack{arr: make([]string, 0)}
	for _, s := range tokens {
		if s != "+" && s != "-" && s != "*" && s != "/" {
			st.push(s)
		} else {
			op1, _ := strconv.Atoi(st.pop())
			op2, _ := strconv.Atoi(st.pop())
			var temp int
			switch s {
			case "+":
				temp = op2 + op1
			case "-":
				temp = op2 - op1
			case "*":
				temp = op2 * op1
			case "/":
				temp = op2 / op1
			}
			st.push(strconv.Itoa(temp))
		}
	}

	res, _ := strconv.Atoi(st.pop())

	return res
}
