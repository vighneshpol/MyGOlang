/*Write a program to generate the Fibonacci sequence starting from 0. The function accepts
any number as an argument (length of the output sequence) and generates the list of
Fibonacci sequence.

Fibonacci series using recursion*/

package main

import "fmt"

func Fibonacii(a int, b int, n int) {
	var sum int = 0
	if n > 0 {
		sum = a + b
		fmt.Printf("%d ", sum)
		a = b
		b = sum
		Fibonacii(a, b, n-1)
	}
}

func main() {
	var a, b, n int

	a = 0
	b = 1

	fmt.Printf("Enter total range   :")
	fmt.Scanf("%d", &n)

	fmt.Printf("Fibonacii series of above range : ")
	fmt.Printf("%d\t%d ", a, b)

	Fibonacii(a, b, n-2)
	fmt.Printf("\n")
}
