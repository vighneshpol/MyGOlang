/* Create Calculator using OOPs concept which allows user to perform basic
operations(Addition, Subtraction, Multiplication and Division) */

package main

import "fmt"

func main() {
	var operator string
	var number1, number2 int
	fmt.Print("Please enter the operation and please give space before and after of the operator \n")
	fmt.Scan(&number1, &operator, &number2)
	// fmt.Print("Please enter Operator (+,-,/,x):")
	// fmt.Scanln(
	// // fmt.Print("Please enter Second number: ")
	// fmt.Scanln()

	output := 0
	switch operator {
	case "+":
		output = number1 + number2
	case "-":
		output = number1 - number2
	case "*":
		output = number1 * number2
	case "/":
		output = number1 / number2
		break
	default:
		fmt.Println("Wrong operator you have used")
	}
	fmt.Println("Answer :", number1, operator, number2, "=", output)
}
