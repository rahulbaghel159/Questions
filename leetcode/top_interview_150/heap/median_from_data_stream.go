package heap

import "container/heap"

// https://leetcode.com/problems/find-median-from-data-stream/
type MedianFinder struct {
	lo MaxHeap
	hi MinHeap
}

func Constructor() MedianFinder {
	m := MedianFinder{
		lo: MaxHeap{},
		hi: MinHeap{},
	}

	heap.Init(&m.lo)
	heap.Init(&m.hi)

	return m
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.lo, num)

	l := heap.Pop(&this.lo)
	heap.Push(&this.hi, l)

	if this.hi.Len() > this.lo.Len() {
		r := heap.Pop(&this.hi)
		heap.Push(&this.lo, r)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() == this.hi.Len() {
		l, r := heap.Pop(&this.lo).(int), heap.Pop(&this.hi).(int)
		heap.Push(&this.lo, l)
		heap.Push(&this.hi, r)

		return (float64(l) + float64(r)) / 2
	}

	l := heap.Pop(&this.lo).(int)
	heap.Push(&this.lo, l)

	return float64(l)
}
