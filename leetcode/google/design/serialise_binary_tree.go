package design

import (
	"strconv"
	"strings"
)

// https://leetcode.com/explore/interview/card/google/65/design-4/3092/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor2() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return dfsSerialise(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s_arr := strings.Split(data, ",")
	return dfsDeserialise(&s_arr)
}

func dfsSerialise(node *TreeNode, s string) string {
	if node == nil {
		s += "null,"
	} else {
		s += strconv.Itoa(node.Val) + ","

		s = dfsSerialise(node.Left, s)
		s = dfsSerialise(node.Right, s)
	}

	return s
}

func dfsDeserialise(s_arr *[]string) *TreeNode {
	if len(*s_arr) == 0 {
		return nil
	}

	if (*s_arr)[0] != "null" {
		val, _ := strconv.Atoi((*s_arr)[0])

		node := &TreeNode{Val: val}
		*s_arr = (*s_arr)[1:]
		node.Left = dfsDeserialise(s_arr)
		node.Right = dfsDeserialise(s_arr)

		return node
	} else {
		*s_arr = (*s_arr)[1:]
		return nil
	}
}
