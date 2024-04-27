package atlassian

import (
	"reflect"
	"testing"
)

func Test_findHighAccessEmployees(t *testing.T) {
	type args struct {
		access_times [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Case 1",
			args: args{
				access_times: [][]string{
					{"a", "0549"},
					{"b", "0457"},
					{"a", "0532"},
					{"a", "0621"},
					{"b", "0540"},
				},
			},
			want: []string{"a"},
		},
		{
			name: "Case 2",
			args: args{
				access_times: [][]string{
					{"d", "0002"},
					{"c", "0808"},
					{"c", "0829"},
					{"e", "0215"},
					{"d", "1508"},
					{"d", "1444"},
					{"d", "1410"},
					{"c", "0809"},
				},
			},
			want: []string{"c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findHighAccessEmployees(tt.args.access_times); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findHighAccessEmployees() = %v, want %v", got, tt.want)
			}
		})
	}
}
