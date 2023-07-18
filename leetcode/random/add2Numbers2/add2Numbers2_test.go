package add2numbers2

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Case 1",
			args: args{
				l1: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 7,
						},
					},
				},
			},
		},
		{
			name: "Case 2",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 7,
					},
				},
			},
		},
		{
			name: "Case 3",
			args: args{
				l1: &ListNode{Val: 0},
				l2: &ListNode{Val: 0},
			},
			want: &ListNode{Val: 0},
		},
		{
			name: "Case 4",
			args: args{
				l1: &ListNode{Val: 1},
				l2: &ListNode{
					Val:  9,
					Next: &ListNode{Val: 9},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  0,
					Next: &ListNode{Val: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
