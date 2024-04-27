package google

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				seats: []int{1, 0, 0, 0, 1, 0, 1},
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				seats: []int{1, 0, 0, 0},
			},
			want: 3,
		},
		{
			name: "Case 3",
			args: args{
				seats: []int{0, 1},
			},
			want: 1,
		},
		{
			name: "Case 4",
			args: args{
				seats: []int{0, 0, 1, 0, 1, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
