package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/256/complexity-analysis/2375/

func maxDepth(root *TreeNode) int {
	//breaking condition
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	//processing and recursing
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(n1, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}
