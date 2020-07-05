package recursion1

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Case 1",
			args: args{
				n: 10,
				x: 2.00000,
			},
			want: 1024.000000,
		},
		{
			name: "Test Case 2",
			args: args{
				n: 3,
				x: 2.10000,
			},
			want: 9.261000000000001,
		},
		{
			name: "Test Case 3",
			args: args{
				n: -2,
				x: 2.00000,
			},
			want: 0.25000,
		},
		{
			name: "Test Case 4",
			args: args{
				n: 3,
				x: 8.88023,
			},
			want: 700.2814829452681,
		},
		{
			name: "Test Case 5",
			args: args{
				n: 2147483647,
				x: 0.00001,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
