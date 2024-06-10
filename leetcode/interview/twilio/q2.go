package main

import (
	"fmt"
	"sync"
)

func main() {
	producer()
}

type Task struct {
	name     string
	interval string
	handler  func() error
	error    string
}

func producer() {
	ch := make(chan (Task))

	wg := sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go consumer(ch, wg)
	}

	fn := func() error {
		fmt.Println("Hello")
		return nil
	}

	ch <- Task{
		name:     "One",
		interval: "10:00:00",
		handler:  fn,
	}

	ch <- Task{
		name:     "Two",
		interval: "10:00:15",
		handler:  fn,
	}

	ch <- Task{
		name:     "Three",
		interval: "10:00:30",
		handler:  fn,
	}

	close(ch)

	wg.Wait()
}

func consumer(ch <-chan Task, wg sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case n := <-ch:
			fmt.Println(n.name)
			err := n.handler()
			n.error = err.Error()
		}
	}
}
