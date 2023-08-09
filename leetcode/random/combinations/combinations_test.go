package combinations

import (
	"reflect"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Case 1",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		{
			name: "Case 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{
				{1},
			},
		},
		{
			name: "Case 3",
			args: args{
				n: 6,
				k: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 2, 5},
				{1, 2, 6},
				{1, 3, 4},
				{1, 3, 5},
				{1, 3, 6},
				{1, 4, 5},
				{1, 4, 6},
				{1, 5, 6},
				{2, 3, 4},
				{2, 3, 5},
				{2, 3, 6},
				{2, 4, 5},
				{2, 4, 6},
				{2, 5, 6},
				{3, 4, 5},
				{3, 4, 6},
				{3, 5, 6},
				{4, 5, 6},
			},
		},
		{
			name: "Case 4",
			args: args{
				n: 5,
				k: 4,
			},
			want: [][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 5},
				{1, 2, 4, 5},
				{1, 3, 4, 5},
				{2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
