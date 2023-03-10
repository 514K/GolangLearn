package main

import (
	"fmt"
	"strconv"
)

type item struct {
	name  string
	model string
	year  int
}

func (i item) getItemStringInfo() string {
	return i.name + " " + i.model + " " + strconv.Itoa(i.year)
}

func (i item) setName(n string) {
	i.name = n
}

func (i *item) setName2(n string) {
	i.name = n
}

func main() {
	phone := struct {
		name  string
		model string
		year  int
	}{}

	phone.year = 2020

	fmt.Printf("%+v\n", phone)

	type car struct {
		name  string
		model string
		year  int
	}

	myCar := car{
		name: "bmw",
	}

	fmt.Printf("%+v\n", myCar)

	myCar.name = "bmw"
	myCar.model = "x3"
	myCar.year = 2019
	fmt.Printf("my car %+v\n", myCar)

	otherCar := car{
		name:  "suzuki",
		model: "grand vitara",
		year:  2011,
	}
	fmt.Printf("other car %+v\n", otherCar)

	keyboard := item{
		name:  "keyboard",
		model: "bloody",
		year:  2023,
	}

	mouse := item{
		name:  "mouse",
		model: "logitech",
		year:  2022,
	}
	fmt.Printf("Keyboard %+v\n", keyboard)
	fmt.Printf("Mouse %+v\n", mouse)

	fmt.Printf("Monitor %+v\n", itemTest())
	fmt.Printf("Monitor %+v\n", itemTest().getItemStringInfo())

	keyboard.setName("kb")  //not work bs thats a copy of this object
	keyboard.setName2("kb") //work bs thats NOT a copy of this object
	fmt.Printf("Keyboard %+v\n", keyboard)
}

func itemTest() item {
	return item{
		name:  "acer",
		model: "n/a",
		year:  2020,
	}
}
