package validsplit

func minimumIndex(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	freq := make(map[int][]int)
	for i, num := range nums {
		if _, ok := freq[num]; !ok {
			freq[num] = make([]int, 0)
		}
		freq[num] = append(freq[num], i)
	}

	dominant, maxLen := -1, -1
	for k, v := range freq {
		if len(v) > maxLen {
			maxLen = len(v)
			dominant = k
		}
	}

	indexes := freq[dominant]
	count1 := 0
	for i := 0; i < len(indexes); i++ {
		count1++
		count2 := len(indexes) - count1
		if count1*2 > indexes[i]+1 && count2*2 > len(nums)-indexes[i]-1 {
			return indexes[i]
		}
	}

	return -1
}
