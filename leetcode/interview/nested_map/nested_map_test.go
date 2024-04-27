package nestedmap

import (
	"reflect"
	"testing"
)

func Test_generateNestedMap(t *testing.T) {
	type args struct {
		dict map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Case 1",
			args: args{
				dict: map[string]int{
					"X_a_one": 10,
					"X_a_two": 20,
					"X_b_one": 10,
					"X_b_two": 20,
					"Y_a_one": 10,
					"Y_a_two": 20,
					"Y_b_one": 10,
					"Y_b_two": 20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateNestedMap(tt.args.dict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateNestedMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
