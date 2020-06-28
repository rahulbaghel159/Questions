package recursion1

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "Test Case 2",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{
				n: 6,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
