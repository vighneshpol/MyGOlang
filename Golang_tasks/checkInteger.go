/* Write a program to check whether number is positive or negative.
 */

package main

import "fmt"

func main() {

	var checkInteger int

	fmt.Print("Please enter any Integer = ")
	fmt.Scanf("%d", &checkInteger)
	

	if checkInteger >= 0 {
		fmt.Println("==> You entered an Positive Integer",checkInteger)
	} else {
		fmt.Println("==> You entered an Negative Integer")
	}
}
