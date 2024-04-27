package google

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "Case 2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "Case 3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "Case 4",
			args: args{
				s: "ab",
				t: "a",
			},
			want: "a",
		},
		{
			name: "Case 5",
			args: args{
				s: "ab",
				t: "b",
			},
			want: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
