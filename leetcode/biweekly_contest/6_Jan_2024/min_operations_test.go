package jan2024

import "testing"

func Test_minimumOperationsToMakeEqual(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				x: 26,
				y: 1,
			},
			want: 3,
		},
		// {
		// 	name: "Case 2",
		// 	args: args{
		// 		x: 54,
		// 		y: 2,
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "Case 3",
		// 	args: args{
		// 		x: 25,
		// 		y: 30,
		// 	},
		// 	want: 5,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperationsToMakeEqual(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("minimumOperationsToMakeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
