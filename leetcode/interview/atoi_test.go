package interview

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "-42",
			},
			want: -42,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "Test Case 4",
			args: args{
				s: "words and 987",
			},
			want: 0,
		},
		{
			name: "Test Case 5",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "Test Case 6",
			args: args{
				s: "91283472332",
			},
			want: 2147483647,
		},
		{
			name: "Test Case 7",
			args: args{
				s: "3.14159",
			},
			want: 3,
		},
		{
			name: "Test Case 8",
			args: args{
				s: "9223372036854775808",
			},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
