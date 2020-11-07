package recursion2

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				matrix: [][]int{
					[]int{1, 4, 7, 11, 15},
					[]int{2, 5, 8, 12, 19},
					[]int{3, 6, 9, 16, 22},
					[]int{10, 13, 14, 17, 24},
					[]int{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				matrix: [][]int{
					[]int{1, 4, 7, 11, 15},
					[]int{2, 5, 8, 12, 19},
					[]int{3, 6, 9, 16, 22},
					[]int{10, 13, 14, 17, 24},
					[]int{18, 21, 23, 26, 30},
				},
				target: 20,
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				matrix: [][]int{
					[]int{17, 24},
					[]int{26, 30},
				},
				target: 20,
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				matrix: [][]int{
					[]int{10, 13, 14},
					[]int{18, 21, 23},
				},
				target: 20,
			},
			want: false,
		},
		{
			name: "Test Case 5",
			args: args{
				matrix: [][]int{
					[]int{-1, 3},
				},
				target: -1,
			},
			want: true,
		},
		{
			name: "Test Case 6",
			args: args{
				matrix: [][]int{
					[]int{1, 2, 3, 4, 5},
					[]int{6, 7, 8, 9, 10},
					[]int{11, 12, 13, 14, 15},
					[]int{16, 17, 18, 19, 20},
					[]int{21, 22, 23, 24, 25},
				},
				target: 25,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
