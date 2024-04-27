package atlassian

// https://leetcode.com/problems/design-hit-counter/
type HitCounter struct {
	total_hit int
	queue     *Queue
}

func Constructor2() HitCounter {
	return HitCounter{
		total_hit: 0,
		queue: &Queue{
			arr: make([]int, 0),
		},
	}
}

func (this *HitCounter) Hit(timestamp int) {
	this.queue.Enqueue(timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	for this.queue.Size() > 0 {
		if timestamp-this.queue.Peek() >= 300 {
			this.queue.Dequeue()
		} else {
			break
		}
	}

	return this.queue.Size()
}

type Queue struct {
	arr []int
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) Enqueue(x int) {
	q.arr = append(q.arr, x)
}

func (q *Queue) Dequeue() int {
	first := -1
	if q.Size() > 0 {
		first = q.arr[0]
	}
	if q.Size() > 1 {
		q.arr = q.arr[1:]
	} else {
		q.arr = []int{}
	}

	return first
}

func (q *Queue) Peek() int {
	first := -1
	if q.Size() > 0 {
		first = q.arr[0]
	}

	return first
}
