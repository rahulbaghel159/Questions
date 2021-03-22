package main

import "fmt"

/*
1, 32, 5, 6, 3, 4, 9
arr1=1, 3, 5
arr2=32, 6, 9
*/

func main() {
	arr := []int{1, 32, 5, 6, 3, 4, 9}
	N := 5
	var arr1, arr2, arr3 []int

	for _, val := range arr {
		if val <= N {
			arr3 = []int{}
			if len(arr1) == 0 {
				arr1 = append(arr1, val)
			} else {
				for _, v := range arr1 {
					if v < val {
						arr3 = append(arr3, v)
					} else {
						arr3 = append(arr3, val)
					}
				}
				arr1 = arr3
			}
		} else {
			arr2 = append(arr2, val)
		}
	}
	arr1 = append(arr1, arr2...)
	fmt.Println(arr1)
}
