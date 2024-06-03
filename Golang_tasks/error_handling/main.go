package main

import "fmt"

func divide(a, b int) (int, error) {
	// return a/b
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be 0")
	}
	return a / b, nil

}

func main() {

	// fmt.Println("Started Error handling in function")
	data, err := divide(6, 2)
	if err != nil {
		fmt.Println("0 cannot be used")
	}
	fmt.Println("Data is :", data)
}
