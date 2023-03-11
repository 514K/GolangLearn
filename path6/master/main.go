package main

import (
	"fmt"

	st "stor/storage"
)

// "storage"

func main() {
	sotr := *new(st.Sotrudnik)
	ms := st.NewMemoryStorage()
	fmt.Printf("%v\n%v\n", sotr, ms)
	// ms.insert()
}
