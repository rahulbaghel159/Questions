package google

import "testing"

func Test_expressiveWords(t *testing.T) {
	type args struct {
		s     string
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
				s:     "heeellooo",
				words: []string{"hello", "hi", "helo"},
			},
			want: 1,
		},
		{
			name: "Case 2",
			args: args{
				s:     "abcd",
				words: []string{"abc"},
			},
			want: 0,
		},
		{
			name: "Case 3",
			args: args{
				s:     "aaa",
				words: []string{"aaaa"},
			},
			want: 0,
		},
		{
			name: "Case 4",
			args: args{
				s:     "heeelllooo",
				words: []string{"hellllo"},
			},
			want: 0,
		},
		{
			name: "Case 5",
			args: args{
				s:     "sass",
				words: []string{"sa"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWords(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
