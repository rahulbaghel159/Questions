package random

//LRUCache defines internal data structures for cache
type LRUCache struct {
	capacity int
	data     map[int]int
	lruCache list
}

type list struct {
	data int
	next *list
}

//Constructor initialises cache with given capacity
func Constructor(capacity int) LRUCache {
	data := make(map[int]int)

	return LRUCache{
		capacity: capacity,
		data:     data,
	}
}

//Get fetches element from cache
func (lru *LRUCache) Get(key int) int {
	var (
		result int
		ok     bool
	)

	if result, ok = lru.data[key]; !ok {
		return -1
	}

	//update usage in list

	return result
}

//Put adds element in cache
func (lru *LRUCache) Put(key int, value int) {
	var (
		count   int
		element list
	)

	//add element at tail
	element = lru.lruCache
	for element.next != nil {
		count++
	}

	//if cache is already full, evict an element
	if count > lru.capacity {

	}

	element.next = &list{
		data: key,
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
