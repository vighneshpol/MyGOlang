/*Write a program to do Sum of N series. N value should be entered by user.
Input :- Please enter number to get sum of N Series Output :- 5 Sum of 1 to 5 :- 15
*/

package main

import "fmt"

func main() {
	var N, sum int
	fmt.Print(" Please enter number to get sum of N Series:")
	fmt.Scan(&N)
	for i := 1; i <= N; i++ { // assigning 1 to i
		sum += i // sum = sum + i
	}
	fmt.Print("Sum of 1 to", N, ":-", sum, "\n")
}
