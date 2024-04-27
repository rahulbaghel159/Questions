package google

import "testing"

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: 2,
		},
		{
			name: "Case 2",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			want: 1,
		},
		{
			name: "Case 3",
			args: args{
				intervals: [][]int{
					{13, 15},
					{1, 13},
				},
			},
			want: 1,
		},
		{
			name: "Case 4",
			args: args{
				intervals: [][]int{
					{2, 11},
					{6, 16},
					{11, 16},
				},
			},
			want: 2,
		},
		{
			name: "Case 5",
			args: args{
				intervals: [][]int{
					{6, 17},
					{8, 9},
					{11, 12},
					{6, 9},
				},
			},
			want: 3,
		},
		{
			name: "Case 6",
			args: args{
				intervals: [][]int{
					{2, 15},
					{36, 45},
					{9, 29},
					{16, 23},
					{4, 9},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMeetingRooms(tt.args.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
