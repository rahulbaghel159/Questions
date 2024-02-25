package google

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
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
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Test Case 4",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Test Case 5",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Test Case 6",
			args: args{
				s: "abcabdef",
			},
			want: 6,
		},
		{
			name: "Test Case 7",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "Test Case 8",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
