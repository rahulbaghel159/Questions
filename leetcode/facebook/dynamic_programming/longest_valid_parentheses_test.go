package dynamicprogramming

import "testing"

func Test_longestValidParentheses(t *testing.T) {
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
				s: "(()",
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "Case 3",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Case 4",
			args: args{
				s: "(",
			},
			want: 0,
		},
		{
			name: "Case 5",
			args: args{
				s: "()(())",
			},
			want: 6,
		},
		{
			name: "Case 6",
			args: args{
				s: "()(()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
