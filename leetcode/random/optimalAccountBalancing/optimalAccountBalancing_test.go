package optimalaccountbalancing

import "testing"

func Test_minTransfers(t *testing.T) {
	type args struct {
		transactions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "Case 1",
		// 	args: args{
		// 		transactions: [][]int{
		// 			{0, 1, 10},
		// 			{2, 0, 5},
		// 		},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "Case 2",
		// 	args: args{
		// 		transactions: [][]int{
		// 			{0, 1, 10},
		// 			{1, 0, 1},
		// 			{1, 2, 5},
		// 			{2, 0, 5},
		// 		},
		// 	},
		// 	want: 1,
		// },
		// {
		// 	name: "Case 3",
		// 	args: args{
		// 		transactions: [][]int{
		// 			{0, 1, 2},
		// 			{1, 2, 1},
		// 			{1, 3, 1},
		// 		},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "Case 4",
		// 	args: args{
		// 		transactions: [][]int{
		// 			{1, 5, 8},
		// 			{8, 9, 8},
		// 			{2, 3, 9},
		// 			{4, 3, 1},
		// 		},
		// 	},
		// 	want: 4,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTransfers(tt.args.transactions); got != tt.want {
				t.Errorf("minTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
