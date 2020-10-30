package recursion2

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2

	left := make([]int, mid)
	right := make([]int, len(nums)-mid)

	for i := 0; i < len(nums); i++ {
		if i < mid {
			left[i] = nums[i]
		} else {
			right[i-mid] = nums[i]
		}
	}

	return mergeSortedArray(sortArray(left), sortArray(right))
}

func mergeSortedArray(arr1 []int, arr2 []int) []int {
	var (
		index, i, j, small int
	)

	arr := make([]int, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			small = arr1[i]
			i++
		} else {
			small = arr2[j]
			j++
		}
		arr[index] = small
		index++
	}

	for ; i < len(arr1); i++ {
		arr[index] = arr1[i]
		index++
	}

	for ; j < len(arr2); j++ {
		arr[index] = arr2[j]
		index++
	}

	return arr
}
