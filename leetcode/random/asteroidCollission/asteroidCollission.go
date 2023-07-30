package asteroidcollission

//https://leetcode.com/problems/asteroid-collision

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) == 0 {
		return []int{}
	}

	s := stack{arr: []int{}}
	for _, v := range asteroids {
		flag := true

		for !s.isEmpty() && (s.peek() > 0 && v < 0) {
			if abs(v) > abs(s.peek()) {
				s.pop()
				continue
			} else if abs(v) == abs(s.peek()) {
				s.pop()
			}

			flag = false
			break
		}

		if flag {
			s.push(v)
		}
	}

	result, index := make([]int, s.size()), s.size()-1
	for !s.isEmpty() {
		result[index] = s.pop()
		index--
	}

	return result
}

type stack struct {
	arr []int
}

func (s *stack) pop() int {
	if len(s.arr) == 0 {
		return 0
	}

	a := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return a
}

func (s *stack) peek() int {
	if len(s.arr) == 0 {
		return 0
	}
	return s.arr[len(s.arr)-1]
}

func (s *stack) push(a int) {
	s.arr = append(s.arr, a)
}

func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func (s *stack) size() int {
	return len(s.arr)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
