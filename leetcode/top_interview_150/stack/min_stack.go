package stack

type MinStack struct {
	arr [][]int
}

func Constructor() MinStack {
	return MinStack{arr: make([][]int, 0)}
}

func (this *MinStack) Push(val int) {
	prev_m := val
	if len(this.arr) != 0 {
		prev_m = this.arr[len(this.arr)-1][1]
	}
	this.arr = append(this.arr, []int{val, min(val, prev_m)})
}

func (this *MinStack) Pop() {
	if len(this.arr) != 0 {
		this.arr = this.arr[:len(this.arr)-1]
	}
}

func (this *MinStack) Top() int {
	res := 0

	if len(this.arr) != 0 {
		res = this.arr[len(this.arr)-1][0]
	}

	return res
}

func (this *MinStack) GetMin() int {
	if len(this.arr) != 0 {
		return this.arr[len(this.arr)-1][1]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
