package main

import (
	"fmt"
)

func main() {
	age := 29
	name := "Max"
	height := 6.234

	fmt.Println("age:", age, "height:", height, "name:", name)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.3f\n", height)
	fmt.Printf("Type of the age is %T\n", height)

}
