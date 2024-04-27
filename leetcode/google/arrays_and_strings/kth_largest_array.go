package google

import "math/rand"

// https://leetcode.com/explore/interview/card/google/59/array-and-strings/361/
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -999999999999
	}

	return quickSelect(nums, k)
}

// func kth_largest(arr []int, low, high, k int) int {
// 	partition := partition(arr, low, high)

// 	if partition == k-1 {
// 		return arr[partition]
// 	} else if partition > k-1 {
// 		return kth_largest(arr, low, partition-1, k)
// 	} else {
// 		return kth_largest(arr, partition+1, high, k)
// 	}
// }

// func partition(arr []int, low, high int) int {
// 	random := rand.Intn(high+1-low) + low
// 	arr[random], arr[high] = arr[high], arr[random]

// 	pivot, pivotLoc := arr[high], low

// 	for i := low; i <= high; i++ {
// 		if arr[i] < pivot {
// 			arr[i], arr[pivotLoc] = arr[pivotLoc], arr[i]
// 			pivotLoc++
// 		}
// 	}

// 	arr[high], arr[pivotLoc] = arr[pivotLoc], arr[high]

// 	return pivotLoc
// }

func quickSelect(nums []int, k int) int {
	pivot := nums[rand.Intn(len(nums))]

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
