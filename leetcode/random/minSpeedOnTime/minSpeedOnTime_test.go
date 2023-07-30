package minspeedontime

import "testing"

func Test_minSpeedOnTime(t *testing.T) {
	type args struct {
		dist []int
		hour float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				dist: []int{1, 3, 2},
				hour: float64(6),
			},
			want: 1,
		},
		{
			name: "Case 2",
			args: args{
				dist: []int{1, 3, 2},
				hour: float64(2.7),
			},
			want: 3,
		},
		{
			name: "Case 3",
			args: args{
				dist: []int{1, 3, 2},
				hour: float64(1.9),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSpeedOnTime(tt.args.dist, tt.args.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
