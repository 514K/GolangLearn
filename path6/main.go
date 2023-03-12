package main

import (
	"fmt"

	"github.com/514K/GolangLearn/path6/storage"
	"github.com/514K/GolangLearn/path6/test"
)

func main() {
	sotr := *new(storage.Sotrudnik)
	ms := storage.NewMemoryStorage()
	fmt.Printf("%v\n%v\n", sotr, ms)
	test.SayHi()
}
