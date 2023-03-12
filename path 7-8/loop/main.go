package main

import (
	"fmt"
	"time"
)

func main() {
	num := make(chan int)

	go func() {
		num <- 13
	}()

	time.Sleep(time.Millisecond * 100)

	select {
	case n := <-num:
		fmt.Printf("%v\n", n)
	default:
		fmt.Printf("Thats all")
	}

}
