package atlassian

import "testing"

func Test_longestStrChain(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				words: []string{"a", "b", "ba", "bca", "bda", "bdca"},
			},
			want: 4,
		},
		{
			name: "Case 2",
			args: args{
				words: []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			},
			want: 5,
		},
		{
			name: "Case 3",
			args: args{
				words: []string{"abcd", "dbqca"},
			},
			want: 1,
		},
		{
			name: "Case 4",
			args: args{
				words: []string{"a", "ab", "ac", "bd", "abc", "abd", "abdd"},
			},
			want: 4,
		},
		{
			name: "Case 5",
			args: args{
				words: []string{"wnyxmflkf", "xefx", "usqhb", "ttmdvv", "hagmmn", "tmvv", "pttmdvv", "nmzlhlpr", "ymfk", "uhpaglmmnn", "zckgh", "hgmmn", "isqxrk", "isqrk", "nmzlhpr", "uysyqhxb", "haglmmn", "xfx", "mm", "wymfkf", "tmdvv", "uhaglmmn", "mf", "uhaglmmnn", "mfk", "wnymfkf", "powttkmdvv", "kwnyxmflkf", "xx", "rnqbhxsj", "uysqhb", "pttkmdvv", "hmmn", "iq", "m", "ymfkf", "zckgdh", "zckh", "hmm", "xuefx", "mv", "iqrk", "tmv", "iqk", "wnyxmfkf", "uysyqhb", "v", "m", "pwttkmdvv", "rnqbhsj"},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestStrChain(tt.args.words); got != tt.want {
				t.Errorf("longestStrChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
