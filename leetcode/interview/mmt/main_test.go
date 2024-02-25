package mmt

import "testing"

func Test_countTank(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				str: "_H_H_H_H_H_",
			},
			want: 3,
		},
		{
			name: "Case 2",
			args: args{
				str: "_H_H_HH_",
			},
			want: 3,
		},
		{
			name: "Case 3",
			args: args{
				str: "_HH_",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTank(tt.args.str); got != tt.want {
				t.Errorf("countTank() = %v, want %v", got, tt.want)
			}
		})
	}
}
