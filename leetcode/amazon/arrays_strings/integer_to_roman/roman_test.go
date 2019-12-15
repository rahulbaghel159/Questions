package roman

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				num: 3,
			},
			want: "III",
		},
		{
			name: "Test Case 2",
			args: args{
				num: 4,
			},
			want: "IV",
		},
		{
			name: "Test Case 3",
			args: args{
				num: 1,
			},
			want: "I",
		},
		{
			name: "Test Case 4",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "Test Case 5",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
		{
			name: "Test Case 6",
			args: args{
				num: 20,
			},
			want: "XX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
