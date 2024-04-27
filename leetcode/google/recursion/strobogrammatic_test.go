package recursion

import (
	"reflect"
	"testing"
)

func Test_findStrobogrammatic(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case 1",
			args: args{
				n: 1,
			},
			want: []string{"0", "1", "8"},
		},
		{
			name: "Case 2",
			args: args{
				n: 2,
			},
			want: []string{"11", "69", "88", "96"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStrobogrammatic(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findStrobogrammatic() = %v, want %v", got, tt.want)
			}
		})
	}
}
