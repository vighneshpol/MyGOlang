package main

import "fmt"

func add() {
	a := 2
	b := 3
	result := a / b

	if result>1{
		fmt.Println("result is greater than 1")
	} else{
		fmt.Println("Result is less than 1")
	}
}

func main() {
	x := 10
	if x == 0 {
		fmt.Println("x could not be 0")

	} else {
		fmt.Println("X is not 0")
	}

	z := 11
	if z < 10 {
		fmt.Println("z values is less than 10")
	} else if z == 10 {
		fmt.Println("z values is equal to 10")
	} else {
		fmt.Println("z values is greater than 10")
	}
	add()
}
