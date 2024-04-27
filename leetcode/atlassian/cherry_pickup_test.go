package atlassian

import "testing"

func Test_cherryPickup(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				grid: [][]int{
					{0, 1, -1},
					{1, 0, -1},
					{1, 1, 1},
				},
			},
			want: 5,
		},
		{
			name: "Case 2",
			args: args{
				grid: [][]int{
					{1, 1, -1},
					{1, -1, 1},
					{-1, 1, 1},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cherryPickup(tt.args.grid); got != tt.want {
				t.Errorf("cherryPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}
