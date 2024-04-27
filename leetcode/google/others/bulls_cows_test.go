package others

import "testing"

func Test_getHint(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				secret: "1807",
				guess:  "7810",
			},
			want: "1A3B",
		},
		{
			name: "Case 2",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want: "1A1B",
		},
		{
			name: "Case 3",
			args: args{
				secret: "1122",
				guess:  "1222",
			},
			want: "3A0B",
		},
		{
			name: "Case 4",
			args: args{
				secret: "1122",
				guess:  "2211",
			},
			want: "0A4B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
