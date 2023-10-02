package arrays

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
			name: "Case 1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "Case 2",
			args: args{
				s: "LVIII",
			},
			want: 58,
		},
		{
			name: "Case 3",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
		},
		{
			name: "Case 3",
			args: args{
				s: "MCDLXXVI",
			},
			want: 1476,
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
