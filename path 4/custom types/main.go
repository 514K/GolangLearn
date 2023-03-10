package main

import "fmt"

type name string

func (n name) valid() bool {
	if len(n) >= 2 {
		return true
	} else {
		return false
	}
}

func main() {
	myName := name("Alex")
	notName := name("z")
	fmt.Printf("%v\n", myName.valid())
	fmt.Printf("%v\n", notName.valid())
}
