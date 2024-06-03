package main

import "fmt"

//we dont need to call explicitly in our entire go program go handles execution implicitly
var globalVar int

func init() {
	globalVar = 42
	fmt.Println("init function called")
}

func main() {
	slcie := make([]int, 2, 3)
	fmt.Println("slice is:",slcie)
	fmt.Println("main function called")
	fmt.Println("Global variable:", globalVar)
}
