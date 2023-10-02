package arrays

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "Case 2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			want: []int{1},
		},
		{
			name: "Case 3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
			want: []int{1},
		},
		{
			name: "Case 4",
			args: args{
				nums1: []int{1, 0},
				m:     1,
				nums2: []int{2},
				n:     1,
			},
			want: []int{1, 2},
		},
		{
			name: "Case 5",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{4, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n); !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
