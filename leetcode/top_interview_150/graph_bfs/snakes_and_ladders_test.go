package graphbfs

import "testing"

func Test_snakesAndLadders(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 0",
			args: args{
				board: [][]int{
					{-1, -1},
					{-1, 3},
				},
			},
			want: 1,
		},
		{
			name: "Case 1",
			args: args{
				board: [][]int{
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 35, -1, -1, 13, -1},
					{-1, -1, -1, -1, -1, -1},
					{-1, 15, -1, -1, -1, -1},
				},
			},
			want: 4,
		},
		{
			name: "Case 2",
			args: args{
				board: [][]int{
					{-1, 4, -1},
					{6, 2, 6},
					{-1, 3, -1},
				},
			},
			want: 2,
		},
		{
			name: "Case 3",
			args: args{
				board: [][]int{
					{-1, -1, 19, 10, -1},
					{2, -1, -1, 6, -1},
					{-1, 17, -1, 19, -1},
					{25, -1, 20, -1, -1},
					{-1, -1, -1, -1, 15},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := snakesAndLadders(tt.args.board); got != tt.want {
				t.Errorf("snakesAndLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
