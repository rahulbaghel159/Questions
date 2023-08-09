package uniquebst

import "fmt"

//https://leetcode.com/problems/unique-binary-search-trees-ii/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return allPossibleBST(1, n, make(map[string][]*TreeNode))
}

func allPossibleBST(start, end int, memo map[string][]*TreeNode) []*TreeNode {
	res := make([]*TreeNode, 0)

	if start > end {
		res = append(res, nil)
		return res
	}

	if _, ok := memo[fmt.Sprint(start)+fmt.Sprint(end)]; ok {
		return memo[fmt.Sprint(start)+fmt.Sprint(end)]
	}

	for i := start; i <= end; i++ {
		leftSubTree := allPossibleBST(start, i-1, memo)
		rightSubTree := allPossibleBST(i+1, end, memo)

		for _, u := range leftSubTree {
			for _, v := range rightSubTree {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  u,
					Right: v,
				})
			}
		}
	}

	memo[fmt.Sprint(start)+fmt.Sprint(end)] = res

	return res
}
