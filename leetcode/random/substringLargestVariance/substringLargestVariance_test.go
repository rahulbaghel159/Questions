package substringlargestvariance

import "testing"

func Test_largestVariance(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				s: "aababbb",
			},
			want: 3,
		},
		{
			name: "Case 2",
			args: args{
				s: "abcde",
			},
			want: 0,
		},
		{
			name: "Case 3",
			args: args{
				s: "aaaaa",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestVariance(tt.args.s); got != tt.want {
				t.Errorf("largestVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}
