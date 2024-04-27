package removeduplicates

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {
	if len(s) == 0 {
		return ""
	}

	st := stack{
		arr: make([]rune, 0),
	}

	for _, r := range s {
		if st.peek() == r {
			for st.peek() == r {
				st.pop()
			}
		} else {
			st.push(r)
		}
	}

	ans, index := make([]rune, st.size()), st.size()-1
	for !st.isEmpty() {
		ans[index] = st.pop()
		index--
	}

	return string(ans)
}

type stack struct {
	arr []rune
}

func (s *stack) push(r rune) {
	s.arr = append(s.arr, r)
}

func (s *stack) pop() rune {
	if len(s.arr) == 0 {
		return -1
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
