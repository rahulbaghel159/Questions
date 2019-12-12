package threesum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Positive Case-1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		{
			name: "Case with all 0's",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				[]int{0, 0, 0},
			},
		},
		{
			name: "Case with array length less than 3",
			args: args{
				nums: []int{0, 0},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
