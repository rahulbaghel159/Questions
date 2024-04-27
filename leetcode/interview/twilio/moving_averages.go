package main

import (
	"fmt"
)

/*
You are tasked with creating a Go function that calculates the Simple Moving Average (SMA) of a stream of numbers. The SMA is a commonly used statistical calculation in finance and time-series analysis, and it provides a smoothed average of a set of values over a specified window size.
Function Signature:

func worker(in <-chan float64, out chan<- float64, windowSize int)

*/

func main() {
	in, out := make(chan float64), make(chan float64)
	windowSize := 3

	go worker(in, out, windowSize)

	for i := 1; i <= 5; i++ {
		in <- float64(i)
		select {
		case b := <-out:
			fmt.Println("b from main", b)
		}
	}

	close(in)
	close(out)
}

/*
  in -> 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
  windowSize -> 1, 1, 2, 2, 3, 1, 5

  array -> dynamically (not)
  linked list ->

    count = 1, window = [1], ma = 1
    count = 2, window = [1, 2], ma = 1.5

    Time := O(1), Space := O(windowSize)
*/

func worker(in <-chan float64, out chan<- float64, windowSize int) {
	if windowSize <= 0 {
		return
	}

	var (
		window []float64
		sum    float64
		count  int
	)

	for {
		select {
		case a := <-in:
			count++
			if count > windowSize {
				sum -= window[0]
				window = window[1:]
			}
			window = append(window, a)
			sum += a
			ma := sum / float64(min(count, windowSize))
			out <- ma
			// case :
			//     return
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
