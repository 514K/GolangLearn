package main

import (
	"errors"
	"fmt"
)

func main() {
	message, err := test(10)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf(message)
}

func test(year int) (string, error) {
	if year < 18 {
		return "", errors.New("Вы слишком молодой, Вам тут не рады!!!")
	}

	return "Добро пожаловать", nil
}
