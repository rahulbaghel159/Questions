package process

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		s string
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
				s: "5F3Z-2e-9-w",
				k: 4,
			},
			want: "5F3Z-2E9W",
		},
		{
			name: "Case 2",
			args: args{
				s: "2-5g-3-J",
				k: 2,
			},
			want: "2-5G-3J",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
