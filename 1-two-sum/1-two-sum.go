func twoSum(nums []int, target int) []int {
    for index1, num1:=range nums{
        for index2, num2:=range nums[index1+1:]{
            if (num1+num2)==target{
                return []int{index1, index2+index1+1}
            }
        }
    }
    return []int{}
}