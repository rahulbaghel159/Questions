package treesandgraphs

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3071/
// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	left_count, right_count := 0, 0
// 	if root.Left != nil {
// 		left_count = countNodes(root.Left)
// 	}

// 	if root.Right != nil {
// 		right_count = countNodes(root.Right)
// 	}

// 	return left_count + right_count + 1
// }

// func countNodes(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	queue := NodeQueue{
// 		arr: make([]*TreeNode, 0),
// 	}

// 	queue.Enqueue(root)

// 	total_count := 0
// 	for queue.Size() > 0 {
// 		total_count++
// 		node := queue.Dequeue()
// 		if node.Left != nil {
// 			queue.Enqueue(node.Left)
// 		}
// 		if node.Right != nil {
// 			queue.Enqueue(node.Right)
// 		}
// 	}

// 	return total_count
// }

// type NodeQueue struct {
// 	arr []*TreeNode
// }

// func (q *NodeQueue) Size() int {
// 	return len(q.arr)
// }

// func (q *NodeQueue) Enqueue(x *TreeNode) {
// 	q.arr = append(q.arr, x)
// }

// func (q *NodeQueue) Dequeue() *TreeNode {
// 	first := &TreeNode{}
// 	if q.Size() > 0 {
// 		first = q.arr[0]
// 	}
// 	if q.Size() > 1 {
// 		q.arr = q.arr[1:]
// 	} else {
// 		q.arr = []*TreeNode{}
// 	}

// 	return first
// }

// func (q *NodeQueue) Peek() *TreeNode {
// 	first := &TreeNode{}
// 	if q.Size() > 0 {
// 		first = q.arr[0]
// 	}

// 	return first
// }

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d := depth(root)
	if d == 0 {
		return 1
	}

	left, right, pivot := 1, pow(2, d)-1, 0
	for left <= right {
		pivot = left + (right-left)/2
		if exists(pivot, d, root) {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}

	return pow(2, d) - 1 + left
}

func depth(root *TreeNode) int {
	d := 0

	for root.Left != nil {
		d++
		root = root.Left
	}

	return d
}

func exists(idx, d int, node *TreeNode) bool {
	left, right, pivot := 0, pow(2, d)-1, 0
	for i := 0; i < d; i++ {
		pivot = left + (right-left)/2
		if idx <= pivot {
			node = node.Left
			right = pivot
		} else {
			node = node.Right
			left = pivot + 1
		}
	}

	return node != nil
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		res *= a
		b--
	}

	return res
}
