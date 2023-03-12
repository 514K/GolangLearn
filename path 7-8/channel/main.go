package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Printf("%v\n", t.Format(time.RFC3339))

	result1 := make(chan int)
	result2 := make(chan int)

	go calcSome(2000, result1)

	go calcSome(1000, result2)

	fmt.Printf("%v\n", <-result1)
	fmt.Printf("%v\n", <-result2)

	fmt.Printf("End in: %v\n", time.Since(t))

	// ch := make(chan int)
	// ch <- 1
	// num := <-ch
	// fmt.Printf("%v\n", num)
}

func calcSome(n int, res chan int) {
	t := time.Now()

	result := 0
	for i := 0; i <= n; i++ {
		result++
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Res: %v; Need time: %v\n", result, time.Since(t))
	res <- result
}
