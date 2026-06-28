package main

import "testing"

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Case 1",
			args: args{
				arr: []int{1, 0, 2, 3, 0, 4, 5, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.args.arr)
		})
	}
}
