package chess

import (
	"reflect"
	"testing"

	slice "github.com/bradfitz/slice"
)

func Test_possibleQueens(t *testing.T) {
	type args struct {
		queens []Position
		king   Position
	}
	tests := []struct {
		name               string
		args               args
		wantPossibleQueens []Position
	}{
		{
			name: "Test Case 1",
			args: args{
				queens: []Position{
					{0, 0},
					{1, 1},
					{2, 2},
					{3, 4},
					{3, 5},
					{4, 4},
					{4, 5},
				},
				king: Position{
					i: 3, j: 3,
				},
			},
			wantPossibleQueens: []Position{
				{2, 2},
				{3, 4},
				{4, 4},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				queens: []Position{
					{0, 1},
					{1, 0},
					{4, 0},
					{0, 4},
					{3, 3},
					{2, 4},
				},
				king: Position{
					i: 0, j: 0,
				},
			},
			wantPossibleQueens: []Position{
				{0, 1},
				{1, 0},
				{3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPossibleQueens := possibleQueens(tt.args.queens, tt.args.king)
			slice.Sort(gotPossibleQueens[:], func(i, j int) bool {
				return (gotPossibleQueens[i].i < gotPossibleQueens[j].i && gotPossibleQueens[i].j < gotPossibleQueens[j].j)
			})
			slice.Sort(tt.wantPossibleQueens[:], func(i, j int) bool {
				return (tt.wantPossibleQueens[i].i < tt.wantPossibleQueens[j].i && tt.wantPossibleQueens[i].j < tt.wantPossibleQueens[j].j)
			})
			if !reflect.DeepEqual(gotPossibleQueens, tt.wantPossibleQueens) {
				t.Errorf("possibleQueens() = %v, want %v", gotPossibleQueens, tt.wantPossibleQueens)
			}
		})
	}
}
