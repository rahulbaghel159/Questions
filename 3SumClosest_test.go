package threesum

import "testing"

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{
				nums:   []int{0, 1, 2},
				target: 0,
			},
			want: 3,
		},
		{
			name: "Test Case 3",
			args: args{
				nums:   []int{1, 1, 1, 1},
				target: 0,
			},
			want: 3,
		},
		{
			name: "Test Case 4",
			args: args{
				nums:   []int{1, 1, 1, 1},
				target: -100,
			},
			want: 3,
		},
		{
			name: "Test Case 4",
			args: args{
				nums:   []int{0, 2, 1, -3},
				target: 1,
			},
			want: 0,
		},
		{
			name: "Test Case 5",
			args: args{
				nums:   []int{0, 2, 1, -3, 4},
				target: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
