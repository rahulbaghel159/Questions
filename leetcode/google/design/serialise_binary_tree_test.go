package design

import (
	"reflect"
	"testing"
)

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 2},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 4},
						Right: &TreeNode{Val: 5},
					},
				},
			},
			this: &Codec{},
			want: "1,2,null,null,3,4,null,null,5,null,null,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want *TreeNode
	}{
		{
			name: "Case 1",
			args: args{
				data: "1,2,null,null,3,4,null,null,5,null,null,",
			},
			this: &Codec{},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Codec.deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
