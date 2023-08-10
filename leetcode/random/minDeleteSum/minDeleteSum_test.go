package mindeletesum

import "testing"

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			name: "Case 2",
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
		{
			name: "Case 3",
			args: args{
				s1: "nbzvshnmtlioe",
				s2: "rudntqxfmslvpib",
			},
			want: 2206,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
