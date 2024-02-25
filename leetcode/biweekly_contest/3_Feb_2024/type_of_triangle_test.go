package feb2024

import "testing"

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{3, 3, 3},
			},
			want: "equilateral",
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{3, 4, 5},
			},
			want: "scalene",
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{3, 4, 8},
			},
			want: "none",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
