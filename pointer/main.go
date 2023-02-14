package main

import (
	"fmt"
)

func main() {
	x := 10
	p := &x
	fmt.Printf("%v\n", *p)
	x = 13
	fmt.Printf("%v\n", *p)
}
