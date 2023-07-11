package mindepthbinarytree

import "fmt"

//https://leetcode.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q, depth := &queue{arr: []*TreeNode{}}, 1

	q.push(root)

	for !q.isEmpty() {
		qSize := q.size()

		for qSize > 0 {
			qSize--
			node := q.pop()

			if node == nil {
				continue
			}

			if node.Left == nil && node.Right == nil {
				fmt.Println(depth)
				return depth
			}

			q.push(node.Left)
			q.push(node.Right)
		}
		depth++
	}

	return depth
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type queue struct {
	arr []*TreeNode
}

func (q *queue) push(node *TreeNode) {
	q.arr = append(q.arr, node)
}

func (q *queue) pop() *TreeNode {
	node := q.arr[0]
	q.arr = q.arr[1:len(q.arr)]

	return node
}

func (q *queue) isEmpty() bool {
	return len(q.arr) == 0
}

func (q *queue) size() int {
	return len(q.arr)
}
