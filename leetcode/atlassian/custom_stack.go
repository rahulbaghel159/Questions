package atlassian

// https://leetcode.com/problems/design-a-stack-with-increment-operation/
type CustomStack struct {
	arr     []int
	maxSize int
}

func Constructor1(maxSize int) CustomStack {
	return CustomStack{
		arr:     make([]int, 0),
		maxSize: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.arr) != this.maxSize {
		this.arr = append(this.arr, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.arr) == 0 {
		return -1
	}

	top := this.arr[len(this.arr)-1]
	this.arr = this.arr[:len(this.arr)-1]

	return top
}

func (this *CustomStack) Increment(k int, val int) {
	i := 0
	for ; i < len(this.arr) && i < k; i++ {
		this.arr[i] = this.arr[i] + val
	}
}
