package dynamicprogramming

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{2, 3, -2, 4},
			},
			want: 6,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{-2, 0, -1},
			},
			want: 0,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{-2},
			},
			want: -2,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{-3, -1, -1},
			},
			want: 3,
		},
		{
			name: "Case 5",
			args: args{
				nums: []int{-4, -3, -2},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
