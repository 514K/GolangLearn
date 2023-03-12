package main

import (
	"fmt"
)

func main() {
	nums := make(chan int)

	go genNums(1000, nums)

	// for n := range nums {
	// 	fmt.Printf("%v\n", n)
	// }
	for {
		num, ok := <-nums
		fmt.Printf("%v; %v\n", num, ok)

		if !ok {
			break
		}
	}
}

func genNums(n int, res chan int) {
	for i := 0; i < n; i++ {
		res <- i * 2
		// time.Sleep(time.Millisecond * 100)
	}
	close(res)
}
