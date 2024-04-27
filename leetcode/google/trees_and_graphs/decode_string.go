package treesandgraphs

import "strconv"

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3073/
func decodeString(s string) string {
	if len(s) == 0 {
		return ""
	}

	stack := stack{
		arr: make([]string, 0),
	}

	for _, r := range s {
		stack.push(string(r))

		if stack.peek() == "]" {
			stack.pop() //pop ending square bracket

			var arr []string
			for stack.peek() != "[" {
				arr = append(arr, stack.pop()) //find string
			}

			stack.pop() //pop opening square bracket

			k, power := 0, 0
			for isDigit(stack.peek()) {
				n, _ := strconv.Atoi(stack.pop())
				k = n*pow(10, power) + k
				power++
			}

			str, new_str := "", ""

			for i := len(arr) - 1; i >= 0; i-- {
				str += arr[i]
			}

			for i := 0; i < k; i++ {
				new_str += str
			}

			stack.push(new_str)
		}
	}

	res := ""
	for stack.size() > 0 {
		res = stack.pop() + res
	}

	return res
}

func isDigit(r string) bool {
	_, err := strconv.Atoi(r)
	return err == nil
}

type stack struct {
	arr []string
}

func (s *stack) push(r string) {
	s.arr = append(s.arr, r)
}

func (s *stack) pop() string {
	if len(s.arr) == 0 {
		return ""
	}

	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return r
}

func (s *stack) peek() string {
	if len(s.arr) == 0 {
		return ""
	}

	return s.arr[len(s.arr)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack) size() int {
	return len(s.arr)
}
