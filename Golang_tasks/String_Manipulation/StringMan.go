// // Go program to illustrate how to
// // count the elements of the string
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Kashyap"
	K := strings.Count(str, "K")
	A := strings.Count(str, "a")
	S := strings.Count(str, "s")
	H := strings.Count(str, "h")
	Y := strings.Count(str, "y")

	fmt.Printf("[%v] string has %d characters of [K] \n", str, K)
	fmt.Printf("[%v] string has %d  characters of [A] \n ", str, A)
	fmt.Printf("[%v] string has %d  characters of [S] \n", str, S)
	fmt.Printf("[%v] string has %d  characters of [H] \n", str, H)
	fmt.Printf("[%v] string has %d  characters of [Y] \n", str, Y)

}
