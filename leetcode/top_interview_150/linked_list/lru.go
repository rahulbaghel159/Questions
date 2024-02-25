package linkedlist

// https://leetcode.com/problems/lru-cache/
type DoubleLinkedListNode struct {
	key   int
	value int
	next  *DoubleLinkedListNode
	prev  *DoubleLinkedListNode
}

type LRUCache struct {
	dic      map[int]*DoubleLinkedListNode
	head     *DoubleLinkedListNode
	tail     *DoubleLinkedListNode
	capacity int
}

func (cache *LRUCache) add(node *DoubleLinkedListNode) {
	tail_prev := cache.tail.prev
	tail_prev.next = node
	node.prev = tail_prev
	node.next = cache.tail
	cache.tail.prev = node
}

func (cache *LRUCache) remove(node *DoubleLinkedListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func Constructor(capacity int) LRUCache {
	head, tail := &DoubleLinkedListNode{value: -1}, &DoubleLinkedListNode{value: -1}
	head.next = tail
	tail.prev = head

	return LRUCache{
		dic:      make(map[int]*DoubleLinkedListNode),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.dic[key]; !ok {
		return -1
	}

	node := this.dic[key]
	this.remove(node)
	this.add(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.dic[key]; ok {
		this.remove(node)
	}

	node := &DoubleLinkedListNode{
		key:   key,
		value: value,
	}

	this.dic[key] = node
	this.add(node)

	if len(this.dic) > this.capacity {
		nodeToDelete := this.head.next
		this.remove(nodeToDelete)
		delete(this.dic, nodeToDelete.key)
	}
}
