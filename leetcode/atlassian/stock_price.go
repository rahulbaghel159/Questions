package atlassian

import "container/heap"

// https://leetcode.com/problems/stock-price-fluctuation/
type StockPrice struct {
	minHeap      StockMinHeap
	maxHeap      StockMaxHeap
	latest       int
	priceHistory map[int]int
}

func Constructor6() StockPrice {
	minHeap := StockMinHeap{}
	maxHeap := StockMaxHeap{}

	heap.Init(&minHeap)
	heap.Init(&maxHeap)

	return StockPrice{
		minHeap:      minHeap,
		maxHeap:      maxHeap,
		priceHistory: make(map[int]int),
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	heap.Push(&this.minHeap, StockRecord{
		price:     price,
		timestamp: timestamp,
	})

	heap.Push(&this.maxHeap, StockRecord{
		price:     price,
		timestamp: timestamp,
	})

	this.priceHistory[timestamp] = price
	this.latest = max(this.latest, timestamp)
}

func (this *StockPrice) Current() int {
	return this.priceHistory[this.latest]
}

func (this *StockPrice) Maximum() int {
	record := heap.Pop(&this.maxHeap).(StockRecord)

	for this.maxHeap.Len() > 0 && record.price != this.priceHistory[record.timestamp] {
		record = heap.Pop(&this.maxHeap).(StockRecord)
	}

	heap.Push(&this.maxHeap, record)

	return record.price
}

func (this *StockPrice) Minimum() int {
	record := heap.Pop(&this.minHeap).(StockRecord)

	for this.minHeap.Len() > 0 && record.price != this.priceHistory[record.timestamp] {
		record = heap.Pop(&this.minHeap).(StockRecord)
	}

	heap.Push(&this.minHeap, record)

	return record.price
}

type StockRecord struct {
	price, timestamp int
}

type StockMaxHeap []StockRecord

func (h StockMaxHeap) Len() int {
	return len(h)
}

func (h StockMaxHeap) Less(i, j int) bool {
	return h[i].price > h[j].price
}

func (h StockMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StockMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(StockRecord))
}

func (h *StockMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type StockMinHeap []StockRecord

func (h StockMinHeap) Len() int {
	return len(h)
}

func (h StockMinHeap) Less(i, j int) bool {
	return h[i].price < h[j].price
}

func (h StockMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StockMinHeap) Push(x interface{}) {
	*h = append(*h, x.(StockRecord))
}

func (h *StockMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
