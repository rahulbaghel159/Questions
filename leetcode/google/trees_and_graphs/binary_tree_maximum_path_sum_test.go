package treesandgraphs

import "testing"

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: 6,
		},
		{
			name: "Case 2",
			args: args{
				root: &TreeNode{
					Val:  -10,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: 42,
		},
		{
			name: "Case 3",
			args: args{
				root: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: -1},
				},
			},
			want: 2,
		},
		{
			name: "Case 4",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: -2},
					Right: &TreeNode{Val: 3},
				},
			},
			want: 4,
		},
		{
			name: "Case 5",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: -2,
						Left: &TreeNode{
							Val:  1,
							Left: &TreeNode{Val: -1},
						},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:  -3,
						Left: &TreeNode{Val: -2},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
