package others

import "testing"

func Test_minAreaRect(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				points: [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAreaRect(tt.args.points); got != tt.want {
				t.Errorf("minAreaRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
