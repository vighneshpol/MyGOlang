package main

import "fmt"

func main() {
a:= "Vighnesh"
b:= "Kavya"
c:= "Jighnesh"
d:= "Kavya"

if a==d||b==c||a==b {

	fmt.Println("one of the strings is matching")
	
}else{
	fmt.Println("a and b strings are not matching")
	error:=fmt.Errorf("%q is not matching with %q ", a, b)
	fmt.Println(error.Error())
}



}