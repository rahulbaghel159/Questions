package wraplines

import (
	"reflect"
	"testing"
)

func Test_wrapLines(t *testing.T) {
	type args struct {
		words []string
		m     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case 1",
			args: args{
				words: []string{"The", "day", "began", "as", "still", "as", "the", "night", "abruptly", "lighted", "with", "brilliant", "flame"},
				m:     20,
			},
			want: []string{
				"The-day-began-as",
				"still-as-the-night",
				"abruptly-lighted",
				"with-brilliant-flame",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapLines(tt.args.words, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrapLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
