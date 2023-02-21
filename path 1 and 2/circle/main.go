package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	r := float32(3.4)

	ans, err := square(r)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", ans)
	}

	ans, err = per(r)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", ans)
	}
}

func square(r float32) (float32, error) {
	if r < 0 {
		return float32(0), errors.New("Радиус должен быть больше нуля")
	} else {
		return float32(math.Pi * math.Pow(float64(r), 2)), nil
	}
}

func per(r float32) (float32, error) {
	if r < 0 {
		return float32(0), errors.New("Радиус должен быть больше нуля")
	} else {
		return float32(2 * math.Pi * r), nil
	}
}
