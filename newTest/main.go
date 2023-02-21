package main

import "fmt"

func main() {
	p1 := 11
	p2 := &p1
	p3 := *p2
	p4 := &p1
	a := new(int)
	*a = int(*p4)
	fmt.Printf("%v\n", p1)
	fmt.Printf("%v\n", p2)
	fmt.Printf("%v\n", p3)
	fmt.Printf("%v\n", p4)
	fmt.Printf("%v\n", *a)
}
