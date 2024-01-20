package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
				left:  2,
				right: 4,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			name: "Case 2",
			args: args{
				head: &ListNode{
					Val: 5,
				},
				left:  1,
				right: 1,
			},
			want: &ListNode{
				Val: 5,
			},
		},
		{
			name: "Case 3",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
				left:  1,
				right: 1,
			},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
		{
			name: "Case 4",
			args: args{
				head: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
				left:  1,
				right: 2,
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
