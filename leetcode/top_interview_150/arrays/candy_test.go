package arrays

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				ratings: []int{1, 0, 2},
			},
			want: 5,
		},
		{
			name: "Case 2",
			args: args{
				ratings: []int{1, 2, 2},
			},
			want: 4,
		},
		{
			name: "Case 3",
			args: args{
				ratings: []int{1, 2, 87, 87, 87, 2, 1},
			},
			want: 13,
		},
		{
			name: "Case 4",
			args: args{
				ratings: []int{29, 51, 87, 87, 72, 12},
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
