package main

import (
	"fmt"
	"time"
)

func factorial(n int, ch chan bool) int {
	if n == 1 {
		fmt.Printf("sas\n")
		ch <- true
		close(ch)
		return 1

	}
	// time.Sleep(time.Millisecond * 100)
	// ch <- factorial(n-1, ch) * n

	return factorial(n-1, ch) * n
}

func load() {

	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func main() {
	// ch := make(chan bool)
	// ch <- false

	result := make(chan bool)

	go factorial(5000000, result)
	go load()

	fmt.Printf("%v\n", <-result)

	// time.Sleep(time.Second * 10)
	// for {
	// 	num, ok := <-result
	// 	fmt.Printf("%v; %v\n", num, ok)

	// 	if !ok {
	// 		break
	// 	}
	// }
}
