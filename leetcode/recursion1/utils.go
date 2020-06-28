package recursion1

import (
	"fmt"

	"github.com/emirpasic/gods/stacks/arraystack"
)

//PrintNode prints linked list node
func PrintNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%+v", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

//PrintTreeNode prints tree node
func PrintTreeNode(root *TreeNode) {
	if root == nil {
		return
	}
	stack := arraystack.New()

	stack.Push(root)

	node, ok := stack.Pop()
	for ok {
		fmt.Printf("%+v", node.(*TreeNode).Val)

		if node.(*TreeNode).Left != nil {
			stack.Push(node.(*TreeNode).Left)
		}

		if node.(*TreeNode).Left != nil {
			stack.Push(node.(*TreeNode).Right)
		}

		node, ok = stack.Pop()
	}
	fmt.Println()
}
