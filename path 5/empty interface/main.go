package main

import "fmt"

func main() {
	printType("sas")
	printType(4)
	printType(map[int]float64{})

	printType2("sas")
	printType2(4)
	printType2(map[int]float64{})
}

func printType(val interface{}) {
	if _, err := val.(int); err {
		fmt.Println("int")
	} else if _, err := val.(string); err {
		fmt.Println("string")
	} else {
		fmt.Println("not int and string")
	}
}

func printType2(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("not int and string")
	}
}
