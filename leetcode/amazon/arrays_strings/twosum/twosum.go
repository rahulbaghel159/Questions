package twosum

//https://leetcode.com/explore/interview/card/amazon/76/array-and-strings/508/
func twoSum(nums []int, target int) []int {
	//prepare result array
	result := make([]int, 2)

	//put array into map while checking target-num exists or not
	numbers := make(map[int]int)
	for index, num := range nums {
		index2, ok := numbers[target-num]
		if ok {
			if index < index2 {
				result[0] = index
				result[1] = index2
			} else {
				result[0] = index2
				result[1] = index
			}
			break
		} else {
			numbers[num] = index
		}
	}
	return result
}
