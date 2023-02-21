package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!!!")
	fmt.Printf("100%%\n")
	test := "13TestTFdGdbVV"
	fmt.Printf("%X\n", test)
	print(hello("Hello", "World!!!\n"))
}

func hello(str1 string, str2 string) string {
	print(str1 + " " + str2)
	if (str1 + " " + str2) == "Hello world!!!" {
		return "Hello"
	}
	return str1 + " " + str2
}
