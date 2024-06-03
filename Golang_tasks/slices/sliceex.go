package main

import "fmt"

func main() {
    // Declare a nil slice
    var s [3]int

    // Create a new slice with a specific capacity
    // s = make([]int, 3) // slice of length 3 with capacity 3

    // Assign values directly to each element
    s[0] = 1
    s[1] = 2
    s[2] = 3

    // Printing the updated slice
    fmt.Println("Updated slice:", s) // Output: Updated slice: [1, 2, 3]

    // Printing the length and capacity of the slice
    fmt.Println("Length:", len(s))   // Output: Length: 3
    fmt.Println("Capacity:", cap(s)) // Output: Capacity: 3
}
