package treesandgraphs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/explore/interview/card/google/61/trees-and-graphs/3067/
var max_sum int

func maxPathSum(root *TreeNode) int {
	INF := -999999999999999999

	max_sum = INF
	gain_from_subtree(root)
	return max_sum
}

func gain_from_subtree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left_gain := max(gain_from_subtree(root.Left), 0)
	right_gain := max(gain_from_subtree(root.Right), 0)

	max_sum = max(max_sum, left_gain+right_gain+root.Val)

	return max(left_gain+root.Val, right_gain+root.Val)
}
