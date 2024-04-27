package main

import (
	"fmt"
	"sync"
)

type Shared struct {
	arr   []int
	index int
	m     sync.Mutex
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg := &sync.WaitGroup{}

	s := &Shared{
		arr: arr,
	}
	ch, done := make(chan interface{}), make(chan interface{})

	wg.Add(1)
	go CustomPrint(done, ch, s, 1, wg)
	go CustomPrint(done, ch, s, 2, wg)
	go CustomPrint(done, ch, s, 3, wg)

	ch <- 0

	wg.Wait()
}

func CustomPrint(done chan interface{}, ch chan interface{}, shared *Shared, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ch:
			shared.m.Lock()
			fmt.Println(id, shared.arr[shared.index])
			shared.index++
			shared.m.Unlock()

			// time.Sleep(10000)

			ch <- 0

			shared.m.Lock()
			if shared.index == len(shared.arr) {
				close(done)
				break
			}
			shared.m.Unlock()
		case <-done:
			return
		}
	}
}
