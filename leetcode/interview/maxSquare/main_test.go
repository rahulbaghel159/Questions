package maxsquare

import "testing"

func Test_maxSquareSize(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				matrix: [][]int{
					{1, 0, 1, 0, 0},
					{1, 0, 1, 1, 1},
					{1, 1, 1, 1, 1},
					{1, 0, 0, 1, 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSquareSize(tt.args.matrix); got != tt.want {
				t.Errorf("maxSquareSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
