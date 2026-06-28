package main

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				arr: []int{2, 1},
			},
			want: false,
		},
		{
			name: "Case 2",
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				arr: []int{0, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "Case 4",
			args: args{
				arr: []int{0, 2, 3, 4, 5, 2, 1, 0},
			},
			want: true,
		},
		{
			name: "Case 5",
			args: args{
				arr: []int{0, 2, 3, 3, 5, 2, 1, 0},
			},
			want: false,
		},
		{
			name: "Case 6",
			args: args{
				arr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
