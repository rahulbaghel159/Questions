package design

import "math"

type StackNode struct {
	val, min_so_far int
}

type MinStack struct {
	arr []StackNode
}

func Constructor1() MinStack {
	return MinStack{
		arr: make([]StackNode, 0),
	}
}

func (this *MinStack) Push(val int) {
	prev_min := math.MaxInt

	if len(this.arr) > 0 {
		prev_min = this.arr[len(this.arr)-1].min_so_far
	}
	min_so_far := min(val, prev_min)

	this.arr = append(this.arr, StackNode{
		val:        val,
		min_so_far: min_so_far,
	})
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
}

func (this *MinStack) Top() int {
	return this.arr[len(this.arr)-1].val
}

func (this *MinStack) GetMin() int {
	return this.arr[len(this.arr)-1].min_so_far
}
