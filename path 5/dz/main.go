package main

import (
	"fmt"
	"math"
)

type figure interface {
	getPerimetr() float64
	getSquare() float64
}

type circle struct {
	radius float64
}

func newCircle(r float64) *circle {
	return &circle{radius: r}
}

func (c *circle) getPerimetr() float64 {
	return 2 * math.Pi * c.radius
}

func (c *circle) getSquare() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type rect struct {
	width  float64
	height float64
}

func newRect(w float64, h float64) *rect {
	return &rect{
		width:  w,
		height: h,
	}
}

func (r *rect) getPerimetr() float64 {
	return 2 * (r.width + r.height)
}

func (r *rect) getSquare() float64 {
	return r.width * r.height
}

func main() {
	myFig := *new(figure)

	fmt.Printf("%v\n", myFig)

	myFig = newCircle(5)
	fmt.Printf("%v\n", myFig.getPerimetr())
	fmt.Printf("%v\n", myFig.getSquare())

	myFig = newRect(2, 3)
	fmt.Printf("%v\n", myFig.getPerimetr())
	fmt.Printf("%v\n", myFig.getSquare())
}
