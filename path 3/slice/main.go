package main

import "fmt"

func main() {
	someSlice := []int{1, 2, 3, 4, 5, 6}

	someSlice = append(someSlice, 7)

	fmt.Printf("Len %v\n", len(someSlice))
	fmt.Printf("Capacity %v\n", cap(someSlice))

	for i, v := range someSlice {
		fmt.Printf("someSlice[%v] = %v; ", i, v)
	}
	fmt.Printf("\n")

	emptySlice := []int{}
	emptySlice = someSlice[1:3]
	for i := range emptySlice {
		fmt.Printf("%v; ", emptySlice[i])
	}
	fmt.Print("\n")

	fmt.Printf("%v\n", emptySlice)

	wtf(emptySlice)
	fmt.Printf("%v\n", emptySlice)
	fmt.Printf("%v\n", someSlice)

	forwardSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%v\n", forwardSlice)

	forwardSlice = reverseSlice(forwardSlice)
	fmt.Printf("%v\n", forwardSlice)

}

func reverseSlice(arr []int) []int {
	var tmpSlice = []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		tmpSlice = append(tmpSlice, arr[i])
	}
	// fmt.Printf("%v\n", tmpSlice)
	// arr = tmpSlice
	// fmt.Printf("%v\n", arr)
	return tmpSlice
}

func wtf(arr []int) {
	arr[0] = 1337
}
