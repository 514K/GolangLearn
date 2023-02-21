package main

import (
	"errors"
	"fmt"
)

func main() {
	x := 5
	y := 2

	ans, err := square(float32(x), float32(y))
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", ans)
	}

	ans, err = per(float32(x), float32(y))
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", ans)
	}

	// fmt.Printf()
}

func square(x float32, y float32) (float32, error) {
	if x < 0 || y < 0 {
		return float32(0), errors.New("Стороны должны быть больше нуля")
	}

	fmt.Printf("%v\n\n\n\n", x)
	// err := nil

	// x = float32(x)

	// if err != nil {
	// 	return float32(0), errors.New("x должен быть числом")
	// }

	// y = float32(y)

	// if err != nil {
	// 	return float32(0), errors.New("y должен быть числом")
	// }

	return x * y, nil
}

func per(x float32, y float32) (float32, error) {
	if x < 0 || y < 0 {
		return float32(0), errors.New("Стороны должны быть больше нуля")
	} else {
		x, y = float32(x), float32(y)
		return (x + y) * 2, nil
	}
}
