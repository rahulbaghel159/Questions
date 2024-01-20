package hashmap

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "Case 2",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "Case 3",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "Case 4",
			args: args{
				n: 4,
			},
			want: false,
		},
		{
			name: "Case 4",
			args: args{
				n: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
