package dynamicprogramming

import "testing"

func Test_checkSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{23, 2, 4, 6, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    6,
			},
			want: true,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{23, 2, 6, 4, 7},
				k:    13,
			},
			want: false,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{23, 2, 4, 6, 6},
				k:    7,
			},
			want: true,
		},
		{
			name: "Case 5",
			args: args{
				nums: []int{5, 0, 0, 0},
				k:    3,
			},
			want: true,
		},
		{
			name: "Case 6",
			args: args{
				nums: []int{2, 4, 3},
				k:    3,
			},
			want: true,
		},
		{
			name: "Case 7",
			args: args{
				nums: []int{0, 0},
				k:    1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
