package concurrent

import "testing"

func Test_print(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print()
		})
	}
}
