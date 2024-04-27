package twilio

import (
	"reflect"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{1, 1, 2, 2, 2, 3},
			},
			want: []int{3, 1, 1, 2, 2, 2},
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{2, 3, 1, 3, 2},
			},
			want: []int{1, 3, 3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
