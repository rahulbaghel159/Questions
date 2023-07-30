package lrucache

import "fmt"

//https://leetcode.com/problems/lru-cache/description/

type DoubleListNode struct {
	key   int
	value int
	next  *DoubleListNode
	prev  *DoubleListNode
}

func (node *DoubleListNode) print() {
	if node == nil {
		return
	}
	fmt.Print(node.key)
	for node.next != nil {
		fmt.Print(",", node.next.key)
		node = node.next
	}
	fmt.Println()
}

type LRUCache struct {
	dic      map[int]*DoubleListNode
	head     *DoubleListNode
	tail     *DoubleListNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &DoubleListNode{value: -1}
	tail := &DoubleListNode{value: -1}

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		dic:      make(map[int]*DoubleListNode),
		head:     head,
		tail:     tail,
	}
}

func (cache *LRUCache) Get(key int) int {
	if _, ok := cache.dic[key]; !ok {
		return -1
	}

	node := cache.dic[key]
	cache.remove(node)
	cache.add(node)

	return node.value
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.dic[key]; ok {
		cache.remove(node)
	}

	node := &DoubleListNode{
		key:   key,
		value: value,
	}

	cache.dic[key] = node
	cache.add(node)

	if len(cache.dic) > cache.capacity {
		nodeToDelete := cache.head.next
		cache.remove(nodeToDelete)
		delete(cache.dic, nodeToDelete.key)
	}
}

//adds node on tail
func (cache *LRUCache) add(node *DoubleListNode) {
	previousEnd := cache.tail.prev
	previousEnd.next = node
	node.prev = previousEnd
	node.next = cache.tail
	cache.tail.prev = node
}

//delete node
func (cache *LRUCache) remove(node *DoubleListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
