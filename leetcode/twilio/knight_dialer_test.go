package twilio

import "testing"

func Test_knightDialer(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				n: 1,
			},
			want: 10,
		},
		{
			name: "Case 2",
			args: args{
				n: 2,
			},
			want: 20,
		},
		{
			name: "Case 3",
			args: args{
				n: 3,
			},
			want: 46,
		},
		{
			name: "Case 4",
			args: args{
				n: 3131,
			},
			want: 136006598,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knightDialer(tt.args.n); got != tt.want {
				t.Errorf("knightDialer() = %v, want %v", got, tt.want)
			}
		})
	}
}
