package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	for i := 1; i <= n; i++ {
		fmt.Printf("%v\n", math.Pow(2, float64(i)))
	}

	firstSlice := make([]int, 3, 5)
	fmt.Printf("%v %v\n", len(firstSlice), cap(firstSlice))
	fmt.Printf("%v\n", firstSlice)
	num := copy(firstSlice, []int{1, 2, 3, 4, 5, 6, 7})
	fmt.Printf("%v\n", firstSlice)

	fmt.Printf("%v\n", num)
}
