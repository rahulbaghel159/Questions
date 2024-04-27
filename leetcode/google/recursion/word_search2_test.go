package recursion

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
					[]byte("oaan"),
					[]byte("etae"),
					[]byte("ihkr"),
					[]byte("iflv"),
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"oath", "eat"},
		},
		{
			name: "Case 2",
			args: args{
				board: [][]byte{
					[]byte("oabn"),
					[]byte("otae"),
					[]byte("ahkr"),
					[]byte("aflv"),
				},
				words: []string{"oa", "oaa"},
			},
			want: []string{"oa", "oaa"},
		},
		{
			name: "Case 3",
			args: args{
				board: [][]byte{
					[]byte("a"),
					[]byte("a"),
				},
				words: []string{"aaa"},
			},
			want: []string{},
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
