package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two two two three four six five"
	count := strings.Count(str, "two")
	fmt.Println("Count",count)

	str="                                 Hello World!"
	trimmed:= strings.TrimSpace(str)
	fmt.Println("Trimmed sace is",trimmed)

}
