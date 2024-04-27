package wonderfulsubstring

import "testing"

func Test_wonderfulSubstrings(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Case 1",
			args: args{
				word: "aba",
			},
			want: int64(4),
		},
		{
			name: "Case 2",
			args: args{
				word: "aabb",
			},
			want: int64(9),
		},
		{
			name: "Case 3",
			args: args{
				word: "he",
			},
			want: int64(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wonderfulSubstrings(tt.args.word); got != tt.want {
				t.Errorf("wonderfulSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
