package main

import (
	"fmt"
)

type Person struct {
	name     string
	age      int
	hAddress Address //nested struct
}

type Address struct {
	street string
	city   string
}

func main() {
	newPerson := Person{
		name: "Vighnesh Pol",
		age:  25,
		hAddress: Address{
			street: "Main Road", 
			city:   "Bangalore",
		},
	}
	fmt.Println("Name:", newPerson.name)
	fmt.Println("Age:", newPerson.age)
	fmt.Println("Street:", newPerson.hAddress.street)
	fmt.Println("City:", newPerson.hAddress.city)
}
