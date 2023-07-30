package soupservings

import (
	"math"
	"testing"
)

func Test_soupServings(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Case 1",
			args: args{
				n: 50,
			},
			want: float64(0.625),
		},
		{
			name: "Case 2",
			args: args{
				n: 100,
			},
			want: float64(0.71875),
		},
		{
			name: "Case 3",
			args: args{
				n: 125,
			},
			want: float64(0.74219),
		},
		{
			name: "Case 4",
			args: args{
				n: 850,
			},
			want: float64(0.96612),
		},
		{
			name: "Case 5",
			args: args{
				n: 660295675,
			},
			want: float64(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soupServings(tt.args.n); math.Abs(got-tt.want) > float64(0.00001) {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}
