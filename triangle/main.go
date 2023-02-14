package main

import (
	"errors"
	"fmt"
	"math"
)

// Добавить треугольнику проверку на тругольность)
func main() {
	// var x, y, z float32 = 2, 2, 3
	x := float32(2)
	y := float32(2)
	z := float32(3)

	ans, err := square(x, y, z)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", ans)
	}

	ans, err = per(x, y, z)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", ans)
	}
}

func square(x float32, y float32, z float32) (float32, error) {
	if x < 0 || y < 0 || z < 0 {
		return float32(0), errors.New("Все стороны треугольника должны быть больше нуля")
	} else {
		p, _ := per(x, y, z)

		return float32(math.Sqrt(float64(p * (p - x) * (p - y) * (p - z)))), nil
	}
}

func per(x float32, y float32, z float32) (float32, error) {
	if x < 0 || y < 0 || z < 0 {
		return float32(0), errors.New("Все стороны треугольника должны быть больше нуля")
	} else {
		return x + y + z, nil
	}
}
