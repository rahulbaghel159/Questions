package main

import "fmt"

//1<0=false=0, 0<=0=true=0
func toggle(input int) bool {
	return input <= 0
}

func main() {
	fmt.Println(toggle(1))
	fmt.Println(toggle(0))
}
