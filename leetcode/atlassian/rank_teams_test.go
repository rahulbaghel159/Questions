package atlassian

import "testing"

func Test_rankTeams(t *testing.T) {
	type args struct {
		votes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				votes: []string{"ABC", "ACB", "ABC", "ACB", "ACB"},
			},
			want: "ACB",
		},
		{
			name: "Case 2",
			args: args{
				votes: []string{"WXYZ", "XYZW"},
			},
			want: "XWYZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rankTeams(tt.args.votes); got != tt.want {
				t.Errorf("rankTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
