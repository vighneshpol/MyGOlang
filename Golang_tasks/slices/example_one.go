package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60}

	s = s[2:5]
	fmt.Println(s)
//30,40,50
	s = append(s, 70)
	fmt.Println(s)
//30,40,50,70
	s = s[:3]
	fmt.Println(s)
//30,40,50
	s = s[1:]
	fmt.Println(s)
//40,50
}
