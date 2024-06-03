/*
Write a program to find largest number from given 3 number.
*/

package main

import "fmt"

func main() {

	var number1, number2, number3 int

	fmt.Print("\nEnter the Three Numbers to find Largest = ")
	fmt.Scanln(&number1)
	fmt.Scanln(&number2)
	fmt.Scanln(&number3)

	if number1 > number2 && number1 > number3 {
		fmt.Println(number1, "is large than ", number2, " and ", number3)
	} else if number2 > number1 && number2 > number3 {
		fmt.Println(number2, "is large than", number1, " and ", number3)
	} else if number3 > number1 && number3 > number2 {
		fmt.Println(number3, " is large tha  ", number1, " and ", number2)
	} else {
		fmt.Println("All of them are equal")
	}
}
