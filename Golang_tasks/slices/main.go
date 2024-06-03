package main

import (
	"fmt"
)

func main() {

	// data := make([]int, 3, 5)

	// data = append(data, 4)
	// fmt.Println("Slice:", data)
	// fmt.Println("Length:", len(data))
	// fmt.Println("Capacity:", cap(data))
	// data = append(data, 45, 5)
	// fmt.Println("new Capacity:", cap(data))

	// // Declare array and convert to slice
	// var arr = []int{1, 2, 3, 4, 5}
	// // Convert slice to array
	// arr = arr[2:4]
	// fmt.Println("Converted slice is:", arr)

	//Declare a slice without initialization
	 slice:= []int{
		111,2222,3333,45,
	}
    fmt.Println("New Slice is:", slice)
    fmt.Println("Length of slice is:", len(slice))
    fmt.Println("Capacity of slice is:", cap(slice))
    slice = append(slice, 1, 2, 3, 4, 5)
    fmt.Println("Slice is:", slice)
    fmt.Println("Length of slice is:", len(slice))
    fmt.Println("Capacity of slice is:", cap(slice))
    slice = append(slice, 1, 2, 3, 4, 5)
    fmt.Println("Slice is:", slice)
    fmt.Println("Length of slice is:", len(slice))
    fmt.Println("Capacity of slice is:", cap(slice))
    slice = append(slice, 1, 2, 3, 4, 5)
}
