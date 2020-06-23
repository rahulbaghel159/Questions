package recursion1

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test Case 1",
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if reverseString(tt.args.s); !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("threeSum() = %v, want %v", string(tt.args.s), string(tt.want))
			}
		})
	}
}

func Test_reverseString2(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test Case 1",
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseString2() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
