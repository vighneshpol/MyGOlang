/*Write a program to triples each integer present in list [55, 22, 77,44, 58, 90, 30, 72]
which are even. Use Continue statement.*/

package main

import "fmt"

func main() {
	var Arrsize, i int

	fmt.Print("Enter the Array Size = ")
	fmt.Scan(&Arrsize)

	Arr := make([]int, Arrsize)

	fmt.Print("Enter the Array values = ")
	for i = 0; i < Arrsize; i++ {
		fmt.Scan(&Arr[i])
	}
	fmt.Print("\nTriple of even multiple number is = ")
	for i = 0; i < Arrsize; i++ {
		if Arr[i]%2 == 0 {
			fmt.Print(Arr[i]*Arr[i]*Arr[i], "  ")
			continue
		}

	}
	fmt.Println()
}
