package minsizesubarraysum

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				nums:   []int{2, 3, 1, 2, 4, 3},
				target: 7,
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				nums:   []int{1, 4, 4},
				target: 4,
			},
			want: 1,
		},
		{
			name: "Case 3",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
				target: 11,
			},
			want: 0,
		},
		{
			name: "Case 4",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 15,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
