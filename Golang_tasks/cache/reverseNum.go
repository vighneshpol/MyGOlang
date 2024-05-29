package main

import {
	"fmt"
}

func reverseNumber(num int) int{
	rev:= 0
	for num > 0{
		rem:= num % 10
		rev= rev*10+ rem
		num /=10
	}
	return rev
}

func main(){
	num:= 12345
	rev:= reverseNumber(num)
	fmt.Print("original"+num)
	fmt.Print()
}