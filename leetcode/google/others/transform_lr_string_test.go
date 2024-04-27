package others

import "testing"

func Test_canTransform(t *testing.T) {
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				start: "RXXLRXRXL",
				end:   "XRLXXRRLX",
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				start: "X",
				end:   "L",
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				start: "XXXXXLXXXX",
				end:   "LXXXXXXXXX",
			},
			want: true,
		},
		{
			name: "Case 4",
			args: args{
				start: "LXXLXRLXXL",
				end:   "XLLXRXLXLX",
			},
			want: false,
		},
		{
			name: "Case 5",
			args: args{
				start: "XRXXXLXXXR",
				end:   "XXRLXXXRXX",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canTransform(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("canTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
