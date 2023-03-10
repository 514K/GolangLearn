package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%v\n", t.Format(time.RFC3339))

	go calcSome(3000)

	go calcSome(1000)

	go load()

	time.Sleep(time.Second * 50)

	fmt.Printf("End in: %v\n", time.Since(t))

}

func calcSome(n int) {
	t := time.Now()
	result := 0
	for i := 0; i <= n; i++ {
		result++
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Res: %v; Need time: %v\n", result, time.Since(t))
}

func load() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
