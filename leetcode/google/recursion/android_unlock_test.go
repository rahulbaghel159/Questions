package recursion

import "testing"

func Test_numberOfPatterns(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				m: 1,
				n: 1,
			},
			want: 9,
		},
		{
			name: "Case 2",
			args: args{
				m: 2,
				n: 2,
			},
			want: 56,
		},
		{
			name: "Case 3",
			args: args{
				m: 1,
				n: 2,
			},
			want: 65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPatterns(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("numberOfPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}
