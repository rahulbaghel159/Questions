package feb2024

import "testing"

func Test_maximumSubarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    1,
			},
			want: 11,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{-1, 3, 2, 4, 5},
				k:    3,
			},
			want: 11,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{-1, -2, -3, -4},
				k:    2,
			},
			want: -6,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{1, 5},
				k:    2,
			},
			want: 0,
		},
		{
			name: "Case 5",
			args: args{
				nums: []int{3, 3, 2},
				k:    1,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maximumSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
