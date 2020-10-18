package recursion1

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateTreeList(1, n)
}

func generateTreeList(start, end int) []*TreeNode {
	list := make([]*TreeNode, 0)
	if start > end {
		list = append(list, nil)
	}

	for i := start; i <= end; i++ {
		leftList := generateTreeList(start, i-1)
		rightList := generateTreeList(i+1, end)
		for _, left := range leftList {
			for _, right := range rightList {
				root := &TreeNode{Val: i}
				root.Left = left
				root.Right = right
				list = append(list, root)
			}
		}
	}

	return list
}
