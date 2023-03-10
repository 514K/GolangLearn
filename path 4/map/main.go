package main

import "fmt"

func main() {
	myMap := make(map[string]int)
	fmt.Printf("%v\n", myMap)
	myMap["Alex"] = 22
	myMap["Sas"] = 31
	fmt.Printf("%v\n", myMap)

	for key, val := range myMap {
		fmt.Printf("myMap[\"%v\"] = %v\n", key, val)
	}

	authors := map[string]string{
		"J.R.": "Harry Potter",
	}

	a, b := authors["saske"]
	fmt.Printf("author is %v and its exist? %v\n", a, b)

	delete(authors, "w.R.")
	delete(authors, "J.R.")
	fmt.Printf("%v\n", authors)

}
