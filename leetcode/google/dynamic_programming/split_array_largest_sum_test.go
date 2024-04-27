package dynamicprogramming

import "testing"

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{7, 2, 5, 10, 8},
				k:    2,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
