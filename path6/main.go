package main

import (
	"fmt"

	"storage"
)

// "storage"

func main() {
	sotr := *new(storage.Sotrudnik)
	ms := storage.NewMemoryStorage()
	fmt.Printf("%v\n%v\n", sotr, ms)
	// ms.insert()
}
