package groupanagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// {
		// 	name: "Test Case 1",
		// 	args: args{
		// 		strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
		// 	},
		// 	want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		// },
		// {
		// 	name: "Test Case 2",
		// 	args: args{
		// 		strs: []string{""},
		// 	},
		// 	want: [][]string{{""}},
		// },
		// {
		// 	name: "Test Case 3",
		// 	args: args{
		// 		strs: []string{"a"},
		// 	},
		// 	want: [][]string{{"a"}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				s1: "eat",
				s2: "tea",
			},
			want: true,
		},
		{
			name: "Test Case 1",
			args: args{
				s1: "eat",
				s2: "bat",
			},
			want: false,
		},
		{
			name: "Test Case 1",
			args: args{
				s1: "test",
				s2: "estt",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
