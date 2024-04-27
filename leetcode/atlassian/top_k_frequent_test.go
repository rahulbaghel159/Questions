package atlassian

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		words []string
		k     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case 1",
			args: args{
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     2,
			},
			want: []string{"i", "love"},
		},
		{
			name: "Case 2",
			args: args{
				words: []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
				k:     4,
			},
			want: []string{"the", "is", "sunny", "day"},
		},
		{
			name: "Case 3",
			args: args{
				words: []string{"i", "love", "leetcode", "i", "love", "coding"},
				k:     1,
			},
			want: []string{"i"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.words, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
