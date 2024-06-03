// write a simple function in go program
package main

import "fmt"

func simpleFunction() {
	fmt.Println("We are inside a simple function")

}

func add(a, b int) (result int) {
	result = a + b
	return
}
func main() {
	fmt.Println("We are inside main function")

	simpleFunction()
	ans := add(10, 12)
	fmt.Println("Addition of two numbers is ", ans)
}
