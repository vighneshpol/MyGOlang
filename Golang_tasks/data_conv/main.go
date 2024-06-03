package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 42
	fmt.Printf("Type of num is %T\n", num)

	dataConv := float64(num)
	fmt.Printf("Type of num is %T\n", dataConv)

	dataConv = 123
	str:= strconv.Itoa(num)
	fmt.Printf("Type of str is %T\n", str)
}
