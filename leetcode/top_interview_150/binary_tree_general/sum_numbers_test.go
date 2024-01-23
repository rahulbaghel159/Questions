package binarytreegeneral

import "testing"

func Test_sumNumbers(t *testing.T) {
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
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 25,
		},
		{
			name: "Case 2",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{
						Val: 0,
					},
				},
			},
			want: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
