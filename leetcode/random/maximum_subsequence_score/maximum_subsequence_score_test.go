package maximumsubsequencescore

import "testing"

func Test_maxScore(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				nums1: []int{1, 3, 3, 2},
				nums2: []int{2, 1, 3, 4},
				k:     3,
			},
			want: int64(12),
		},
		{
			name: "",
			args: args{
				nums1: []int{4, 2, 3, 1, 1},
				nums2: []int{7, 5, 10, 9, 6},
				k:     1,
			},
			want: int64(30),
		},
		{
			name: "",
			args: args{
				nums1: []int{2, 1, 14, 12},
				nums2: []int{11, 7, 13, 6},
				k:     3,
			},
			want: int64(168),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
