package design

type LRUCache struct {
	head, tail *DoubleLinkedListNode
	dict       map[int]*DoubleLinkedListNode
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &DoubleLinkedListNode{key: -1}
	tail := &DoubleLinkedListNode{key: -1}

	head.next = tail
	tail.prev = head

	return LRUCache{
		head:     head,
		tail:     tail,
		dict:     make(map[int]*DoubleLinkedListNode, 0),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.dict[key]; !ok {
		return -1
	}

	node := this.dict[key]
	this.delete(node)
	this.insert(node)

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.dict[key]; !ok {
		if this.capacity == len(this.dict) {
			delete(this.dict, this.head.next.key)
			this.delete(this.head.next)
		}

		node := &DoubleLinkedListNode{
			key: key,
			val: value,
		}

		this.insert(node)
		this.dict[key] = node
	} else {
		node := this.dict[key]

		this.delete(node)

		node.val = value
		this.insert(node)
	}
}

type DoubleLinkedListNode struct {
	key, val int
	next     *DoubleLinkedListNode
	prev     *DoubleLinkedListNode
}

func (this *LRUCache) insert(node *DoubleLinkedListNode) {
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *LRUCache) delete(node *DoubleLinkedListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next, node.prev = nil, nil
}
