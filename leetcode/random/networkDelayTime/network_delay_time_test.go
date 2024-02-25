package networkdelaytime

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				times: [][]int{
					{2, 1, 1},
					{2, 3, 1},
					{3, 4, 1},
				},
				n: 4,
				k: 2,
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				times: [][]int{
					{1, 2, 1},
					{2, 1, 3},
				},
				n: 2,
				k: 2,
			},
			want: 3,
		},
		{
			name: "Case 3",
			args: args{
				times: [][]int{
					{1, 2, 1},
					{2, 3, 2},
					{1, 3, 2},
				},
				n: 3,
				k: 1,
			},
			want: 2,
		},
		{
			name: "Case 4",
			args: args{
				times: [][]int{
					{1, 2, 1},
					{2, 3, 7},
					{1, 3, 4},
					{2, 1, 2},
				},
				n: 3,
				k: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := networkDelayTime(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
