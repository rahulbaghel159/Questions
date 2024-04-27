package treesandgraphs

import "testing"

func Test_crackSafe(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				n: 1,
				k: 2,
			},
			want: "01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crackSafe(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("crackSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
