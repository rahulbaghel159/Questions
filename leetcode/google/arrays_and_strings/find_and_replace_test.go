package google

import "testing"

func Test_findReplaceString(t *testing.T) {
	type args struct {
		s       string
		indices []int
		sources []string
		targets []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"a", "cd"},
				targets: []string{"eee", "ffff"},
			},
			want: "eeebffff",
		},
		{
			name: "Case 2",
			args: args{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"ab", "ec"},
				targets: []string{"eee", "ffff"},
			},
			want: "eeecd",
		},
		{
			name: "Case 3",
			args: args{
				s:       "wreorttvosuidhrxvmvo",
				indices: []int{14, 12, 10, 5, 0, 18},
				sources: []string{"rxv", "dh", "ui", "ttv", "wreor", "vo"},
				targets: []string{"frs", "c", "ql", "qpir", "gwbeve", "n"},
			},
			want: "gwbeveqpirosqlcfrsmn",
		},
		{
			name: "Case 4",
			args: args{
				s:       "abcde",
				indices: []int{2, 2},
				sources: []string{"cde", "cdef"},
				targets: []string{"fe", "f"},
			},
			want: "abfe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findReplaceString(tt.args.s, tt.args.indices, tt.args.sources, tt.args.targets); got != tt.want {
				t.Errorf("findReplaceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
