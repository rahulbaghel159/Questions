package binarysearch

// https://leetcode.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nA, nB := len(nums1), len(nums2)
	n := nA + nB

	if n%2 == 1 {
		return float64(solve(nums1, nums2, n/2, 0, nA-1, 0, nB-1))
	} else {
		return (float64(solve(nums1, nums2, n/2, 0, nA-1, 0, nB-1)) + float64(solve(nums1, nums2, n/2-1, 0, nA-1, 0, nB-1))) / 2
	}
}

func solve(A, B []int, k, a_start, a_end, b_start, b_end int) int {
	if a_start > a_end {
		return B[k-a_start]
	}

	if b_start > b_end {
		return A[k-b_start]
	}

	a_mid, b_mid := (a_start+a_end)/2, (b_start+b_end)/2
	a_val, b_val := A[a_mid], B[b_mid]

	if a_mid+b_mid < k {
		if a_val < b_val {
			return solve(A, B, k, a_mid+1, a_end, b_start, b_end)
		} else {
			return solve(A, B, k, a_start, a_end, b_mid+1, b_end)
		}
	} else {
		if a_val > b_val {
			return solve(A, B, k, a_start, a_mid-1, b_start, b_end)
		} else {
			return solve(A, B, k, a_start, a_end, b_start, b_mid-1)
		}
	}
}
