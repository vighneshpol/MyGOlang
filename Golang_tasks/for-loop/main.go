package main

import "fmt"

func main() {
	// for i := 3; i < 4; i++ {
	// 	fmt.Println("Values is", i)
	// 	if i == 100 {
	// 		break
	// 	}
	// }

	for counter := 10; counter > 10; counter++ {
		fmt.Println("Infinite loop", counter)
		// counter++
		if counter == 15 {
			break

		}

	}

	numbers:= []int{1,2,3,4,5}
	for index,value:= range numbers{
		fmt.Printf("Index of number %d, and value is %d\n",index,value)
	}
	data:= "Hello World!"
	for index,char:= range data{
		fmt.Printf("Index of Data is %d, and Value is %c\n",index,char)
	}

}
