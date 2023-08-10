package numbermusicplaylist

import "testing"

func Test_numMusicPlaylists(t *testing.T) {
	type args struct {
		n    int
		goal int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				n:    3,
				goal: 3,
				k:    1,
			},
			want: 6,
		},
		{
			name: "Case 2",
			args: args{
				n:    2,
				goal: 3,
				k:    0,
			},
			want: 6,
		},
		{
			name: "Case 3",
			args: args{
				n:    2,
				goal: 3,
				k:    1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMusicPlaylists(tt.args.n, tt.args.goal, tt.args.k); got != tt.want {
				t.Errorf("numMusicPlaylists() = %v, want %v", got, tt.want)
			}
		})
	}
}
