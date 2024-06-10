package agoda

import "testing"

func Test_leastBricks(t *testing.T) {
	type args struct {
		wall [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				wall: [][]int{
					{1, 2, 2, 1},
					{3, 1, 2},
					{1, 3, 2},
					{2, 4},
					{3, 1, 2},
					{1, 3, 1, 1},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastBricks(tt.args.wall); got != tt.want {
				t.Errorf("leastBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}