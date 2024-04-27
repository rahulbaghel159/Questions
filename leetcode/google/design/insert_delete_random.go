package design

import (
	"math/rand"
)

type RandomizedSet struct {
	dict map[int]int
	arr  []int
}

func Constructor4() RandomizedSet {
	return RandomizedSet{
		dict: make(map[int]int, 0),
		arr:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.dict[val] = len(this.arr) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dict[val]; !ok {
		return false
	}

	index, last := this.dict[val], len(this.arr)-1
	this.arr[index], this.arr[last] = this.arr[last], this.arr[index]

	this.dict[this.arr[index]] = index
	this.arr = this.arr[:last]
	delete(this.dict, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}
