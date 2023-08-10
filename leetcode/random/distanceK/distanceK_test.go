package distancek

import (
	"reflect"
	"sort"
	"testing"
)

func Test_distanceK(t *testing.T) {
	type args struct {
		root   *TreeNode
		target *TreeNode
		k      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "Case 1",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 3,
		// 			Left: &TreeNode{
		// 				Val: 5,
		// 				Left: &TreeNode{
		// 					Val: 6,
		// 				},
		// 				Right: &TreeNode{
		// 					Val:   2,
		// 					Left:  &TreeNode{Val: 7},
		// 					Right: &TreeNode{Val: 4},
		// 				},
		// 			},
		// 			Right: &TreeNode{
		// 				Val:   1,
		// 				Left:  &TreeNode{Val: 0},
		// 				Right: &TreeNode{Val: 8},
		// 			},
		// 		},
		// 		target: &TreeNode{
		// 			Val: 5,
		// 			Left: &TreeNode{
		// 				Val: 6,
		// 			},
		// 			Right: &TreeNode{
		// 				Val:   2,
		// 				Left:  &TreeNode{Val: 7},
		// 				Right: &TreeNode{Val: 4},
		// 			},
		// 		},
		// 		k: 2,
		// 	},
		// 	want: []int{7, 4, 1},
		// },
		// {
		// 	name: "Case 2",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 1,
		// 		},
		// 		target: &TreeNode{
		// 			Val: 5,
		// 		},
		// 		k: 3,
		// 	},
		// 	want: []int{},
		// },
		// {
		// 	name: "Case 3",
		// 	args: args{
		// 		root: &TreeNode{
		// 			Val: 0,
		// 			Left: &TreeNode{
		// 				Val:   1,
		// 				Left:  &TreeNode{Val: 3},
		// 				Right: &TreeNode{Val: 2},
		// 			},
		// 		},
		// 		target: &TreeNode{
		// 			Val: 2,
		// 		},
		// 		k: 1,
		// 	},
		// 	want: []int{1},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distanceK(tt.args.root, tt.args.target, tt.args.k)

			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i] < tt.want[j]
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
