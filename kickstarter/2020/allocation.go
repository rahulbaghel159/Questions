package twentytwenty

func MaxHouses(arr []int, b int) int {
	//sort array: Counting Sort
	c := make([]int, 1000)
	for i := 0; i < len(arr); i++ {
		c[arr[i]] = 1
	}
	i := 0
	for index, v := range c {
		if v == 1 {
			arr[i] = index
			i++
		}
	}

	//loop from beginning counting until sum becomes larger than b
	count, sum := 0, 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > b {
			break
		}
		count++
	}

	return count
}
