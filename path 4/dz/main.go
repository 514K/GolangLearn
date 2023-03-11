package main

import (
	"fmt"
	"math"
)

type circle struct {
	r  float64
	d  float64
	ln float64
	s  float64
}

func newCircle(rad float64) circle {
	return circle{
		r:  rad,
		d:  2 * rad,
		ln: 2 * math.Pi * rad,
		s:  2 * math.Pi * math.Pow(rad, 2),
	}
}

func main() {
	myCircle := newCircle(math.Sqrt(1 / math.Pi))
	fmt.Printf("%v\n", myCircle.ln)
	fmt.Printf("%v\n", myCircle.s)
}
