package predictwinner

import "testing"

func TestPredictTheWinner(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{1, 5, 2},
			},
			want: false,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{1, 5, 233, 7},
			},
			want: true,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{0},
			},
			want: true,
		},
		{
			name: "Case 4",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PredictTheWinner(tt.args.nums); got != tt.want {
				t.Errorf("PredictTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
