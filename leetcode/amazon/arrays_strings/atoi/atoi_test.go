package atoi

import "testing"

func Test_atoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				str: "42",
			},
			want: 42,
		},
		{
			name: "Test Case 2",
			args: args{
				str: "-42",
			},
			want: -42,
		},
		{
			name: "Test Case 3",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "Test Case 4",
			args: args{
				str: "words and 987",
			},
			want: 0,
		},
		{
			name: "Test Case 5",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "Test Case 6",
			args: args{
				str: "+1",
			},
			want: 1,
		},
		{
			name: "Test Case 7",
			args: args{
				str: "9223372036854775808",
			},
			want: 2147483647,
		},
		{
			name: "Test Case 8",
			args: args{
				str: "words",
			},
			want: 0,
		},
		{
			name: "Test Case 8",
			args: args{
				str: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atoi(tt.args.str); got != tt.want {
				t.Errorf("atoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
