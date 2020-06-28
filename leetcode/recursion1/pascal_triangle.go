package recursion1

//https://leetcode.com/explore/learn/card/recursion-i/251/scenario-i-recurrence-relation/3234/

func getRow(rowIndex int) []int {
	//breaking condition
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}

	//recursing
	arr := getRow(rowIndex - 1)

	//processing
	result := make([]int, rowIndex+1)
	result[0] = 1
	result[rowIndex] = 1
	i := 1
	for i < rowIndex {
		result[i] = arr[i-1] + arr[i]
		i++
	}
	return result
}
