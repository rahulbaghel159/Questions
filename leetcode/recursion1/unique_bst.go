package recursion1

func generateTrees(n int) []*TreeNode {
	//breaking condition
	//if n==1, only 1 tree is possible
	if n == 1 {
		return []*TreeNode{
			&TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		}
	}
	//if n==2, 2 trees are possible
	if n == 2 {
		return []*TreeNode{
			&TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			&TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		}
	}
	return nil
}
