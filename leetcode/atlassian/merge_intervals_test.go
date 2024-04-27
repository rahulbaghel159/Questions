package atlassian

import (
	"reflect"
	"testing"
)

func Test_mergeInterval(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Case 1",
			args: args{
				intervals: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			name: "Case 2",
			args: args{
				intervals: [][]int{
					{2, 3},
					{4, 5},
					{6, 7},
					{8, 9},
					{1, 10},
				},
			},
			want: [][]int{
				{1, 10},
			},
		},
		{
			name: "Case 3",
			args: args{
				intervals: [][]int{
					{1, 4},
					{5, 6},
				},
			},
			want: [][]int{
				{1, 4},
				{5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInterval(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
