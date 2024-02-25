package heap

import "math/rand"

// https://leetcode.com/problems/kth-largest-element-in-an-array
// func findKthLargest(nums []int, k int) int {
// 	if len(nums) < k {
// 		return -1
// 	}

// 	h := &MinHeap{}
// 	for _, num := range nums {
// 		heap.Push(h, num)
// 		if len(*h) > k {
// 			heap.Remove(h, 0)
// 		}
// 	}

// 	return heap.Pop(h).(int)
// }

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, k)
}

func quickSelect(nums []int, k int) int {
	pivotIndex := rand.Intn(len(nums))
	pivot := nums[pivotIndex]

	left, mid, right := make([]int, 0), make([]int, 0), make([]int, 0)

	for _, num := range nums {
		if num > pivot {
			left = append(left, num)
		} else if num < pivot {
			right = append(right, num)
		} else {
			mid = append(mid, num)
		}
	}

	if k <= len(left) {
		return quickSelect(left, k)
	}

	if len(left)+len(mid) < k {
		return quickSelect(right, k-len(left)-len(mid))
	}

	return pivot
}
