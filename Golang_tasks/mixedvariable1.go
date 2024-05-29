/*Write a program to frind out type of variable name*/

package main

import "fmt"

func main() {
	var a, b, c = 3, 4.0, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	fmt.Printf("What is abc %T\n",a)
}
