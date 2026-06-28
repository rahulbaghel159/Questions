package main

import "fmt"

// https://leetcode.com/explore/learn/card/fun-with-arrays/525/introduction/3245/
func duplicateZeros(arr []int) {
	if len(arr) == 0 {
		return
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			if i+1 < len(arr) {
				arr[i+1] = 0
			}
			i++
		}
		fmt.Println(arr)
	}
}
