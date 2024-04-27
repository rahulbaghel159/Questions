package google

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/467/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := stack{
		arr: make([]rune, 0),
	}

	for _, r := range s {
		if r == '(' || r == '{' || r == '[' {
			stack.push(r)
		} else {
			p := stack.pop()
			if p == '0' || (r == ')' && p != '(') || (r == '}' && p != '{') || (r == ']' && p != '[') {
				return false
			}
		}
	}

	return stack.isEmpty()
}

type stack struct {
	arr []rune
}

func (s *stack) push(r rune) {
	s.arr = append(s.arr, r)
}

func (s *stack) pop() rune {
	if len(s.arr) == 0 {
		return 0
	}

	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return r
}

func (s *stack) peek() rune {
	if len(s.arr) == 0 {
		return 0
	}

	return s.arr[len(s.arr)-1]
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack) size() int {
	return len(s.arr)
}
