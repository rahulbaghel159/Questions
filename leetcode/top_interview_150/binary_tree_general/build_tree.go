package binarytreegeneral

import (
	"github.com/emirpasic/gods/queues/arrayqueue"
)

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	queue := arrayqueue.New()
	for i := 0; i < len(preorder); i++ {
		queue.Enqueue(preorder[i])
	}

	root, _ := recursiveBuild(queue, inorder)

	return root
}

func recursiveBuild(queue *arrayqueue.Queue, inorder []int) (*TreeNode, *arrayqueue.Queue) {
	if len(inorder) == 0 {
		return nil, queue
	}

	if len(inorder) == 1 {
		queue.Dequeue()
		return &TreeNode{Val: inorder[0]}, queue
	}

	v, _ := queue.Dequeue()
	root := &TreeNode{Val: v.(int)}

	i := 0
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == v.(int) {
			break
		}
	}
	right := inorder[i:]
	if len(right) != 0 {
		right = right[1:]
	}
	root.Left, queue = recursiveBuild(queue, inorder[:i])
	root.Right, queue = recursiveBuild(queue, right)

	return root, queue
}
