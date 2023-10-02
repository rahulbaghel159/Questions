package arrays

import "testing"

func Test_jump(t *testing.T) {
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
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
