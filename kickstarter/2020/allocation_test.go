package twentytwenty

import "testing"

func TestMaxHouses(t *testing.T) {
	type args struct {
		arr []int
		b   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				arr: []int{20, 90, 40, 90},
				b:   100,
			},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{
				arr: []int{30, 30, 10, 10},
				b:   50,
			},
			want: 3,
		},
		{
			name: "Test Case 3",
			args: args{
				arr: []int{999, 999, 999},
				b:   300,
			},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{
				arr: []int{30, 30, 10, 10},
				b:   80,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxHouses(tt.args.arr, tt.args.b); got != tt.want {
				t.Errorf("MaxHouses() = %v, want %v", got, tt.want)
			}
		})
	}
}
