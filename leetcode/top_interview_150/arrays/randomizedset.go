package arrays

import (
	"math/rand"
)

//https://leetcode.com/problems/insert-delete-getrandom-o1

type RandomizedSet struct {
	hashMap map[int]int
	slice   []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		hashMap: make(map[int]int),
		slice:   make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hashMap[val]; ok {
		return false
	}

	this.hashMap[val] = len(this.slice)
	this.slice = append(this.slice, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.hashMap[val]; !ok {
		return false
	}

	n, i := len(this.slice), this.hashMap[val]
	this.slice[i], this.slice[n-1] = this.slice[n-1], this.slice[i]
	this.hashMap[this.slice[i]] = i
	this.slice = this.slice[:n-1]
	delete(this.hashMap, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.slice[rand.Intn(len(this.slice))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
