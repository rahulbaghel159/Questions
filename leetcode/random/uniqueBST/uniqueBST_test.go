package uniquebst

import (
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "Case 1",
			args: args{
				n: 3,
			},
			want: []*TreeNode{
				{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
				{
					Val: 1,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 2,
						},
					},
				},
				{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
				},
				{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
		},
		{
			name: "Case 2",
			args: args{
				n: 1,
			},
			want: []*TreeNode{
				{
					Val: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !isEqual(got, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}

func isEqual(arr1, arr2 []*TreeNode) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		if !AreNodesEqual(arr1[i], arr2[i]) {
			return false
		}
	}

	return true
}

func AreNodesEqual(n1, n2 *TreeNode) bool {
	//both nil
	if n1 == nil && n2 == nil {
		return true
	}

	//1 of the 2 node is nil
	if n1 == nil || n2 == nil {
		return false
	}

	return n1.Val == n2.Val && AreNodesEqual(n1.Left, n2.Left) && AreNodesEqual(n1.Right, n2.Right)
}
