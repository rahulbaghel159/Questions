package maxrequest

import "testing"

func Test_maximumRequests(t *testing.T) {
	type args struct {
		n        int
		requests [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				n: 5,
				requests: [][]int{
					{0, 1},
					{1, 0},
					{0, 1},
					{1, 2},
					{2, 0},
					{3, 4},
				},
			},
			want: 5,
		},
		{
			name: "Case 2",
			args: args{
				n: 3,
				requests: [][]int{
					{0, 0},
					{1, 2},
					{2, 1},
				},
			},
			want: 3,
		},
		{
			name: "Case 3",
			args: args{
				n: 4,
				requests: [][]int{
					{0, 3},
					{3, 1},
					{1, 2},
					{2, 0},
				},
			},
			want: 4,
		},
		{
			name: "Case 4",
			args: args{
				n: 3,
				requests: [][]int{
					{1, 2},
					{1, 2},
					{2, 2},
					{0, 2},
					{2, 1},
					{1, 1},
					{1, 2},
				},
			},
			want: 4,
		},
		{
			name: "Case 5",
			args: args{
				n: 3,
				requests: [][]int{
					{0, 0},
					{1, 1},
					{0, 0},
					{2, 0},
					{2, 2},
					{1, 1},
					{2, 1},
					{0, 1},
					{0, 1},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRequests(tt.args.n, tt.args.requests); got != tt.want {
				t.Errorf("maximumRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
