package google

import "testing"

func Test_nextClosestTime(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				time: "19:34",
			},
			want: "19:39",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextClosestTime(tt.args.time); got != tt.want {
				t.Errorf("nextClosestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
