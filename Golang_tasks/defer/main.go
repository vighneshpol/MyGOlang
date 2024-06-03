package main

import(
	"fmt"
)

func add(a,b int)int{
	return a+b
}

func main(){
	fmt.Println("Starting of the prograM")
	data:= add(6,7)
	defer fmt.Println("Data is :",data)
	defer fmt.Println("Middle of the progra")
	fmt.Println("End of the program")

}