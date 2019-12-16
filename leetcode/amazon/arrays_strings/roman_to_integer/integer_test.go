package integer

import "testing"

func Test_romanToInt(t *testing.T) {
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
				s: "III",
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "IV",
			},
			want: 4,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "Test Case 4",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "Test Case 5",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
		{
			name: "Test Case 6",
			args: args{
				s: "DCXXI",
			},
			want: 621,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
