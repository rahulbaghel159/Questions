package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/251/scenario-i-recurrence-relation/3233/

func searchBST(root *TreeNode, val int) *TreeNode {
	//breaking condition
	if root.Val == val {
		return root
	}

	//recurse in left sub-tree
	if val < root.Val && root.Left != nil {
		return searchBST(root.Left, val)
	}

	//recurse in right sub-tree
	if val > root.Val && root.Right != nil {
		return searchBST(root.Right, val)
	}

	return nil
}
