package fairDistributionCookies

import "testing"

func Test_distributeCookies(t *testing.T) {
	type args struct {
		cookies []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				cookies: []int{8, 15, 10, 20, 8},
				k:       2,
			},
			want: 31,
		},
		{
			name: "Case 2",
			args: args{
				cookies: []int{6, 1, 3, 2, 2, 4, 1, 2},
				k:       3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCookies(tt.args.cookies, tt.args.k); got != tt.want {
				t.Errorf("distributeCookies() = %v, want %v", got, tt.want)
			}
		})
	}
}
