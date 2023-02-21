package main

import "fmt"

func main() {
	someList := [3]int{4, 7, 2}
	fmt.Printf("%v\n", someList)

	for i := 0; i < len(someList); i++ {
		fmt.Printf("%v; ", someList[i])
	}
	fmt.Print("\n")

	for i, v := range someList {
		if i == 1 {
			continue
		}
		fmt.Printf("someList[%v] = %v; ", i, v)
	}
	fmt.Print("\n")

	emptyList := [3]int{}
	fmt.Printf("%v\n", emptyList)

	emptyList = fillArray(emptyList)
	fmt.Printf("%v\n", emptyList)

}

func fillArray(arr [3]int) [3]int {
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Printf("%v\n", arr)
	return arr
}
