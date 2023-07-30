package maxruntime

import "testing"

func Test_maxRunTime(t *testing.T) {
	type args struct {
		n         int
		batteries []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Case 1",
			args: args{
				n:         2,
				batteries: []int{3, 3, 3},
			},
			want: int64(4),
		},
		{
			name: "Case 2",
			args: args{
				n:         2,
				batteries: []int{1, 1, 1, 1},
			},
			want: int64(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRunTime(tt.args.n, tt.args.batteries); got != tt.want {
				t.Errorf("maxRunTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
