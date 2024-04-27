package equaltreepartition

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var seen stack

func checkEqualTree(root *TreeNode) bool {
	seen = stack{
		arr: make([]int, 0),
	}

	total := sum(root)

	seen.pop()

	if total%2 == 0 {
		for !seen.isEmpty() {
			if seen.pop() == total/2 {
				return true
			}
		}
	}

	return false
}

func sum(node *TreeNode) int {
	if node == nil {
		return 0
	}

	seen.push(sum(node.Left) + sum(node.Right) + node.Val)

	return seen.peek()
}

type stack struct {
	arr []int
}

func (s *stack) push(r int) {
	s.arr = append(s.arr, r)
}

func (s *stack) pop() int {
	if len(s.arr) == 0 {
		return -1
	}

	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]

	return r
}

func (s *stack) peek() int {
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
