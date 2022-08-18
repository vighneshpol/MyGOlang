/*write
 */

package main

import "fmt"

func main1() {

	var checkInteger int

	fmt.Print("Please enter any Integer = ")
	fmt.Scanf("%d", &checkInteger)

	if checkInteger >= 0 {
		fmt.Println("==> You entered an Positive Integer")
	} else {
		fmt.Println("==> You entered an Negative Integer")
	}
}
