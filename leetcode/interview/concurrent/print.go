package concurrent

import (
	"fmt"
	"sync"
	"time"
)

type Container struct {
	arr   []int
	index int
	mu    sync.Mutex
}

func print() {
	c := Container{
		arr:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		index: 0,
	}

	wg := sync.WaitGroup{}
	wg.Add(3)

	go printAlternate("first", &wg, &c)
	go printAlternate("second", &wg, &c)
	go printAlternate("third", &wg, &c)

	wg.Wait()
}

func printAlternate(label string, wg *sync.WaitGroup, c *Container) {
	defer wg.Done()

	for c.index < len(c.arr) {
		c.mu.Lock()
		if c.index >= len(c.arr) {
			break
		}
		fmt.Println(label)
		fmt.Println(c.arr[c.index])
		c.index++
		c.mu.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}
