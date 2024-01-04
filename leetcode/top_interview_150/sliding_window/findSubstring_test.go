package slidingwindow

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			want: []int{0, 9},
		},
		{
			name: "Case 2",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			want: []int{},
		},
		{
			name: "Case 3",
			args: args{
				s:     "barfoofoobarthefoobarman",
				words: []string{"bar", "foo", "the"},
			},
			want: []int{6, 9, 12},
		},
		{
			name: "Case 4",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "good"},
			},
			want: []int{8},
		},
		{
			name: "Case 5",
			args: args{
				s:     "lingmindraboofooowingdingbarrwingmonkeypoundcake",
				words: []string{"fooo", "barr", "wing", "ding", "wing"},
			},
			want: []int{13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
