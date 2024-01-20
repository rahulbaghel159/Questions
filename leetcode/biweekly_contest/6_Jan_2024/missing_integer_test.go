package jan2024

import "testing"

func Test_missingInteger(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				nums: []int{1, 2, 3, 2, 5},
			},
			want: 6,
		},
		{
			name: "Case 2",
			args: args{
				nums: []int{3, 4, 5, 1, 12, 14, 13},
			},
			want: 15,
		},
		{
			name: "Case 3",
			args: args{
				nums: []int{4, 5, 6, 7, 8, 8, 9, 4, 3, 2, 7},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingInteger(tt.args.nums); got != tt.want {
				t.Errorf("missingInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
