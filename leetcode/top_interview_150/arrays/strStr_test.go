package arrays

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "Case 2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "Case 3",
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "Case 4",
			args: args{
				haystack: "a",
				needle:   "a",
			},
			want: 0,
		},
		{
			name: "Case 5",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			name: "Case 6",
			args: args{
				haystack: "aabaaabaaac",
				needle:   "aabaaac",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
