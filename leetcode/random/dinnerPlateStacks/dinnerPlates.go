package dinnerplatestacks

import "fmt"

//DinnerPlates acts as a abstraction for dinner plates stack
type DinnerPlates struct {
	platesArr               [][]int
	maxCapacity, last, next int
}

//Constructor used for creating DinnerPlates reference
func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		platesArr:   [][]int{},
		maxCapacity: capacity,
		last:        0,
		next:        0,
	}
}

//Push used for pushing plate in leftmost stack which is not full
func (plates *DinnerPlates) Push(val int) {
	var (
		found bool
	)

	if len(plates.platesArr) <= plates.next {
		plates.platesArr = append(plates.platesArr, []int{})
	}

	plates.platesArr[plates.next] = append(plates.platesArr[plates.next], val)

	if len(plates.platesArr[plates.next]) == plates.maxCapacity {
		for l := plates.next; l < len(plates.platesArr); l++ {
			if len(plates.platesArr[l]) < plates.maxCapacity {
				plates.next = l
				found = true
			}
		}

		if !found {
			plates.next = len(plates.platesArr)
		}
	}

	if plates.last <= plates.next {
		for l := len(plates.platesArr) - 1; l >= 0; l-- {
			if len(plates.platesArr[l]) > 0 {
				plates.last = l
			}
		}
	}
}

//Pop used for removing plate from rightmost non-empty stack
func (plates *DinnerPlates) Pop() int {
	if len(plates.platesArr) == 0 || len(plates.platesArr[plates.last]) == 0 {
		return -1
	}

	var (
		result int
	)

	result = plates.platesArr[plates.last][len(plates.platesArr[plates.last])-1]
	plates.platesArr[plates.last] = plates.platesArr[plates.last][:len(plates.platesArr[plates.last])-1]

	if len(plates.platesArr[plates.last]) == 0 {
		for l := len(plates.platesArr) - 1; l >= 0; l-- {
			if len(plates.platesArr[l]) > 0 {
				plates.last = l
				break
			}
		}
	}

	return result
}

//PopAtStack used for removing plate from given stack
func (plates *DinnerPlates) PopAtStack(index int) int {
	if index > len(plates.platesArr)-1 {
		return -1
	}

	var (
		result int
	)

	result = plates.platesArr[index][len(plates.platesArr[index])-1]
	plates.platesArr[index] = plates.platesArr[index][:len(plates.platesArr[index])-1]

	if plates.last <= index && len(plates.platesArr[index]) == 0 {
		for l := len(plates.platesArr) - 1; l >= 0; l-- {
			if len(plates.platesArr[l]) > 0 {
				plates.last = l
				break
			}
		}
	}

	if plates.next > index {
		plates.next = index
	}

	return result
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */

//PrintStack prints internal stack
func (plates *DinnerPlates) PrintStack() {
	fmt.Println(plates.platesArr)
}

//PrintLast prints internal last element
func (plates *DinnerPlates) PrintLast() {
	fmt.Println(plates.last)
}
