package main

import "testing"

func Test_thirdMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Case 1",
			nums: []int{3, 2, 1},
			want: 1,
		},
		{
			name: "Case 2",
			nums: []int{1, 2},
			want: 2,
		},
		{
			name: "Case 3",
			nums: []int{2, 2, 3, 1},
			want: 1,
		},
		{
			name: "Case 4",
			nums: []int{1, 1, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := thirdMax(tt.nums)
			if got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
