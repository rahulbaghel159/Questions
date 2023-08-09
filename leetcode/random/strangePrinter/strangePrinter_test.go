package strangeprinter

import "testing"

func Test_strangePrinter(t *testing.T) {
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
				s: "aaabbb",
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				s: "aba",
			},
			want: 2,
		},
		{
			name: "Case 3",
			args: args{
				s: "tbgtgb",
			},
			want: 4,
		},
		{
			name: "Case 4",
			args: args{
				s: "abcabc",
			},
			want: 5,
		},
		{
			name: "Case 5",
			args: args{
				s: "abcdefghijklmn",
			},
			want: 14,
		},
		{
			name: "Case 6",
			args: args{
				s: "baacdddaaddaaaaccbddbcabdaabdbbcdcbbbacbddcabcaaa",
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strangePrinter(tt.args.s); got != tt.want {
				t.Errorf("strangePrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}
