package dynamicprogramming

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				s: "babad",
			},
			want: "bab",
		},
		{
			name: "Case 2",
			args: args{
				s: "cbbd",
			},
			want: "bb",
		},
		{
			name: "Case 3",
			args: args{
				s: "xaabacxcabaaxcabaax",
			},
			want: "xaabacxcabaax",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
