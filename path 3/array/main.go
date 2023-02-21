package main

import "fmt"

func main() {
	someList := [3]int{
		3, 1, 2,
	}

	var otherList [4]float64 = [4]float64{3.1, 3.2, 3.3, 3.4}

	fmt.Printf("%d\n", len(someList))
	fmt.Printf("%d\n", len(otherList))

	randomList := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("%v\n", len(randomList))

}
