package recursion1

import "testing"

func Test_kthGrammar(t *testing.T) {
	type args struct {
		N int
		K int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				N: 1,
				K: 1,
			},
			want: 0,
		},
		{
			name: "Test Case 2",
			args: args{
				N: 2,
				K: 1,
			},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{
				N: 2,
				K: 2,
			},
			want: 1,
		},
		{
			name: "Test Case 4",
			args: args{
				N: 4,
				K: 5,
			},
			want: 1,
		},
		{
			name: "Test Case 5",
			args: args{
				N: 3,
				K: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("kthGrammar() = %v, want %v", got, tt.want)
			}
		})
	}
}
