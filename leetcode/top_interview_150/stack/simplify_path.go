package stack

import "strings"

// https://leetcode.com/problems/simplify-path
func simplifyPath(path string) string {
	if len(path) == 0 {
		return "/"
	}

	s := str_stack{arr: make([]string, 0)}
	for _, d := range strings.Split(path, "/") {
		if d == "." || d == "" {
			continue
		} else if d == ".." {
			if !s.isEmpty() {
				s.pop()
			}
		} else {
			s.push(d)
		}
	}

	res := ""
	for _, d := range s.arr {
		res += "/" + d
	}

	if len(res) == 0 {
		res = "/"
	}

	return res
}

type str_stack struct {
	arr []string
}

func (s *str_stack) push(str string) {
	s.arr = append(s.arr, str)
}

func (s *str_stack) pop() string {
	if len(s.arr) == 0 {
		return ""
	}

	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return r
}

func (s *str_stack) peek() string {
	if len(s.arr) == 0 {
		return ""
	}

	return s.arr[len(s.arr)-1]
}

func (s *str_stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *str_stack) size() int {
	return len(s.arr)
}
