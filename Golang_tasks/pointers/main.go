package main

import "fmt"

/**
When a variable is paired with the * operator, 
that variable holds a memory address. 
When it is paired with the & operator, 
it returns the address at which the variable is held.
*/

func modifyValueByReference(a *int) {
	*a = *a+20
}
func main() {

	num := 2

	ptr := &num

	fmt.Println("ptr contains", ptr)
	fmt.Println("Val contains through pointer", *ptr)

	var pointer *int //Default pointer == nil
	if pointer == nil {
		fmt.Println("pointer is nil")
	}

	value := 5
	modifyValueByReference(&value)//passing adress using &
	fmt.Println("Value contains", value)
}
