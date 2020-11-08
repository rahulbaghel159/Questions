package kthlargestelement

//Method 1: Using Max Heap
func findKthLargest(nums []int, k int) int {
	//construct max heap and call getMax k times
	heap := createHeap(nums)
	for i := 1; i < k; i++ {
		_, heap = extractMax(heap)
	}
	return getMax(heap)
}

func createHeap(nums []int) []int {
	var heap []int
	//insert each element at end and call heapify
	for _, val := range nums {
		heap = append(heap, val)
		upheapify(heap, len(heap)-1)
	}

	return heap
}

func extractMax(heap []int) (int, []int) {
	num := heap[0]

	heap[0] = heap[len(heap)-1]
	heap = heap[:len(heap)-1]

	downheapify(heap, 0)

	return num, heap
}

func downheapify(heap []int, index int) {
	largest := index
	l := 2*largest + 1
	r := 2*largest + 2

	if l < len(heap) && heap[l] > heap[largest] {
		largest = l
	}

	if r < len(heap) && heap[r] > heap[largest] {
		largest = r
	}

	if largest != index {
		temp := heap[index]
		heap[index] = heap[largest]
		heap[largest] = temp

		downheapify(heap, largest)
	}
}

func getMax(heap []int) int {
	return heap[0]
}

func upheapify(heap []int, index int) {
	parent := (index - 1) / 2

	if heap[index] > heap[parent] {
		temp := heap[index]
		heap[index] = heap[parent]
		heap[parent] = temp

		upheapify(heap, parent)
	}
}

func isLeaf(heap []int, index int) bool {
	if index >= len(heap)/2 && index < len(heap) {
		return true
	}
	return false
}

//Method 2: Using Quickselect
func findKthLargest2(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(arr []int, low, high, k int) int {
	partition := partition(arr, low, high)

	if partition == k {
		return arr[partition]
	}

	if partition < k {
		return quickSelect(arr, partition+1, high, k)
	}

	return quickSelect(arr, low, partition-1, k)
}

//partition function returns pivot
func partition(arr []int, low, high int) int {
	//select last element as pivot, and first element as pivot location
	pivot := arr[high]
	pivotLoc := low

	//for each element smaller than pivot at left, swap with element at pivot location and increment pivot location
	for i := low; i <= high; i++ {
		if arr[i] < pivot {
			temp := arr[i]
			arr[i] = arr[pivotLoc]
			arr[pivotLoc] = temp

			pivotLoc++
		}
	}

	//swap pivot with final pivot location
	temp := arr[high]
	arr[high] = arr[pivotLoc]
	arr[pivotLoc] = temp

	return pivotLoc
}
