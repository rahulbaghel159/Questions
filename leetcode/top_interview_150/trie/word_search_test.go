package trie

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case 1",
			args: args{
				board: [][]byte{
					[]byte("oath"),
					[]byte("etae"),
					[]byte("ihkr"),
					[]byte("iflv"),
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"eat", "oath"},
		},
		{
			name: "Case 2",
			args: args{
				board: [][]byte{
					[]byte("ab"),
					[]byte("cd"),
				},
				words: []string{"abcb"},
			},
			want: []string{},
		},
		{
			name: "Case 3",
			args: args{
				board: [][]byte{
					[]byte("ab"),
					[]byte("cd"),
				},
				words: []string{"abdc"},
			},
			want: []string{"abdc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
